import { ColumnDef } from '@tanstack/react-table'

import { Checkbox } from '@/components/ui/checkbox'
import { PushPinSlashIcon } from '@/components/icons'

import { Badge } from '@/components/ui/badge'
import UserAvatar from '@/components/user-avatar'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import DataTableMultiSelectFilter from '@/components/data-table/data-table-filters/multi-select-filter'
import { IGlobalSearchTargets } from '@/components/data-table/data-table-filters/data-table-global-search'

import { toReadableDate } from '@/utils'
import logger from '@/helpers/loggers/logger'
import { MemberResource } from '@/horizon-corp/types/profile'

export const membersGlobalSearchTargets: IGlobalSearchTargets<MemberResource>[] =
    [
        { field: 'username', displayText: 'Username' },
        { field: 'firstName', displayText: 'First Name' },
        { field: 'middleName', displayText: 'Middle Name' },
        { field: 'lastName', displayText: 'Last Name' },
        { field: 'email', displayText: 'Email' },
        { field: 'gender', displayText: 'Gender' },
        { field: 'permanentAddress', displayText: 'Address' },
        { field: 'contactNumber', displayText: 'Contact Number' },
    ]

const memberTableColumns: ColumnDef<MemberResource>[] = [
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
        id: 'Photo',
        accessorKey: 'media',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Photo">
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
        id: 'Username',
        accessorKey: 'username',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Username">
                <ColumnActions {...props}>
                    <TextFilter<MemberResource>
                        field="username"
                        displayText="Username"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { username },
            },
        }) => <div>{username}</div>,
    },
    {
        id: 'First Name',
        accessorKey: 'firstName',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="First Name">
                <ColumnActions {...props}>
                    <TextFilter<MemberResource>
                        field="firstName"
                        displayText="First Name"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { firstName },
            },
        }) => <div>{firstName}</div>,
    },
    {
        id: 'Middle Name',
        accessorKey: 'middleName',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Middle Name">
                <ColumnActions {...props}>
                    <TextFilter<MemberResource>
                        field="middleName"
                        displayText="Middle Name"
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
        id: 'Last Name',
        accessorKey: 'lastName',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Last Name">
                <ColumnActions {...props}>
                    <TextFilter<MemberResource>
                        field="lastName"
                        displayText="Last Name"
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
        id: 'Email',
        accessorKey: 'email',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Email">
                <ColumnActions {...props}>
                    <TextFilter<MemberResource>
                        field="email"
                        displayText="Email"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { email },
            },
        }) => <div>{email}</div>,
    },
    {
        id: 'Gender',
        accessorKey: 'gender',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Gender">
                <ColumnActions {...props}>
                    <DateFilter<MemberResource>
                        field="gender"
                        displayText="Gender"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { gender },
            },
        }) => <div>{gender?.name ?? '-'}</div>,
    },
    {
        id: 'Address',
        accessorKey: 'permanentAddress',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Address">
                <ColumnActions {...props}>
                    <TextFilter<MemberResource>
                        field="permanentAddress"
                        displayText="Address"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { permanentAddress },
            },
        }) => <div>{permanentAddress}</div>,
    },
    {
        id: 'Birth Date',
        accessorKey: 'birthDate',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Birth Date">
                <ColumnActions {...props}>
                    <DateFilter<MemberResource>
                        field="birthDate"
                        displayText="Birth Date"
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { birthDate },
            },
        }) => <div>{toReadableDate(birthDate)}</div>,
    },
    {
        id: 'Contact Number',
        accessorKey: 'contactNumber',
        header: (props) => (
            <DataTableColumnHeader
                {...props}
                isResizable
                title="Contact Number"
            >
                <ColumnActions {...props}>
                    <TextFilter<MemberResource>
                        field="contactNumber"
                        displayText="Contact"
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
        id: 'Email Status',
        accessorKey: 'isEmailVerified',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Email Status">
                <ColumnActions {...props}>
                    <DataTableMultiSelectFilter<MemberResource, boolean>
                        dataType="boolean"
                        field="isEmailVerified"
                        displayText="Email Status"
                        multiSelectOptions={[
                            { label: 'Verified', value: true },
                            { label: 'Not Verified', value: false },
                        ]}
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { isEmailVerified },
            },
        }) => (
            <div>
                {isEmailVerified ? (
                    <Badge variant="success">Verified</Badge>
                ) : (
                    <Badge variant="warning">Not Verified</Badge>
                )}
            </div>
        ),
    },
    {
        id: 'Contact Status',
        accessorKey: 'isContactVerified',
        header: (props) => (
            <DataTableColumnHeader
                {...props}
                isResizable
                title="Contact Status"
            >
                <ColumnActions {...props}>
                    <DataTableMultiSelectFilter<MemberResource>
                        field="isContactVerified"
                        displayText="Contact Status"
                        multiSelectOptions={[
                            { label: 'Verified', value: 'true' },
                            { label: 'Not Verified', value: 'false' },
                        ]}
                    />
                </ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { isContactVerified },
            },
        }) => (
            <div>
                {isContactVerified ? (
                    <Badge variant="success">Verified</Badge>
                ) : (
                    <Badge variant="warning">Not Verified</Badge>
                )}
            </div>
        ),
    },
    {
        id: 'Join Date',
        accessorKey: 'createdAt',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Join Date">
                <ColumnActions {...props}>
                    <DateFilter<MemberResource>
                        field="createdAt"
                        displayText="Last Update"
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
        id: 'Updated At',
        accessorKey: 'updatedAt',
        header: (props) => (
            <DataTableColumnHeader {...props} isResizable title="Updated At">
                <ColumnActions {...props}>
                    <DateFilter<MemberResource>
                        field="updatedAt"
                        displayText="Updated At"
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

export default memberTableColumns
