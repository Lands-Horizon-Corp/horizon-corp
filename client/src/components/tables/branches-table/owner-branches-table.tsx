import {
    useReactTable,
    getCoreRowModel,
    getSortedRowModel,
} from '@tanstack/react-table'
import { useMemo } from 'react'

import DataTable from '@/components/data-table'
import DataTableToolbar, {
    IDataTableToolbarProps,
} from '@/components/data-table/data-table-toolbar'
import DataTablePagination from '@/components/data-table/data-table-pagination'

import branchColumns, {
    IBranchesTableColumnProps,
    branchesGlobalSearchTargets,
} from './columns'

import { TableProps } from '../types'
import { IBranchResource, TEntityId } from '@/server/types'

import { cn } from '@/lib'
import useFilterState from '@/hooks/use-filter-state'
import { usePagination } from '@/hooks/use-pagination'
import BranchService from '@/server/api-service/branch-service'
import FilterContext from '@/contexts/filter-context/filter-context'
import { useOwnerPaginatedBranch } from '@/hooks/api-hooks/use-owner'
import useDataTableState from '@/hooks/data-table-hooks/use-datatable-state'
import { useDataTableSorting } from '@/hooks/data-table-hooks/use-datatable-sorting'

export interface BranchesTableProps
    extends TableProps<IBranchResource>,
        IBranchesTableColumnProps {
    toolbarProps?: Omit<
        IDataTableToolbarProps<IBranchResource>,
        | 'table'
        | 'refreshActionProps'
        | 'globalSearchProps'
        | 'scrollableProps'
        | 'filterLogicProps'
        | 'exportActionProps'
        | 'deleteActionProps'
    >
    ownerId: TEntityId
}

const OwnerBranchesTable = ({
    ownerId,
    className,
    onSelectData,
    toolbarProps,
    defaultFilter,
    actionComponent,
}: BranchesTableProps) => {
    const { pagination, setPagination } = usePagination()
    const { sortingState, tableSorting, setTableSorting } =
        useDataTableSorting()
    const columns = useMemo(
        () =>
            branchColumns({
                actionComponent,
            }),
        [actionComponent]
    )

    const {
        getRowIdFn,
        columnOrder,
        setColumnOrder,
        isScrollable,
        setIsScrollable,
        columnVisibility,
        setColumnVisibility,
        rowSelectionState,
        createHandleRowSelectionChange,
    } = useDataTableState<IBranchResource>({
        columnOrder: columns.map((c) => c.id!),
        onSelectData,
    })

    const filterState = useFilterState({
        defaultFilter,
        onFilterChange: () => setPagination({ ...pagination, pageIndex: 0 }),
    })

    const {
        isPending,
        isRefetching,
        data: { data, totalPage, pageSize, totalSize },
        refetch,
    } = useOwnerPaginatedBranch({
        ownerId,
        pagination,
        sort: sortingState,
        filterPayload: filterState.finalFilterPayload,
    })

    const handleRowSelectionChange = createHandleRowSelectionChange(data)

    const table = useReactTable({
        columns,
        data: data,
        initialState: {
            columnPinning: { left: ['select'] },
        },
        state: {
            sorting: tableSorting,
            pagination,
            columnOrder,
            rowSelection: rowSelectionState.rowSelection,
            columnVisibility,
        },
        rowCount: pageSize,
        pageCount: totalPage,
        manualSorting: true,
        manualFiltering: true,
        manualPagination: true,
        getRowId: getRowIdFn,
        onSortingChange: setTableSorting,
        onPaginationChange: setPagination,
        getCoreRowModel: getCoreRowModel(),
        onColumnOrderChange: setColumnOrder,
        getSortedRowModel: getSortedRowModel(),
        onColumnVisibilityChange: setColumnVisibility,
        onRowSelectionChange: handleRowSelectionChange,
    })

    return (
        <FilterContext.Provider value={filterState}>
            <div className={cn('flex h-full flex-col gap-y-2', className)}>
                <DataTableToolbar
                    globalSearchProps={{
                        defaultMode: 'equal',
                        targets: branchesGlobalSearchTargets,
                    }}
                    table={table}
                    refreshActionProps={{
                        onClick: () => refetch(),
                        isLoading: isPending || isRefetching,
                    }}
                    deleteActionProps={{
                        onDelete: (selectedData) =>
                            BranchService.deleteMany(
                                selectedData.map((data) => data.id)
                            ),
                    }}
                    scrollableProps={{ isScrollable, setIsScrollable }}
                    exportActionProps={{
                        pagination,
                        isLoading: isPending,
                        filters: filterState.finalFilterPayload,
                        disabled: isPending || isRefetching,
                        exportAll: BranchService.exportAll,
                        exportAllFiltered: BranchService.exportAllFiltered,
                        exportCurrentPage: (ids) =>
                            BranchService.exportSelected(
                                ids.map((data) => data.id)
                            ),
                        exportSelected: (ids) =>
                            BranchService.exportSelected(
                                ids.map((data) => data.id)
                            ),
                    }}
                    filterLogicProps={{
                        filterLogic: filterState.filterLogic,
                        setFilterLogic: filterState.setFilterLogic,
                    }}
                    {...toolbarProps}
                />
                <DataTable
                    table={table}
                    isStickyHeader
                    isStickyFooter
                    isScrollable={isScrollable}
                    setColumnOrder={setColumnOrder}
                    className="mb-2"
                />
                <DataTablePagination table={table} totalSize={totalSize} />
            </div>
        </FilterContext.Provider>
    )
}

export default OwnerBranchesTable
