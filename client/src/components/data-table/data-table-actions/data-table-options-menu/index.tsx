import { Table } from '@tanstack/react-table'

import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuTrigger,
    DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import { MixerHorizontalIcon } from '@/components/icons'

import DataTableScrollOption, {
    IDataTableScrollableOptionProps,
} from './scroll-option'
import ColumnVisibilityOption from './column-visibility-option'

interface Props<T> {
    table: Table<T>
    scrollOption?: IDataTableScrollableOptionProps
}

const DataTableOptionsMenu = <T,>({ table, scrollOption }: Props<T>) => {
    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button size="icon" variant="secondary" className="rounded-md">
                    <MixerHorizontalIcon className="size-4" />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent
                align="end"
                className="ecoop-scroll max-h-[70vh] min-w-[180px] overflow-y-scroll"
            >
                <ColumnVisibilityOption table={table} />
                {scrollOption && (
                    <>
                        <DropdownMenuSeparator />
                        <DataTableScrollOption {...scrollOption} />
                    </>
                )}
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default DataTableOptionsMenu
