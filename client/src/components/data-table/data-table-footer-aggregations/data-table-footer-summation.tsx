import { ReactNode, useMemo } from 'react'
import { Column, Header, Table } from '@tanstack/react-table'

import { cn } from '@/lib'
import { formatNumber } from '@/utils'
import { IBaseCompNoChild } from '@/types'

interface DataTableFooterSummationProps<TData, TValue>
    extends IBaseCompNoChild {
    totalLabel?: string | ReactNode
    table: Table<TData>
    column: Column<TData, TValue>
    header: Header<TData, TValue>
}

const DataTableFooterSummation = <TData, TValue>({
    table,
    column,
    className,
    totalLabel = 'Total',
}: DataTableFooterSummationProps<TData, TValue>) => {
    const sum = useMemo(() => {
        const rows = table.getCoreRowModel().rows

        return rows.reduce((total, row) => {
            const value = row.getValue<TValue>(column.id)
            return total + (typeof value === 'number' ? value : 0)
        }, 0)
    }, [table, column.id])

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
