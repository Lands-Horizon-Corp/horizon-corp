import { Editor, JSONContent, Node } from '@tiptap/react';
import { create } from 'zustand'

interface PageProps {
    htmlTemplate: string;
    style: string;
    JsonPage?: JSONContent,
}

interface DocumentBuilderState {
    pages: PageProps[];
    currentPage: number;
    height: number,
}

interface DocumentBuilderActions {
    addPage: (editor: Editor) => void;
    deletePage: (index: number) => void;
    switchToPage: (index: number, editor: Editor) => void;
    setCurrentPage: (index: number) => void;
    updateJsonPageContent: (index:number, content: Editor) => void;
    handleScrollFocusView: (pageIndex: number) => void;
    setHeight: (height: number) => void;
}

interface DocumentBuilderStore extends DocumentBuilderState, DocumentBuilderActions {
    editorRefFocus: Array<HTMLDivElement | null>;
}

export const useDocumentBuilderStore = create<DocumentBuilderStore>((set, get) => ({
    pages: [
        { htmlTemplate: '<div>start Typing...</div>', style: '', JsonPage:[{type: 'doc', content:[{type:"text", text:"hello"}]}] },
    ],
    currentPage: 0,
    height: 0,
    editorRefFocus: [],
    addPage: (editor: Editor) => {
        const newPage: PageProps = {
            htmlTemplate: '',
            style: '',
            JsonPage: [{type: 'doc', content:[{type:"text", text:"hello..."}]}]
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
        const updatedContent = editor?.getHTML()
        console.log("switch")
        if (updatedContent) {
            pages[currentPage].htmlTemplate = updatedContent;
            set((state) => ({
                ...state,
                pages,
                currentPage: index
            }));
            editor?.commands.setContent(pages[index].htmlTemplate)
        }
    },
    setCurrentPage: (index: number) => {
        set((state) => ({
            ...state,
            currentPage: index,
        }));
    },
    updateJsonPageContent: (index:number, editor: Editor) => {
        set((state) => {
            const content = editor?.getJSON()
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
    setHeight: (height: number)=>{
        set((state)=>({
            ...state,
            height:height 
        })) 
    },  

}));
