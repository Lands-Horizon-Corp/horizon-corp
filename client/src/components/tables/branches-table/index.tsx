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
import useDataTableState from '@/components/data-table/hooks/use-datatable-state'

import branchColumns, {
    branchesGlobalSearchTargets,
    IBranchesTableColumnProps,
} from './columns'

import { cn } from '@/lib'
import { TableProps } from '../types'
import { BranchResource } from '@/horizon-corp/types'
import useFilterState from '@/hooks/use-filter-state'
import FilterContext from '@/contexts/filter-context/filter-context'
import BranchService from '@/horizon-corp/server/admin/BranchService'
import { useFilteredPaginatedBranch } from '@/hooks/api-hooks/use-branch'

export interface BranchesTableProps
    extends TableProps<BranchResource>,
        IBranchesTableColumnProps {
    toolbarProps?: Omit<
        IDataTableToolbarProps<BranchResource>,
        | 'table'
        | 'refreshActionProps'
        | 'globalSearchProps'
        | 'scrollableProps'
        | 'filterLogicProps'
        | 'exportActionProps'
        | 'deleteActionProps'
    >
}

const BranchesTable = ({
    className,
    onSelectData,
    toolbarProps,
    defaultFilter,
    actionComponent,
}: BranchesTableProps) => {
    const columns = useMemo(
        () =>
            branchColumns({
                actionComponent,
            }),
        [actionComponent]
    )

    const {
        sorting,
        setSorting,
        getRowIdFn,
        pagination,
        setPagination,
        columnOrder,
        setColumnOrder,
        isScrollable,
        setIsScrollable,
        rowSelectionState,
        createHandleRowSelectionChange,
        columnVisibility,
        setColumnVisibility,
    } = useDataTableState<BranchResource>({
        columnOrder: columns.map((c) => c.id!),
        onSelectData,
    })

    const filterState = useFilterState({
        defaultFilter,
        onFilterChange: () => setPagination({ ...pagination, pageIndex: 0 }),
    })

    const {
        data: { data, totalPage, pageSize, totalSize },
        isPending,
        isRefetching,
        refetch,
    } = useFilteredPaginatedBranch({
        filterPayload: filterState.finalFilterPayload,
        pagination,
    })

    const handleRowSelectionChange = createHandleRowSelectionChange(data)

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
            rowSelection: rowSelectionState.rowSelection,
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

export default BranchesTable
