import Document from '@tiptap/extension-document'
import Gapcursor from '@tiptap/extension-gapcursor'
import Paragraph from '@tiptap/extension-paragraph'
import TableCell from '@tiptap/extension-table-cell'
import TableHeader from '@tiptap/extension-table-header'
import TableRow from '@tiptap/extension-table-row'
import Text from '@tiptap/extension-text'
import TextAlign from '@tiptap/extension-text-align'
import { EditorContent, useEditor } from '@tiptap/react'
import { useRef, useState } from 'react'
import StarterKit from '@tiptap/starter-kit'
import { CustomTable, CustomTableCell } from './custom-table-editor'
import { SidebarProvider } from '../ui/sidebar'
import DocumenetSidePanel from './document-side-panel'
import DocumentBuilderTools from './document-builder-tools'
import DOMPurify from 'isomorphic-dompurify'

export type PageProps = {
    htmlTemplate: string
    style: string
}

export interface addPageProps {
    setCurrentPage: (value: React.SetStateAction<number>) => void
    setPages: (value: React.SetStateAction<PageProps[]>) => void
    handleScrollFocusView: (pageIndex: number) => void
}

export default () => {
    const [pages, setPages] = useState<PageProps[]>([
        { htmlTemplate: '<div></div>', style: '' },
    ])
    // const pageHeight = 1056
    const [currentPage, setCurrentPage] = useState<number>(0)
    const editorRefFocus = useRef<Array<HTMLDivElement | null>>([])

    const editor = useEditor({
        extensions: [
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
            TableRow,
            TableHeader,
            TableCell,
            CustomTableCell,
            TextAlign.configure({
                types: ['heading', 'paragraph'],
            }),
        ],
        content: pages[currentPage].htmlTemplate,
        onUpdate: ({ editor }) => {
            const updatedContent = editor.getHTML()

            setPages((prevPages) => {
                const updatedPages = [...prevPages]
                updatedPages[currentPage].htmlTemplate = updatedContent
                return updatedPages
            })
        },
        editorProps: {
            attributes: {
                class: `table-toolbar-custom !w-full dark:text-black  `,
            },
        },
        autofocus: true,
        parseOptions: {
            preserveWhitespace: 'full',
        },
    })

    const NonEditableHTML = ({ htmlContent }: { htmlContent: string }) => {
        const sanitizedHTML = DOMPurify.sanitize(htmlContent)

        return (
            <div
                className="tiptap ProseMirror table-toolbar-custom !w-full"
                dangerouslySetInnerHTML={{ __html: sanitizedHTML }}
            />
        )
    }

    // const handlePageUpdate = () => {
    //     const container = editorRef.current;

    //     if (!container || !editor) return;

    //     const children = Array.from(container.children); // All child elements of the editor container
    //     let currentHeight = 0;
    //     let currentPageContent: HTMLElement[] = [];
    //     const newPages: PageProps[] = []; // Updated to use PageProps[]

    //     children.forEach((child) => {
    //         const element = child as HTMLElement;
    //         const elementHeight = element.offsetHeight;

    //         if (currentHeight + elementHeight > pageHeight) {
    //             // Add current page to the newPages array
    //             newPages.push({
    //                 htmlTemplate: currentPageContent.map((e) => e.outerHTML).join(''),
    //                 style: '', // Add custom styles if needed
    //             });

    //             // Reset current page content and height
    //             currentPageContent = [];
    //             currentHeight = 0;

    //             // Check if the current element itself exceeds one page
    //             if (elementHeight > pageHeight) {
    //                 const splitContent = splitElementContent(element, pageHeight);
    //                 if (splitContent) {
    //                     // Add the filled part to the current page
    //                     currentPageContent.push(splitContent.filled);

    //                     // Add the remaining part to a new page immediately
    //                     newPages.push({
    //                         htmlTemplate: [splitContent.remaining.outerHTML].join(''),
    //                         style: '',
    //                     });
    //                 }
    //             } else {
    //                 // If the element fits in a new page, push it to the next page
    //                 currentPageContent.push(element);
    //                 currentHeight += elementHeight;
    //             }
    //         } else {
    //             // Add the element to the current page
    //             currentPageContent.push(element);
    //             currentHeight += elementHeight;
    //         }
    //     });

    //     // Add the last page's content
    //     if (currentPageContent.length) {
    //         newPages.push({
    //             htmlTemplate: currentPageContent.map((e) => e.outerHTML).join(''),
    //             style: '',
    //         });
    //     }

    //     // Ensure pages are updated only if they differ
    //     if (JSON.stringify(newPages) !== JSON.stringify(pages)) {
    //         setPages(newPages);
    //     }
    // };

    // Utility to split element content
    // const splitElementContent = (element: HTMLElement, maxHeight: number) => {
    //     if (element.tagName === 'P' || element.tagName === 'DIV') {
    //         const clone = element.cloneNode(true) as HTMLElement;
    //         const originalContent = element.textContent || '';
    //         let currentContent = '';
    //         let remainingContent = originalContent;

    //         while (clone.scrollHeight > maxHeight && remainingContent.length) {
    //             currentContent += remainingContent.charAt(0);
    //             remainingContent = remainingContent.slice(1);
    //             clone.textContent = currentContent;
    //         }

    //         const remaining = element.cloneNode(true) as HTMLElement;
    //         remaining.textContent = remainingContent;

    //         return { filled: clone, remaining };
    //     }
    //     return null; // Return null if splitting isn't supported
    // };

    // const handleChange = (newContent: string) => {
    //     console.log('Editor onChange:', newContent);
    //     setPages((prevPages) => {
    //         const updatedPages = [...prevPages];
    //         updatedPages[currentPage].htmlTemplate = newContent;
    //         return updatedPages;
    //     });
    // };

    // Attach in `useEffect`
    // useEffect(() => {
    //     if (editor) {
    //         editor.on('update', ({ editor }) => {
    //             handleChange(editor);
    //         });

    //         return () => {
    //             editor.off('update');
    //         };
    //     }
    // }, [editor]);

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

    const addPage = () => {
        let newPage: PageProps = {
            style: '',
            htmlTemplate: '',
        }
        const newPageIndex = pages.length
        newPage = {
            style: '',
            htmlTemplate: '',
        }
        setPages((prevPages) => [...prevPages, newPage])
        setCurrentPage(newPageIndex)
        setTimeout(() => {
            handleScrollFocusView(newPageIndex)
        }, 0)
    }

    const switchToPage = (index: number) => {
        const updatedContent = editor?.getHTML()
        if (updatedContent !== undefined) {
            setPages((prevPages) => {
                const updatedPages = [...prevPages]
                updatedPages[currentPage].htmlTemplate = updatedContent
                return updatedPages
            })
            setCurrentPage(index)
            editor?.commands.setContent(pages[index].htmlTemplate)
        }
    }

    const handleScrollFocusView = (pageIndex: number) => {
        if (editorRefFocus.current[pageIndex]) {
            editorRefFocus.current[pageIndex]?.scrollIntoView({
                behavior: 'smooth',
                block: 'start',
            })
        }
    }

    if (!editor) {
        return null
    }

    return (
        <SidebarProvider>
            <div className="flex w-full">
                <div className="relative w-full">
                    <DocumentBuilderTools addPage={addPage} editor={editor} />
                    <div className="absolute flex w-full flex-col gap-10 overflow-auto bg-gray-100 pb-10 pt-24 dark:bg-black">
                        {pages.map((page, index) => {
                            const isCurrentPage = currentPage === index
                            return (
                                <div
                                    key={index}
                                    ref={(el) =>
                                        (editorRefFocus.current[index] = el)
                                    }
                                    onClick={() =>
                                        !isCurrentPage
                                            ? switchToPage(index)
                                            : ''
                                    }
                                    className={`page relative mx-auto h-[11in] w-[8.5in] cursor-pointer overflow-hidden rounded-lg bg-white p-[1in] shadow-md ${
                                        isCurrentPage
                                            ? 'shadow-xs border-4'
                                            : ''
                                    }`}
                                >
                                    {isCurrentPage ? (
                                        <EditorContent
                                            className="h-full w-full border-0 bg-white"
                                            editor={editor}
                                        />
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
                    </div>
                </div>
                <DocumenetSidePanel editor={editor} />
            </div>
        </SidebarProvider>
    )
}
