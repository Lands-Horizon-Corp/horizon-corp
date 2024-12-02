import { toast } from 'sonner'
import { useQuery } from '@tanstack/react-query'
import {
    useReactTable,
    getCoreRowModel,
    getSortedRowModel,
} from '@tanstack/react-table'

import DataTable from '@/components/data-table'
import DataTableToolbar from '@/components/data-table/data-table-toolbar'
import DataTablePagination from '@/components/data-table/data-table-pagination'
import useDataTableState from '@/components/data-table/hooks/use-datatable-state'
import useDatableFilterState from '@/components/data-table/hooks/use-datatable-filter-state'
import DataTableFilterContext from '@/components/data-table/data-table-filters/data-table-filter-context'

import { cn, toBase64, withCatchAsync } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { serverRequestErrExtractor } from '@/helpers'
import { companiesTableColumns as columns } from './columns'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'
import logger from '@/helpers/loggers/logger'
import { CompanyPaginatedResource } from '@/horizon-corp/types'

const CompaniesTable = ({ className }: IBaseCompNoChild) => {
    const {
        sorting,
        setSorting,
        pagination,
        setPagination,
        columnOrder,
        setColumnOrder,
        isScrollable,
        setIsScrollable,
        rowSelection,
        setRowSelection,
        columnVisibility,
        setColumnVisibility,
    } = useDataTableState({
        columnOrder: columns.map((c) => c.id!),
    })

    const filterState = useDatableFilterState({
        name: 'preload.name',
    })

    const {
        data: { data, totalPage, pageSize },
        isPending,
        refetch,
    } = useQuery<CompanyPaginatedResource, string>({
        queryKey: ['company-list', filterState.filters],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                CompanyService.filter(
                    toBase64({ filters: filterState.filters, pagination })
                )
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            if (!result) throw "Something wen't wroing"

            return result
        },
        initialData: {
            data: [],
            pages: [],
            totalSize: 0,
            totalPage: 1,
            ...pagination,
        },
        retry: 1,
    })

    const table = useReactTable({
        columns,
        data: data,
        initialState: {
            columnPinning: { left: ['select'] },
        },
        state: {
            sorting,
            pagination,
            columnOrder,
            rowSelection,
            columnVisibility,
        },
        rowCount: pageSize,
        pageCount: totalPage,
        manualSorting: true,
        manualFiltering: true,
        manualPagination: true,
        onSortingChange: setSorting,
        onPaginationChange: setPagination,
        getCoreRowModel: getCoreRowModel(),
        onColumnOrderChange: setColumnOrder,
        onRowSelectionChange: setRowSelection,
        getSortedRowModel: getSortedRowModel(),
        onColumnVisibilityChange: setColumnVisibility,
    })

    logger.log(data)

    return (
        <DataTableFilterContext.Provider value={filterState}>
            <div className={cn('flex h-full flex-col gap-y-2', className)}>
                <DataTableToolbar
                    globalSearchProps={{
                        defaultMode: 'equal',
                        keysToSearch: ['name', 'contactNumber'],
                    }}
                    table={table}
                    refreshActionProps={{
                        isLoading: isPending,
                        onClick: () => refetch(),
                    }}
                    deleteActionProps={{
                        isLoading: false,
                        onClick: () => {},
                    }}
                    scrollableProps={{ isScrollable, setIsScrollable }}
                    exportActionProps={{
                        columnsToExport: [
                            'name',
                            'owner',
                            'address',
                            'branches',
                            'contactNumber',
                            'isAdminVerified',
                        ],
                    }}
                />
                <DataTable
                    table={table}
                    isStickyHeader
                    isStickyFooter
                    isScrollable={isScrollable}
                    setColumnOrder={setColumnOrder}
                    className="mb-2 max-h-96 flex-1"
                />
                <DataTablePagination table={table} totalSize={data.length} />
            </div>
        </DataTableFilterContext.Provider>
    )
}

export default CompaniesTable
