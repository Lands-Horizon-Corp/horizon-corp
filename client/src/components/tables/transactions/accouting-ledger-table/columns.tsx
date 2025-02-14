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
import { IAccountingLedgerRequest } from '@/server/types/accounts/accounting-ledger'

export const IAccountingLedgerRequestGlobalSearchTargets: IGlobalSearchTargets<IAccountingLedgerRequest>[] =
    [
        { field: 'description', displayText: 'Description' },
        { field: 'or_number', displayText: 'OR Number' },
        { field: 'transaction_source', displayText: 'Transaction Source' },
    ]

export interface IIAccountingLedgerRequestTableActionComponentProp {
    row: Row<IAccountingLedgerRequest>
}

export interface IIAccountingLedgerRequestTableColumnProps {
    actionComponent?: (
        props: IIAccountingLedgerRequestTableActionComponentProp
    ) => ReactNode
}

const IAccountingLedgerRequestTableColumns = (
    opts?: IIAccountingLedgerRequestTableColumnProps
): ColumnDef<IAccountingLedgerRequest>[] => {
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
            id: 'description',
            accessorKey: 'description',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="Description"
                >
                    <ColumnActions {...props}>
                        <TextFilter<IAccountingLedgerRequest>
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
            id: 'or_number',
            accessorKey: 'or_number',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="OR Number">
                    <ColumnActions {...props}>
                        <TextFilter<IAccountingLedgerRequest>
                            displayText="OR Number"
                            field="or_number"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { or_number },
                },
            }) => <div>{or_number}</div>,
            enableMultiSort: true,
        },
        {
            id: 'transaction_date',
            accessorKey: 'transaction_date',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="Transaction Date"
                >
                    <ColumnActions {...props}>
                        <DateFilter<IAccountingLedgerRequest>
                            displayText="Transaction Date"
                            field="transaction_date"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { transaction_date },
                },
            }) => <div>{toReadableDate(transaction_date)}</div>,
            enableMultiSort: true,
        },
    ]
}

export default IAccountingLedgerRequestTableColumns
