import { cn } from '@/lib'
import { Editor } from '@tiptap/react'
import { Table2, SplitIcon } from 'lucide-react'
import {
    CgBorderAll,
    CgBorderLeft,
    CgBorderRight,
    CgBorderTop,
    CgBorderBottom,
} from 'react-icons/cg'
import { FaTable } from 'react-icons/fa'
import { MdOutlineBorderAll, MdOutlineBorderClear } from 'react-icons/md'
import {
    TrashIcon,
} from '../icons'
import { Input } from '../ui/input'
import { Button } from '../ui/button'
import insertColumnRight from '@/assets/svg/insertColumnRight.svg'
import insertColumnLeft from '@/assets/svg/insertColumnLeft.svg'
import insertRowBefore from '@/assets/svg/insertRowUp.svg'
import insertRowAfter from '@/assets/svg/insertRowDown.svg'
import deleteSelectedRows from '@/assets/svg/deleteSelectedRows.svg'
import deleteSelectedColumns from '@/assets/svg/deleteSelectedColumns.svg'
import mergeCells from '@/assets/svg/mergeCells.svg'
import { useState } from 'react'
import { THeadingLevel } from '../text-editor'
import Toolbar from '../text-editor/toolbar'
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
} from '../ui/dropdown-menu'
import { SidebarTrigger } from '../ui/sidebar'
//import { useDocumentBuilderStore } from '@/store/document-builder-store'

interface TableToolBarProps extends React.ComponentProps<'div'> {
    editor: Editor | null
}

export const ReportBuilderToolbar = ({
    editor,
    className,
}: TableToolBarProps) => {

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
        <div
            className={cn(
                'flex h-full items-center justify-center',
                className ?? className
            )}
        >
            {editor && (
                <>
                    <TableToolbar
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
                    <SidebarTrigger />
                </>
            )}
        </div>
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

export const TableToolbar = ({
    editor,
    dimensions,
    handleInputChange,
}: TableToolsProps) => {
    if (!editor) {
        return null
    }

    const { height, setHeight } = useDocumentBuilderStore()

    const handleHeightOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { value } = e.target
        if (parseInt(value) < 0) return
        setHeight(parseInt(value))
    }

    const insertTable = () => {
        console.log(editor.commands)
        editor.commands.insertTable({
            rows: dimensions.row,
            cols: dimensions.col,
            withHeaderRow: true,
        })
    }

    const handleTableHeight = (height: number) => {
        editor?.chain().focus().updateAttributes('tableRow', {
            rowHeight: height, 
          }).run()
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
                <DropdownMenuTrigger >
                    <Button variant={'ghost'} className={cn('px-2.5')}>
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
                            <MdOutlineBorderAll size={19} className="mr-2" />{' '}
                            Border
                        </DropdownMenuSubTrigger>
                        <DropdownMenuPortal>
                            <DropdownMenuSubContent className="grid grid-cols-2">
                                <Button
                                    className={cn(
                                        'justify-start dark:bg-black'
                                    )}
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
                    <DropdownMenuSub>
                        <DropdownMenuSubTrigger>
                            <FaTable size={19} className="mr-2" /> Table Height
                        </DropdownMenuSubTrigger>
                        <DropdownMenuPortal>
                            <DropdownMenuSubContent className="flex flex-col gap-2 p-2 [&>div>span]:text-sm [&>div>span]:text-muted-foreground [&>div]:flex [&>div]:items-center [&>div]:space-x-2">
                                <div>
                                    <span>Height</span>
                                    <Input
                                        name="height"
                                        value={height}
                                        type="number"
                                        onChange={handleHeightOnChange}
                                        onClick={(e) => e.preventDefault()}
                                        onPointerDown={(e) =>
                                            e.stopPropagation()
                                        }
                                        className={cn(tableInputClass)}
                                    ></Input>
                                </div>
                                <Button onClick={()=> handleTableHeight(height)}>set Height</Button>
                            </DropdownMenuSubContent>
                        </DropdownMenuPortal>
                    </DropdownMenuSub>
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
                        <TrashIcon size={20} className="mr-2" />
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

export const TextAlignTool = ({ editor }: TextAlignToolProps) => {
    if (!editor) return null
    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button variant="ghost" className={cn('px-2.5')}>
                    <TextAlignLeftIcon size={17} />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="[&>div>svg]:mr-2 [&>div]:text-xs">
                <DropdownMenuLabel>Text Align</DropdownMenuLabel>
                <DropdownMenuSeparator />
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
