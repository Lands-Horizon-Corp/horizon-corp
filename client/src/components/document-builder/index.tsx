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
import DOMPurify from 'isomorphic-dompurify'
import { useDocumentBuilderStore } from '@/store/document-builder-store'
import TableRow from '@tiptap/extension-table-row'
import TableHeader from '@tiptap/extension-table-header'
import { Pagination } from './pagination'
// import { Editor, Node, NodeViewRendererOptions, mergeAttributes } from '@tiptap/core'
// import { DOMSerializer } from 'prosemirror-model'
// import { PageNode, PaginationExtension } from './pagination'


export type PageProps = {
    htmlTemplate: string
    style: string
}


// export const CustomNode = Node.create({
//   name: 'customNode',

//   group: 'block',

//   content: 'inline*',

//   parseHTML() {
//     return [{ tag: 'div[data-type="customNode"]' }];
//   },

//   renderHTML({ HTMLAttributes }) {
//     return ['div', { ...HTMLAttributes, 'data-type': 'customNode' }, 0];
//   },

//   addNodeView() {
//     return ({ node }: NodeViewRendererOptions) => {
//       const dom = document.createElement('div');
//       dom.setAttribute('data-type', 'customNode');
//       dom.textContent = node.textContent || 'Hello, custom node!';

//       return {
//         dom,
//       };
//     };
//   },
// });



