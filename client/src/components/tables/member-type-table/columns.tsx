import { ReactNode } from 'react'
import { ColumnDef, Row } from '@tanstack/react-table'

import { PushPinSlashIcon } from '@/components/icons'
import { Checkbox } from '@/components/ui/checkbox'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import { IGlobalSearchTargets } from '@/components/data-table/data-table-filters/data-table-global-search'

import { toReadableDate } from '@/utils'
import { IMemberTypeResource } from '@/server/types'

export const memberTypeGlobalSearchTargets: IGlobalSearchTargets<IMemberTypeResource>[] =
    [
        { field: 'name', displayText: 'Name' },
        { field: 'description', displayText: 'Description' },
        { field: 'prefix', displayText: 'Prefix' },
    ]

export interface IMemberTypeTableActionComponentProp {
    row: Row<IMemberTypeResource>
}

export interface IMemberTypeTableColumnProps {
    actionComponent?: (props: IMemberTypeTableActionComponentProp) => ReactNode
}

const memberTypeTableColumns = (
    opts?: IMemberTypeTableColumnProps
): ColumnDef<IMemberTypeResource>[] => {
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
                <div className="flex w-fit max-w-[80] items-center gap-x-1 px-0">
                    {opts?.actionComponent?.({ row })}
                    <Checkbox
                        checked={row.getIsSelected()}
                        onCheckedChange={(value) => row.toggleSelected(!!value)}
                        aria-label="Select row"
                    />
                </div>
            ),
            maxSize: 40,
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
                        <TextFilter<IMemberTypeResource>
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
            enableResizing: false,
            minSize: 100,
        },
        {
            id: 'prefix',
            accessorKey: 'prefix',
            header: (props) => (
                <DataTableColumnHeader {...props} title="Prefix">
                    <ColumnActions {...props}>
                        <TextFilter<IMemberTypeResource>
                            displayText="Prefix"
                            field="prefix"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { prefix },
                },
            }) => <div>{prefix}</div>,
            enableMultiSort: true,
            enableResizing: false,
            minSize: 150,
        },
        {
            id: 'description',
            accessorKey: 'description',
            header: (props) => (
                <DataTableColumnHeader {...props} title="Description">
                    <ColumnActions {...props}>
                        <TextFilter<IMemberTypeResource>
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
            enableResizing: false,
            minSize: 150,
        },
        {
            id: 'createdAt',
            accessorKey: 'createdAt',
            header: (props) => (
                <DataTableColumnHeader {...props} title="Date Created">
                    <ColumnActions {...props}>
                        <DateFilter<IMemberTypeResource>
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
            enableResizing: false,
            minSize: 150,
        },
    ]
}

export default memberTypeTableColumns
