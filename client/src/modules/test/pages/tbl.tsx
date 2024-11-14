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
import DataTableExportButton from '@/components/data-table/data-table-export-button'
import { DataTableViewOptions } from '@/components/data-table/data-table-column-toggle'
import { DataTableColumnHeader } from '@/components/data-table/data-table-column-header'
import DataTableFilterContext from '@/components/data-table/data-table-filter-context'
import useDatableFilterState from '@/components/data-table/hooks/use-datatable-filter-state'

type TData = {
    name: string
    age: number
    bday: Date
    name1: string
    age1: number
    bday1: Date
}

const generateRandomDate = (start: Date, end: Date): Date => {
    return new Date(
        start.getTime() + Math.random() * (end.getTime() - start.getTime())
    )
}

const data = () => {
    // return []

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
            name1: name,
            age1: age,
            bday1: bday,
        }
    })

    return mockData
}

const defaultColumns: ColumnDef<TData>[] = [
    {
        id: 'select',
        header: ({ table }) => (
            <div className={'flex w-fit items-center px-2'}>
                <Checkbox
                    checked={table.getIsAllPageRowsSelected()}
                    onCheckedChange={(value) =>
                        table.toggleAllPageRowsSelected(!!value)
                    }
                    aria-label="Select all"
                />
            </div>
        ),
        cell: ({ row }) => (
            <div className={'w-fit items-center px-2'}>
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
                dataType="text"
                isResizable
                title="Name"
                {...props}
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
                {...props}
            />
        ),
        cell: ({
            row: {
                original: { age1 },
            },
        }) => <div>{age1}</div>,
    },
]

const Tbl = () => {
    const {
        sorting,
        setSorting,
        pagination,
        setPagination,
        columnVisibility,
        setColumnVisibility,
        rowSelection,
        setRowSelection,
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
            <div className="flex h-full max-h-screen flex-col gap-y-2">
                <p>Table Desu</p>
                <pre>{JSON.stringify(dataTableFilterState.filters, null, 4)}</pre>
                <div className="flex items-center justify-end gap-x-2">
                    <DataTableExportButton table={table} />
                    <DataTableViewOptions table={table} />
                </div>
                <DataTable
                    table={table}
                    isStickyHeader
                    className="mb-2 flex-1"
                    setColumnOrder={setColumnOrder}
                />
            </div>
        </DataTableFilterContext.Provider>
    )
}

export default Tbl
