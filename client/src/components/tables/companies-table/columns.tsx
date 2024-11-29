import { ColumnDef } from '@tanstack/react-table'

import { Checkbox } from '@/components/ui/checkbox'
import { PushPinSlashIcon } from '@/components/icons'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'

import logger from '@/helpers/loggers/logger'
import DataTableColumnHeader from '@/components/data-table/data-table-column-header'
import ColumnActions from '@/components/data-table/data-table-column-header/column-actions'

export const companiesTableColumns: ColumnDef<{
    name: string
    address: string
    owner: string
}>[] = [
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
            <DataTableColumnHeader
                {...props}
                isResizable
                title="Name"
                dataType="text"
            >
                <ColumnActions {...props}></ColumnActions>
            </DataTableColumnHeader>
        ),
        cell: ({
            row: {
                original: { name },
            },
        }) => <div>{name}</div>,
    },
]
