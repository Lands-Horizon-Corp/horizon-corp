import { useMemo, useState } from 'react'

import {
    ColumnDef,
    getCoreRowModel,
    getSortedRowModel,
    useReactTable,
} from '@tanstack/react-table'

import DataTable from '@/components/data-table'
import { Checkbox } from '@/components/ui/checkbox'
import useDataTableState from '@/components/data-table/hooks/use-datatable-state'
import DataTableExportButton from '@/components/data-table/data-table-actions/data-table-export-button'
import DataTableFilterContext from '@/components/data-table/data-table-filters/data-table-filter-context'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import useDatableFilterState from '@/components/data-table/hooks/use-datatable-filter-state'
import { PushPinSlashIcon } from '@/components/icons'
import DataTablePagination from '@/components/data-table/data-table-pagination'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'
import RefreshButton from '@/components/refresh-button'
import DataTableActiveFilters from '@/components/data-table/data-table-actions/data-table-active-filters'
import DataTableDeleteSelected from '@/components/data-table/data-table-actions/data-table-delete-selected'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import MultiSelectFilter from '@/components/data-table/data-table-filters/multi-select-filter'
import DataTableOptionsMenu from '@/components/data-table/data-table-actions/data-table-options-menu'
import { Separator } from '@/components/ui/separator'
import DataTableFooterSummation from '@/components/data-table/data-table-footer-aggregations/data-table-footer-summation'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import NumberFilter from '@/components/data-table/data-table-filters/number-filter'
import DataTableGlobalSearch from '@/components/data-table/data-table-filters/data-table-global-search'
import logger from '@/helpers/loggers/logger'

type TData = {
    name: string
    age: number
    bday: Date
    gender: 'Male' | 'Female' | 'Other'
}

const generateRandomDate = (start: Date, end: Date): Date => {
    return new Date(
        start.getTime() + Math.random() * (end.getTime() - start.getTime())
    )
}

const data = () => {
    const names = [
        'Alice',
        'Bob',
        'Charlie',
        'Diana',
        'Eve',
        'Frank',
        'Grace',
        'Hank',
        'Ivy',
        'Jack',
    ]

    const genders: TData['gender'][] = ['Male', 'Female', 'Other'] // Gender options

    const mockData: TData[] = Array.from({ length: 20 }).map(() => {
        const age = Math.floor(Math.random() * (65 - 18 + 1)) + 18
        const name = names[Math.floor(Math.random() * names.length)]
        const bday = generateRandomDate(
            new Date(1959, 0, 1),
            new Date(2006, 0, 1)
        )
        const gender = genders[Math.floor(Math.random() * genders.length)] // Randomly select gender

        return {
            name,
            age,
            bday,
            gender,
        }
    })

    return mockData
}

const defaultColumns: ColumnDef<TData>[] = [
    {
        id: 'select',
        header: ({ table, column }) => (
            <div className={'flex w-fit items-center gap-x-1 px-2'}>
                <Checkbox
                    checked={table.getIsAllPageRowsSelected()}
                    onCheckedChange={(value) =>
                        table.toggleAllPageRowsSelected(!!value)
                    }
                    aria-label="Select all"
                />
                {!column.getIsPinned() && (
                    <PushPinSlashIcon
                        onClick={() => column.pin('left')}
                        className="mr-2 size-3.5 cursor-pointer"
                    />
                )}
            </div>
        ),
        cell: ({ row }) => (
            <div className="flex w-fit items-center gap-x-1 px-0">
                <RowActionsGroup
                    onDelete={{
                        text: 'Delete',
                        isAllowed: true,
                        onClick: () => {
                            logger.log('delete')
                        },
                    }}
                    onEdit={{
                        text: 'Edit',
                        isAllowed: false,
                        onClick: () => {
                            logger.log('edit')
                        },
                    }}
                    onView={{
                        text: 'View',
                        isAllowed: true,
                        onClick: () => {
                            logger.log('view')
                        },
                    }}
                />
                <Checkbox
                    checked={row.getIsSelected()}
                    onCheckedChange={(value) => row.toggleSelected(!!value)}
                    aria-label="Select row"
                />
            </div>
        ),
        enableSorting: false,
        enableHiding: false,
        maxSize: 80,
    },
    {
        id: 'Name',
        accessorKey: 'name',
        header: (props) => (
            <DataTableColumnHeader
                {...props}
                isResizable
                title="Name"
                dataType="text"
            >
                <ColumnActions {...props}>
                    <TextFilter<TData> accessorKey={'name'} />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { name },
            },
        }) => <div>{name}</div>,
    },
    {
        id: 'Bday',
        accessorKey: 'bday',
        header: (props) => (
            <DataTableColumnHeader
                isResizable
                title="bday"
                tooltipDescription="Birth Date of the member"
                {...props}
            >
                <ColumnActions {...props}>
                    <DateFilter<TData> accessorKey={'bday'} />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { bday },
            },
        }) => <div>{bday.toDateString()}</div>,
    },
    {
        id: 'Age',
        accessorKey: 'age',
        header: (props) => (
            <DataTableColumnHeader
                isResizable
                title="Age"
                tooltipDescription="Age of the member"
                {...props}
            >
                <ColumnActions {...props}>
                    <NumberFilter<TData> accessorKey="age" />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { age },
            },
        }) => <div>{age}</div>,
        footer: (props) => <DataTableFooterSummation {...props} />,
    },
    {
        id: 'Gender',
        accessorKey: 'gender',
        header: (props) => (
            <DataTableColumnHeader
                isResizable
                title="Gender"
                tooltipDescription="Gender of the user"
                {...props}
            >
                <ColumnActions {...props}>
                    <MultiSelectFilter<TData>
                        accessorKey={'gender'}
                        multiSelectOptions={[
                            {
                                label: 'Male',
                                value: 'male',
                            },
                            {
                                label: 'Femal',
                                value: 'female',
                            },
                            {
                                label: 'Other',
                                value: 'other',
                            },
                        ]}
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { gender },
            },
        }) => <div>{gender}</div>,
    },
]

const Tbl = () => {
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

    const [columns] = useState<typeof defaultColumns>(() => [...defaultColumns])
    const [columnOrder, setColumnOrder] = useState<string[]>(() =>
        columns.map((c) => c.id!)
    )

    const memoizedData = useMemo(() => data(), [])

    const dummyPageSize = 300

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
        rowCount: dummyPageSize,
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
            <div className="flex w-full flex-col gap-y-4 py-8">
                <div className="flex h-full flex-col gap-y-2">
                    <div className="flex w-full max-w-full items-center justify-between gap-x-2">
                        <div className="flex items-center gap-x-2">
                            <DataTableGlobalSearch<TData>
                                keysToSearch={['name', 'gender']}
                                defaultMode="contains"
                            />
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
                            <RefreshButton
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
                                columnsToExport={['name', 'age', 'bday']}
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
                </div>
                <DataTablePagination table={table} totalSize={dummyPageSize} />
                <pre>{JSON.stringify(dataTableFilterState, null, 2)}</pre>
            </div>
        </DataTableFilterContext.Provider>
    )
}

export default Tbl
