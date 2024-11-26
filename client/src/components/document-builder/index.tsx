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

const CustomTable = Table.extend({
    addAttributes() {
        return {
            ...this.parent?.(), // Keep existing attributes
            class: {
                default: 'table',
            },
        }
    },
})

export default () => {

    const [pages, setPages] = useState<string[]>(['<p>Start typing...</p>']);
    const pageHeight = 1056;
    const editorRef = useRef<HTMLDivElement | null>(null);
    const [currentPage, setCurrentPage] = useState<number>(0);

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
        content: pages[currentPage],
        onUpdate: ({ editor }) => {
            // Update the content of the current page whenever it changes
            const updatedContent = editor.getHTML();
            setPages((prevPages) => {
              const updatedPages = [...prevPages];
              updatedPages[currentPage] = updatedContent;
              return updatedPages;
            });
        }, // Load the content of the active page
        editorProps: {
            attributes: {
                class: `table-toolbar-custom !w-full dark:text-white `,
            },
        },
        
    })


    const handlePageUpdate = () => {
        const container = editorRef.current;
    
        if (!container || !editor) return;
    
        const children = Array.from(container.children);
        let currentHeight = 0;
        let currentPageContent: HTMLElement[] = [];
        const newPages: string[] = [];
    
        children.forEach((child) => {
          const element = child as HTMLElement;
          const elementHeight = element.offsetHeight;
    
          if (currentHeight + elementHeight > pageHeight) {
            // Save current page content and start a new page
            newPages.push(currentPageContent.map((e) => e.outerHTML).join(''));
            currentPageContent = [];
            currentHeight = 0;
          }
    
          currentPageContent.push(element);
          currentHeight += elementHeight;
        });
    
        // Add the last page's content
        if (currentPageContent.length) {
          newPages.push(currentPageContent.map((e) => e.outerHTML).join(''));
        }
    
        if (newPages.length !== pages.length) {
          setPages(newPages);
        }
      };
    
      useEffect(() => {
        if (editor) {
          editor.on('update', handlePageUpdate);
        }
        return () => {
          editor?.off('update', handlePageUpdate);
        };
      }, [editor]);
    
      if (!editor) {
        return null;
      }

    const addPage = () => {
        setPages((prevPages) => [...prevPages, '<p>New page content...</p>']);
        setCurrentPage(pages.length); // Switch to the new page
      };

    const switchToPage = (index: number) => {
        // Save the current page's content before switching
        const updatedContent = editor?.getHTML();
        if (updatedContent !== undefined) {
          setPages((prevPages) => {
            const updatedPages = [...prevPages];
            updatedPages[currentPage] = updatedContent;
            return updatedPages;
          });
        }
    
        // Switch to the clicked page
        setCurrentPage(index);
        editor?.commands.setContent(pages[index]);
      };
    
    // const testEditor = () => {
    //     console.log(
    //         editor?.commands.insertContent(generateTableHTML(headers, Person))
    //     )
    // }

    if (!editor) {
        return null
    }

    console.log(pages)

    return (
        <div className="editor-wrapper">
              <TableToolbar editor={editor} />
              <Button
          className="bg-blue-500 text-white px-4 py-2 rounded-md shadow-md hover:bg-blue-600"
          onClick={addPage}
        >
          Add Page
        </Button>
        <div className='flex flex-col gap-2'>
        {pages.map((content, index) => (
          <div
            key={index}
            onClick={() => switchToPage(index)} // Switch page on click
            className={`page w-[8.5in] h-[11in] mx-auto p-[1in] bg-white shadow-md rounded-lg overflow-hidden relative cursor-pointer ${
              currentPage === index ? 'border-4 border-blue-500' : ''
            }`}
          >
            {/* Display non-editable content if not the active page */}
            {currentPage === index ? (
              <EditorContent className='bg-white w-full h-full' editor={editor} />
            ) : (
              <div dangerouslySetInnerHTML={{ __html: content }} />
            )}
            <div className='absolute flex bottom-0 w-[80%] justify-center'><span className=''>{index}</span></div>
          </div>
        ))}
        </div>
      </div>
    )

    return (
        <div className="mx-auto h-full w-[90%]">
            <div className="w-full">
                {/* <Button onClick={testEditor}>Add table</Button> */}
              
            </div>
            {/* <EditorContent
                onContextMenu={(event) => {
                    console.log('hello')
                    event.preventDefault()
                }}
                className="w-full"
                editor={editor}
            /> */}
          
        </div>
    )
}
