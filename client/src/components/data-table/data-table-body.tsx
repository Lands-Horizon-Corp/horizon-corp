import { flexRender, Row } from "@tanstack/react-table"
import { TableBody, TableCell, TableRow } from "@/components/ui/table"

const DataTableBody = <TData,>({
    rows,
    targetGroup,
}: {
    rows: Row<TData>[]
    targetGroup?: 'left' | 'right'
}) => (
    <TableBody>
        {rows.map((row) => (
            <TableRow
                key={row.id}
                className="h-14 w-fit bg-secondary align-middle hover:bg-popover"
                data-state={row.getIsSelected() && 'selected'}
            >
                {(targetGroup === undefined
                    ? row.getCenterVisibleCells()
                    : targetGroup === 'left'
                      ? row.getLeftVisibleCells()
                      : row.getRightVisibleCells()
                ).map((cell) => (
                    <TableCell key={cell.id} className="size-fit text-nowrap">
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

export default DataTableBody