export const DocumentBuilder = () => {
    const pageHeight = 1056

    const editorRefFocus = useRef<Array<HTMLDivElement | null>>([])
    const editorRef = useRef<HTMLDivElement | null>()

    const { pages, currentPage, switchToPage, updateJsonPageContent, handleAddOrSwitchPage, addPage} = useDocumentBuilderStore()

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
            // PageNode,
            // PaginationExtension,
            CustomTableCell,
            TextAlign.configure({
                types: ['heading', 'paragraph'],
            }),
            
        ],
        content: pages[currentPage].htmlTemplate,
        onUpdate: ({ editor }) => {
            //  handleAddOrSwitchPage(editor);
            //   console.log(editor)
            // const contentHeight = editor.view.dom.scrollHeight
            // const { state } = editor

            // updateJsonPageContent(currentPage, editor)
            // const pages = organizeContentIntoPages(editor, 420)
            // console.log('pages', pages)
            // console.log(contentHeight)
            // console.log('contentHeight', contentHeight)
            // if(contentHeight === 844) return
            // const nextPage = currentPage + 1
            // const hasNextPage = !!pages[currentPage + 1]

            // if(hasNextPage && contentHeight > 844){
            //     switchToPage(nextPage, editor)
            // }else{
            //     if(contentHeight > 844){
            //         addPage()
            //     }
            // }
        },
        editorProps: {
            attributes: {
                class: `table-toolbar-custom !w-full dark:text-black `,
            },
        },
        autofocus: true,
        parseOptions: {
            preserveWhitespace: 'full',
        },
    })
    // const organizeContentIntoPages = (
    //     editor: Editor,
    //     maxPageHeight: number
    // ) => {
    //     const updatePagesInStore = useDocumentBuilderStore.getState().setPages; 
    //     const { state } = editor
    //     const { schema, doc } = state

    //     const pages: Array<{
    //         htmlTemplate: string
    //         style: string
    //         jsonPage: any[]
    //     }> = []

    //     let currentPageNodes: any[] = []
    //     let currentHeight = 0

    //     // Iterate over the document nodes
    //     doc.content.forEach((node, offset) => {
    //         const nodeHeight = estimateNodeHeight(node)
    //         console.log(currentHeight, maxPageHeight, offset)
    //         // If adding the current node exceeds the max page height, finalize the current page
    //         if (currentHeight + nodeHeight > maxPageHeight) {
    //             const pageHtml = createHtmlFromNodes(currentPageNodes, editor)
    //             const pageJson = currentPageNodes.map((n) => n.toJSON())

    //             pages.push({
    //                 htmlTemplate: pageHtml,
    //                 style: '', // Add styles if needed
    //                 jsonPage: pageJson,
    //             })

    //             // Reset for the next page
    //             currentPageNodes = []
    //             currentHeight = 0
    //         }

    //         // Add the node to the current page
    //         currentPageNodes.push(node)
    //         currentHeight += nodeHeight
    //     })

    //     // Push the final page if any nodes remain
    //     if (currentPageNodes.length) {
    //         const pageHtml = createHtmlFromNodes(currentPageNodes, editor)
    //         const pageJson = currentPageNodes.map((n) => n.toJSON())

    //         pages.push({
    //             htmlTemplate: pageHtml,
    //             style: '', // Add styles if needed
    //             jsonPage: pageJson,
    //         })
    //     }

    //     console.log('Generated Pages:', pages)
    //     updatePagesInStore(pages)
    //     return pages
    // }

    // const createHtmlFromNodes = (nodes: any[], editor: any): string => {
    //     const { schema } = editor.state

    //     // Create a temporary document fragment from the nodes
    //     const fragment = DOMSerializer.fromSchema(schema).serializeFragment(
    //         schema.node('doc', null, nodes).content
    //     )

    //     // Convert the fragment into an HTML string
    //     return Array.from(fragment.childNodes)
    //         .map((node) => node.outerHTML)
    //         .join('')
    // }

    const estimateNodeHeight = (node: any) => {
        const baseHeight = 20 // Base height for a textblock
        if (node.type.name === 'paragraph' || node.type.name === 'heading') {
            return (
                baseHeight +
                Math.ceil(node.textContent.length / 50) * baseHeight
            )
        }
        if (node.type.name === 'table') {
            const rows = node.content.childCount
            return rows * 40 // Example row height
        }
        return baseHeight
    }

    const NonEditableHTML = ({ htmlContent }: { htmlContent: string }) => {
        const sanitizedHTML = DOMPurify.sanitize(htmlContent)
        return (
            <div
                className=""
                dangerouslySetInnerHTML={{ __html: sanitizedHTML }}
            />
        )
    }

    // console.log(editor.state)

    // const handlePageUpdate = () => {
    //     const container = editorRef.current

    //     if (!container || !editor) return
    //     const children = Array.from(container.children) // All child elements of the editor container
    //     let currentHeight = 0
    //     let currentPageContent: HTMLElement[] = []
    //     const newPages: PageProps[] = [] // Updated to use PageProps[]

    //     children.forEach((child) => {
    //         const element = child as HTMLElement
    //         const elementHeight = element.offsetHeight

    //         console.log(currentPageContent)

    //         if (currentHeight + elementHeight > pageHeight) {
    //             // Add current page to the newPages array
    //             newPages.push({
    //                 htmlTemplate: currentPageContent
    //                     .map((e) => e.outerHTML)
    //                     .join(''),
    //                 style: '', // Add custom styles if needed
    //             })

    //             // Reset current page content and height
    //             currentPageContent = []
    //             currentHeight = 0

    //             // Check if the current element itself exceeds one page
    //             if (elementHeight > pageHeight) {
    //                 const splitContent = splitElementContent(
    //                     element,
    //                     pageHeight
    //                 )
    //                 if (splitContent) {
    //                     // Add the filled part to the current page
    //                     currentPageContent.push(splitContent.filled)

    //                     // Add the remaining part to a new page immediately
    //                     newPages.push({
    //                         htmlTemplate: [
    //                             splitContent.remaining.outerHTML,
    //                         ].join(''),
    //                         style: '',
    //                     })
    //                 }
    //             } else {
    //                 // If the element fits in a new page, push it to the next page
    //                 currentPageContent.push(element)
    //                 currentHeight += elementHeight
    //             }
    //         } else {
    //             // Add the element to the current page
    //             currentPageContent.push(element)
    //             currentHeight += elementHeight
    //         }
    //     })

    //     // Add the last page's content
    //     if (currentPageContent.length) {
    //         newPages.push({
    //             htmlTemplate: currentPageContent
    //                 .map((e) => e.outerHTML)
    //                 .join(''),
    //             style: '',
    //         })
    //     }

    //     // Ensure pages are updated only if they differ
    //     if (JSON.stringify(newPages) !== JSON.stringify(pages)) {
    //         setPages(newPages)
    //     }
    // }

    // Utility to split element content
    // const splitElementContent = (element: HTMLElement, maxHeight: number) => {
    //     if (element.tagName === 'P' || element.tagName === 'DIV') {
    //         const clone = element.cloneNode(true) as HTMLElement
    //         const originalContent = element.textContent || ''
    //         let currentContent = ''
    //         let remainingContent = originalContent

    //         while (clone.scrollHeight > maxHeight && remainingContent.length) {
    //             currentContent += remainingContent.charAt(0)
    //             remainingContent = remainingContent.slice(1)
    //             clone.textContent = currentContent
    //         }

    //         const remaining = element.cloneNode(true) as HTMLElement
    //         remaining.textContent = remainingContent

    //         return { filled: clone, remaining }
    //     }
    //     return null // Return null if splitting isn't supported
    // }

    // const checkCursorOverflow = (editor: Editor) => {
    //     const { from } = editor.state.selection;
    //     const contentHeight = editor.view.dom.scrollHeight;

    //     console.log(editor.view.state.selection.$anchor)
    //     console.log('contentHeight', contentHeight)
    //     if (contentHeight > 844) {
    //         console.log('Cursor is at the bottom, adding a new page.');
    //         addPage();
    //     }
    // };

    // // Attach in `useEffect`
    // useEffect(() => {
    //     if (editor) {
    //         editor.on('update', ({editor}) => {
    //             checkCursorOverflow(editor)
    //         })
    //         // editor.on('update', handlePageUpdate);

    //         return () => {
    //             editor.off('update')
    //         }
    //     }
    // }, [editor])

    // useEffect(() => {
    //     if (editor) {
    //         editor.on('selectionUpdate', ({ editor }) => {
    //             const { from } = editor.state.selection; // Current cursor position
    //             const contentHeight = editor.view.dom.scrollHeight; // Get content height

    //             if (contentHeight >= pageHeight) {
    //                 // If content exceeds the current page height, add a new page
    //                 console.log('Cursor reached the end of the page. Adding new page.');
    //                 addPage();
    //             }
    //         });

    //         return () => {
    //             editor.off('selectionUpdate');
    //         };
    //     }
    // }, [editor, pages]);

    // const addPage = () => {
    //     let newPage: PageProps = {
    //         style: '',
    //         htmlTemplate: '',
    //     }
    //     const newPageIndex = pages.length
    //     newPage = {
    //         style: '',
    //         htmlTemplate: '',
    //     }
    //     setPages((prevPages) => [...prevPages, newPage])
    //     handleScrollFocusView(newPageIndex)
    // }

    // const deletePage = (index: number) => {
    //     setPages(pages.filter((page, idx)=> idx === index))
    // }

    // const switchToPage = (index: number) => {
    //     const updatedContent = editor?.getHTML()
    //     if (updatedContent !== undefined) {
    //         setPages((prevPages) => {
    //             const updatedPages = [...prevPages]
    //             updatedPages[currentPage].htmlTemplate = updatedContent
    //             return updatedPages
    //         })
    //         setCurrentPage(index)
    //         editor?.commands.setContent(pages[index].htmlTemplate)
    //     }
    // }

    // const handleScrollFocusView = (pageIndex: number) => {
    //     if (editorRefFocus.current[pageIndex]) {
    //         editorRefFocus.current[pageIndex]?.scrollIntoView({
    //             behavior: 'smooth',
    //             block: 'start',
    //         })
    //     }
    // }


    if (!editor) {
        return null
    }
    return (
        <SidebarProvider>
            <div className="flex w-full">
                <div className="relative w-full">
                    <DocumentBuilderTools editor={editor} />
                    <div className=' !mt-10 flex justify-center gap-3 page relative mx-auto h-fit w-[8.5in] cursor-pointer overflow-hidden rounded-lg border-0 bg-white p-[1in] '>
                    <EditorContent className='' editor={editor} />
                    </div>
                    {/* <div
                        ref={(el) => {
                            editorRef.current = el
                        }}
                        className="absolute flex w-full flex-col gap-10 overflow-auto bg-gray-100 pb-10 pt-24 dark:bg-black"
                    >
                        {pages.map((page, index) => {
                            const isCurrentPage = currentPage === index
                            return (
                                <div
                                    key={index}
                                    ref={(el) => {
                                        editorRefFocus.current[index] = el
                                    }}
                                    onClick={() =>
                                        !isCurrentPage
                                            ? switchToPage(index, editor)
                                            : ''
                                    }
                                    className={`page relative mx-auto h-fit w-[8.5in] cursor-pointer overflow-hidden rounded-lg border-0 bg-white p-[1in] ${
                                        isCurrentPage ? 'shadow-md' : ''
                                    }`}
                                >
                                    {isCurrentPage ? (
                                        <EditorContent className='border h-fit' editor={editor} />
                                    ) : (
                                        <NonEditableHTML
                                            htmlContent={page.htmlTemplate}
                                        />
                                    )}
                                    <div className="absolute bottom-5 right-5 flex justify-center">
                                        <span className="text-xs text-muted-foreground">
                                            {index + 1}
                                        </span>
                                    </div>
                                </div>
                            )
                        })}
                    </div> */}
                </div>
                <DocumenetSidePanel editor={editor} />
            </div>
        </SidebarProvider>
    )
}

export default DocumentBuilder
