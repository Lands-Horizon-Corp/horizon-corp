import {
    SortableContext,
    horizontalListSortingStrategy,
} from '@dnd-kit/sortable'
import { flexRender, HeaderGroup } from '@tanstack/react-table'

import { TableHead, TableHeader, TableRow } from '@/components/ui/table'

import { cn } from '@/lib'

const DataTableHeaderGroup = <TData,>({
    columnOrder,
    headerGroups,
    headerClassName,
    isStickyHeader = true,
}: {
    columnOrder: string[]
    headerClassName?: string
    isStickyHeader?: boolean
    headerGroups: HeaderGroup<TData>[]
}) => (
    <TableHeader className={isStickyHeader ? 'sticky bottom-0 z-50' : ''}>
        {headerGroups.map((headerGroup) => (
            <TableRow
                key={headerGroup.id}
                className="text-nowrap bg-popover hover:bg-popover"
            >
                <SortableContext
                    items={columnOrder}
                    strategy={horizontalListSortingStrategy}
                >
                    {headerGroup.headers.map((header) => (
                        <TableHead
                            key={header.id}
                            colSpan={header.colSpan}
                            className={cn(
                                'relative bg-popover',
                                headerClassName
                            )}
                        >
                            {header.isPlaceholder
                                ? null
                                : flexRender(
                                      header.column.columnDef.header,
                                      header.getContext()
                                  )}
                        </TableHead>
                    ))}
                </SortableContext>
            </TableRow>
        ))}
    </TableHeader>
)
export default DataTableHeaderGroup
