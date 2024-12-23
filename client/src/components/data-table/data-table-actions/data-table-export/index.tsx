import { toast } from 'sonner'
import { PaginationState, Table } from '@tanstack/react-table'
import { useQuery } from '@tanstack/react-query'

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

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import { TFilterPayload } from '../../data-table-filters/data-table-filter-context'

export interface IDataTableExportProps<TData> {
    table: Table<TData>
    disabled?: boolean
    isLoading?: boolean

    filters: TFilterPayload
    pagination: PaginationState

    exportAll?: () => Promise<void>
    exportCurrentPage?: (pageIndex: number) => Promise<void>
    exportSelected?: (rowsSelected: TData[]) => Promise<void>
    exportAllFiltered?: (base64Filter: string) => Promise<void>
}

const DataTableExport = <TData,>({
    table,
    filters,
    pagination,
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

    const onSuccessToast = (message?: string) =>
        toast.success(message ?? 'Data has been exported successfully.')

    const onFailedToast = (message?: string) =>
        toast.error(message ?? 'Export Failed!')

    const { refetch: exportSelectedData, isFetching: isExportingSelected } =
        useQuery<void, string>({
            queryKey: ['export', selectedData],
            queryFn: async () => {
                if (!exportSelected) return

                const [error] = await withCatchAsync(
                    exportSelected(selectedData)
                )

                if (error) {
                    const errorMessage = serverRequestErrExtractor({ error })
                    onFailedToast(errorMessage)
                    throw errorMessage
                }

                onSuccessToast()
            },
            enabled: false,
        })

    const {
        refetch: exportAllFilteredData,
        isFetching: isExportingAllFiltered,
    } = useQuery<void, string, TData[]>({
        queryKey: ['export', 'all-filtered', filters],
        queryFn: async () => {
            if (!exportAllFiltered) return

            const [error] = await withCatchAsync(
                exportAllFiltered?.(toBase64({ ...filters }))
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                onFailedToast(errorMessage)
                throw errorMessage
            }

            onSuccessToast()
        },
        enabled: false,
    })

    const { refetch: exportAllData, isFetching: isExportingAll } = useQuery<
        void,
        string,
        TData[]
    >({
        queryKey: ['export', 'all'],
        queryFn: async () => {
            if (!exportAll) return

            const [error] = await withCatchAsync(exportAll?.())

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                onFailedToast(errorMessage)
                throw errorMessage
            }

            onSuccessToast()
        },
        enabled: false,
    })

    const {
        refetch: exportCurrentPageData,
        isFetching: isExportingCurrentPage,
    } = useQuery<void, string, TData[]>({
        queryKey: ['export', /* filter ,*/ pagination],

        queryFn: async () => {
            if (!exportCurrentPage) return

            const [error] = await withCatchAsync(
                /* not sure I think this should have a filter since the page result will also affected by filters 
                and also pageSize is important so the server knows how many row in current page desyo*/
                exportCurrentPage?.(pagination.pageIndex)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                onFailedToast(errorMessage)
                throw errorMessage
            }

            onSuccessToast()
        },
        enabled: false,
    })

    const forceDisabled = !(
        exportAll ||
        exportSelected ||
        exportCurrentPage ||
        exportAllFiltered
    )

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button
                    size="sm"
                    variant={'secondary'}
                    disabled={disabled || isLoading || forceDisabled}
                    className="gap-x-1 rounded-md"
                >
                    <ExportIcon className="size-4 mr-2" />
                    {isLoading ? <LoadingSpinner /> : 'Export'}
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="min-w-44">
                <DropdownMenuGroup>
                    {selectedData.length > 0 && exportSelected && (
                        <DropdownMenuItem
                            disabled={
                                isExportingSelected || selectedData.length === 0
                            }
                            onClick={() => exportSelectedData()}
                        >
                            <CsvIcon className="mr-2 size-4 text-emerald-600" />
                            <span>Export Selected</span>
                        </DropdownMenuItem>
                    )}
                    {exportAllFiltered && (
                        <DropdownMenuItem
                            disabled={isExportingAllFiltered}
                            onClick={() => exportAllFilteredData()}
                        >
                            <CsvIcon className="mr-2 size-4 text-emerald-600" />
                            <span>Export All Filtered</span>
                            {/* No ID */}
                        </DropdownMenuItem>
                    )}
                    {exportAll && (
                        <DropdownMenuItem
                            disabled={isExportingAll}
                            onClick={() => exportAllData()}
                        >
                            <CsvIcon className="mr-2 size-4 text-emerald-600" />
                            <span>Export All</span>
                            {/* No ID, no filters */}
                        </DropdownMenuItem>
                    )}
                    {exportCurrentPage && (
                        <DropdownMenuItem
                            disabled={isExportingCurrentPage}
                            onClick={() => exportCurrentPageData()}
                        >
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

export default DataTableExport
