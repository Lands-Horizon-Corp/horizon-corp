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
import LoadingSpinner from '@/components/spinners/loading-spinner'

export interface IDataTableExportProps<TData> {
    table: Table<TData>
    disabled?: boolean
    isLoading?: boolean
    exportAll?: () => void
    exportSelected?: (rowsSelected: TData[]) => void
    exportCurrentPage?: () => void
    exportAllFiltered?: () => void
}

const DataTableExportButton = <TData,>({
    table,
    disabled = false,
    isLoading = false,
    exportAll,
    exportSelected,
    exportCurrentPage,
    exportAllFiltered,
}: IDataTableExportProps<TData>) => {
    const selectedData = table
        .getSelectedRowModel()
        .flatRows.map(({ original }) => original)

    const forceDisabled =
        selectedData.length === 0 ||
        !(exportAll || exportSelected || exportCurrentPage || exportAllFiltered)

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button
                    size="sm"
                    variant={'secondary'}
                    disabled={disabled || isLoading || forceDisabled}
                    className="gap-x-1 rounded-md"
                >
                    <ExportIcon className="size-4" />
                    {isLoading ? <LoadingSpinner /> : 'Export'}
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="min-w-44">
                <DropdownMenuGroup>
                    {selectedData.length > 0 && exportSelected && (
                        <DropdownMenuItem
                            onClick={() => exportSelected(selectedData)}
                        >
                            <CsvIcon className="mr-2 size-4 text-emerald-600" />
                            <span>Export Selected</span>
                        </DropdownMenuItem>
                    )}
                    {exportAllFiltered && (
                        <DropdownMenuItem onClick={() => exportAllFiltered()}>
                            <CsvIcon className="mr-2 size-4 text-emerald-600" />
                            <span>Export All Filtered</span>
                            {/* No ID */}
                        </DropdownMenuItem>
                    )}
                    {exportAll && (
                        <DropdownMenuItem onClick={() => exportAll}>
                            <CsvIcon className="mr-2 size-4 text-emerald-600" />
                            <span>Export All</span>
                            {/* No ID, no filters */}
                        </DropdownMenuItem>
                    )}
                    {exportCurrentPage && (
                        <DropdownMenuItem onClick={() => exportCurrentPage()}>
                            <CsvIcon className="mr-2 size-4 text-emerald-600" />
                            <span>Export Current Page</span>
                            {/* With ID's */}
                        </DropdownMenuItem>
                    )}
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default DataTableExportButton
