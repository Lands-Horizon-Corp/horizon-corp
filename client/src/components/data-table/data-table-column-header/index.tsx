import { CSSProperties } from 'react'
import { CSS } from '@dnd-kit/utilities'
import { useSortable } from '@dnd-kit/sortable'
import { Column, Header, Table } from '@tanstack/react-table'

import { Button } from '@/components/ui/button'
import ColumnResizeHandle from './column-drag-resize'
import ActionTooltip from '@/components/action-tooltip'

import { cn } from '@/lib/utils'

interface DataTableColumnHeaderProps<TData, TValue>
    extends React.HTMLAttributes<HTMLDivElement> {
    title: string
    tooltipDescription?: string
    table: Table<TData>
    isResizable?: boolean
    column: Column<TData, TValue>
    header: Header<TData, TValue>
}

const DataTableColumnHeader = <TData, TValue>({
    title,
    table,
    header,
    children,
    className,
    isResizable,
    tooltipDescription,
}: DataTableColumnHeaderProps<TData, TValue>) => {
    const { attributes, isDragging, listeners, setNodeRef, transform } =
        useSortable({
            id: header.column.id,
        })

    const style : CSSProperties = {
        opacity: isDragging ? 0.8 : 1,
        position: 'relative',
        transform: CSS.Translate.toString(transform),
        transition: 'width transform 0.2s ease-in-out',
        whiteSpace: 'nowrap',
        width: header.column.getSize(),
        zIndex: isDragging ? 1 : 0,
    }

    const finalTitle = tooltipDescription ? (
        <ActionTooltip tooltipContent={tooltipDescription ?? ''}>
            <span className="relative font-bold">{title}</span>
        </ActionTooltip>
    ) : (
        <span className="relative font-bold">{title}</span>
    )

    return (
        <div className="inline-block flex-1 space-y-1 py-2">
            <div
                className={cn(
                    'flex w-fit items-center gap-x-2',
                    className,
                    isResizable && 'pr-1'
                )}
            >
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
                        {finalTitle}
                    </Button>
                    {children}
                    {isResizable && (
                        <ColumnResizeHandle table={table} header={header} />
                    )}
                </div>
            </div>
        </div>
    )
}

export default DataTableColumnHeader
