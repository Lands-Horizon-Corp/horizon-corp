import { useMemo } from 'react'
import { Column, Header, Table } from '@tanstack/react-table'

import { cn, formatNumber } from '@/lib'
import { IBaseCompNoChild } from '@/types'

interface DataTableFooterSummationProps<TData, TValue>
    extends IBaseCompNoChild {
    table: Table<TData>
    sumAll?: boolean
    totalLabel?: string
    column: Column<TData, TValue>
    header: Header<TData, TValue>
}

const DataTableFooterSummation = <TData, TValue>({
    table,
    column,
    className,
    sumAll = false,
    totalLabel = 'Total',
}: DataTableFooterSummationProps<TData, TValue>) => {
    const sum = useMemo(() => {
        const { pageIndex, pageSize } = table.getState().pagination

        const startRow = pageIndex * pageSize
        const endRow = startRow + pageSize

        const rows = sumAll
            ? table.getCoreRowModel().rows
            : table.getCoreRowModel().rows.slice(startRow, endRow)

        return rows.reduce((total, row) => {
            const value = row.getValue<TValue>(column.id)
            return total + (typeof value === 'number' ? value : 0)
        }, 0)
    }, [table, column.id, sumAll])

    return (
        <div>
            <span className="text-muted-foreground">{totalLabel} : </span>
            <span className={cn('font-medium', className)}>
                {formatNumber(sum)}
            </span>
        </div>
    )
}

export default DataTableFooterSummation
