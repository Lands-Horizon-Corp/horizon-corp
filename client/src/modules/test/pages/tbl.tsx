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
import { DataTableViewOptions } from '@/components/data-table/data-table-actions/data-table-column-toggle'
import { DataTableColumnHeader } from '@/components/data-table/data-table-column-header'
import useDatableFilterState from '@/components/data-table/hooks/use-datatable-filter-state'
import { PushPinSlashIcon } from '@/components/icons'
import DataTablePagination from '@/components/data-table/data-table-pagination'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'
import DataTableRefreshButton from '@/components/data-table/data-table-actions/data-table-refresh-button'
import DataTableActiveFilters from '@/components/data-table/data-table-actions/data-table-active-filters'

type TData = {
    name: string
    age: number
    bday: Date
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

    const mockData: TData[] = Array.from({ length: 100 }).map(() => {
        const age = Math.floor(Math.random() * (65 - 18 + 1)) + 18
        const name = names[Math.floor(Math.random() * names.length)]
        const bday = generateRandomDate(
            new Date(1959, 0, 1),
            new Date(2006, 0, 1)
        )

        return {
            name,
            age,
            bday,
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
            />
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
                dataType="date"
                title="bday"
                tooltipDescription="Birth Date of the member"
                {...props}
            />
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
                dataType="number"
                title="Age"
                tooltipDescription="Age of the member"
                {...props}
            />
        ),
        cell: ({
            row: {
                original: { age },
            },
        }) => <div>{age}</div>,
    },
]

const Tbl = () => {
    const {
        sorting,
        setSorting,
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
            columnOrder,
            rowSelection,
            columnVisibility,
            pagination,
        },
        manualSorting: true,
        manualFiltering: true,
        manualPagination: true,
        rowCount: 1000, // dito papasok ang total rows result galing sa paginated response
        onSortingChange: setSorting,
        onPaginationChange: setPagination,
        getCoreRowModel: getCoreRowModel(),
        onColumnOrderChange: setColumnOrder,
        onRowSelectionChange: setRowSelection,
        getSortedRowModel: getSortedRowModel(),
        onColumnVisibilityChange: setColumnVisibility,
    })

    /*

    ✅ - Add Row Action (with icons)
        ✅ - Delete (permission for this action)
        ✅ - View (permission for this action)
        ✅ - Edit (permission for this action)
    ✅ - Refresh Button
    - Delete Selected (left side nalangs)
   ✅ - Toggle Column Visibility add show all option
   ✅ - remove delay on filter change 500 
        - (Optional) Transfer debounce to the input itself?

    - Add reset all filter button
        - Filter as Chips with close
    
   ✅ - Pin default the action column to left 
   ✅ - remove delay on filter change 500 
    
    ✅ - Tooltip per column

    ✅ - Pagination
        ✅ - page size
        ✅ - total selected / total Data

    - On export, send Add id's, if all, pass no id's

    */

    return (
        <DataTableFilterContext.Provider value={dataTableFilterState}>
            <div className="flex w-full flex-col gap-y-4 py-8">
                <div className="flex h-full flex-col gap-y-2">
                    <div className="flex w-full max-w-full items-center justify-between gap-x-2">
                        <div className="flex items-center w-1/2 gap-x-2">
                            <DataTableActiveFilters />
                        </div>
                        <div className="flex items-center gap-x-2">
                            <DataTableRefreshButton
                                onClick={() => {}}
                                isLoading={false}
                            />
                            <DataTableExportButton table={table} />
                            <DataTableViewOptions table={table} />
                        </div>
                    </div>
                    <DataTable
                        table={table}
                        isStickyHeader
                        className="mb-2 flex-1 max-h-96"
                        setColumnOrder={setColumnOrder}
                    />
                </div>
                <DataTablePagination table={table} />
            </div>
        </DataTableFilterContext.Provider>
    )
}

export default Tbl
