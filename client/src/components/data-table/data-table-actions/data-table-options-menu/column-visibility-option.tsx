import { Table } from '@tanstack/react-table'

import {
    DropdownMenuCheckboxItem,
    DropdownMenuGroup,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu'
import { EyeIcon } from '@/components/icons'

interface DataTableViewOptionsProps<TData> {
    table: Table<TData>
}

const ColumnVisibilityOption = <TData,>({
    table,
}: DataTableViewOptionsProps<TData>) => {
    return (
        <DropdownMenuGroup>
            <DropdownMenuLabel>Toggle columns</DropdownMenuLabel>
            <DropdownMenuSeparator />
            {table.getAllColumns().filter((col) => !col.getIsVisible()).length >
                0 && (
                <DropdownMenuItem
                    onClick={() =>
                        table
                            .getAllColumns()
                            .forEach((col) => col.toggleVisibility(true))
                    }
                >
                    <EyeIcon className="mr-2" />
                    Show All
                </DropdownMenuItem>
            )}
            {table
                .getAllColumns()
                .filter(
                    (column) =>
                        typeof column.accessorFn !== 'undefined' &&
                        column.getCanHide()
                )
                .map((column) => {
                    return (
                        <DropdownMenuCheckboxItem
                            key={column.id}
                            className="capitalize"
                            checked={column.getIsVisible()}
                            onCheckedChange={(value) =>
                                column.toggleVisibility(!!value)
                            }
                        >
                            {column.id}
                        </DropdownMenuCheckboxItem>
                    )
                })}
        </DropdownMenuGroup>
    )
}

export default ColumnVisibilityOption
