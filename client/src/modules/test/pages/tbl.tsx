import { useMemo, useState } from 'react'

import {
    ColumnDef,
    getCoreRowModel,
    getSortedRowModel,
    PaginationState,
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

    const mockData: TData[] = Array.from({ length: 3 }).map(() => {
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

const columns: ColumnDef<TData>[] = [
    {
        id: 'select',
        header: ({ table }) => (
            <div className={'w-fit items-center px-2'}>
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
        accessorKey: 'age',
        header: (props) => (
            <DataTableColumnHeader isResizable title="Age" {...props} />
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
        header: (props) => (
            <DataTableColumnHeader isResizable title="Birth Date" {...props} />
        ),
        cell: ({
            row: {
                original: { bday },
            },
        }) => <div className="text-nowrap">{bday.toLocaleString()}</div>,
    },
    {
        id: 'Name1',
        accessorKey: 'name1',
        header: (props) => (
            <DataTableColumnHeader isResizable title="Name1" {...props} />
        ),
        cell: ({
            row: {
                original: { name1 },
            },
        }) => <div>{name1}</div>,
    },
    {
        id: 'Age1',
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
        id: 'Birth Date 1',
        accessorKey: 'bday1',
        header: (props) => (
            <DataTableColumnHeader isResizable title="Birth Date" {...props} />
        ),
        cell: ({
            row: {
                original: { bday1 },
            },
        }) => <div className="text-nowrap">{bday1.toLocaleString()}</div>,
    },
]

const Tbl = () => {
    const [rowSelection, setRowSelection] = useState({})
    const [sorting, setSorting] = useState<SortingState>([])
    const [columnVisibility, setColumnVisibility] = useState({})
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
        debugColumns: true,
        debugCells: true,
        manualPagination: true,
        onSortingChange: setSorting,
        columnResizeMode: 'onChange',
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
                wrapperClassName="flex-1 mb-2 rounded-lg"
            />
        </div>
    )
}

export default Tbl
