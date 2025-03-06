import { flexRender, Row } from '@tanstack/react-table'

import { MagnifyingGlassIcon } from '../icons'
import { TableBody, TableCell, TableRow } from '@/components/ui/table'

import { getPinningStyles } from './data-table-utils'

const DataTableBody = <TData,>({
    rows,
    colCount,
}: {
    rows: Row<TData>[]
    rowClassName?: string
    colCount?: number
}) => {
    return (
        <TableBody>
            {rows.map((row) => (
                <TableRow
                    key={row.id}
                    data-row-id={row.id}
                    onClick={() => {
                        row.toggleSelected()
                    }}
                    className="h-14"
                    data-state={row.getIsSelected() && 'selected'}
                >
                    {row.getVisibleCells().map((cell) => {
                        const { column } = cell
                        const isPinned = column.getIsPinned()
                        const isLastLeftPinned =
                            isPinned === 'left' &&
                            column.getIsLastColumn('left')
                        const isFirstRightPinned =
                            isPinned === 'right' &&
                            column.getIsFirstColumn('right')

                        return (
                            <TableCell
                                key={cell.id}
                                className="truncate data-[pinned]:bg-muted/60 data-[pinned]:backdrop-blur-sm [&[data-pinned=left][data-last-col=left]]:border-r [&[data-pinned=right][data-last-col=right]]:border-l [&[data-pinned][data-last-col]]:border-border"
                                style={{
                                    ...getPinningStyles(column),
                                }}
                                data-pinned={isPinned || undefined}
                                data-last-col={
                                    isLastLeftPinned
                                        ? 'left'
                                        : isFirstRightPinned
                                          ? 'right'
                                          : undefined
                                }
                            >
                                {flexRender(
                                    cell.column.columnDef.cell,
                                    cell.getContext()
                                )}
                            </TableCell>
                        )
                    })}
                </TableRow>
            ))}
            {rows.length === 0 && (
                <TableRow>
                    <TableCell colSpan={colCount} className="h-24 text-center">
                        <span className="w-full text-center text-xs text-foreground/60">
                            <MagnifyingGlassIcon />
                        </span>
                    </TableCell>
                </TableRow>
            )}
        </TableBody>
    )
}

export default DataTableBody
