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
import { useState } from 'react'
import { THeadingLevel } from '../text-editor'
import Toolbar from '../text-editor/toolbar'
import ReactDOMServer from 'react-dom/server'
import { Button } from '../ui/button'
import { FaTable } from 'react-icons/fa'
import { Ghost, Table2 } from 'lucide-react'
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
    [key: string]: string | number | null | undefined // Key-value pairs for each row
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

const CustomTable = Table.extend({
    addAttributes() {
      return {
        ...this.parent?.(), // Keep existing attributes
        class: {
          default: 'table',
        },
      };
    },
  });

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
            {/* <Button onClick={testEditor}>test</Button> */}

            {/*
            <Button onClick={() => editor.chain().focus().mergeCells().run()}>
                <MergeIcon className="mr-2" /> Merge Cells
            </Button>
            <Button onClick={() => editor.chain().focus().splitCell().run()}>
                <SplitIcon className="mr-2" /> Split Cells
            </Button>
            <Button
                onClick={() =>
                    editor.chain().focus().toggleHeaderColumn().run()
                }
            >
                Toggle Header Column
            </Button>
           
            <Button
                onClick={() => editor.chain().focus().toggleHeaderCell().run()}
            >
                Toggle Header Cell
            </Button>
            <Button onClick={() => editor.chain().focus().mergeOrSplit().run()}>
                Merge Or Split
            </Button>
            <Button
                onClick={() =>
                    editor.chain().focus().setCellAttribute('colspan', 2).run()
                }
            >
                Set Cell Attribute
            </Button>
            <Button onClick={() => editor.chain().focus().fixTables().run()}>
                Fix Tables
            </Button>
            <Button onClick={() => editor.chain().focus().goToNextCell().run()}>
                Go To Next Cell
            </Button>
            <Button
                onClick={() => editor.chain().focus().goToPreviousCell().run()}
            >
                Go To Previous Cell
            </Button> */}
        </div>
    )
}

export default () => {
    const [filterText, setFilterText] = useState('')

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
            CustomTable,
            TableRow,
            TableHeader,
            TableCell,
            TextAlign.configure({
                types: ['heading', 'paragraph'],
            }),
        ],
        content: '',
        editorProps: {
            attributes: {
                class: `table-toolbar-custom !w-full dark:text-white `,
            },
        },
    })

    const json = editor?.getJSON()
    // console.log(json)

    const applyFilter = () => {
        if (!editor) return

        const { doc, tr } = editor.state


        const tableNode = doc.descendants((node, pos) => {
            console.log('node', node)
            if (node.type.name === 'table') {
                // Iterate over rows and apply the filter
                node.descendants((rowNode, rowPos) => {
                    if (rowNode.type.name === 'tableRow') {
                        const rowText = rowNode.textContent.toLowerCase()
                        const shouldShow = rowText.includes(
                            filterText.toLowerCase()
                        )
                        const newAttrs = {
                            ...rowNode.attrs,
                            'data-hidden': !shouldShow ? 'true' : null,
                        }

                        tr.setNodeMarkup(pos + rowPos + 1, undefined, newAttrs)
                    }
                })
                return true
            }
            return false
        })
        console.log('tableNode', tableNode)
        editor.view.dispatch(tr)
    }
    const testEditor = () => {
        console.log(
            editor?.commands.insertContent(generateTableHTML(headers, Person))
        )
    }

    if (!editor) {
        return null
    }

    return (
        <div className="mx-auto h-full w-[90%]">
            <div className="">
                <label>
                    Filter:
                    <input
                        type="text"
                        value={filterText}
                        onChange={(e) => setFilterText(e.target.value)}
                    />
                </label>
                <Button onClick={() => applyFilter()}>Apply Filter</Button>
            </div>
            <div className="w-full">
                <Button onClick={testEditor}>Add table</Button>

                <TableToolbar editor={editor} />
            </div>
            <EditorContent
                onContextMenu={(event) => {
                    console.log('hello')
                    event.preventDefault()
                }}
                className="w-full"
                editor={editor}
            />
        </div>
    )
}
