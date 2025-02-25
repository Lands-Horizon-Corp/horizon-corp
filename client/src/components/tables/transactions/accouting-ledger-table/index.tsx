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

import IAccountingLedgerRequestTableColumns, {
    IAccountingLedgerTableColumnProps,
} from './columns'

import { cn } from '@/lib'
import { usePagination } from '@/hooks/use-pagination'
import useDatableFilterState from '@/hooks/use-filter-state'
import FilterContext from '@/contexts/filter-context/filter-context'
import useDataTableState from '@/hooks/data-table-hooks/use-datatable-state'
import { useDataTableSorting } from '@/hooks/data-table-hooks/use-datatable-sorting'

import { IAccountingLedgerRequest } from '@/server/types/accounts/accounting-ledger'
import { TableProps } from '../../types'
import { useFilteredPaginatedIAccountingLedger } from '@/hooks/api-hooks/transactions/use-accounting-ledger'
import { sampleLedgerData } from '@/components/forms/transactions/transaction-dummy-data'

export interface IAccountingLedgerRequestTableProps
    extends TableProps<IAccountingLedgerRequest>,
        IAccountingLedgerTableColumnProps {
    toolbarProps?: Omit<
        IDataTableToolbarProps<IAccountingLedgerRequest>,
        | 'table'
        | 'refreshActionProps'
        | 'globalSearchProps'
        | 'scrollableProps'
        | 'filterLogicProps'
        | 'exportActionProps'
        | 'deleteActionProps'
    >
}

const AccountsLedgerTable = ({
    className,
    toolbarProps,
    defaultFilter,
    onSelectData,
    actionComponent,
}: IAccountingLedgerRequestTableProps) => {
    const { pagination, setPagination } = usePagination()
    const { sortingState, tableSorting, setTableSorting } =
        useDataTableSorting()

    const columns = useMemo(
        () => IAccountingLedgerRequestTableColumns(),
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
    } = useDataTableState<IAccountingLedgerRequest>({
        columnOrder: columns.map((c) => c.id!),
        onSelectData,
    })

    const filterState = useDatableFilterState({
        defaultFilter,
        onFilterChange: () => setPagination({ ...pagination, pageIndex: 0 }),
    })

    const {
        isPending,
        isRefetching,
        data: { data, totalPage, pageSize, totalSize },
        refetch,
    } = useFilteredPaginatedIAccountingLedger({
        pagination,
        sort: sortingState,
        filterPayload: filterState.finalFilterPayload,
    })

    const handleRowSelectionChange = createHandleRowSelectionChange(data)

    const table = useReactTable({
        columns,
        data: sampleLedgerData,
        initialState: {
            columnPinning: { left: ['select', 'description', 'balance'] },
        },
        state: {
            sorting: tableSorting,
            pagination,
            columnOrder,
            rowSelection: rowSelectionState.rowSelection,
            columnVisibility,
        },
        rowCount: pageSize,
        manualSorting: true,
        pageCount: totalPage,
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
            <div
                className={cn(
                    'flex h-full w-full flex-col gap-y-2',
                    className,
                    !isScrollable && 'h-fit !max-h-none'
                )}
            >
                <DataTableToolbar
                    table={table}
                    refreshActionProps={{
                        onClick: () => refetch(),
                        isLoading: isPending || isRefetching,
                    }}
                    scrollableProps={{ isScrollable, setIsScrollable }}
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

export default AccountsLedgerTable
