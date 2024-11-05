import { useMemo, useState } from 'react'

import {
    ColumnDef,
    getCoreRowModel,
    getSortedRowModel,
    SortingState,
    useReactTable,
} from '@tanstack/react-table'
import DataTable from '@/components/data-table'
import { Checkbox } from '@/components/ui/checkbox'
import DataTableExportButton from '@/components/data-table/data-table-export-button'
import { DataTableViewOptions } from '@/components/data-table/data-table-column-toggle'
import { DataTableColumnHeader } from '@/components/data-table/data-table-column-header'


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

    const mockData: TData[] = Array.from({ length: 50 }).map(() => {
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

const columns: ColumnDef<TData>[] = [
    {
        id: 'select',
        header: ({ table }) => (
            <Checkbox
                checked={table.getIsAllPageRowsSelected()}
                onCheckedChange={(value) =>
                    table.toggleAllPageRowsSelected(!!value)
                }
                aria-label="Select all"
            />
        ),
        cell: ({ row }) => (
            <Checkbox
                checked={row.getIsSelected()}
                onCheckedChange={(value) => row.toggleSelected(!!value)}
                aria-label="Select row"
            />
        ),
        enableSorting: false,
        enableHiding: false,
    },
    {
        id: 'Name',
        accessorKey: 'name',
        header: ({ column }) => (
            <DataTableColumnHeader title="Name" column={column} />
        ),
        cell: ({
            row: {
                original: { name },
            },
        }) => <div>{name}</div>,
    },
    {
        id: 'Age',
        accessorKey: 'age',
        header: ({ column }) => (
            <DataTableColumnHeader title="Age" column={column} />
        ),
        cell: ({
            row: {
                original: { age },
            },
        }) => <div>{age}</div>,
    },
    {
        id: 'Birth Date',
        accessorKey: 'bday',
        header: ({ column }) => (
            <DataTableColumnHeader title="Birth Date" column={column} />
        ),
        cell: ({
            row: {
                original: { bday },
            },
        }) => <div>{bday.toLocaleString()}</div>,
    },
]

const Tbl = () => {
    const [rowSelection, setRowSelection] = useState({})
    const [sorting, setSorting] = useState<SortingState>([])
    const [columnVisibility, setColumnVisibility] = useState({})

    const memoizedData = useMemo(() => data(), [])

    const table = useReactTable({
        columns,
        data: memoizedData,
        state: {
            sorting,
            rowSelection,
            columnVisibility,
        },
        onSortingChange: setSorting,
        getCoreRowModel: getCoreRowModel(),
        onRowSelectionChange: setRowSelection,
        getSortedRowModel: getSortedRowModel(),
        onColumnVisibilityChange: setColumnVisibility,
    })

    return (
        <div className="flex h-full max-h-screen flex-col gap-y-2">
            <p>Table Desu</p>
            <div className="flex items-center justify-end gap-x-2">
                <DataTableExportButton table={table} />
                <DataTableViewOptions table={table} />
            </div>
            <DataTable
                table={table}
                isStickyHeader
                wrapperClassName="flex-1 mb-2 rounded-lg"
            />
        </div>
    )
}

export default Tbl
