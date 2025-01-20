import { ColumnDef, Row } from "@tanstack/react-table"
import { Link, ReactNode } from "@tanstack/react-router"

import { Checkbox } from "@/components/ui/checkbox"
import { PushPinSlashIcon } from "@/components/icons"
import TextFilter from "@/components/data-table/data-table-filters/text-filter"
import DataTableColumnHeader from "@/components/data-table/data-table-column-header"
import ColumnActions from "@/components/data-table/data-table-column-header/column-actions"
import { IGlobalSearchTargets } from "@/components/data-table/data-table-filters/data-table-global-search"

import { FeedbackResource } from "@/horizon-corp/types"

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
    ]
}

export default AdminCompaniesFeedbackTableColumns