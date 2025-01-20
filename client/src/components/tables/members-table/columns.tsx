import { ReactNode } from 'react'
import { ColumnDef, Row } from '@tanstack/react-table'

import {
    PushPinSlashIcon,
    BadgeQuestionIcon,
    BadgeCheckFillIcon,
    BadgeExclamationIcon,
} from '@/components/icons'
import { Badge } from '@/components/ui/badge'
import { Checkbox } from '@/components/ui/checkbox'
import ImageDisplay from '@/components/image-display'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import DataTableMultiSelectFilter from '@/components/data-table/data-table-filters/multi-select-filter'
import { IGlobalSearchTargets } from '@/components/data-table/data-table-filters/data-table-global-search'

import { toReadableDate } from '@/utils'
import { TAccountStatus, IMemberResource } from '@/server/types'

export const memberGlobalSearchTargets: IGlobalSearchTargets<IMemberResource>[] =
    [
        { field: 'name', displayText: 'Name' },
        { field: 'address', displayText: 'Address' },
        { field: 'Owner.username', displayText: 'Owner' },
        { field: 'contactNumber', displayText: 'Contact' },
        { field: 'isAdminVerified', displayText: 'Verify Status' },
    ]

export interface IMemberTableActionComponentProp {
    row: Row<IMemberResource>
}

export interface IMembersTableColumnProps {
    actionComponent?: (props: IMemberTableActionComponentProp) => ReactNode
}

const membersTableColumns = (
    opts?: IMembersTableColumnProps
): ColumnDef<IMemberResource>[] => {
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
            id: 'Media',
            accessorKey: 'media',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Picture">
                    <ColumnActions {...props} />
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { media },
                },
            }) => (
                <div>
                    <ImageDisplay src={media?.downloadURL} className="size-7" />
                </div>
            ),
            enableSorting: false,
        },
        {
            id: 'Firstname',
            accessorKey: 'firstName',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="First Name"
                >
                    <ColumnActions {...props}>
                        <TextFilter
                            field="firstName"
                            displayText="Firstname"
                            defaultMode="contains"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { firstName },
                },
            }) => <div onClick={(e) => e.stopPropagation()}>{firstName}</div>,
            enableMultiSort: true,
        },
        {
            id: 'middleName',
            accessorKey: 'middleName',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="Middle Name"
                >
                    <ColumnActions {...props}>
                        <TextFilter
                            field="middleName"
                            displayText="Middle Name"
                            defaultMode="contains"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { middleName },
                },
            }) => <div onClick={(e) => e.stopPropagation()}>{middleName}</div>,
            enableMultiSort: true,
        },
        {
            id: 'lastName',
            accessorKey: 'lastName',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Last Name">
                    <ColumnActions {...props}>
                        <TextFilter
                            field="lastName"
                            displayText="Last Name"
                            defaultMode="contains"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { middleName },
                },
            }) => <div onClick={(e) => e.stopPropagation()}>{middleName}</div>,
            enableMultiSort: true,
        },
        {
            id: 'permanentAddress',
            accessorKey: 'permanentAddress',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Address">
                    <ColumnActions {...props}>
                        <TextFilter
                            field="permanentAddress"
                            displayText="Address"
                            defaultMode="contains"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { permanentAddress },
                },
            }) => <div>{permanentAddress}</div>,
            enableMultiSort: true,
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
                        <TextFilter<IMemberResource>
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
            enableMultiSort: true,
        },
        {
            id: 'status',
            accessorKey: 'status',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Status">
                    <ColumnActions {...props}>
                        <DataTableMultiSelectFilter<
                            IMemberResource,
                            TAccountStatus
                        >
                            mode="equal"
                            field="status"
                            dataType="text"
                            displayText="Status"
                            multiSelectOptions={[
                                {
                                    label: 'Pending',
                                    value: 'Pending',
                                },
                                {
                                    label: 'Verified',
                                    value: 'Verified',
                                },
                                {
                                    label: 'Not Allowed',
                                    value: 'Not Allowed',
                                },
                            ]}
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { status },
                },
            }) => (
                <div>
                    {status === 'Verified' && (
                        <Badge
                            className="flex w-fit items-center gap-x-1"
                            variant="success"
                        >
                            <BadgeCheckFillIcon className="inline" /> Verified
                        </Badge>
                    )}
                    {status === 'Pending' && (
                        <Badge
                            className="flex w-fit items-center gap-x-1"
                            variant="warning"
                        >
                            <BadgeQuestionIcon className="inline" /> Pending
                        </Badge>
                    )}
                    {status === 'Not Allowed' && (
                        <Badge
                            className="flex w-fit items-center gap-x-1"
                            variant="destructive"
                        >
                            <BadgeExclamationIcon className="inline" /> Not
                            Allowed
                        </Badge>
                    )}
                </div>
            ),
            enableMultiSort: true,
        },
        {
            id: 'isEmailVerified',
            accessorKey: 'isEmailVerified',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="Email Verified"
                >
                    <ColumnActions {...props}>
                        <DataTableMultiSelectFilter<IMemberResource, boolean>
                            mode="equal"
                            dataType="boolean"
                            field="isEmailVerified"
                            displayText="Email Verified"
                            multiSelectOptions={[
                                {
                                    label: 'Verified',
                                    value: true,
                                },
                                {
                                    label: 'Not Verified',
                                    value: false,
                                },
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
                        <Badge
                            className="flex w-fit items-center gap-x-1"
                            variant="success"
                        >
                            <BadgeCheckFillIcon className="inline" /> Verified
                        </Badge>
                    ) : (
                        <Badge
                            className="flex w-fit items-center gap-x-1"
                            variant="warning"
                        >
                            <BadgeQuestionIcon className="inline" /> Not
                            Verified
                        </Badge>
                    )}
                </div>
            ),
            enableMultiSort: true,
        },
        {
            id: 'isContactVerified',
            accessorKey: 'isContactVerified',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="Contact Verified"
                >
                    <ColumnActions {...props}>
                        <DataTableMultiSelectFilter<IMemberResource, boolean>
                            mode="equal"
                            dataType="boolean"
                            field="isContactVerified"
                            displayText="Contact Verified"
                            multiSelectOptions={[
                                {
                                    label: 'Verified',
                                    value: true,
                                },
                                {
                                    label: 'Not Verified',
                                    value: false,
                                },
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
                        <Badge
                            className="flex w-fit items-center gap-x-1"
                            variant="success"
                        >
                            <BadgeCheckFillIcon className="inline" />
                        </Badge>
                    ) : (
                        <Badge
                            className="flex w-fit items-center gap-x-1"
                            variant="warning"
                        >
                            <BadgeQuestionIcon className="inline" /> Not
                            Verified
                        </Badge>
                    )}
                </div>
            ),
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
                        <DateFilter<IMemberResource>
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
    ]
}

export default membersTableColumns
