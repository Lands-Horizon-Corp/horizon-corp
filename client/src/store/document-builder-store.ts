import { Editor } from '@tiptap/react';
import { create } from 'zustand'

interface PageProps {
    htmlTemplate: string;
    style: string;
}

interface DocumentBuilderState {
    pages: PageProps[];
    currentPage: number;
}

interface DocumentBuilderActions {
    addPage: () => void;
    deletePage: (index: number) => void;
    switchToPage: (index: number, editor: Editor) => void;
    setCurrentPage: (index: number) => void;
    updatePageContent: (index: number, content: string) => void;
    handleScrollFocusView: (pageIndex: number) => void;
}

interface DocumentBuilderStore extends DocumentBuilderState, DocumentBuilderActions {
    editorRefFocus: Array<HTMLDivElement | null>;
}

export const useDocumentBuilderStore = create<DocumentBuilderStore>((set, get) => ({
    pages: [
        { htmlTemplate: '<div></div>', style: '' },
    ],
    currentPage: 0,
    editorRefFocus: [],
    addPage: () => {
        const newPage: PageProps = {
            htmlTemplate: '',
            style: '',
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

    updatePageContent: (index: number, content: string) => {
        set((state) => {
            const updatedPages = [...state.pages];
            updatedPages[index].htmlTemplate = content;
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

}));
