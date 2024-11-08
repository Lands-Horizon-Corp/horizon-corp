import { useMemo, useState } from 'react'

import {
    ColumnDef,
    getCoreRowModel,
    getSortedRowModel,
    PaginationState,
    SortingState,
    useReactTable,
} from '@tanstack/react-table'
import { Checkbox } from '@/components/ui/checkbox'
import DataTableExportButton from '@/components/data-table/data-table-export-button'
import { DataTableViewOptions } from '@/components/data-table/data-table-column-toggle'
import { DataTableColumnHeader } from '@/components/data-table/data-table-column-header'
import DataTable from '@/components/data-table'

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
            <DataTableColumnHeader isResizable title="Name" {...props} />
        ),
        cell: ({
            row: {
                original: { name },
            },
        }) => <div>{name}</div>,
    },
    {
        id: 'Age',
        accessorKey: 'age1',
        header: (props) => (
            <DataTableColumnHeader isResizable title="Age1" {...props} />
        ),
        cell: ({
            row: {
                original: { age1 },
            },
        }) => <div>{age1}</div>,
    },
    {
        id: 'Bday',
        accessorKey: 'bday',
        header: (props) => (
            <DataTableColumnHeader isResizable title="bday" {...props} />
        ),
        cell: ({
            row: {
                original: { bday },
            },
        }) => <div>{bday.toDateString()}</div>,
    },

    {
        id: 'Name2',
        accessorKey: 'name',
        header: (props) => (
            <DataTableColumnHeader isResizable title="Name2" {...props} />
        ),
        cell: ({
            row: {
                original: { name },
            },
        }) => <div>{name}</div>,
    },
    {
        id: 'Age2',
        accessorKey: 'age1',
        header: (props) => (
            <DataTableColumnHeader isResizable title="Age2" {...props} />
        ),
        cell: ({
            row: {
                original: { age1 },
            },
        }) => <div>{age1}</div>,
    },
    {
        id: 'Bday1',
        accessorKey: 'bday',
        header: (props) => (
            <DataTableColumnHeader isResizable title="bday1" {...props} />
        ),
        cell: ({
            row: {
                original: { bday },
            },
        }) => <div>{bday.toDateString()}</div>,
    },

    {
        id: 'Name3',
        accessorKey: 'name',
        header: (props) => (
            <DataTableColumnHeader isResizable title="Name3" {...props} />
        ),
        cell: ({
            row: {
                original: { name },
            },
        }) => <div>{name}</div>,
    },
    {
        id: 'Age3',
        accessorKey: 'age1',
        header: (props) => (
            <DataTableColumnHeader isResizable title="Age3" {...props} />
        ),
        cell: ({
            row: {
                original: { age1 },
            },
        }) => <div>{age1}</div>,
    },
    {
        id: 'Bday2',
        accessorKey: 'bday',
        header: (props) => (
            <DataTableColumnHeader isResizable title="bday2" {...props} />
        ),
        cell: ({
            row: {
                original: { bday },
            },
        }) => <div>{bday.toDateString()}</div>,
    },
]

const Tbl = () => {
    const [rowSelection, setRowSelection] = useState({})
    const [sorting, setSorting] = useState<SortingState>([])
    const [columnVisibility, setColumnVisibility] = useState({})
    const [columns] = useState<typeof defaultColumns>(() => [...defaultColumns])
    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: 0,
        pageSize: 10,
    })

    const memoizedData = useMemo(() => data(), [])

    const table = useReactTable({
        columns,
        data: memoizedData,
        state: {
            sorting,
            rowSelection,
            columnVisibility,
            pagination,
        },
        manualPagination: true,
        onSortingChange: setSorting,
        onPaginationChange: setPagination,
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
                className="flex-1 mb-2"
            />
        </div>
    )
}

export default Tbl
