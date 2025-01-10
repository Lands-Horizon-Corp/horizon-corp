import DataTableColumnHeader from "@/components/data-table/data-table-column-header"
import ColumnActions from "@/components/data-table/data-table-column-header/column-actions"
import { IGlobalSearchTargets } from "@/components/data-table/data-table-filters/data-table-global-search"
import TextFilter from "@/components/data-table/data-table-filters/text-filter"
import { PushPinSlashIcon } from "@/components/icons"
import { Checkbox } from "@/components/ui/checkbox"
import { FeedbackResource } from "@/horizon-corp/types"
import { Link, ReactNode } from "@tanstack/react-router"
import { ColumnDef, Row } from "@tanstack/react-table"

export const FeedbackGlobalSearchTargets: IGlobalSearchTargets<FeedbackResource>[] =
[
    { field: 'email', displayText: 'email' },
    { field: 'feedbackType', displayText: 'feedbackType' },
    { field: 'description', displayText: 'description' },
]

export interface IFeedBackTableActionComponentProp {
    row: Row<FeedbackResource>
}

export interface IFeedbackTableColumnProps {
    actionComponent?: (props: IFeedBackTableActionComponentProp) => ReactNode 
}

const AdminCompaniesFeedbackTableColumns = (opts: IFeedbackTableColumnProps): ColumnDef<FeedbackResource>[] => {
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
            id: 'email',
            accessorKey: 'email',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Name">
                    <ColumnActions {...props}>
                        <TextFilter
                            field="email"
                            displayText="email"
                            defaultMode="contains"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { id, email },
                },
            }) => (
                <div onClick={(e) => e.stopPropagation()}>
                    <Link
                        params={{ companyId: id }}
                        to="/admin"
                        className="hover:underline"
                    >
                        {email}
                    </Link>
                </div>
            ),
        },
        {
            id: 'description',
            accessorKey: 'description',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Description">
                    <ColumnActions {...props} />
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { description },
                },
            }) => (
                <div>
                    {description}
                </div>
            ),
        },
        {
            id: 'feedbackType',
            accessorKey: 'feedbackType',
            header: (props) => (
                <DataTableColumnHeader {...props} isResizable title="Feedback Type">
                    <ColumnActions {...props}>
                        <TextFilter
                            field="feedbackType"
                            displayText="feedbackType"
                            defaultMode="contains"
                        />
                    </ColumnActions>
                </DataTableColumnHeader>
            ),
            cell: ({
                row: {
                    original: { feedbackType },
                },
            }) => <div>{feedbackType}</div>,
         },
        // {
        //     id: 'branches',
        //     accessorKey: 'branches',
        //     header: (props) => (
        //         <DataTableColumnHeader {...props} isResizable title="Branches">
        //             <ColumnActions {...props}>
        //                 <NumberFilter
        //                     displayText="# Branches"
        //                     field="branches"
        //                 />
        //             </ColumnActions>
        //         </DataTableColumnHeader>
        //     ),
        //     cell: ({
        //         row: {
        //             original: { branches },
        //         },
        //     }) => <div>{branches?.length ?? 0}</div>,
        // },
        // {
        //     id: 'owner',
        //     accessorKey: 'owner.username',
        //     header: (props) => (
        //         <DataTableColumnHeader {...props} isResizable title="Owner">
        //             <ColumnActions {...props}>
        //                 <TextFilter<CompanyResource>
        //                     displayText="Owner"
        //                     field="Owner.username"
        //                 />
        //             </ColumnActions>
        //         </DataTableColumnHeader>
        //     ),
        //     cell: ({
        //         row: {
        //             original: { owner },
        //         },
        //     }) => (
        //         <div>
        //             {owner?.username ?? (
        //                 <span className="text-center italic text-foreground/40">
        //                     -
        //                 </span>
        //             )}
        //         </div>
        //     ),
        // },
        // {
        //     id: 'Contact Number',
        //     accessorKey: 'contactNumber',
        //     header: (props) => (
        //         <DataTableColumnHeader
        //             {...props}
        //             isResizable
        //             title="Contact Number"
        //         >
        //             <ColumnActions {...props}>
        //                 <TextFilter<CompanyResource>
        //                     displayText="Contact"
        //                     field="contactNumber"
        //                 />
        //             </ColumnActions>
        //         </DataTableColumnHeader>
        //     ),
        //     cell: ({
        //         row: {
        //             original: { contactNumber },
        //         },
        //     }) => <div>{contactNumber}</div>,
        // },
        // {
        //     id: 'Verify Status',
        //     accessorKey: 'isAdminVerified',
        //     header: (props) => (
        //         <DataTableColumnHeader {...props} isResizable title="Verified">
        //             <ColumnActions {...props}>
        //                 <DataTableMultiSelectFilter<CompanyResource, boolean>
        //                     mode="equal"
        //                     dataType="boolean"
        //                     field="isAdminVerified"
        //                     displayText="Verify Status"
        //                     multiSelectOptions={[
        //                         {
        //                             label: 'Verified',
        //                             value: true,
        //                         },
        //                         {
        //                             label: 'Not Verified',
        //                             value: false,
        //                         },
        //                     ]}
        //                 />
        //             </ColumnActions>
        //         </DataTableColumnHeader>
        //     ),
        //     cell: ({
        //         row: {
        //             original: { isAdminVerified },
        //         },
        //     }) => (
        //         <div>
        //             {isAdminVerified ? (
        //                 <Badge
        //                     className="flex w-fit items-center gap-x-1"
        //                     variant="success"
        //                 >
        //                     <BadgeCheckFillIcon className="inline" /> Verified
        //                 </Badge>
        //             ) : (
        //                 <Badge
        //                     className="flex w-fit items-center gap-x-1"
        //                     variant="warning"
        //                 >
        //                     <BadgeQuestionIcon className="inline" /> Pending
        //                 </Badge>
        //             )}
        //         </div>
        //     ),
        // },
        // {
        //     id: 'Created At',
        //     accessorKey: 'createdAt',
        //     header: (props) => (
        //         <DataTableColumnHeader
        //             {...props}
        //             isResizable
        //             title="Date Created"
        //         >
        //             <ColumnActions {...props}>
        //                 <DateFilter<CompanyResource>
        //                     displayText="Date Created"
        //                     field="createdAt"
        //                 />
        //             </ColumnActions>
        //         </DataTableColumnHeader>
        //     ),
        //     cell: ({
        //         row: {
        //             original: { createdAt },
        //         },
        //     }) => <div>{toReadableDate(createdAt)}</div>,
        // },
    ]
}

export default AdminCompaniesFeedbackTableColumns