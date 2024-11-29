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
                id: 1,
                name: 'User1',
                contactNumber: '0991234567',
                isAdminVerified: true,
                createdAt: '2023-11-28T12:00:00.000Z',
                updatedAt: '2023-11-29T12:00:00.000Z',
            },
            {
                id: 2,
                name: 'User2',
                contactNumber: '0992345678',
                isAdminVerified: false,
                createdAt: '2023-11-27T11:30:00.000Z',
                updatedAt: '2023-11-28T12:45:00.000Z',
            },
            {
                id: 3,
                name: 'User3',
                contactNumber: '0993456789',
                isAdminVerified: true,
                createdAt: '2023-11-26T10:15:00.000Z',
                updatedAt: '2023-11-27T11:15:00.000Z',
            },
            {
                id: 4,
                name: 'User4',
                contactNumber: '0994567890',
                isAdminVerified: false,
                createdAt: '2023-11-25T09:45:00.000Z',
                updatedAt: '2023-11-26T10:45:00.000Z',
            },
            {
                id: 5,
                name: 'User5',
                contactNumber: '0995678901',
                isAdminVerified: true,
                createdAt: '2023-11-24T08:30:00.000Z',
                updatedAt: '2023-11-25T09:30:00.000Z',
            },
            {
                id: 6,
                name: 'User6',
                contactNumber: '0996789012',
                isAdminVerified: false,
                createdAt: '2023-11-23T07:00:00.000Z',
                updatedAt: '2023-11-24T08:15:00.000Z',
            },
            {
                id: 7,
                name: 'User7',
                contactNumber: '0997890123',
                isAdminVerified: true,
                createdAt: '2023-11-22T06:45:00.000Z',
                updatedAt: '2023-11-23T07:15:00.000Z',
            },
            {
                id: 8,
                name: 'User8',
                contactNumber: '0998901234',
                isAdminVerified: false,
                createdAt: '2023-11-21T05:30:00.000Z',
                updatedAt: '2023-11-22T06:00:00.000Z',
            },
            {
                id: 9,
                name: 'User9',
                contactNumber: '0999012345',
                isAdminVerified: true,
                createdAt: '2023-11-20T04:15:00.000Z',
                updatedAt: '2023-11-21T05:00:00.000Z',
            },
            {
                id: 10,
                name: 'User10',
                contactNumber: '0990123456',
                isAdminVerified: false,
                createdAt: '2023-11-19T03:00:00.000Z',
                updatedAt: '2023-11-20T04:15:00.000Z',
            },
            {
                id: 11,
                name: 'User11',
                contactNumber: '0993456721',
                isAdminVerified: true,
                createdAt: '2023-11-18T11:00:00.000Z',
                updatedAt: '2023-11-19T12:00:00.000Z',
            },
            {
                id: 12,
                name: 'User12',
                contactNumber: '0994567891',
                isAdminVerified: false,
                createdAt: '2023-11-17T10:00:00.000Z',
                updatedAt: '2023-11-18T11:00:00.000Z',
            },
            {
                id: 13,
                name: 'User13',
                contactNumber: '0995678932',
                isAdminVerified: true,
                createdAt: '2023-11-16T09:00:00.000Z',
                updatedAt: '2023-11-17T10:00:00.000Z',
            },
            {
                id: 14,
                name: 'User14',
                contactNumber: '0996789132',
                isAdminVerified: false,
                createdAt: '2023-11-15T08:00:00.000Z',
                updatedAt: '2023-11-16T09:00:00.000Z',
            },
            {
                id: 15,
                name: 'User15',
                contactNumber: '0997899134',
                isAdminVerified: true,
                createdAt: '2023-11-14T07:00:00.000Z',
                updatedAt: '2023-11-15T08:00:00.000Z',
            },
            {
                id: 16,
                name: 'User16',
                contactNumber: '0998912342',
                isAdminVerified: false,
                createdAt: '2023-11-13T06:00:00.000Z',
                updatedAt: '2023-11-14T07:00:00.000Z',
            },
            {
                id: 17,
                name: 'User17',
                contactNumber: '0999012321',
                isAdminVerified: true,
                createdAt: '2023-11-12T05:00:00.000Z',
                updatedAt: '2023-11-13T06:00:00.000Z',
            },
            {
                id: 18,
                name: 'User18',
                contactNumber: '0992348124',
                isAdminVerified: false,
                createdAt: '2023-11-11T04:00:00.000Z',
                updatedAt: '2023-11-12T05:00:00.000Z',
            },
            {
                id: 19,
                name: 'User19',
                contactNumber: '0997458932',
                isAdminVerified: true,
                createdAt: '2023-11-10T03:00:00.000Z',
                updatedAt: '2023-11-11T04:00:00.000Z',
            },
            {
                id: 20,
                name: 'User20',
                contactNumber: '0998213124',
                isAdminVerified: false,
                createdAt: '2023-11-09T02:00:00.000Z',
                updatedAt: '2023-11-10T03:00:00.000Z',
            },
            {
                id: 21,
                name: 'User21',
                contactNumber: '0998798321',
                isAdminVerified: true,
                createdAt: '2023-11-08T01:00:00.000Z',
                updatedAt: '2023-11-09T02:00:00.000Z',
            },
            {
                id: 22,
                name: 'User22',
                contactNumber: '0991328932',
                isAdminVerified: false,
                createdAt: '2023-11-07T12:00:00.000Z',
                updatedAt: '2023-11-08T01:00:00.000Z',
            },
            {
                id: 23,
                name: 'User23',
                contactNumber: '0994321342',
                isAdminVerified: true,
                createdAt: '2023-11-06T11:00:00.000Z',
                updatedAt: '2023-11-07T12:00:00.000Z',
            },
            {
                id: 24,
                name: 'User24',
                contactNumber: '0995321892',
                isAdminVerified: false,
                createdAt: '2023-11-05T10:00:00.000Z',
                updatedAt: '2023-11-06T11:00:00.000Z',
            },
            {
                id: 25,
                name: 'User25',
                contactNumber: '0996328931',
                isAdminVerified: true,
                createdAt: '2023-11-04T09:00:00.000Z',
                updatedAt: '2023-11-05T10:00:00.000Z',
            },
            {
                id: 26,
                name: 'User26',
                contactNumber: '0997321942',
                isAdminVerified: false,
                createdAt: '2023-11-03T08:00:00.000Z',
                updatedAt: '2023-11-04T09:00:00.000Z',
            },
            {
                id: 27,
                name: 'User27',
                contactNumber: '0998321924',
                isAdminVerified: true,
                createdAt: '2023-11-02T07:00:00.000Z',
                updatedAt: '2023-11-03T08:00:00.000Z',
            },
            {
                id: 28,
                name: 'User28',
                contactNumber: '0999328124',
                isAdminVerified: false,
                createdAt: '2023-11-01T06:00:00.000Z',
                updatedAt: '2023-11-02T07:00:00.000Z',
            },
            {
                id: 29,
                name: 'User29',
                contactNumber: '0990123184',
                isAdminVerified: true,
                createdAt: '2023-10-31T05:00:00.000Z',
                updatedAt: '2023-11-01T06:00:00.000Z',
            },
            {
                id: 30,
                name: 'User30',
                contactNumber: '0992189321',
                isAdminVerified: false,
                createdAt: '2023-10-30T04:00:00.000Z',
                updatedAt: '2023-10-31T05:00:00.000Z',
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
