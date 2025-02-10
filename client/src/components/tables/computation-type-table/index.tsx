import {
    useReactTable,
    getCoreRowModel,
    getSortedRowModel,
} from '@tanstack/react-table'
import { useMemo } from 'react'
// import { useQueryClient } from '@tanstack/react-query'

import DataTable from '@/components/data-table'
import DataTableToolbar, {
    IDataTableToolbarProps,
} from '@/components/data-table/data-table-toolbar'
import DataTablePagination from '@/components/data-table/data-table-pagination'


import { cn } from '@/lib'
import { usePagination } from '@/hooks/use-pagination'
import useDatableFilterState from '@/hooks/use-filter-state'
import FilterContext from '@/contexts/filter-context/filter-context'
import useDataTableState from '@/hooks/data-table-hooks/use-datatable-state'
import { useDataTableSorting } from '@/hooks/data-table-hooks/use-datatable-sorting'

import { TableProps } from '../types'
import AccountsComputationTypeTableColumns, { IAccountsComputationTypeTableColumnProps } from './column'
import { genderGlobalSearchTargets } from '../genders-table/columns'
import { dummyAccountComputationTypeData, IAccountsComputationTypeResource } from '@/server/types/accounts/computation-type'

export interface AccountsComputationTypeTableProps
    extends TableProps<IAccountsComputationTypeResource>,
        IAccountsComputationTypeTableColumnProps {
    toolbarProps?: Omit<
        IDataTableToolbarProps<IAccountsComputationTypeResource>,
        | 'table'
        | 'refreshActionProps'
        | 'globalSearchProps'
        | 'scrollableProps'
        | 'filterLogicProps'
        | 'exportActionProps'
        | 'deleteActionProps'
    >
}

const AccountsComputationTypeTable = ({
    className,
    toolbarProps,
    defaultFilter,
    onSelectData,
    actionComponent,
}: AccountsComputationTypeTableProps) => {
    // const queryClient = useQueryClient()
    const { pagination, setPagination } = usePagination()
    const { sortingState, tableSorting, setTableSorting } =
        useDataTableSorting()

    const columns = useMemo(
        () =>
            AccountsComputationTypeTableColumns({
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
    } = useDataTableState<IAccountsComputationTypeResource>({
        columnOrder: columns.map((c) => c.id!),
        onSelectData,
    })

    const filterState = useDatableFilterState({
        defaultFilter,
        onFilterChange: () => setPagination({ ...pagination, pageIndex: 0 }),
    })

    // const {
    //     isPending,
    //     isRefetching,
    //     data: { data, totalPage, pageSize, totalSize },
    //     refetch,
    // } = useFilteredPaginatedAccountComputationTypes({
    //     pagination,
    //     sort: sortingState,
    //     filterPayload: filterState.finalFilterPayload,
    // })

    const handleRowSelectionChange = createHandleRowSelectionChange(dummyAccountComputationTypeData)

    const table = useReactTable({
        columns,
        data: dummyAccountComputationTypeData,
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
        // rowCount: pageSize,
        manualSorting: true,
        // pageCount: totalPage,
        enableMultiSort: false,
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
                        targets: genderGlobalSearchTargets,
                    }}
                    table={table}
                    refreshActionProps={{
                        onClick: () => {},
                        // isLoading: isPending || isRefetching,
                    }}
                    // deleteActionProps={{
                    //     onDeleteSuccess: () =>
                    //         queryClient.invalidateQueries({
                    //             queryKey: ['account-computation-type', 'resource-query'],
                    //         }),
                    //     onDelete: (selectedData) =>
                    //         AccountComputationTypeService.deleteMany(
                    //             selectedData.map((data) => data.id)
                    //         ),
                    // }}
                    scrollableProps={{ isScrollable, setIsScrollable }}
                    exportActionProps={{
                        pagination,
                        // isLoading: isPending,
                        filters: filterState.finalFilterPayload,
                        // disabled: isPending || isRefetching,
                        // exportAll: AccountComputationTypeService.exportAll,
                        // exportCurrentPage: (ids) =>
                        //     AccountComputationTypeService.exportSelected(
                        //         ids.map((data) => data.id)
                        //     ),
                        // exportSelected: (ids) =>
                        //     AccountComputationTypeService.exportSelected(
                        //         ids.map((data) => data.id)
                        //     ),
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
                <DataTablePagination table={table} totalSize={0} />
            </div>
        </FilterContext.Provider>
    )
}

export default AccountsComputationTypeTable
