import Document from '@tiptap/extension-document'
import Gapcursor from '@tiptap/extension-gapcursor'
import Paragraph from '@tiptap/extension-paragraph'
import Table from '@tiptap/extension-table'
import TableCell from '@tiptap/extension-table-cell'
import TableHeader from '@tiptap/extension-table-header'
import TableRow from '@tiptap/extension-table-row'
import Text from '@tiptap/extension-text'
import TextAlign from '@tiptap/extension-text-align'

import { Editor, EditorContent, useEditor } from '@tiptap/react'
import { useEffect, useRef, useState } from 'react'
import { THeadingLevel } from '../text-editor'
import Toolbar from '../text-editor/toolbar'
import ReactDOMServer from 'react-dom/server'
import { Button } from '../ui/button'
import { FaTable } from 'react-icons/fa'
import { Ghost, SplitIcon, Table2 } from 'lucide-react'
import StarterKit from '@tiptap/starter-kit'
import insertColumnRight from '@/assets/svg/insertColumnRight.svg'
import insertColumnLeft from '@/assets/svg/insertColumnLeft.svg'
import insertRowBefore from '@/assets/svg/insertRowUp.svg'
import insertRowAfter from '@/assets/svg/insertRowDown.svg'
import deleteSelectedRows from '@/assets/svg/deleteSelectedRows.svg'
import deleteSelectedColumns from '@/assets/svg/deleteSelectedColumns.svg'
import mergeCells from '@/assets/svg/mergeCells.svg'

import { Input } from '@/components/ui/input'

import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuPortal,
    DropdownMenuSeparator,
    DropdownMenuSub,
    DropdownMenuSubContent,
    DropdownMenuSubTrigger,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import {
    RxTrashIcon,
    TextAlignCenterIcon,
    TextAlignJustifyIcon,
    TextAlignLeftIcon,
    TextAlignRightIcon,
} from '../icons'
import { cn } from '@/lib'

import { CgBorderLeft } from 'react-icons/cg'
import { CgBorderRight } from 'react-icons/cg'
import { CgBorderBottom } from 'react-icons/cg'
import { CgBorderTop } from 'react-icons/cg'
import { CgBorderAll } from 'react-icons/cg'
import { MdOutlineBorderClear } from 'react-icons/md'

const headers = ['Id', 'Name', 'Bday', 'Age', 'Gender']
const Person: any[] = []

for (let i = 1; i <= 100; i++) {
    const randomYear = Math.floor(Math.random() * (2005 - 1970 + 1)) + 1970 // Random year between 1970 and 2005
    const randomMonth = Math.floor(Math.random() * 12) + 1 // Random month between 1 and 12
    const randomDay = Math.floor(Math.random() * 28) + 1 // Random day between 1 and 28
    const bday = `${randomYear}-${String(randomMonth).padStart(2, '0')}-${String(randomDay).padStart(2, '0')}`
    const age = new Date().getFullYear() - randomYear // Calculate age based on the year
    const gender = Math.random() > 0.5 ? 'Male' : 'Female' // Random gender
    const name = `Person ${i}` // Generate a placeholder name

    Person.push({
        id: i,
        bday: bday,
        name: name,
        age: age,
        gender: gender,
    })
}

interface TableRowData {
    [key: string]: string | number | null | undefined
}

interface TableContentProps {
    headers: string[]
    data: TableRowData[]
}

const TableContent: React.FC<TableContentProps> = ({ headers, data }) => {
    return (
        <table>
            <thead>
                <tr>
                    {headers.map((header) => (
                        <th key={header}>{header}</th>
                    ))}
                </tr>
            </thead>
            <tbody>
                {data.map((row, rowIndex) => (
                    <tr key={rowIndex}>
                        {headers.map((header) => (
                            <td key={header}>
                                {row[header.toLowerCase()] || ''}
                            </td>
                        ))}
                    </tr>
                ))}
            </tbody>
        </table>
    )
}

export const generateTableHTML = (
    headers: string[],
    data: TableRowData[]
): string => {
    return ReactDOMServer.renderToStaticMarkup(
        <TableContent headers={headers} data={data} />
    )
}

interface TableToolsProps {
    editor: Editor | null
    dimensions: {
        row: number
        col: number
    }
    handleInputChange: (event: React.ChangeEvent<HTMLInputElement>) => void
}

