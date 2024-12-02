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

import { cn, withCatchAsync } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { serverRequestErrExtractor } from '@/helpers'
import { CompanyResource } from '@/horizon-corp/types'
import { companiesTableColumns as columns } from './columns'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'
import logger from '@/helpers/loggers/logger'

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

    const dataTableFilterState = useDatableFilterState()

    logger.log(dataTableFilterState.finalFilters)

    const { data, isPending, refetch } = useQuery<CompanyResource[], string>({
        queryKey: ['company-list' /* FILTER DITO DESU */],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                CompanyService.getAll()
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            if (!result) throw "Something wen't wroing"

            return result ?? []
        },
        initialData: [
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
        ],
        retry: 1,
    })

    const table = useReactTable({
        columns,
        data,
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
        rowCount: data.length,
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
