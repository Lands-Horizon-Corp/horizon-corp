import { Editor, JSONContent } from '@tiptap/react';
import StarterKit from '@tiptap/starter-kit';
import { create } from 'zustand';
import { Node as ProseMirrorNode } from 'prosemirror-model';

interface PageProps {
  htmlTemplate: string;
  style: string;
  JsonPage: JSONContent;
}

interface DocumentBuilderState {
  pages: PageProps[];
  currentPage: number;
  height: number;
}

interface DocumentBuilderActions {
  addPage: (editor: Editor) => void;
  deletePage: (index: number) => void;
  switchToPage: (index: number, editor: Editor) => void;
  setCurrentPage: (index: number) => void;
  updateJsonPageContent: (index: number, editor: Editor) => void;
  handleScrollFocusView: (pageIndex: number) => void;
  setHeight: (height: number) => void;
  setPages: (pages: PageProps[]) => void;
  handleAddOrSwitchPage: (editor: Editor) => void;
}

interface DocumentBuilderStore extends DocumentBuilderState, DocumentBuilderActions {
  editorRefFocus: Array<HTMLDivElement | null>;
}

const createHtmlFromJson = (jsonContent: JSONContent) => {
  const editor = new Editor({
    extensions: [StarterKit],
    content: jsonContent,
  });
  return editor.getHTML();
};


// Helper function to estimate the height of a ProseMirrorNode
const estimateNodeHeightWithToDOM = (node: ProseMirrorNode): number => {
  if (!node || !node.type?.spec?.toDOM) {
    console.log('No toDOM method available');
    return 0; // Fallback for invalid or undefined nodes
  }

  // Generate a temporary DOM element using `toDOM`
  const toDOMResult = node.type.spec.toDOM(node);
  let dom: HTMLElement | null = null;

  if (Array.isArray(toDOMResult)) {
    // Handle the array structure returned by `toDOM`
    const [tagName, attrs, content] = toDOMResult;
    dom = document.createElement(tagName);

    // Apply attributes
    if (attrs && typeof attrs === 'object') {
      Object.entries(attrs).forEach(([key, value]) => {
        dom?.setAttribute(key, value as string);
      });
    }

    // Add content if applicable
    if (typeof content === 'string' && dom) {
      dom.textContent = content;
    }
  } else if (toDOMResult instanceof HTMLElement) {
    dom = toDOMResult;
  }

  if (!dom) {
    console.log('Failed to create DOM element');
    return 0; // Fallback for invalid DOM generation
  }

  // Create an element and attach it temporarily to the document to measure
  const wrapper = document.createElement('div');
  wrapper.style.position = 'absolute'; // Prevent affecting layout
  wrapper.style.visibility = 'hidden'; // Prevent visibility issues
  wrapper.appendChild(dom);

  document.body.appendChild(wrapper); // Attach to DOM to calculate dimensions

  const height = wrapper.getBoundingClientRect().height;

  // Clean up the temporary element
  document.body.removeChild(wrapper);

  return height;
};


// Split content by height
export const splitContentByHeight = (
  editor: Editor,
  content: JSONContent,
  maxHeight: number
): [JSONContent[], JSONContent[]] => {
  const doc = editor.schema.nodeFromJSON(content);
  const state = editor.state;

  const currentContent: JSONContent[] = [];
  const overflowContent: JSONContent[] = [];
  let currentHeight = 0;
  
  console.log(currentContent)

  doc.content.forEach((node: ProseMirrorNode) => {
    const nodeHeight = estimateNodeHeightWithToDOM(node);

    if (currentHeight + nodeHeight <= maxHeight) {
      currentContent.push(node.toJSON());
      currentHeight += nodeHeight;
    } else {
      overflowContent.push(node.toJSON());
    }
  });

  return [currentContent, overflowContent];
};

