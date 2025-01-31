import { flexRender, Row } from '@tanstack/react-table'

import { TableBody, TableCell, TableRow } from '@/components/ui/table'

import { cn } from '@/lib'

type TTargetGroup = 'left' | 'right'

const DataTableBody = <TData,>({
    rows,
    colCount,
    targetGroup,
    rowClassName,
}: {
    rows: Row<TData>[]
    targetGroup?: TTargetGroup
    rowClassName?: string
    colCount?: number
}) => {
    const getVisibleCells = (row: Row<TData>, targetGroup?: TTargetGroup) => {
        switch (targetGroup) {
            case 'left':
                return row.getLeftVisibleCells()
            case 'right':
                return row.getRightVisibleCells()
            default:
                return row.getCenterVisibleCells()
        }
    }

    return (
        <TableBody>
            {rows.map((row) => (
                <TableRow
                    key={row.id}
                    data-row-id={row.id}
                    onClick={() => {
                        row.toggleSelected()
                    }}
                    className={cn(
                        'h-14 w-fit bg-background/90 align-middle cursor-pointer hover:bg-secondary/90 data-[state=selected]:bg-secondary dark:bg-secondary/90 dark:hover:bg-popover/95 dark:data-[state=selected]:bg-popover',
                        rowClassName
                    )}
                    data-state={row.getIsSelected() && 'selected'}
                >
                    {getVisibleCells(row, targetGroup).map((cell) => (
                        <TableCell
                            key={cell.id}
                            className="size-fit px-4 py-2.5 text-nowrap"
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
                    <TableCell colSpan={colCount} className="h-24 text-center">
                        <span className="w-full text-center text-xs text-foreground/60">
                            🫥 no data
                        </span>
                    </TableCell>
                </TableRow>
            )}
        </TableBody>
    )
}

export default DataTableBody
