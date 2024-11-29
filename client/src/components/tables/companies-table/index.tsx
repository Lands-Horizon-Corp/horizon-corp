import {
    getCoreRowModel,
    getSortedRowModel,
    useReactTable,
} from '@tanstack/react-table'
import { useMemo, useState } from 'react'

import DataTable from '@/components/data-table'
import { Separator } from '@/components/ui/separator'
import DataTableRefreshButton from '@/components/refresh-button'
import DataTablePagination from '@/components/data-table/data-table-pagination'
import useDataTableState from '@/components/data-table/hooks/use-datatable-state'
import DataTableOptionsMenu from '@/components/data-table/data-table-actions/data-table-options-menu'
import DataTableExportButton from '@/components/data-table/data-table-actions/data-table-export-button'
import DataTableActiveFilters from '@/components/data-table/data-table-actions/data-table-active-filters'
import DataTableDeleteSelected from '@/components/data-table/data-table-actions/data-table-delete-selected'

import logger from '@/helpers/loggers/logger'
import useDatableFilterState from '@/components/data-table/hooks/use-datatable-filter-state'
import DataTableFilterContext from '@/components/data-table/data-table-filters/data-table-filter-context'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { CompanyResource } from '@/horizon-corp/types'
import { companiesTableColumns as columns } from './columns'

const CompaniesTable = ({ className }: IBaseCompNoChild) => {
    const {
        sorting,
        setSorting,
        isScrollable,
        setIsScrollable,
        pagination,
        setPagination,
        rowSelection,
        setRowSelection,
        columnVisibility,
        setColumnVisibility,
    } = useDataTableState()

    const dataTableFilterState = useDatableFilterState()

    const [columnOrder, setColumnOrder] = useState<string[]>(() =>
        columns.map((c) => c.id!)
    )

    const memoizedData: CompanyResource[] = useMemo(
        () => [
            {
                id: 0,
                name: 'Johny',
                contactNumber: '0999999999',
                isAdminVerified: false,
                createdAt: '2024-11-29T04:03:53.492Z',
                updatedAt: '2024-11-29T04:03:53.492Z',
            },
        ],
        []
    )

    const table = useReactTable({
        columns,
        data: memoizedData,
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
        rowCount: memoizedData.length,
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
        <DataTableFilterContext.Provider value={dataTableFilterState}>
            <div className={cn('flex h-full flex-col gap-y-2', className)}>
                <div className="flex w-full max-w-full items-center justify-between gap-x-2">
                    <div className="flex items-center gap-x-2">
                        <DataTableActiveFilters />
                    </div>
                    <div className="flex items-center gap-x-2">
                        <DataTableDeleteSelected
                            table={table}
                            canDelete={true}
                            onClick={(rows) => {
                                logger.log('Deleting these -> ', rows)
                            }}
                        />
                        <DataTableRefreshButton
                            onClick={() => {}}
                            isLoading={false}
                        />
                        <DataTableOptionsMenu
                            table={table}
                            scrollOption={{
                                isScrollable,
                                setIsScrollable,
                            }}
                        />
                        <Separator
                            orientation="vertical"
                            className="h-full min-h-7"
                        />
                        <DataTableExportButton
                            table={table}
                            columnsToExport={['name']}
                        />
                    </div>
                </div>
                <DataTable
                    table={table}
                    isStickyHeader
                    isStickyFooter
                    isScrollable={isScrollable}
                    setColumnOrder={setColumnOrder}
                    className="mb-2 max-h-96 flex-1"
                />
                <DataTablePagination
                    table={table}
                    totalSize={memoizedData.length}
                />
            </div>
        </DataTableFilterContext.Provider>
    )
}

export default CompaniesTable