const TableTools = ({
    editor,
    dimensions,
    handleInputChange,
}: TableToolsProps) => {
    if (!editor) {
        return null
    }

    const insertTable = () => {
        console.log(editor.commands)
        editor.commands.insertTable({
            rows: dimensions.row,
            cols: dimensions.col,
            withHeaderRow: true,
        })
    }

    const tableInputClass =
        'rounded-[10px] border border-[#4D4C4C]/20 bg-white/50 dark:bg-secondary/70 focus:border-none focus-visible:ring-1 focus-visible:ring-primary focus-visible:ring-offset-0 placeholder:text-[#838383]'

    const toggleBorder = (side: string, style: string) => {
        const currentBorder = editor.getAttributes('tableCell')[side]
        editor.commands.setCellAttribute(
            side,
            currentBorder === style ? null : style
        )
    }

    return (
        <>
            <DropdownMenu>
                <DropdownMenuTrigger>
                    <Button variant={'ghost'} className="">
                        <Table2 size={20} />
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="[&>div>img]:mr-2 [&>div]:text-xs">
                    <DropdownMenuLabel>Create Table</DropdownMenuLabel>
                    <DropdownMenuSeparator />
                    <DropdownMenuSub>
                        <DropdownMenuSubTrigger>
                            <FaTable size={19} className="mr-2" /> Insert Table
                        </DropdownMenuSubTrigger>
                        <DropdownMenuPortal>
                            <DropdownMenuSubContent className="flex flex-col gap-2 p-2 [&>div>span]:text-sm [&>div>span]:text-muted-foreground [&>div]:flex [&>div]:items-center [&>div]:space-x-2">
                                <div>
                                    <span>row</span>
                                    <Input
                                        name="row"
                                        value={dimensions.row}
                                        type="number"
                                        onChange={handleInputChange}
                                        onClick={(e) => e.preventDefault()}
                                        onPointerDown={(e) =>
                                            e.stopPropagation()
                                        }
                                        className={cn(tableInputClass)}
                                    ></Input>
                                </div>
                                <div>
                                    <span>col</span>
                                    <Input
                                        name="col"
                                        type="number"
                                        value={dimensions.col}
                                        onChange={handleInputChange}
                                        onClick={(e) => e.preventDefault()}
                                        onPointerDown={(e) =>
                                            e.stopPropagation()
                                        }
                                        className={cn(tableInputClass)}
                                    ></Input>
                                </div>
                                <Button onClick={insertTable}>create</Button>
                            </DropdownMenuSubContent>
                        </DropdownMenuPortal>
                    </DropdownMenuSub>
                    <DropdownMenuSeparator />
                    <DropdownMenuSub>
                        <DropdownMenuSubTrigger>
                            <FaTable size={19} className="mr-2" /> Border
                        </DropdownMenuSubTrigger>
                        <DropdownMenuPortal>
                            <DropdownMenuSubContent className="grid grid-cols-2">
                                <Button
                                    className={cn('justify-start')}
                                    variant={'ghost'}
                                    onClick={() => {
                                        editor.commands.setCellAttribute(
                                            'borderLeft',
                                            '1px solid black'
                                        )
                                        editor.commands.setCellAttribute(
                                            'borderRight',
                                            '1px solid black'
                                        )
                                        editor.commands.setCellAttribute(
                                            'borderTop',
                                            '1px solid black'
                                        )
                                        editor.commands.setCellAttribute(
                                            'borderBottom',
                                            '1px solid black'
                                        )
                                    }}
                                >
                                    <CgBorderAll size={24} className="mr-2" />
                                    border all
                                </Button>
                                <Button
                                    variant={'ghost'}
                                    className={cn('justify-start')}
                                    onClick={() => {
                                        // Remove all borders from the selected cell
                                        editor.commands.setCellAttribute(
                                            'borderLeft',
                                            null
                                        )
                                        editor.commands.setCellAttribute(
                                            'borderRight',
                                            null
                                        )
                                        editor.commands.setCellAttribute(
                                            'borderTop',
                                            null
                                        )
                                        editor.commands.setCellAttribute(
                                            'borderBottom',
                                            null
                                        )
                                    }}
                                >
                                    <MdOutlineBorderClear
                                        size={18}
                                        className="mr-2"
                                    />
                                    Unborder all
                                </Button>
                                <Button
                                    variant={'ghost'}
                                    size={'sm'}
                                    className={cn('justify-start')}
                                    onClick={() =>
                                        toggleBorder(
                                            'borderLeft',
                                            '1px solid black'
                                        )
                                    }
                                >
                                    <CgBorderLeft size={24} className="mr-2" />
                                    border left
                                </Button>
                                <Button
                                    size={'sm'}
                                    variant={'ghost'}
                                    className={cn('justify-start')}
                                    onClick={() =>
                                        toggleBorder(
                                            'borderRight',
                                            '1px solid black'
                                        )
                                    }
                                >
                                    <CgBorderRight size={24} className="mr-2" />
                                    border Right
                                </Button>
                                <Button
                                    size={'sm'}
                                    variant={'ghost'}
                                    className={cn('justify-start')}
                                    onClick={() =>
                                        toggleBorder(
                                            'borderTop',
                                            '1px solid black'
                                        )
                                    }
                                >
                                    <CgBorderTop size={24} className="mr-2" />
                                    border Top
                                </Button>
                                <Button
                                    size={'sm'}
                                    variant={'ghost'}
                                    className={cn('justify-start')}
                                    onClick={() =>
                                        toggleBorder(
                                            'borderBottom',
                                            '1px solid black'
                                        )
                                    }
                                >
                                    <CgBorderBottom
                                        size={24}
                                        className="mr-2"
                                    />
                                    border Bottom
                                </Button>
                            </DropdownMenuSubContent>
                        </DropdownMenuPortal>
                    </DropdownMenuSub>
                    <DropdownMenuSeparator />
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().addColumnBefore().run()
                        }
                    >
                        <img src={insertColumnRight} />
                        Add Column Before
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().addColumnAfter().run()
                        }
                    >
                        <img src={insertColumnLeft} />
                        Add Column After
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().addRowBefore().run()
                        }
                    >
                        <img src={insertRowBefore} />
                        Add Row Before
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().addRowAfter().run()
                        }
                    >
                        <img src={insertRowAfter} />
                        Add Row After
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().deleteColumn().run()
                        }
                    >
                        <img src={deleteSelectedColumns} />
                        Delete Column
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() => editor.chain().focus().deleteRow().run()}
                    >
                        <img src={deleteSelectedRows} />
                        Delete Row
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().mergeCells().run()
                        }
                    >
                        <img src={mergeCells} />
                        Merge cells
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().deleteTable().run()
                        }
                    >
                        <RxTrashIcon size={20} className="mr-2" />
                        Delete Table
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() => editor.chain().focus().splitCell().run()}
                    >
                        <SplitIcon size={18} className="mr-2" />
                        Split Cells
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() => {
                            const currentColor =
                                editor.getAttributes(
                                    'tableCell'
                                ).backgroundColor
                            editor.commands.setCellAttribute(
                                'backgroundColor',
                                currentColor === 'yellow' ? null : 'yellow'
                            )
                        }}
                    >
                        Toggle Highlight
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().toggleHeaderColumn().run()
                        }
                    >
                        Toggle header Column
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().toggleHeaderCell().run()
                        }
                    >
                        Toggle Header Cell
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() =>
                            editor.chain().focus().mergeOrSplit().run()
                        }
                    >
                        Merge Or Split
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() => editor.chain().focus().fixTables().run()}
                    >
                        Fix Tables
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </>
    )
}

