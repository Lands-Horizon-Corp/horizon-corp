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
import DataTableFilterContext from '@/components/data-table/data-table-filter-context'
import { DataTableColumnHeader } from '@/components/data-table/data-table-column-header'
import useDatableFilterState from '@/components/data-table/hooks/use-datatable-filter-state'
import { PushPinSlashIcon } from '@/components/icons'
import DataTablePagination from '@/components/data-table/data-table-pagination'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'
import DataTableRefreshButton from '@/components/data-table/data-table-actions/data-table-refresh-button'
import DataTableActiveFilters from '@/components/data-table/data-table-actions/data-table-active-filters'
import DataTableDeleteSelected from '@/components/data-table/data-table-actions/data-table-delete-selected'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import { ColumnFilterProvider } from '@/components/data-table/data-table-column-header/column-filters/column-filter-state-context'
import TextFilter from '@/components/data-table/data-table-column-header/column-filters/text-filter'
import NumberFilter from '@/components/data-table/data-table-column-header/column-filters/number-filter'
import DateFilter from '@/components/data-table/data-table-column-header/column-filters/date-filter'
import MultiSelectFilter from '@/components/data-table/data-table-column-header/column-filters/multi-select-filter'
import DataTableOptionsMenu from '@/components/data-table/data-table-actions/data-table-options-menu'
import { Separator } from '@/components/ui/separator'

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

    const mockData: TData[] = Array.from({ length: 100 }).map(() => {
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
                            console.log('delete')
                        },
                    }}
                    onEdit={{
                        text: 'Edit',
                        isAllowed: false,
                        onClick: () => {
                            console.log('edit')
                        },
                    }}
                    onView={{
                        text: 'View',
                        isAllowed: true,
                        onClick: () => {
                            console.log('view')
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
                <ColumnFilterProvider dataType="text" fieldName="name">
                    <ColumnActions {...props}>
                        <TextFilter />
                    </ColumnActions>
                </ColumnFilterProvider>
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
                <ColumnFilterProvider dataType="date" fieldName="bday">
                    <ColumnActions {...props}>
                        <DateFilter />
                    </ColumnActions>
                </ColumnFilterProvider>
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
                <ColumnFilterProvider dataType="number" fieldName="age">
                    <ColumnActions {...props}>
                        <NumberFilter />
                    </ColumnActions>
                </ColumnFilterProvider>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { age },
            },
        }) => <div>{age}</div>,
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
                <ColumnFilterProvider dataType="enum" fieldName="gender">
                    <ColumnActions {...props}>
                        <MultiSelectFilter
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
                </ColumnFilterProvider>
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
        rowCount: 1000,
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
                            <DataTableActiveFilters />
                        </div>
                        <div className="flex items-center gap-x-2">
                            <DataTableDeleteSelected
                                table={table}
                                canDelete={true}
                                onClick={(rows) => {
                                    console.log('Deleting these -> ', rows)
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
                                fileName="sampletable"
                                columnsToExport={['name', 'age', 'bday']}
                            />
                        </div>
                    </div>
                    <DataTable
                        table={table}
                        isStickyHeader
                        isScrollable={isScrollable}
                        className="mb-2 max-h-96 flex-1"
                        setColumnOrder={setColumnOrder}
                    />
                </div>
                <DataTablePagination table={table} />
                <pre>{JSON.stringify(dataTableFilterState, null, 2)}</pre>
            </div>
        </DataTableFilterContext.Provider>
    )
}

export default Tbl
