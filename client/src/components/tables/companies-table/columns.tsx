import { ColumnDef, Row } from '@tanstack/react-table'
import { useRouter, Link } from '@tanstack/react-router'

import {
    BadgeCheckFillIcon,
    BadgeCheckIcon,
    BadgeQuestionIcon,
    PushPinSlashIcon,
} from '@/components/icons'
import { Badge } from '@/components/ui/badge'
import { Checkbox } from '@/components/ui/checkbox'
import ImageDisplay from '@/components/image-display'
import { DropdownMenuItem } from '@/components/ui/dropdown-menu'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import NumberFilter from '@/components/data-table/data-table-filters/number-filter'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import DataTableMultiSelectFilter from '@/components/data-table/data-table-filters/multi-select-filter'
import { IGlobalSearchTargets } from '@/components/data-table/data-table-filters/data-table-global-search'

import { toReadableDate } from '@/utils'
import { CompanyResource } from '@/horizon-corp/types'
import useConfirmModalStore from '@/store/confirm-modal-store'

import {
    useApproveCompany,
    useDeleteCompany,
} from '@/hooks/api-hooks/use-company'

export const companyGlobalSearchTargets: IGlobalSearchTargets<CompanyResource>[] =
    [
        { field: 'name', displayText: 'Name' },
        { field: 'address', displayText: 'Address' },
        { field: 'Owner.username', displayText: 'Owner' },
        { field: 'contactNumber', displayText: 'Contact' },
        { field: 'isAdminVerified', displayText: 'Verify Status' },
    ]

interface ICompaniesTableActionProps {
    row: Row<CompanyResource>
    onDeleteSuccess?: () => void
    onCompanyUpdate?: () => void
}

const CompaniesTableAction = ({
    row,
    onDeleteSuccess,
    onCompanyUpdate,
}: ICompaniesTableActionProps) => {
    const company = row.original

    const router = useRouter()
    const { onOpen } = useConfirmModalStore()

    const { isPending: isDeletingCompany, mutate: deleteCompany } =
        useDeleteCompany({
            onSuccess: onDeleteSuccess,
        })

    const { mutate: approveCompany, isPending: isApprovingCompany } =
        useApproveCompany({
            onSuccess: onCompanyUpdate,
        })

    return (
        <RowActionsGroup
            onDelete={{
                text: 'Delete',
                isAllowed: !isDeletingCompany,
                onClick: () =>
                    onOpen({
                        title: 'Delete Company',
                        description: 'Are you sure to delete this company?',
                        onConfirm: () => deleteCompany(company.id),
                    }),
            }}
            onEdit={{
                text: 'Edit',
                isAllowed: true,
                onClick: () => {
                    router.navigate({
                        to: '/admin/companies-management/$companyId/edit',
                        params: { companyId: company.id },
                    })
                },
            }}
            onView={{
                text: 'View',
                isAllowed: true,
                onClick: () => {
                    router.navigate({
                        to: '/admin/companies-management/$companyId/view',
                        params: { companyId: company.id },
                    })
                },
            }}
            otherActions={
                <>
                    {!company.isAdminVerified && (
                        <DropdownMenuItem
                            onClick={() => {
                                onOpen({
                                    title: 'Approve Company',
                                    description:
                                        'Are you sure you want to approve this company? Approval will enable them to begin their operations.',
                                    onConfirm: () => approveCompany(company.id),
                                })
                            }}
                        >
                            {isApprovingCompany ? (
                                <LoadingSpinner />
                            ) : (
                                <BadgeCheckIcon className="mr-2" />
                            )}{' '}
                            Approve
                        </DropdownMenuItem>
                    )}
                </>
            }
        />
    )
}

const companiesTableColumns = (
    props: Omit<ICompaniesTableActionProps, 'row'>
): ColumnDef<CompanyResource>[] => {
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
                    <CompaniesTableAction row={row} {...props} />
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
                            field="name"
                            displayText="Name"
                            defaultMode="contains"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { id, name },
                },
            }) => (
                <div onClick={(e) => e.stopPropagation()}>
                    <Link
                        params={{ companyId: id }}
                        to="/admin/companies-management/$companyId/view"
                        className="hover:underline"
                    >
                        {name}
                    </Link>
                </div>
            ),
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
        },
        {
            id: 'address',
            accessorKey: 'address',
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
                        <NumberFilter
                            displayText="# Branches"
                            field="branches"
                        />
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
            id: 'Contact Number',
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
            id: 'Verify Status',
            accessorKey: 'isAdminVerified',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Verified">
                    <ColumnActions {...props}>
                        <DataTableMultiSelectFilter<CompanyResource, boolean>
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
                    original: { isAdminVerified },
                },
            }) => (
                <div>
                    {isAdminVerified ? (
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
                            <BadgeQuestionIcon className="inline" /> Pending
                        </Badge>
                    )}
                </div>
            ),
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
}

export default companiesTableColumns
