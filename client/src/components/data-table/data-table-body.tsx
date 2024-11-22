import { flexRender, Row } from '@tanstack/react-table'

import { TableBody, TableCell, TableRow } from '@/components/ui/table'

const DataTableBody = <TData,>({
    rows,
    targetGroup,
}: {
    rows: Row<TData>[]
    targetGroup?: 'left' | 'right'
}) => {
    return (
        <TableBody>
            {rows.map((row) => (
                <TableRow
                    key={row.id}
                    data-row-id={row.id}
                    className="h-14 w-fit dark:bg-secondary/90 align-middle hover:bg-popover/80 data-[state=selected]:bg-popover"
                    data-state={row.getIsSelected() && 'selected'}
                >
                    {(targetGroup === undefined
                        ? row.getCenterVisibleCells()
                        : targetGroup === 'left'
                          ? row.getLeftVisibleCells()
                          : row.getRightVisibleCells()
                    ).map((cell) => (
                        <TableCell
                            key={cell.id}
                            className="size-fit text-nowrap"
                        >
                            {flexRender(
                                cell.column.columnDef.cell,
                                cell.getContext()
                            )}
                        </TableCell>
                    ))}
                </TableRow>
            ))}
            {rows.length === 0 && (
                <TableRow>
                    <TableCell
                        colSpan={rows.length || 0}
                        className="h-24 text-center"
                    >
                        <span className="w-full text-center text-xs text-foreground/60">
                            no data
                        </span>
                    </TableCell>
                </TableRow>
            )}
        </TableBody>
    )
}

DataTableBody.displayName = 'DataTableBody'

export default DataTableBody
