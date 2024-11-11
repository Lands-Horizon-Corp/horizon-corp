import { CSSProperties, useState } from 'react'
import { CSS } from '@dnd-kit/utilities'
import { useSortable } from '@dnd-kit/sortable'
import { Column, Header, Table } from '@tanstack/react-table'

import {
    ResetIcon,
    ArrowUpIcon,
    EyeNoneIcon,
    PushPinIcon,
    ArrowDownIcon,
    ArrowUpDownIcon,
    PushPinSlashIcon,
} from '@/components/icons'
import {
    DropdownMenu,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuGroup,
    DropdownMenuTrigger,
    DropdownMenuContent,
    DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'

import { cn } from '@/lib/utils'
import DataTableColumnDragResize from './data-table-column-drag-resize'

interface DataTableColumnHeaderProps<TData, TValue>
    extends React.HTMLAttributes<HTMLDivElement> {
    title: string
    table: Table<TData>
    isResizable?: boolean
    column: Column<TData, TValue>
    header: Header<TData, TValue>
}

export function DataTableColumnHeader<TData, TValue>({
    title,
    table,
    column,
    header,
    className,
    isResizable,
}: DataTableColumnHeaderProps<TData, TValue>) {
    const [disableDrag, setDisableDrag] = useState(false)

    const { attributes, isDragging, listeners, setNodeRef, transform } =
        useSortable({
            id: header.column.id,
        })

    const style: CSSProperties = {
        opacity: isDragging ? 0.8 : 1,
        position: 'relative',
        transform: CSS.Translate.toString(transform),
        transition: 'width transform 0.2s ease-in-out',
        whiteSpace: 'nowrap',
        width: header.column.getSize(),
        zIndex: isDragging ? 1 : 0,
    }

    return (
        <div ref={setNodeRef} style={style}>
            <div
                className={cn(
                    'flex w-fit items-center gap-x-2',
                    className,
                    isResizable && 'pr-1'
                )}
                {...attributes}
                {...(disableDrag ? {} : listeners)}
            >
                <span
                    className="relative font-bold"
                    onClick={() => alert('ha')}
                >
                    {title}
                </span>
                <DropdownMenu onOpenChange={(state) => setDisableDrag(state)}>
                    <DropdownMenuTrigger asChild>
                        <Button
                            variant="ghost"
                            size="sm"
                            className="group-hover size-fit gap-x-2 p-1 data-[state=open]:bg-accent"
                        >
                            {column.getCanSort() && (
                                <>
                                    {column.getIsSorted() === 'desc' ? (
                                        <ArrowUpIcon className="size-3 opacity-55 group-hover:opacity-100" />
                                    ) : column.getIsSorted() === 'asc' ? (
                                        <ArrowDownIcon className="size-3 opacity-55 group-hover:opacity-100" />
                                    ) : (
                                        <ArrowUpDownIcon className="size-3 opacity-55 group-hover:opacity-100" />
                                    )}
                                </>
                            )}
                        </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent
                        className="ecoop-scroll max-h-[60vh] min-w-40 overflow-y-scroll border-none shadow-xl"
                        align="start"
                        onClick={(e) => e.stopPropagation()}
                    >
                        {column.getCanPin() && (
                            <>
                                <DropdownMenuSeparator />
                                <DropdownMenuGroup>
                                    <DropdownMenuLabel>
                                        Pin Column
                                    </DropdownMenuLabel>
                                    {column.getIsPinned() !== 'left' && (
                                        <DropdownMenuItem
                                            className="gap-x-2 px-2"
                                            onClick={() => column.pin('left')}
                                        >
                                            <PushPinIcon className="mr-2 size-3.5" />
                                            Pin Left
                                        </DropdownMenuItem>
                                    )}

                                    {column.getIsPinned() !== 'right' && (
                                        <DropdownMenuItem
                                            className="gap-x-2 px-2"
                                            onClick={() => column.pin('right')}
                                        >
                                            <PushPinIcon className="mr-2 size-3.5" />
                                            Pin Right
                                        </DropdownMenuItem>
                                    )}
                                    {column.getIsPinned() && (
                                        <DropdownMenuItem
                                            className="gap-x-2 px-2"
                                            onClick={() => column.pin(false)}
                                        >
                                            <PushPinSlashIcon className="mr-2 size-3.5" />
                                            Unpin
                                        </DropdownMenuItem>
                                    )}
                                    {table.getIsSomeColumnsPinned() && (
                                        <DropdownMenuItem
                                            className="gap-x-2 px-2"
                                            onClick={() =>
                                                table.resetColumnPinning()
                                            }
                                        >
                                            <PushPinSlashIcon className="mr-2 size-3.5" />
                                            Unpin All
                                        </DropdownMenuItem>
                                    )}
                                </DropdownMenuGroup>
                            </>
                        )}
                        {column.getCanSort() && (
                            <>
                                <DropdownMenuSeparator />
                                <DropdownMenuGroup>
                                    <DropdownMenuLabel>
                                        Sorting
                                    </DropdownMenuLabel>
                                    <DropdownMenuItem
                                        className="gap-x-2 px-2"
                                        onClick={() =>
                                            column.toggleSorting(false)
                                        }
                                    >
                                        <ArrowUpIcon className="mr-2 size-3.5" />
                                        Ascending
                                    </DropdownMenuItem>
                                    <DropdownMenuItem
                                        className="gap-x-2 px-2"
                                        onClick={() =>
                                            column.toggleSorting(true)
                                        }
                                    >
                                        <ArrowDownIcon className="mr-2 size-3.5" />
                                        Descending
                                    </DropdownMenuItem>
                                    <DropdownMenuItem
                                        className="gap-x-2 px-2"
                                        onClick={() => column.clearSorting()}
                                    >
                                        <ResetIcon className="mr-2 size-3.5" />
                                        Reset
                                    </DropdownMenuItem>
                                </DropdownMenuGroup>
                            </>
                        )}
                        {column.getCanHide() && (
                            <>
                                <DropdownMenuSeparator />
                                <DropdownMenuItem
                                    className="gap-x-2 px-2"
                                    onClick={() =>
                                        column.toggleVisibility(false)
                                    }
                                >
                                    <EyeNoneIcon className="mr-2 size-3.5" />
                                    Hide
                                </DropdownMenuItem>
                            </>
                        )}
                    </DropdownMenuContent>
                </DropdownMenu>

                {isResizable && (
                    <DataTableColumnDragResize table={table} header={header} />
                )}
            </div>
        </div>
    )
}
