import { toast } from 'sonner'
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
import { exportTableData, FileNameWithoutExtension } from '@/lib/xlsx'
import { format } from 'date-fns'

interface Props<TData> {
    table: Table<TData>
    columnsToExport: Array<keyof TData>
    fileName: FileNameWithoutExtension<string>
}

const DataTableExportButton = <TData,>({
    table,
    fileName,
    columnsToExport,
}: Props<TData>) => {
    if (columnsToExport.length === 0)
        throw new Error('columnsToExport should not be empty')

    const selectedData = table
        .getSelectedRowModel()
        .flatRows.map(({ original }) => original)

    const exportLocalSelected = () => {
        const dataToExport = selectedData.map((row) => {
            if (columnsToExport?.length) {
                const filteredRow: Partial<TData> = {}
                columnsToExport.forEach((key) => {
                    filteredRow[key] = row[key]
                })
                return filteredRow
            }
            return row
        })

        try {
            if (dataToExport.length === 0)
                return toast.warning('Please select data to export.')
            const finalFileName = `${fileName}-${format(new Date(), "yyyy-MM-dd'T'HH:mm:ss'Z'")}`
            exportTableData(dataToExport, finalFileName)
            toast.success(`Selected Data exported to ${finalFileName}`)
        } catch (_err) {
            toast.error('Failed to export')
        }
    }

    // const exportAllFiltered = () => {}

    // const exportAll = () => {}

    const exportCurrentPage = () => {
        try {
            const { pageIndex, pageSize } = table.getState().pagination

            const startRow = pageIndex * pageSize
            const endRow = startRow + pageSize

            const currentPageRows = table
                .getCoreRowModel()
                .rows.slice(startRow, endRow)
                .map((row) => row.original)

            const dataToExport = currentPageRows.map((row) => {
                if (columnsToExport?.length) {
                    const filteredRow: Partial<TData> = {}
                    columnsToExport.forEach((key) => {
                        filteredRow[key] = row[key]
                    })
                    return filteredRow
                }
                return row
            })

            if (dataToExport.length === 0)
                return toast.warning('No data available to export.')

            const finalFileName = `${fileName}-${format(new Date(), "yyyy-MM-dd'T'HH:mm:ss'Z'")}`

            exportTableData(dataToExport, finalFileName)
            toast.success(`Current page data exported to ${finalFileName}`)
        } catch (_err) {
            toast.error('Failed to export')
        }
    }

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
