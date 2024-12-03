import { ColumnDef } from '@tanstack/react-table'

import { Badge } from '@/components/ui/badge'
import { Checkbox } from '@/components/ui/checkbox'
import { PushPinSlashIcon } from '@/components/icons'

import UserAvatar from '@/components/user-avatar'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import NumberFilter from '@/components/data-table/data-table-filters/number-filter'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import DataTableMultiSelectFilter from '@/components/data-table/data-table-filters/multi-select-filter'

import { toReadableDate } from '@/utils'
import logger from '@/helpers/loggers/logger'
import { CompanyResource, GenderResource } from '@/horizon-corp/types'
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
            <DataTableColumnHeader {...props} isResizable title="Name">
                <ColumnActions {...props}>
                    <TextFilter displayText="Name" field="name" />
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
        id: 'Logo',
        accessorKey: 'media',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Logo">
                <ColumnActions {...props} />
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { media },
            },
        }) => (
            <div>
                <UserAvatar src={media?.downloadURL ?? ''} />
            </div>
        ),
    },
    {
        id: 'address',
        accessorKey: 'address',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Address">
                <ColumnActions {...props}>
                    <TextFilter displayText="Address" field="address" />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { address },
            },
        }) => <div>{address}</div>,
    },
    {
        id: 'branches',
        accessorKey: 'branches',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Branches">
                <ColumnActions {...props}>
                    <NumberFilter displayText="# Branches" field="branches" />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { branches },
            },
        }) => <div>{branches?.length ?? 0}</div>,
    },
    {
        id: 'owner',
        accessorKey: 'owner.username',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Owner">
                <ColumnActions {...props}>
                    <TextFilter<CompanyResource>
                        displayText="Owner"
                        field="Owner.username"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { owner },
            },
        }) => (
            <div>
                {owner?.username ?? (
                    <span className="text-center italic text-foreground/40">
                        -
                    </span>
                )}
            </div>
        ),
    },
    {
        id: 'contactNumber',
        accessorKey: 'contactNumber',
        header: (props) => (
            <DataTableColumnHeader
                {...props}
                isResizable
                title="Contact Number"
            >
                <ColumnActions {...props}>
                    <TextFilter<CompanyResource>
                        displayText="Contact"
                        field="contactNumber"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { contactNumber },
            },
        }) => <div>{contactNumber}</div>,
    },
    {
        id: 'isVerified',
        accessorKey: 'isAdminVerified',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Verified">
                <ColumnActions {...props}>
                    <DataTableMultiSelectFilter<CompanyResource>
                        displayText="Verify Status"
                        field="isAdminVerified"
                        multiSelectOptions={[
                            {
                                label: 'Verified',
                                value: 'true',
                            },
                            {
                                label: 'Not Verified',
                                value: 'false',
                            },
                        ]}
                    ></DataTableMultiSelectFilter>
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { isAdminVerified },
            },
        }) => (
            <div>
                {isAdminVerified ? (
                    <Badge variant="success">Verified</Badge>
                ) : (
                    <Badge variant="warning">Not Verified</Badge>
                )}
            </div>
        ),
    },
    {
        id: 'createdAt',
        accessorKey: 'createdAt',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Date Created">
                <ColumnActions {...props}>
                    <DateFilter<CompanyResource>
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
    },
]
