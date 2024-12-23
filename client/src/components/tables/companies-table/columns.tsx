import { toast } from 'sonner'
import { useRouter } from '@tanstack/react-router'
import { useMutation } from '@tanstack/react-query'
import { ColumnDef, Row } from '@tanstack/react-table'

import { Badge } from '@/components/ui/badge'
import UserAvatar from '@/components/user-avatar'
import { Checkbox } from '@/components/ui/checkbox'
import { DropdownMenuItem } from '@/components/ui/dropdown-menu'
import {
    BadgeCheckFillIcon,
    BadgeCheckIcon,
    BadgeQuestionIcon,
    PushPinSlashIcon,
} from '@/components/icons'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'
import TextFilter from '@/components/data-table/data-table-filters/text-filter'
import DateFilter from '@/components/data-table/data-table-filters/date-filter'
import NumberFilter from '@/components/data-table/data-table-filters/number-filter'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'
import DataTableMultiSelectFilter from '@/components/data-table/data-table-filters/multi-select-filter'
import { IGlobalSearchTargets } from '@/components/data-table/data-table-filters/data-table-global-search'

import { serverRequestErrExtractor } from '@/helpers'
import { CompanyResource } from '@/horizon-corp/types'
import { toReadableDate, withCatchAsync } from '@/utils'
import useConfirmModalStore from '@/store/confirm-modal-store'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'

export const companyGlobalSearchTargets: IGlobalSearchTargets<CompanyResource>[] =
    [
        { field: 'name', displayText: 'Name' },
        { field: 'address', displayText: 'Address' },
        { field: 'Owner.username', displayText: 'Owner' },
        { field: 'contactNumber', displayText: 'Contact' },
        { field: 'isAdminVerified', displayText: 'Verify Status' },
    ]

const CompaniesTableAction = ({
    row,
    onDeleteSuccess,
}: {
    row: Row<CompanyResource>
    onDeleteSuccess?: () => void
}) => {
    const company = row.original

    const router = useRouter()
    const { onOpen } = useConfirmModalStore()

    const { isPending: isDeletingCompany, mutate: deleteCompany } = useMutation<
        void,
        string
    >({
        mutationKey: ['delete', 'company', company.id],
        mutationFn: async () => {
            const [error] = await withCatchAsync(
                CompanyService.delete(company.id)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            toast.success('Company deleted')
            onDeleteSuccess?.()
        },
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
                        onConfirm: () => deleteCompany(),
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
                        <DropdownMenuItem>
                            <BadgeCheckIcon className="mr-2" /> Approve
                        </DropdownMenuItem>
                    )}
                </>
            }
        />
    )
}

const companiesTableColumns = ({
    onDeleteSuccess,
}: {
    onDeleteSuccess?: () => void
}): ColumnDef<CompanyResource>[] => {
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
                    <CompaniesTableAction
                        row={row}
                        onDeleteSuccess={onDeleteSuccess}
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
