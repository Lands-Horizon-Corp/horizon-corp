import { Table } from '@tanstack/react-table'

import {
    DropdownMenu,
    DropdownMenuCheckboxItem,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import { EyeIcon, MixerHorizontalIcon } from '@/components/icons'

interface DataTableViewOptionsProps<TData> {
    table: Table<TData>
}

export function DataTableViewOptions<TData>({
    table,
}: DataTableViewOptionsProps<TData>) {
    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button size="icon" variant="secondary" className="rounded-md">
                    <MixerHorizontalIcon className="size-4" />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-[150px]">
                <DropdownMenuLabel>Toggle columns</DropdownMenuLabel>
                <DropdownMenuSeparator />
                {table.getAllColumns().filter((col) => !col.getIsVisible())
                    .length > 0 && (
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
            </DropdownMenuContent>
        </DropdownMenu>
    )
}
