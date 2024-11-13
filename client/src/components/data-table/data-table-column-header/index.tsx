import { CSSProperties } from 'react'
import { CSS } from '@dnd-kit/utilities'
import { useSortable } from '@dnd-kit/sortable'
import { Column, Header, Table } from '@tanstack/react-table'

import ColumnActions from './column-actions'
import { Button } from '@/components/ui/button'
import ColumnResizeHandle from './column-drag-resize'
import { MoveLeftRightIcon } from '@/components/icons'

import { cn } from '@/lib/utils'
import ColumnFilter from './column-filter'
import { TColumnDataTypes } from '../data-table-filter-context'

interface DataTableColumnHeaderProps<TData, TValue>
    extends React.HTMLAttributes<HTMLDivElement> {
    title: string
    table: Table<TData>
    isResizable?: boolean
    column: Column<TData, TValue>
    header: Header<TData, TValue>
    dataType?: TColumnDataTypes
}

export function DataTableColumnHeader<TData, TValue>({
    title,
    table,
    column,
    header,
    dataType,
    className,
    isResizable,
}: DataTableColumnHeaderProps<TData, TValue>) {
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
        <div className="inline-block flex-1 space-y-1 py-2">
            <div
                className={cn(
                    'flex w-fit items-center gap-x-2',
                    className,
                    isResizable && 'pr-1'
                )}
            >
                <span className="relative font-bold">{title}</span>
                <div className="inline-flex items-center gap-x-0.5">
                    <Button
                        size="sm"
                        style={style}
                        {...listeners}
                        variant="ghost"
                        {...attributes}
                        ref={setNodeRef}
                        className="!size-fit cursor-ew-resize gap-x-2 p-1 data-[state=open]:bg-accent"
                    >
                        {isDragging && (
                            <span className="relative font-bold">{title}</span>
                        )}
                        <MoveLeftRightIcon className="size-3 opacity-55 group-hover:opacity-100" />
                    </Button>

                    <ColumnActions table={table} column={column} />
                    {dataType && (
                        <ColumnFilter dataType={dataType} column={column} />
                    )}

                    {isResizable && (
                        <ColumnResizeHandle table={table} header={header} />
                    )}
                </div>
            </div>
        </div>
    )
}
