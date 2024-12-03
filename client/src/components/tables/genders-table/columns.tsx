import { ColumnDef } from '@tanstack/react-table'

import { Checkbox } from '@/components/ui/checkbox'
import { PushPinSlashIcon } from '@/components/icons'

import RowActionsGroup from '@/components/data-table/data-table-row-actions'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'

import { toReadableDate } from '@/utils'
import logger from '@/helpers/loggers/logger'
import { GenderResource } from '@/horizon-corp/types'
import { IGlobalSearchTargets } from '@/components/data-table/data-table-filters/data-table-global-search'

export const gendersGlobalSearchTargets: IGlobalSearchTargets<GenderResource>[] =
    []

export const genderTableColumns: ColumnDef<GenderResource>[] = [
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
                <RowActionsGroup
                    onDelete={{
                        text: 'Delete',
                        isAllowed: true,
                        onClick: () => {
                            logger.log('delete')
                        },
                    }}
                    onEdit={{
                        text: 'Edit',
                        isAllowed: false,
                        onClick: () => {
                            logger.log('edit')
                        },
                    }}
                    onView={{
                        text: 'View',
                        isAllowed: true,
                        onClick: () => {
                            logger.log('view')
                        },
                    }}
                />
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
            <DataTableColumnHeader {...props} isResizable title="Gender">
                <ColumnActions {...props}>
                    <TextFilter<GenderResource>
                        field="name"
                        displayText="Gender"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { name },
            },
        }) => <div>{name}</div>,
    },
    {
        id: 'description',
        accessorKey: 'description',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Description">
                <ColumnActions {...props}>
                    <TextFilter<GenderResource>
                        field="description"
                        displayText="Description"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { description },
            },
        }) => <div>{description}</div>,
    },
    {
        id: 'Created Date',
        accessorKey: 'createdAt',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Created">
                <ColumnActions {...props}>
                    <DateFilter<GenderResource>
                        field="createdAt"
                        displayText="Date Created"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { createdAt },
            },
        }) => <div>{toReadableDate(createdAt)}</div>,
    },
    {
        id: 'Updated Date',
        accessorKey: 'updatedAt',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Updated">
                <ColumnActions {...props}>
                    <DateFilter<GenderResource>
                        field="createdAt"
                        displayText="Last Update"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { updatedAt },
            },
        }) => <div>{toReadableDate(updatedAt)}</div>,
    },
]
