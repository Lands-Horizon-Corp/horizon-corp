import { Column, Table, flexRender } from '@tanstack/react-table'

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
import { CSSProperties } from 'react'

interface Props<TData> extends IBaseCompNoChild, ITableExtraProps {
    table: Table<TData>
    rowClassName?: string
    headerClassName?: string
    isStickyHeader?: boolean
}

const getCommonPinningStyles = <TData,>(
    column: Column<TData>
): CSSProperties => {
    const isPinned = column.getIsPinned()
    const isLastLeftPinnedColumn =
        isPinned === 'left' && column.getIsLastColumn('left')
    const isFirstRightPinnedColumn =
        isPinned === 'right' && column.getIsFirstColumn('right')

    return {
        boxShadow: isLastLeftPinnedColumn
            ? '-4px 0 4px -4px gray inset'
            : isFirstRightPinnedColumn
              ? '4px 0 4px -4px gray inset'
              : undefined,
        left: isPinned === 'left' ? `${column.getStart('left')}px` : undefined,
        right:
            isPinned === 'right' ? `${column.getAfter('right')}px` : undefined,
        opacity: isPinned ? 0.95 : 1,
        position: isPinned ? 'sticky' : 'relative',
        width: column.getSize(),
        zIndex: isPinned ? 1 : 0,
    }
}

const getCommonPinningClasses = <TData,>(column: Column<TData>): string => {
    const isPinned = column.getIsPinned()
    const isLastLeftPinnedColumn =
        isPinned === 'left' && column.getIsLastColumn('left')
    const isFirstRightPinnedColumn =
        isPinned === 'right' && column.getIsFirstColumn('right')

    return cn(
        isPinned ? 'sticky opacity-95 z-10' : 'relative z-0',
        isPinned === 'left' && `left-[${column.getStart('left')}px]`,
        isPinned === 'right' && `right-[${column.getAfter('right')}px]`,
        isLastLeftPinnedColumn && 'shadow-lg shadow-gray-500 inset-l-4',
        isFirstRightPinnedColumn && 'shadow-lg shadow-gray-500 inset-r-4',
        `w-[${column.getSize()}px]`
    )
}

const DataTable = <TData,>({
    table,
    className,
    rowClassName,
    isStickyHeader,
    headerClassName,
    ...other
}: Props<TData>) => {
    if (!table || !table.getRowModel) return null

    return (
        <UITable
            {...other}
            className={className}
            // style={{
            // width: table.getCenterTotalSize(),
            // }}
        >
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
                                    colSpan={header.colSpan}
                                    style={{
                                        width: header.getSize(),
                                        ...getCommonPinningStyles(header.column)
                                    }}
                                    className={cn(
                                        'relative',
                                        headerClassName,
                                        // getCommonPinningClasses(header.column)
                                    )}
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
                {
                    // THIS NEEDS INVESTIGATION LOL, PRODUCING WARNING ON MOUNT WHEN IN DEV MODE ONLY
                    table.getRowModel().rows.map((row) => {
                        return (
                            <TableRow
                                key={row.id}
                                className={cn('w-fit', rowClassName)}
                                data-state={row.getIsSelected() && 'selected'}
                            >
                                {row.getVisibleCells().map((cell) => {
                                    return (
                                        <TableCell
                                            key={cell.id}
                                            style={{
                                                width: cell.column.getSize(),
                                        ...getCommonPinningStyles(cell.column)
                                            }}
                                            // className={getCommonPinningClasses(
                                            //     cell.column
                                            // )}
                                        >
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
                }

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
    )
}

export default DataTable
