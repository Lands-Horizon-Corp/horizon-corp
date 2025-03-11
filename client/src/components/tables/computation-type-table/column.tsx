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
import { IAccountsComputationTypeResource } from '@/server/types/accounts/computation-type'

export const accountComputationTypeGlobalSearchTargets: IGlobalSearchTargets<IAccountsComputationTypeResource>[] =
    [
        { field: 'name', displayText: 'Name' },
        { field: 'description', displayText: 'Description' },
    ]

export interface IAccountsComputationTypeTableActionComponentProp {
    row: Row<IAccountsComputationTypeResource>
}

export interface IAccountsComputationTypeTableColumnProps {
    actionComponent?: (
        props: IAccountsComputationTypeTableActionComponentProp
    ) => ReactNode
}

const AccountsComputationTypeTableColumns = (
    opts?: IAccountsComputationTypeTableColumnProps
): ColumnDef<IAccountsComputationTypeResource>[] => {
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
            maxSize: 60,
            enableSorting: false,
            enableHiding: false,
            enablePinning: false,
            enableResizing: false,
        },
        {
            id: 'name',
            accessorKey: 'name',
            header: (props) => (
                <DataTableColumnHeader {...props} title="Name">
                    <ColumnActions {...props}>
                        <TextFilter<IAccountsComputationTypeResource>
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
            enablePinning: false,
            enableResizing: false,
        },
        {
            id: 'description',
            accessorKey: 'description',
            header: (props) => (
                <DataTableColumnHeader {...props} title="Description">
                    <ColumnActions {...props}>
                        <TextFilter<IAccountsComputationTypeResource>
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
            enablePinning: false,
            enableResizing: false,
        },
        {
            id: 'createdAt',
            accessorKey: 'createdAt',
            header: (props) => (
                <DataTableColumnHeader {...props} title="Date Created">
                    <ColumnActions {...props}>
                        <DateFilter<IAccountsComputationTypeResource>
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
            enablePinning: false,
            enableResizing: false,
        },
        {
            id: 'createdBy',
            accessorKey: 'createdBy',
            header: (props) => (
                <DataTableColumnHeader {...props} title="Created By" />
            ),
            cell: ({
                row: {
                    original: { createdBy },
                },
            }) => <div>{createdBy}</div>,
            enableMultiSort: true,
            enablePinning: false,
            enableResizing: false,
        },
        {
            id: 'updatedAt',
            accessorKey: 'updatedAt',
            header: (props) => (
                <DataTableColumnHeader {...props} title="Last Updated">
                    <ColumnActions {...props}>
                        <DateFilter<IAccountsComputationTypeResource>
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
            enablePinning: false,
            enableResizing: false,
        },
        {
            id: 'updatedBy',
            accessorKey: 'updatedBy',
            header: (props) => (
                <DataTableColumnHeader {...props} title="Updated By" />
            ),
            cell: ({
                row: {
                    original: { updatedBy },
                },
            }) => <div>{updatedBy}</div>,
            enableMultiSort: true,
            enablePinning: false,
            enableResizing: false,
        },
    ]
}

export default AccountsComputationTypeTableColumns