export const calculateContentHeight = (editor: Editor, content: JSONContent) => {
  const doc = editor.schema.nodeFromJSON(content);
  const state = editor.state
  let totalHeight = 0;
  const node = editor.state.doc.child(0);
  console.log( 'height',node, estimateNodeHeightWithToDOM(node))

  doc.content.forEach(node => {
    totalHeight += estimateNodeHeightWithToDOM(node);
  });

  return totalHeight;
};

export const useDocumentBuilderStore = create<DocumentBuilderStore>((set, get) => ({
  pages: [
    {
      htmlTemplate: '<div>start Typing...</div>',
      style: '',
      JsonPage: { type: 'doc', content: [{ type: 'text', text: '' }] },
    },
  ],
  currentPage: 0,
  height: 450,
  editorRefFocus: [],
  setPages: (pages: PageProps[]) => {
    set({ pages });
  },
  addPage: (editor: Editor) => {
    const newPage: PageProps = {
      htmlTemplate: '',
      style: '',
      JsonPage: { type: 'doc', content: [{ type: 'text', text: '' }] },
    };
    const newPageIndex = get().pages.length;
    set((state) => ({
      ...state,
      pages: [...state.pages, newPage],
    }));

    setTimeout(() => {
      get().handleScrollFocusView(newPageIndex);
    }, 0);
  },
  deletePage: (index: number) => {
    set((state) => ({
      ...state,
      pages: state.pages.filter((_, idx) => idx !== index),
      currentPage: state.currentPage > index ? state.currentPage - 1 : state.currentPage,
    }));
  },
  switchToPage: (index: number, editor: Editor) => {
    const { pages, currentPage } = get();
    const updatedContent = editor.getHTML();
    if (updatedContent) {
      pages[currentPage].htmlTemplate = updatedContent;
      set((state) => ({
        ...state,
        pages,
        currentPage: index,
      }));
      editor.commands.setContent(pages[index].JsonPage || { type: 'doc', content: [] });
    }
  },
  setCurrentPage: (index: number) => {
    set((state) => ({
      ...state,
      currentPage: index,
    }));
  },
  handleAddOrSwitchPage: (editor: Editor) => {
    const { pages, currentPage, addPage } = get();
    const nextPageExists = currentPage + 1 < pages.length;
    const height = 780
    const currentPageContent = editor.getJSON();
    const currentPageHeight = calculateContentHeight(editor, currentPageContent);
    const contentHeight = editor.view.dom.scrollHeight

    if (contentHeight > height) {   
      const [currentContent, overflowContent] = splitContentByHeight(editor, currentPageContent, height);

      const updatedPages = [...pages];

      updatedPages[currentPage] = {
        ...updatedPages[currentPage],
        JsonPage: { type: 'doc', content: currentContent },
        htmlTemplate: createHtmlFromJson({ type: 'doc', content: currentContent }) || '',
      };

      if (nextPageExists) {
        console.log('nextpage')
        updatedPages[currentPage + 1].JsonPage = {
          type: 'doc',
          content: [
            ...(updatedPages[currentPage + 1].JsonPage?.content || []),
            ...overflowContent,
          ],
        };
        updatedPages[currentPage + 1].htmlTemplate = createHtmlFromJson(
          updatedPages[currentPage + 1].JsonPage || { type: 'doc', content: [] }
        ) || '';
        set({ pages: updatedPages, currentPage: currentPage + 1 });
      } else {
        addPage(editor)
      }
    }
  },
  updateJsonPageContent: (index: number, editor: Editor) => {
    set((state) => {
      const content = editor.getJSON();
      const updatedPages = [...state.pages];
      updatedPages[index].JsonPage = content;
      return {
        ...state,
        pages: updatedPages,
      };
    });
  },
  handleScrollFocusView: (pageIndex: number) => {
    const editorRefFocus = get().editorRefFocus;
    if (editorRefFocus[pageIndex]) {
      editorRefFocus[pageIndex]?.scrollIntoView({
        behavior: 'smooth',
        block: 'start',
      });
    }
  },
  setHeight: (height: number) => {
    set((state) => ({
      ...state,
      height: height,
    }));
  },
}));
