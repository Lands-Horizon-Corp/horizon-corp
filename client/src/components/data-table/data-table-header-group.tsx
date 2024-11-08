import { flexRender, HeaderGroup } from '@tanstack/react-table'
import { TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { cn } from '@/lib'

const DataTableHeaderGroup = <TData,>({
    headerGroups,
    headerClassName,
    isStickyHeader = true,
}: {
    headerGroups: HeaderGroup<TData>[]
    headerClassName?: string
    isStickyHeader?: boolean
}) => (
    <TableHeader className={isStickyHeader ? 'sticky top-0 z-50' : ''}>
        {headerGroups.map((headerGroup) => (
            <TableRow
                key={headerGroup.id}
                className="text-nowrap bg-popover hover:bg-popover"
            >
                {headerGroup.headers.map((header) => (
                    <TableHead
                        key={header.id}
                        colSpan={header.colSpan}
                        className={cn('relative bg-popover', headerClassName)}
                    >
                        {header.isPlaceholder
                            ? null
                            : flexRender(
                                  header.column.columnDef.header,
                                  header.getContext()
                              )}
                    </TableHead>
                ))}
            </TableRow>
        ))}
    </TableHeader>
)
export default DataTableHeaderGroup
