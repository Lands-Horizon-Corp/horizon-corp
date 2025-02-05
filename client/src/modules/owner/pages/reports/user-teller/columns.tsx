
// config/tellerReportConfig.ts
import { TellerTransactionRecord } from '../../../../../components/reports/types/user-teller';
import { ReactNode } from 'react'
import { ColumnDef, Row } from '@tanstack/react-table'

import {
    PushPinSlashIcon,
} from '@/components/icons'
import { Checkbox } from '@/components/ui/checkbox'

import { IGlobalSearchTargets } from '@/components/data-table/data-table-filters/data-table-global-search'

export const userTellerGlobalSearchTargets: IGlobalSearchTargets<TellerTransactionRecord>[] = [
    { field: 'date', displayText: 'Date' },
    { field: 'orNumber', displayText: 'O.R. No.' },
    { field: 'pbNumber', displayText: 'P.B. No.' },
    { field: 'memberName', displayText: "Member's Name" },
    { field: 'cashOnHand', displayText: 'Cash on Hand' },
    { field: 'savings', displayText: 'Savings' },
    { field: 'timeDeposit', displayText: 'Time Deposit' },
    { field: 'shareCapital', displayText: 'Share Capital' },
    { field: 'regularLoan', displayText: 'Regular Loan' },
    { field: 'salaryLoan', displayText: 'Salary Loan' },
    { field: 'insuranceCash', displayText: 'Insurance Cash' },
    { field: 'pcl', displayText: 'PCL' },
    { field: 'interestOnLoans', displayText: 'Interest On Loans' },
    { field: 'finesPenalties', displayText: 'Fines/Penalties' },
    { field: 'sundries', displayText: 'Sundries' },
];

export interface userReportsTableActionComponentProp {
    row: Row<TellerTransactionRecord>
}

export interface UserReportsTableColumnProps {
    actionComponent?: (props: userReportsTableActionComponentProp) => ReactNode
}

export const userReportsTableColumns = (
    opts?: UserReportsTableColumnProps
): ColumnDef<TellerTransactionRecord>[] => {
    return [
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
            accessorKey: 'date',
            header: 'Date',
            cell: ({ row }) => row.original.date,
        },
        {
            accessorKey: 'orNumber',
            header: 'O.R. No.',
            cell: ({ row }) => row.original.orNumber,
        },
        {
            accessorKey: 'pbNumber',
            header: 'P.B. No.',
            cell: ({ row }) => row.original.pbNumber,
        },
        {
            accessorKey: 'memberName',
            header: "Member's Name",
            cell: ({ row }) => row.original.memberName,
        },
        {
            accessorKey: 'cashOnHand',
            header: 'Cash on Hand',
            cell: ({ row }) => row.original.cashOnHand,
        },
        {
            accessorKey: 'savings',
            header: 'Savings',
            cell: ({ row }) => row.original.savings,
        },
        {
            accessorKey: 'timeDeposit',
            header: 'Time Dep.',
            cell: ({ row }) => row.original.timeDeposit,
        },
        {
            accessorKey: 'shareCapital',
            header: 'ShareCap',
            cell: ({ row }) => row.original.shareCapital,
        },
        {
            accessorKey: 'regularLoan',
            header: 'Reg. Loan',
            cell: ({ row }) => row.original.regularLoan,
        },
        {
            accessorKey: 'salaryLoan',
            header: 'Sal. Loan',
            cell: ({ row }) => row.original.salaryLoan,
        },
        {
            accessorKey: 'insuranceCash',
            header: 'Ins. Cash',
            cell: ({ row }) => row.original.insuranceCash,
        },
        {
            accessorKey: 'pcl',
            header: 'PCL',
            cell: ({ row }) => row.original.pcl,
        },
        {
            accessorKey: 'interestOnLoans',
            header: 'Int. On Loans',
            cell: ({ row }) => row.original.interestOnLoans,
        },
        {
            accessorKey: 'finesPenalties',
            header: 'Fines/Pen.',
            cell: ({ row }) => row.original.finesPenalties,
        },
        {
            accessorKey: 'sundries',
            header: 'Sundries',
            cell: ({ row }) => row.original.sundries,
        },
    ];
};
