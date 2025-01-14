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
import { MemberResource } from '@/horizon-corp/types'

export const memberGlobalSearchTargets: IGlobalSearchTargets<MemberResource>[] =
    [
        { field: 'name', displayText: 'Name' },
        { field: 'address', displayText: 'Address' },
        { field: 'Owner.username', displayText: 'Owner' },
        { field: 'contactNumber', displayText: 'Contact' },
        { field: 'isAdminVerified', displayText: 'Verify Status' },
    ]

export interface ICompanyTableActionComponentProp {
    row: Row<MemberResource>
}

export interface ICompaniesTableColumnProps {
    actionComponent?: (props: ICompanyTableActionComponentProp) => ReactNode
}

const companiesTableColumns = (
    opts?: ICompaniesTableColumnProps
): ColumnDef<MemberResource>[] => {
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
            id: 'Name',
            accessorKey: 'name',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Name">
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
            }) => (
                <div onClick={(e) => e.stopPropagation()}>
                    {/* <Link
                        params={{ companyId: id }}
                        to="/admin/companies-management/$companyId/view"
                        className="hover:underline"
                    > */}
                    {firstName}
                    {/* </Link> */}
                </div>
            ),
            enableMultiSort: true,
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
                    <ImageDisplay src={media?.downloadURL} className="size-7" />
                </div>
            ),
            enableSorting: false,
        },
        {
            id: 'address',
            accessorKey: 'permanentAddress',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Address">
                    <ColumnActions {...props}>
                        <TextFilter
                            field="address"
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
            id: 'Verify Status',
            accessorKey: 'isAdminVerified',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Verified">
                    <ColumnActions {...props}>
                        <DataTableMultiSelectFilter<MemberResource, boolean>
                            mode="equal"
                            dataType="boolean"
                            field="isAdminVerified"
                            displayText="Verify Status"
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
            id: 'Created At',
            accessorKey: 'createdAt',
            header: (props) => (
                <DataTableColumnHeader
                    {...props}
                    isResizable
                    title="Date Created"
                >
                    <ColumnActions {...props}>
                        <DateFilter<MemberResource>
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

export default companiesTableColumns
