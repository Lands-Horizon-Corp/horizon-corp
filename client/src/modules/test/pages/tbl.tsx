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
import DataTableFooterSummation from '@/components/data-table/data-table-footer-aggregations/data-table-footer-summation'

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
    // return [
    //     {
    //         name: 'Ivy',
    //         age: 18,
    //         bday: new Date('1967-04-02T00:37:44.422Z'),
    //         gender: 'Other',
    //     },
    //     {
    //         name: 'Hank',
    //         age: 53,
    //         bday: new Date('1971-02-18T18:16:59.171Z'),
    //         gender: 'Other',
    //     },
    //     {
    //         name: 'Bob',
    //         age: 33,
    //         bday: new Date('2005-03-22T03:32:36.242Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Grace',
    //         age: 22,
    //         bday: new Date('1965-03-01T04:59:23.534Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Hank',
    //         age: 37,
    //         bday: new Date('1964-10-10T21:32:55.882Z'),
    //         gender: 'Male',
    //     },
    //     {
    //         name: 'Eve',
    //         age: 64,
    //         bday: new Date('1961-01-30T18:47:44.231Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Bob',
    //         age: 22,
    //         bday: new Date('1986-05-18T02:14:11.508Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Jack',
    //         age: 43,
    //         bday: new Date('2001-06-23T10:58:59.983Z'),
    //         gender: 'Male',
    //     },
    //     {
    //         name: 'Diana',
    //         age: 27,
    //         bday: new Date('1991-07-20T23:25:38.895Z'),
    //         gender: 'Male',
    //     },
    //     {
    //         name: 'Ivy',
    //         age: 35,
    //         bday: new Date('1992-03-02T08:20:40.745Z'),
    //         gender: 'Other',
    //     },
    //     {
    //         name: 'Grace',
    //         age: 46,
    //         bday: new Date('1970-10-09T01:18:39.321Z'),
    //         gender: 'Other',
    //     },
    //     {
    //         name: 'Hank',
    //         age: 22,
    //         bday: new Date('1964-02-28T09:23:51.803Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Eve',
    //         age: 64,
    //         bday: new Date('2005-04-18T18:44:24.463Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Grace',
    //         age: 31,
    //         bday: new Date('2003-01-19T14:26:38.490Z'),
    //         gender: 'Male',
    //     },
    //     {
    //         name: 'Alice',
    //         age: 59,
    //         bday: new Date('1971-02-15T05:37:29.557Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Ivy',
    //         age: 42,
    //         bday: new Date('1964-06-09T19:38:23.860Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Alice',
    //         age: 29,
    //         bday: new Date('1990-01-24T09:18:35.370Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Jack',
    //         age: 36,
    //         bday: new Date('2003-11-13T01:11:58.565Z'),
    //         gender: 'Female',
    //     },
    //     {
    //         name: 'Eve',
    //         age: 49,
    //         bday: new Date('2001-04-05T16:02:03.476Z'),
    //         gender: 'Male',
    //     },
    //     {
    //         name: 'Frank',
    //         age: 33,
    //         bday: new Date('1971-01-29T05:11:01.950Z'),
    //         gender: 'Female',
    //     },
    // ]
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
                        isStickyFooter
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
