import { Table } from '@tanstack/react-table'

import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuGroup,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import { ExportIcon, CsvIcon } from '@/components/icons'

interface Props<TData> {
    table: Table<TData>
    columnsToExport: Array<keyof TData>
}

const DataTableExportButton = <TData,>({
    table,
    columnsToExport,
}: Props<TData>) => {
    if (columnsToExport.length === 0)
        throw new Error('columnsToExport should not be empty')

    const selectedData = table
        .getSelectedRowModel()
        .flatRows.map(({ original }) => original)

    const exportLocalSelected = () => {}

    // const exportAllFiltered = () => {}

    // const exportAll = () => {}

    const exportCurrentPage = () => {}

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button
                    size="sm"
                    variant={'secondary'}
                    className="gap-x-1 rounded-md"
                >
                    <ExportIcon className="size-4" />
                    Export
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="min-w-44">
                <DropdownMenuGroup>
                    {selectedData.length > 0 && (
                        <DropdownMenuItem onClick={() => exportLocalSelected()}>
                            <CsvIcon className="mr-2 size-4 text-emerald-600" />
                            <span>Export Selected</span>
                        </DropdownMenuItem>
                    )}
                    <DropdownMenuItem>
                        <CsvIcon className="mr-2 size-4 text-emerald-600" />
                        <span>Export All Filtered</span>
                        {/* No ID */}
                    </DropdownMenuItem>
                    <DropdownMenuItem>
                        <CsvIcon className="mr-2 size-4 text-emerald-600" />
                        <span>Export All</span>
                        {/* No ID, no filters */}
                    </DropdownMenuItem>
                    <DropdownMenuItem onClick={() => exportCurrentPage()}>
                        <CsvIcon className="mr-2 size-4 text-emerald-600" />
                        <span>Export Current Page</span>
                        {/* With ID's */}
                    </DropdownMenuItem>
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default DataTableExportButton
