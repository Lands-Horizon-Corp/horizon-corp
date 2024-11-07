import { Table, flexRender } from '@tanstack/react-table'

import {
    Table as UITable,
    TableHeader,
    TableRow,
    TableHead,
    TableBody,
    TableCell,
} from '@/components/ui/table'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'

interface Props<TData> extends IBaseCompNoChild {
    table: Table<TData>
    rowClassName?: string
    headerClassName?: string
    wrapperClassName?: string
    isStickyHeader?: boolean
    cellClassName?: string
}

const DataTable = <TData,>({
    table,
    className,
    isStickyHeader,
    headerClassName,
    wrapperClassName,
}: Props<TData>) => {
    return (
        <div
            className={cn(
                'ecoop-scroll relative max-h-full overflow-y-scroll bg-secondary',
                wrapperClassName
            )}
        >
            <div className="flex h-fit">
                <div className="ecoop-scroll sticky left-0 z-10 w-fit border-r border-popover">
                    <UITable
                        style={{ width: table.getLeftTotalSize() }}
                        className={cn('h-fit !w-fit', className)}
                    >
                        <TableHeader
                            className={cn(
                                '',
                                isStickyHeader && 'sticky top-0 z-50'
                            )}
                        >
                            {table.getLeftHeaderGroups().map((headerGroup) => (
                                <TableRow
                                    key={headerGroup.id}
                                    className="bg-popover hover:bg-popover"
                                >
                                    {headerGroup.headers.map((header) => {
                                        return (
                                            <TableHead
                                                key={header.id}
                                                colSpan={header.colSpan}
                                                style={{
                                                    width: header.column.getSize(),
                                                }}
                                                className={cn(
                                                    'relative bg-popover',
                                                    headerClassName
                                                )}
                                            >
                                                {header.isPlaceholder
                                                    ? null
                                                    : flexRender(
                                                          header.column
                                                              .columnDef.header,
                                                          header.getContext()
                                                      )}
                                            </TableHead>
                                        )
                                    })}
                                </TableRow>
                            ))}
                        </TableHeader>
                        <TableBody>
                            {table.getRowModel().rows.map((row) => {
                                return (
                                    <TableRow
                                        key={row.id}
                                        className={cn('h-14 w-fit')}
                                        data-state={
                                            row.getIsSelected() && 'selected'
                                        }
                                    >
                                        {row
                                            .getLeftVisibleCells()
                                            .map((cell) => {
                                                return (
                                                    <TableCell
                                                        key={cell.id}
                                                        style={{
                                                            width: cell.column.getSize(),
                                                        }}
                                                        className={cn(
                                                            'size-fit backdrop-blur'
                                                        )}
                                                    >
                                                        {flexRender(
                                                            cell.column
                                                                .columnDef.cell,
                                                            cell.getContext()
                                                        )}
                                                    </TableCell>
                                                )
                                            })}
                                    </TableRow>
                                )
                            })}

                            {table.getRowModel().rows.length === 0 && (
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
                </div>
                <div className="z-0 flex-1">
                    <UITable
                        style={{
                            width: table.getCenterTotalSize(),
                        }}
                        className={cn('h-fit', className)}
                    >
                        <TableHeader
                            className={cn(
                                '',
                                isStickyHeader && 'sticky top-0 z-50'
                            )}
                        >
                            {table
                                .getCenterHeaderGroups()
                                .map((headerGroup) => (
                                    <TableRow
                                        key={headerGroup.id}
                                        className="bg-popover hover:bg-popover"
                                    >
                                        {headerGroup.headers.map((header) => {
                                            return (
                                                <TableHead
                                                    key={header.id}
                                                    colSpan={header.colSpan}
                                                    style={{
                                                        width: `calc(var(--header-${header?.id}-size) * 1px)`,
                                                    }}
                                                    className={cn(
                                                        'relative bg-popover',
                                                        headerClassName
                                                    )}
                                                >
                                                    {header.isPlaceholder
                                                        ? null
                                                        : flexRender(
                                                              header.column
                                                                  .columnDef
                                                                  .header,
                                                              header.getContext()
                                                          )}
                                                </TableHead>
                                            )
                                        })}
                                    </TableRow>
                                ))}
                        </TableHeader>
                        <TableBody>
                            {table.getRowModel().rows.map((row) => {
                                return (
                                    <TableRow
                                        key={row.id}
                                        className={cn('h-14 w-fit')}
                                        data-state={
                                            row.getIsSelected() && 'selected'
                                        }
                                    >
                                        {row
                                            .getCenterVisibleCells()
                                            .map((cell) => {
                                                return (
                                                    <TableCell
                                                        key={cell.id}
                                                        style={{
                                                            width: cell.column.getSize(),
                                                        }}
                                                        className={cn(
                                                            'size-fit backdrop-blur'
                                                        )}
                                                    >
                                                        {flexRender(
                                                            cell.column
                                                                .columnDef.cell,
                                                            cell.getContext()
                                                        )}
                                                    </TableCell>
                                                )
                                            })}
                                    </TableRow>
                                )
                            })}

                            {table.getRowModel().rows.length === 0 && (
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
                </div>
                <div className="ecoop-scroll sticky right-0 z-10 w-fit border-l border-popover">
                    <UITable
                        style={{ width: table.getRightTotalSize() }}
                        className={cn('!w-fit', className)}
                    >
                        <TableHeader
                            className={cn(
                                '',
                                isStickyHeader && 'sticky top-0 z-50'
                            )}
                        >
                            {table.getRightHeaderGroups().map((headerGroup) => (
                                <TableRow
                                    key={headerGroup.id}
                                    className="bg-popover hover:bg-popover"
                                >
                                    {headerGroup.headers.map((header) => {
                                        return (
                                            <TableHead
                                                key={header.id}
                                                style={{
                                                    width: header.column.getSize(),
                                                }}
                                                colSpan={header.colSpan}
                                                className={cn(
                                                    'relative bg-popover',
                                                    headerClassName
                                                )}
                                            >
                                                {header.isPlaceholder
                                                    ? null
                                                    : flexRender(
                                                          header.column
                                                              .columnDef.header,
                                                          header.getContext()
                                                      )}
                                            </TableHead>
                                        )
                                    })}
                                </TableRow>
                            ))}
                        </TableHeader>
                        <TableBody>
                            {table.getRowModel().rows.map((row) => {
                                return (
                                    <TableRow
                                        key={row.id}
                                        className={cn('h-14 w-fit')}
                                        data-state={
                                            row.getIsSelected() && 'selected'
                                        }
                                    >
                                        {row
                                            .getRightVisibleCells()
                                            .map((cell) => {
                                                return (
                                                    <TableCell
                                                        key={cell.id}
                                                        style={{
                                                            width: cell.column.getSize(),
                                                        }}
                                                        className={cn(
                                                            'size-fit backdrop-blur'
                                                        )}
                                                    >
                                                        {flexRender(
                                                            cell.column
                                                                .columnDef.cell,
                                                            cell.getContext()
                                                        )}
                                                    </TableCell>
                                                )
                                            })}
                                    </TableRow>
                                )
                            })}

                            {table.getRowModel().rows.length === 0 && (
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
                </div>
            </div>
        </div>
    )
}

export default DataTable
