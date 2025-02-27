import { ReactNode } from 'react'
import { ColumnDef, Row } from '@tanstack/react-table'

import { Checkbox } from '@/components/ui/checkbox'
import { PushPinSlashIcon } from '@/components/icons'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import { IGlobalSearchTargets } from '@/components/data-table/data-table-filters/data-table-global-search'

import { toReadableDate } from '@/utils'
import { ITransactionPaymentTypesResource } from '@/server/types/transactions/transaction-payment-types'

export const transactionTypeGlobalSearchTargets: IGlobalSearchTargets<ITransactionPaymentTypesResource>[] =
    [
        { field: 'name', displayText: 'Name' },
        { field: 'description', displayText: 'Description' },
        { field: 'cheque_id', displayText: 'Cheque ID' },
    ]

export interface ITransactionPaymentTypesTableActionComponentProp {
    row: Row<ITransactionPaymentTypesResource>
}

export interface ITransactionPaymentTypesTableColumnProps {
    actionComponent?: (
        props: ITransactionPaymentTypesTableActionComponentProp
    ) => ReactNode
}

const TransactionPaymentTypesTableColumns = (
    opts?: ITransactionPaymentTypesTableColumnProps
): ColumnDef<ITransactionPaymentTypesResource>[] => {
    return [
        {
            id: 'select',
            header: ({ table, column }) => (
                <div className="flex w-fit items-center gap-x-1 px-2">
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
                    {opts?.actionComponent?.({ row })}
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
            id: 'name',
            accessorKey: 'name',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Name">
                    <ColumnActions {...props}>
                        <TextFilter<ITransactionPaymentTypesResource>
                            displayText="Name"
                            field="name"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { name },
                },
            }) => <div>{name}</div>,
            enableMultiSort: true,
        },
        {
            id: 'description',
            accessorKey: 'description',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="Description"
                >
                    <ColumnActions {...props}>
                        <TextFilter<ITransactionPaymentTypesResource>
                            displayText="Description"
                            field="description"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { description },
                },
            }) => <div>{description}</div>,
            enableMultiSort: true,
        },
        {
            id: 'cheque_id',
            accessorKey: 'cheque_id',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Cheque ID">
                    <ColumnActions {...props}>
                        <TextFilter<ITransactionPaymentTypesResource>
                            displayText="Cheque ID"
                            field="cheque_id"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { cheque_id },
                },
            }) => <div>{cheque_id}</div>,
            enableMultiSort: true,
        },
        {
            id: 'createdAt',
            accessorKey: 'createdAt',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="Date Created"
                >
                    <ColumnActions {...props}>
                        <DateFilter<ITransactionPaymentTypesResource>
                            displayText="Date Created"
                            field="createdAt"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { createdAt },
                },
            }) => <div>{toReadableDate(createdAt)}</div>,
            enableMultiSort: true,
        },
        {
            id: 'updatedAt',
            accessorKey: 'updatedAt',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="Last Updated"
                >
                    <ColumnActions {...props}>
                        <DateFilter<ITransactionPaymentTypesResource>
                            displayText="Last Updated"
                            field="updatedAt"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { updatedAt },
                },
            }) => <div>{toReadableDate(updatedAt)}</div>,
            enableMultiSort: true,
        },
    ]
}

export default TransactionPaymentTypesTableColumns
