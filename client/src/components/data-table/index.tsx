import { Table, flexRender } from '@tanstack/react-table'

import {
    Table as UITable,
    TableHeader,
    TableRow,
    TableHead,
    TableBody,
    TableCell,
    ITableExtraProps,
} from '@/components/ui/table'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'

interface Props<TData> extends IBaseCompNoChild, ITableExtraProps {
    table: Table<TData>
    rowClassName?: string
    headerClassName?: string
    isStickyHeader?: boolean
}

const DataTable = <TData = unknown,>({
    table,
    className,
    rowClassName,
    isStickyHeader,
    headerClassName,
    ...other
}: Props<TData>) => {
    if (!table || !table.getRowModel) return null

    return (
        <UITable {...other} className={className}>
            <TableHeader>
                {table.getHeaderGroups().map((headerGroup) => (
                    <TableRow
                        key={headerGroup.id}
                        className={cn(
                            'bg-popover hover:bg-popover/95',
                            isStickyHeader && 'sticky top-0'
                        )}
                    >
                        {headerGroup.headers.map((header) => {
                            return (
                                <TableHead
                                    key={header.id}
                                    className={headerClassName}
                                >
                                    {header.isPlaceholder
                                        ? null
                                        : flexRender(
                                              header.column.columnDef.header,
                                              header.getContext()
                                          )}
                                </TableHead>
                            )
                        })}
                    </TableRow>
                ))}
            </TableHeader>
            <TableBody>
                {table.getRowModel().rows.length ? ( // THIS NEEDS INVESTIGATION LOL, PRODUCING WARNING ON MOUNT WHEN IN DEV MODE ONLY
                    table.getRowModel().rows.map((row) => {
                        return (
                            <TableRow
                                key={row.id}
                                className={rowClassName}
                                data-state={row.getIsSelected() && 'selected'}
                            >
                                {row.getVisibleCells().map((cell) => {
                                    return (
                                        <TableCell key={cell.id}>
                                            {flexRender(
                                                cell.column.columnDef.cell,
                                                cell.getContext()
                                            )}
                                        </TableCell>
                                    )
                                })}
                            </TableRow>
                        )
                    })
                ) : (
                    <TableRow>
                        <TableCell
                            colSpan={table.getAllColumns().length}
                            className="h-24 text-center"
                        >
                            <span className="w-full text-center text-xs text-foreground/60">
                                no data
                            </span>
                        </TableCell>
                    </TableRow>
                )}
            </TableBody>
        </UITable>
    )
}

export default DataTable
