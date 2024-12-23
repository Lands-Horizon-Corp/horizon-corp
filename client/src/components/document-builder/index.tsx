import Document from '@tiptap/extension-document'
import Gapcursor from '@tiptap/extension-gapcursor'
import Paragraph from '@tiptap/extension-paragraph'
import TableCell from '@tiptap/extension-table-cell'
import Text from '@tiptap/extension-text'
import TextAlign from '@tiptap/extension-text-align'
import { EditorContent, useEditor } from '@tiptap/react'
import { useEffect, useRef } from 'react'
import StarterKit from '@tiptap/starter-kit'
import {
    CustomTable,
    CustomTableCell,
    CustomTableRow,
} from './document-custom-table-editor'
import { SidebarProvider } from '../ui/sidebar'
import DocumenetSidePanel from './document-side-panel'
import DocumentBuilderTools from './document-builder-toolbar-container'
import { useDocumentBuilderStore } from '@/store/document-builder-store'
import TableRow from '@tiptap/extension-table-row'
import TableHeader from '@tiptap/extension-table-header'
import { Pagination } from './pagination'

export type PageProps = {
    htmlTemplate: string
    style: string
}

export const DocumentBuilder = () => {
    const editorRefFocus = useRef<Array<HTMLDivElement | null>>([])

    const { pages, currentPage } = useDocumentBuilderStore()

    useEffect(() => {
        useDocumentBuilderStore.setState({ editorRefFocus: editorRefFocus.current })
    }, [])

    const editor = useEditor({
        extensions: [
            Pagination.configure({
                pageHeight: 1056,
                pageWidth: 816,
                pageMargin: 96,
            }),
            StarterKit.configure({
                bulletList: {
                    keepMarks: false,
                    keepAttributes: false,
                },
            }),
            Document,
            Paragraph,
            Text,
            Gapcursor,
            CustomTable.configure({ resizable: true }),
            CustomTableRow,
            TableRow,
            TableCell,
            TableHeader,
            CustomTableCell,
            TextAlign.configure({
                types: ['heading', 'paragraph'],
            }),
        ],
        content: pages[currentPage].htmlTemplate,
        editorProps: {
            attributes: {
                class: `table-toolbar-custom`,
            },
        },
        autofocus: true,
        parseOptions: {
            preserveWhitespace: 'full',
        },
    })

    if (!editor) {
        return null
    }

    return (
        <SidebarProvider>
            <div className="flex w-full h-full">
                <div className="relative w-full bg-secondary">
                    <DocumentBuilderTools editor={editor} />
                    <div className="h-full w-full justify-center pt-20">
                        <EditorContent
                            className="page relative mx-auto mb-20 shadow-ls bg-white dark:text-white flex min-h-[11.5in] h-fit max-h-fit w-[8.5in] cursor-pointer justify-center gap-3 overflow-hidden rounded-lg p-[1in] text-start"
                            editor={editor}
                        />
                    </div>
                </div>
                <DocumenetSidePanel editor={editor} />
            </div>
        </SidebarProvider>
    )
}

export default DocumentBuilder