interface TextAlignToolProps {
    editor: Editor | null
}

const TextAlignTool = ({ editor }: TextAlignToolProps) => {
    if (!editor) return null
    return (
        <DropdownMenu>
            <DropdownMenuTrigger>
                <Button variant={'ghost'} className="">
                    <TextAlignLeftIcon size={17} />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="[&>div>svg]:mr-2 [&>div]:text-xs">
                <DropdownMenuLabel>Text Align</DropdownMenuLabel>
                <DropdownMenuItem
                    onClick={() => {
                        editor.commands.setTextAlign('left')
                    }}
                >
                    <TextAlignLeftIcon size={17} />
                    align left
                </DropdownMenuItem>
                <DropdownMenuItem
                    onClick={() => {
                        editor.commands.setTextAlign('right')
                    }}
                >
                    <TextAlignRightIcon size={17} />
                    align right
                </DropdownMenuItem>
                <DropdownMenuItem
                    onClick={() => {
                        editor.commands.setTextAlign('center')
                    }}
                >
                    <TextAlignCenterIcon size={17} />
                    align center
                </DropdownMenuItem>
                <DropdownMenuItem
                    onClick={() => {
                        editor.commands.setTextAlign('justify')
                    }}
                >
                    <TextAlignJustifyIcon size={17} />
                    align justify
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

interface TableToolBarProps {
    editor: Editor | null
}

const TableToolbar = ({ editor }: TableToolBarProps) => {
    const [activeHeading, setActiveHeading] = useState<THeadingLevel | null>(
        null
    )

    const [dimensions, setDimensions] = useState({
        row: 0,
        col: 0,
    })

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target
        if (parseInt(value) < 0) return
        setDimensions((prev) => ({
            ...prev,
            [name]: parseInt(value),
        }))
    }

    const toggleHeading = (level: THeadingLevel) => {
        editor?.chain().focus().toggleHeading({ level }).run()
        setActiveHeading(level)
    }

    if (!editor) {
        return null
    }

    return (
        <div className="flex">
            {editor && (
                <>
                    <TableTools
                        editor={editor}
                        dimensions={dimensions}
                        handleInputChange={handleInputChange}
                    />
                    <TextAlignTool editor={editor} />
                    <Toolbar
                        isHeadingDisabled={false}
                        activeHeading={activeHeading}
                        toggleHeading={toggleHeading}
                        editor={editor}
                        className="w-fit"
                    />
                </>
            )}
        </div>
    )
}

const CustomTable = Table.extend({
    addAttributes() {
        return {
            // ...this.parent?.(), // Keep existing attributes
            class: {
                default: 'table',
                parseHTML: (element) =>
                    element.getAttribute('class') || 'table',
                renderHTML: (attributes) => {
                    return {
                        class: attributes.class,
                    }
                },
            },
        }
    },
})

const CustomTableCell = TableCell.extend({
    addAttributes() {
        return {
            ...this.parent?.(),
            backgroundColor: {
                default: null,
                parseHTML: (element) => element.style.backgroundColor || null,
                renderHTML: (attributes) => {
                    console.log('attributes', attributes)
                    if (!attributes.backgroundColor) {
                        return {}
                    }
                    return {
                        style: `background-color: ${attributes.backgroundColor};`,
                    }
                },
            },
            borderLeft: {
                default: null,
                parseHTML: (element) => element.style.borderLeft || null,
                renderHTML: (attributes) => {
                    if (!attributes.borderLeft) {
                        return {}
                    }
                    return {
                        style: `border-left: ${attributes.borderLeft};`,
                    }
                },
            },
            borderRight: {
                default: null,
                parseHTML: (element) => element.style.borderRight || null,
                renderHTML: (attributes) => {
                    if (!attributes.borderRight) {
                        return {}
                    }
                    return {
                        style: `border-right: ${attributes.borderRight};`,
                    }
                },
            },
            borderTop: {
                default: null,
                parseHTML: (element) => element.style.borderTop || null,
                renderHTML: (attributes) => {
                    if (!attributes.borderTop) {
                        return {}
                    }
                    return {
                        style: `border-top: ${attributes.borderTop};`,
                    }
                },
            },
            borderBottom: {
                default: null,
                parseHTML: (element) => element.style.borderBottom || null,
                renderHTML: (attributes) => {
                    if (!attributes.borderBottom) {
                        return {}
                    }
                    return {
                        style: `border-bottom: ${attributes.borderBottom};`,
                    }
                },
            },
        }
    },
})

type PageProps = {
    htmlTemplate: string
    style: string
}

export default () => {
    const [pages, setPages] = useState<PageProps[]>([
        { htmlTemplate: '<div></div>', style: '' },
    ])
    const pageHeight = 1056
    const editorRef = useRef<HTMLDivElement | null>(null)
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
            // const getClass = editor.options.editorProps.attributes

            setPages((prevPages) => {
                const updatedPages = [...prevPages]
                updatedPages[currentPage].htmlTemplate = updatedContent
                return updatedPages
            })
        }, // Load the content of the active page
        editorProps: {
            attributes: {
                class: `table-toolbar-custom !w-full dark:text-white `,
            },
        },
        autofocus: true,
        parseOptions: {
            preserveWhitespace: 'full',
        },
    })

    const handlePageUpdate = () => {
        const container = editorRef.current

        if (!container || !editor) return

        const children = Array.from(container.children) // All child elements of the editor container
        let currentHeight = 0
        let currentPageContent: HTMLElement[] = []
        const newPages: PageProps[] = [] // Updated to use PageProps[]

        console.log(children, 'children')

        children.forEach((child) => {
            const element = child as HTMLElement // Treat each child as an HTMLElement
            const elementHeight = element.offsetHeight

            if (currentHeight + elementHeight > pageHeight) {
                // Save current page content and start a new page
                newPages.push({
                    htmlTemplate: currentPageContent
                        .map((e) => e.outerHTML)
                        .join(''),
                    style: '', // Add custom styles if needed, or leave empty
                })
                currentPageContent = []
                currentHeight = 0
            }

            currentPageContent.push(element)
            currentHeight += elementHeight
        })

        // Add the last page's content
        if (currentPageContent.length) {
            newPages.push({
                htmlTemplate: currentPageContent
                    .map((e) => e.outerHTML)
                    .join(''),
                style: '', // Add custom styles if needed
            })
        }

        // Only update the state if the new pages are different from the current pages
        if (newPages.length !== pages.length) {
            setPages(newPages)
        }
    }

    useEffect(() => {
        console.log(editor)

        if (editor && pages) {
            editor.on('update', handlePageUpdate)
        }
        return () => {
            editor?.off('update', handlePageUpdate)
        }
    }, [editor])

    if (!editor) {
        return null
    }
    // const isObject = (value: unknown): value is { [name: string]: string } =>
    //     value !== null && typeof value === 'object' && !('call' in value)

    const addPage = () => {
        // const newPageStyle = editor?.options?.editorProps?.attributes
        // const getText = editor.getAttributes('text')

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
    // console.log(editor)

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

    // focus on pages when page is added
    const handleScrollFocusView = (pageIndex: number) => {
        if (editorRefFocus.current[pageIndex]) {
            editorRefFocus.current[pageIndex]?.scrollIntoView({
                behavior: 'smooth',
                block: 'start',
            })
        }
    }

    const insertTable = () => {
        console.log(
            editor?.commands.insertContent(generateTableHTML(headers, Person))
        )
    }
    // console.log(pages)
    if (!editor) {
        return null
    }

    return (
        <div className="editor-wrapper relative">
            <div className="fixed z-50 w-full border bg-white">
                <TableToolbar editor={editor} />
                <Button
                    className="rounded-md px-4 py-2 text-white shadow-md"
                    onClick={addPage}
                >
                    Add Page
                </Button>
                <Button
                    className="rounded-md px-4 py-2 text-white shadow-md"
                    onClick={insertTable}
                >
                    insertTable
                </Button>
            </div>
            <div className="absolute mt-16 flex w-full flex-col gap-10 overflow-auto bg-gray-100 py-10">
                {pages.map((page, index) => {
                    console.log(page.htmlTemplate)
                    const isCurrentPage = currentPage === index
                    // console.log('content', content)
                    return (
                        <div
                            key={index}
                            ref={(el) => (editorRefFocus.current[index] = el)}
                            // Switch page on click
                            onClick={() =>
                                !isCurrentPage ? switchToPage(index) : ''
                            }
                            className={`page relative mx-auto h-[11in] w-[8.5in] cursor-pointer overflow-hidden rounded-lg bg-white p-[1in] shadow-md ${
                                isCurrentPage
                                    ? 'border-4 border-blue-500 shadow-md'
                                    : ''
                            }`}
                        >
                            {/* Display non-editable content if not the active page */}
                            {isCurrentPage ? (
                                <EditorContent
                                    className="h-full w-full border-0 bg-white"
                                    editor={editor}
                                />
                            ) : (
                                <div
                                    className="tiptap ProseMirror table-toolbar-custom !w-full dark:text-white"
                                    dangerouslySetInnerHTML={{
                                        __html: page.htmlTemplate,
                                    }}
                                />
                            )}
                            <div className="absolute bottom-5 right-5 flex justify-center">
                                <span className="text-xs text-muted-foreground">
                                    {index}
                                </span>
                            </div>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}
