import {
    useReactTable,
    getCoreRowModel,
    getSortedRowModel,
} from '@tanstack/react-table'
import { toast } from 'sonner'
import { useQuery } from '@tanstack/react-query'
import { forwardRef, useCallback, useImperativeHandle } from 'react'

import DataTable from '@/components/data-table'
import DataTableToolbar from '@/components/data-table/data-table-toolbar'
import DataTablePagination from '@/components/data-table/data-table-pagination'
import useDataTableState from '@/components/data-table/hooks/use-datatable-state'
import useDatableFilterState from '@/components/data-table/hooks/use-datatable-filter-state'
import DataTableFilterContext from '@/components/data-table/data-table-filters/data-table-filter-context'

import columns, { companyGlobalSearchTargets } from './columns'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { withCatchAsync, toBase64 } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'
import { CompanyPaginatedResource, CompanyResource } from '@/horizon-corp/types'

export interface CompaniesTableRef {
    selectedRowIds: Record<string, boolean>
}

const CompaniesTable = forwardRef<CompaniesTableRef, IBaseCompNoChild>(
    ({ className }: IBaseCompNoChild, ref) => {
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
            onFilterChange: () =>
                setPagination({ ...pagination, pageIndex: 0 }),
        })

        const {
            data: { data, totalPage, pageSize, totalSize },
            isPending,
            isRefetching,
            refetch,
        } = useQuery<CompanyPaginatedResource, string>({
            queryKey: [
                'company',
                'table',
                filterState.finalFilters,
                pagination,
            ],
            queryFn: async () => {
                const [error, result] = await withCatchAsync(
                    CompanyService.filter(
                        toBase64({
                            preloads: ['Media'],
                            ...pagination,
                            ...filterState.finalFilters,
                        })
                    )
                )

                if (error) {
                    const errorMessage = serverRequestErrExtractor({ error })
                    toast.error(errorMessage)
                    throw errorMessage
                }

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

        const getRowIdFn = useCallback(
            (row: CompanyResource) => `${row.id}`,
            []
        )

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
            getRowId: getRowIdFn,
            onSortingChange: setSorting,
            onPaginationChange: setPagination,
            getCoreRowModel: getCoreRowModel(),
            onColumnOrderChange: setColumnOrder,
            onRowSelectionChange: setRowSelection,
            getSortedRowModel: getSortedRowModel(),
            onColumnVisibilityChange: setColumnVisibility,
        })

        useImperativeHandle(ref, () => ({
            selectedRowIds: table.getState().rowSelection,
        }))

        return (
            <DataTableFilterContext.Provider value={filterState}>
                <div className={cn('flex h-full flex-col gap-y-2', className)}>
                    <DataTableToolbar
                        globalSearchProps={{
                            defaultMode: 'equal',
                            targets: companyGlobalSearchTargets,
                        }}
                        table={table}
                        refreshActionProps={{
                            isLoading: isPending || isRefetching,
                            onClick: () => refetch(),
                        }}
                        deleteActionProps={{
                            isLoading: false,
                            onClick: () => {},
                        }}
                        scrollableProps={{ isScrollable, setIsScrollable }}
                        exportActionProps={{
                            pagination,
                            isLoading: isPending,
                            filters: filterState.finalFilters,
                            disabled: isPending || isRefetching,
                            exportAll: CompanyService.exportAll,
                            exportAllFiltered: CompanyService.exportAllFiltered,
                            exportCurrentPage: CompanyService.exportCurrentPage,
                            exportSelected: (ids) =>
                                CompanyService.exportSelected(
                                    ids.map(({ id }) => id)
                                ),
                        }}
                        filterLogicProps={{
                            filterLogic: filterState.filterLogic,
                            setFilterLogic: filterState.setFilterLogic,
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
                    <DataTablePagination table={table} totalSize={totalSize} />
                </div>
            </DataTableFilterContext.Provider>
        )
    }
)

CompaniesTable.displayName = 'Companies Table'

export default CompaniesTable
