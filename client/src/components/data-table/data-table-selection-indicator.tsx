import { Table } from '@tanstack/react-table'

interface Props<T> {
    table: Table<T>
    hideSelectedIndicator?: boolean
}

const DataTableSelectionIndicator = <T,>({
    table,
    hideSelectedIndicator,
}: Props<T>) => {
    return (
        <div>
            {!hideSelectedIndicator && (
                <div className="flex-1 text-sm text-muted-foreground">
                    {table.getFilteredSelectedRowModel().rows.length} of{' '}
                    {table.getFilteredRowModel().rows.length} row(s) selected.
                </div>
            )}
        </div>
    )
}

export default DataTableSelectionIndicator
