import {
    useReactTable,
    getCoreRowModel,
    getSortedRowModel,
} from '@tanstack/react-table'
import { useMemo } from 'react'
import { useQueryClient } from '@tanstack/react-query'

import DataTable from '@/components/data-table'
import DataTableToolbar, {
    IDataTableToolbarProps,
} from '@/components/data-table/data-table-toolbar'
import DataTablePagination from '@/components/data-table/data-table-pagination'
import useDataTableState from '@/components/data-table/hooks/use-datatable-state'
import useDatableFilterState from '@/components/data-table/hooks/use-datatable-filter-state'
import DataTableFilterContext from '@/components/data-table/data-table-filters/data-table-filter-context'

import companyColumns, {
    companyGlobalSearchTargets,
    ICompaniesTableColumnProps,
} from './columns'

import { cn } from '@/lib'
import { TableProps } from '../types'
import { CompanyResource } from '@/horizon-corp/types'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'
import { useFilteredPaginatedCompanies } from '@/hooks/api-hooks/use-company'

export interface CompaniesTableProps
    extends TableProps<CompanyResource>,
        ICompaniesTableColumnProps {
    toolbarProps?: Omit<
        IDataTableToolbarProps<CompanyResource>,
        | 'table'
        | 'refreshActionProps'
        | 'globalSearchProps'
        | 'scrollableProps'
        | 'filterLogicProps'
        | 'exportActionProps'
        | 'deleteActionProps'
    >
}

const CompaniesTable = ({
    className,
    onSelectData,
    toolbarProps,
    actionComponent,
}: CompaniesTableProps) => {
    const queryClient = useQueryClient()

    const columns = useMemo(
        () =>
            companyColumns({
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
    } = useDataTableState<CompanyResource>({
        columnOrder: columns.map((c) => c.id!),
        onSelectData,
    })

    const filterState = useDatableFilterState({
        onFilterChange: () => setPagination({ ...pagination, pageIndex: 0 }),
    })

    const {
        data: { data, totalPage, pageSize, totalSize },
        isPending,
        isRefetching,
        refetch,
    } = useFilteredPaginatedCompanies({
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
        <DataTableFilterContext.Provider value={filterState}>
            <div className={cn('flex h-full flex-col gap-y-2', className)}>
                <DataTableToolbar
                    globalSearchProps={{
                        defaultMode: 'equal',
                        targets: companyGlobalSearchTargets,
                    }}
                    table={table}
                    refreshActionProps={{
                        onClick: () => refetch(),
                        isLoading: isPending || isRefetching,
                    }}
                    deleteActionProps={{
                        onDeleteSuccess: () =>
                            queryClient.invalidateQueries({
                                queryKey: ['company', 'table'],
                            }),
                        onDelete: (selectedData) =>
                            CompanyService.deleteMany(
                                selectedData.map((data) => data.id)
                            ),
                    }}
                    scrollableProps={{ isScrollable, setIsScrollable }}
                    exportActionProps={{
                        pagination,
                        isLoading: isPending,
                        filters: filterState.finalFilterPayload,
                        disabled: isPending || isRefetching,
                        exportAll: CompanyService.exportAll,
                        exportAllFiltered: CompanyService.exportAllFiltered,
                        exportCurrentPage: (ids) =>
                            CompanyService.exportSelected(
                                ids.map((data) => data.id)
                            ),
                        exportSelected: (ids) =>
                            CompanyService.exportSelected(
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
        </DataTableFilterContext.Provider>
    )
}

export default CompaniesTable
