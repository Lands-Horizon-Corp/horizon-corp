import { ReactNode } from 'react'
import { ColumnDef, Row } from '@tanstack/react-table'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import { IGlobalSearchTargets } from '@/components/data-table/data-table-filters/data-table-global-search'

import { toReadableDate } from '@/utils'
import { IAccountingLedgerResource } from '@/server/types/accounts/accounting-ledger'

export const accountingLedgerGlobalSearchTargets: IGlobalSearchTargets<IAccountingLedgerResource>[] =
    [
        { field: 'description', displayText: 'Description' },
        { field: 'or_number', displayText: 'OR Number' },
        { field: 'transaction_source', displayText: 'Transaction Source' },
        { field: 'debit', displayText: 'Debit' },
        { field: 'credit', displayText: 'Credit' },
        { field: 'balance', displayText: 'Balance' },
    ]

export interface IAccountingLedgerTableActionComponentProp {
    row: Row<IAccountingLedgerResource>
}

export interface IAccountingLedgerTableColumnProps {
    actionComponent?: (
        props: IAccountingLedgerTableActionComponentProp
    ) => ReactNode
}

const accountingLedgerTableColumns =
    (): ColumnDef<IAccountingLedgerResource>[] => {
        return [
            {
                id: 'description',
                accessorKey: 'description',
                header: (props) => (
                    <DataTableColumnHeader
                        {...props}
                        isResizable
                        title="Account title"
                    >
                        <ColumnActions {...props}>
                            <TextFilter
                                displayText="Account title"
                                field="description"
                            />
                        </ColumnActions>
                    </DataTableColumnHeader>
                ),
                cell: ({ row: { original } }) => (
                    <div>{original.description}</div>
                ),
                enableMultiSort: true,
            },
            {
                id: 'or_number',
                accessorKey: 'or_number',
                header: (props) => (
                    <DataTableColumnHeader
                        {...props}
                        isResizable
                        title="OR Number"
                    >
                        <ColumnActions {...props}>
                            <TextFilter
                                displayText="OR Number"
                                field="or_number"
                            />
                        </ColumnActions>
                    </DataTableColumnHeader>
                ),
                cell: ({ row: { original } }) => (
                    <div>{original.or_number}</div>
                ),
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
                            <DateFilter
                                displayText="Transaction Date"
                                field="transaction_date"
                            />
                        </ColumnActions>
                    </DataTableColumnHeader>
                ),
                cell: ({ row: { original } }) => (
                    <div>{toReadableDate(original.transaction_date)}</div>
                ),
                enableMultiSort: true,
            },
            {
                id: 'debit',
                accessorKey: 'debit',
                header: (props) => (
                    <DataTableColumnHeader {...props} isResizable title="Debit">
                        <ColumnActions {...props}>
                            <TextFilter displayText="Debit" field="debit" />
                        </ColumnActions>
                    </DataTableColumnHeader>
                ),
                cell: ({ row: { original } }) => <div>{original.debit}</div>,
                enableMultiSort: true,
            },
            {
                id: 'credit',
                accessorKey: 'credit',
                header: (props) => (
                    <DataTableColumnHeader
                        {...props}
                        isResizable
                        title="Credit"
                    >
                        <ColumnActions {...props}>
                            <TextFilter displayText="Credit" field="credit" />
                        </ColumnActions>
                    </DataTableColumnHeader>
                ),
                cell: ({ row: { original } }) => <div>{original.credit}</div>,
                enableMultiSort: true,
            },
            {
                id: 'balance',
                accessorKey: 'balance',
                header: (props) => (
                    <DataTableColumnHeader
                        {...props}
                        isResizable
                        title="Balance"
                    >
                        <ColumnActions {...props}>
                            <TextFilter displayText="Balance" field="balance" />
                        </ColumnActions>
                    </DataTableColumnHeader>
                ),
                cell: ({ row: { original } }) => (
                    <div>{original.summary.balance}</div>
                ),
                enableMultiSort: true,
            },
            {
                id: 'transaction_source',
                accessorKey: 'transaction_source',
                header: (props) => (
                    <DataTableColumnHeader
                        {...props}
                        isResizable
                        title="Transaction Source"
                    >
                        <ColumnActions {...props}>
                            <TextFilter
                                displayText="Transaction Source"
                                field="transaction_source"
                            />
                        </ColumnActions>
                    </DataTableColumnHeader>
                ),
                cell: ({ row: { original } }) => (
                    <div>{original.transaction_source}</div>
                ),
                enableMultiSort: true,
            },
            {
                id: 'created_at',
                accessorKey: 'created_at',
                header: (props) => (
                    <DataTableColumnHeader
                        {...props}
                        isResizable
                        title="Created At"
                    >
                        <ColumnActions {...props}>
                            <DateFilter
                                displayText="Created At"
                                field="created_at"
                            />
                        </ColumnActions>
                    </DataTableColumnHeader>
                ),
                cell: ({ row: { original } }) => (
                    <div>{toReadableDate(original.created_at)}</div>
                ),
                enableMultiSort: true,
            },
        ]
    }

export default accountingLedgerTableColumns
