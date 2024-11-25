import { flexRender, Table } from '@tanstack/react-table'

import { TableFooter, TableRow, TableCell } from '@/components/ui/table'
import { cn } from '@/lib'

const DataTableFooter = <TData,>({
    table,
    targetGroup,
    footerTrClassName,
    isStickyFooter = false,
}: {
    table: Table<TData>
    footerTrClassName?: string
    isStickyFooter?: boolean
    targetGroup?: 'left' | 'right' | 'center'
}) => {
    const NFooter = table
        .getFooterGroups()
        .map((group) =>
            group.headers.map((header) => header.column.columnDef.footer)
        )
        .flat()
        .filter(Boolean).length

    if (NFooter === 0) return

    return (
        <TableFooter className={isStickyFooter ? 'sticky bottom-0 z-50' : ''}>
            {table.getFooterGroups().map((footerGroup) => (
                <TableRow
                    data-footer-row-id={footerGroup.id}
                    className={cn(
                        'text-nowrap bg-secondary hover:bg-popover dark:bg-popover',
                        footerTrClassName
                    )}
                >
                    {footerGroup.headers
                        .filter((grp) => {
                            const pinnedGroup = grp.column.getIsPinned()
                            if (!targetGroup) return pinnedGroup === false
                            return targetGroup === pinnedGroup
                        })
                        .map((tc) => (
                            <TableCell
                                key={tc.id}
                                className="size-fit text-nowrap font-medium"
                            >
                                {flexRender(
                                    tc.column.columnDef.footer,
                                    tc.getContext()
                                )}
                            </TableCell>
                        ))}
                </TableRow>
            ))}
        </TableFooter>
    )
}

DataTableFooter.displayName = 'DataTableFooter'

export default DataTableFooter
