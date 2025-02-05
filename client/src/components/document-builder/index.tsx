import Document from '@tiptap/extension-document'
import Gapcursor from '@tiptap/extension-gapcursor'
import Paragraph from '@tiptap/extension-paragraph'
import TableCell from '@tiptap/extension-table-cell'
import Text from '@tiptap/extension-text'
import TextAlign from '@tiptap/extension-text-align'
import { EditorContent, useEditor } from '@tiptap/react'
import { useEffect, useRef, useState } from 'react'
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
import html2canvas from 'html2canvas'
import jsPDF from 'jspdf'
// @ts-ignore
import * as html2pdf from 'html2pdf.js'
// import { useTellerTransactionStore } from '@/store/report-store'
// import { useReactTable, getCoreRowModel } from '@tanstack/react-table'
import ReactDOMServer from 'react-dom/server'
// import Test from './test'
import { useLocation } from '@tanstack/react-router'
import UserTeller from '@/modules/owner/pages/reports/user-teller/user-teller'

const REPORT_COMPONENTS: Record<string, () => JSX.Element> = {
    '/reports/document/user-teller': UserTeller,
}

export const DocumentBuilder = () => {
    
    const location = useLocation()

    console.log(location.pathname)

    const editorRefFocus = useRef<Array<HTMLDivElement | null>>([])
    const reportTemplateRef = useRef<HTMLDivElement>(null)

    // const { pages, currentPage } = useDocumentBuilderStore()
    // const [editorContent, setEditorContent] = useState(pages[currentPage].htmlTemplate)

    useEffect(() => {
        useDocumentBuilderStore.setState({
            editorRefFocus: editorRefFocus.current,
        })
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
        content: ReactDOMServer.renderToStaticMarkup(<UserTeller />),
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

    const [tableWidth, setTableWidth] = useState(0)

    const handleGeneratePdf = async () => {
        const input = reportTemplateRef.current

        if (!input) return

        try {
            const canvas = await html2canvas(input)
            const imgData = canvas.toDataURL('image/png')

            const pdf = new jsPDF({
                orientation: 'landscape',
                unit: 'px',
                format: 'a4',
            })

            const pdfWidth = pdf.internal.pageSize.getWidth()

            const computedPdfWidth = tableWidth - pdfWidth

            const pdfHeight = (canvas.height * pdfWidth) / canvas.width

            // console.log('pdfWidth', pdfWidth);
            // console.log('pdfHeight', pdfHeight);
            // console.log( 'computedPdfWidth', computedPdfWidth);

            // pdf.addImage(imgData, 'PNG', 0, 0, pdfWidth, pdfHeight)
            // pdf.addPage('a4','portrait');
            // pdf.save('report.pdf')

            const opt = {
                margin: 0.1,
                filename: 'myfile.pdf',
                image: { type: 'jpeg', quality: 1 },
                // pagebreak: { avoid: "tr", mode: "css", before: "#nextpage1", after: "1cm" },
                html2canvas: {
                    scale: 4,
                    useCORS: true,
                    dpi: 192,
                    letterRendering: true,
                },
                jsPDF: {
                    unit: 'in',
                    format: 'legal',
                    orientation: 'landscape',
                    putTotalPages: true,
                },
            }
            html2pdf.default(input, opt)
        } catch (error) {
            console.error('Error generating PDF:', error)
        }
    }
    // const location = useLocation();
    // const hasInserted = useRef(false);

    // const { data, config, columns } = useTellerTransactionStore(
    //     (state) => state
    // )

    // const table = useReactTable({
    //     columns: columns,
    //     data: data,
    //     getCoreRowModel: getCoreRowModel(),
    // })

    // const TellerReportConfig = {
    //     ...config,
    //     table: table,
    // }

    return (
        <SidebarProvider defaultOpen={false}>
            <div className="flex h-full w-full ">
                <div className="relative w-full bg-secondary py-5 flex flex-col space-y-5 dark:bg-black">
                    <DocumentBuilderTools
                        handleGeneratePdf={handleGeneratePdf}
                        editor={editor}
                    />
                    <div
                        ref={reportTemplateRef}
                        className="tableWrapper h-full w-full justify-center"
                    >
                        <EditorContent
                            className="page shadow-ls mx relative mb-20 flex h-fit max-h-fit min-h-[11.5in] w-full cursor-pointer justify-center gap-3 overflow-hidden rounded-lg bg-white p-[.5in] text-start dark:bg-secondary dark:text-white"
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
