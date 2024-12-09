import {
    useReactTable,
    getCoreRowModel,
    getSortedRowModel,
} from '@tanstack/react-table'
import { toast } from 'sonner'
import { useQuery } from '@tanstack/react-query'

import DataTable from '@/components/data-table'
import DataTableToolbar from '@/components/data-table/data-table-toolbar'
import DataTablePagination from '@/components/data-table/data-table-pagination'
import useDataTableState from '@/components/data-table/hooks/use-datatable-state'
import useDatableFilterState from '@/components/data-table/hooks/use-datatable-filter-state'
import DataTableFilterContext from '@/components/data-table/data-table-filters/data-table-filter-context'

import columns, { membersGlobalSearchTargets } from './columns'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { withCatchAsync, toBase64 } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
/* TO REPLACE ONCE MEMBER RESOURCE IS MADE */
import { GenderPaginatedResource } from '@/horizon-corp/types'
import GenderService from '@/horizon-corp/server/common/GenderService'

const MembersTable = ({ className }: IBaseCompNoChild) => {
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

    const filterState = useDatableFilterState()

    const {
        data: { data, totalPage, pageSize, totalSize },
        isPending,
        isRefetching,
        refetch,
    } = useQuery<GenderPaginatedResource, string>({
        queryKey: ['members', 'table', filterState.filters],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                GenderService.filter(
                    toBase64({
                        filters: filterState.finalFilters,
                        logic: filterState.filterLogic,
                        preloads: ['Media'],
                        ...pagination,
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

    return (
        <DataTableFilterContext.Provider value={filterState}>
            <div className={cn('flex h-full flex-col gap-y-2', className)}>
                <DataTableToolbar
                    globalSearchProps={{
                        defaultMode: 'equal',
                        targets: membersGlobalSearchTargets,
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
                        filters: filterState.finalFilters,
                        disabled: isPending || isRefetching,
                        /*
                        exportAll: .exportAll,
                        exportAllFiltered: .exportAllFiltered,
                        exportCurrentPage: .exportCurrentPage,
                        exportSelected: (ids) =>
                            .exportSelected(
                                ids.map(({ id }) => id)
                            ),
                            */
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

export default MembersTable
