import { Column, Header, Table } from '@tanstack/react-table'

import {
    ArrowDownIcon,
    ArrowUpDownIcon,
    ArrowUpIcon,
    EyeNoneIcon,
    PushPinIcon,
    PushPinSlashIcon,
    ResetIcon,
} from '@/components/icons'
import {
    DropdownMenu,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuGroup,
    DropdownMenuTrigger,
    DropdownMenuContent,
    DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'

import { cn } from '@/lib/utils'

interface DataTableColumnHeaderProps<TData, TValue>
    extends React.HTMLAttributes<HTMLDivElement> {
    title: string
    table: Table<TData>
    isResizable?: boolean
    column: Column<TData, TValue>
    header: Header<TData, TValue>
}

export function DataTableColumnHeader<TData, TValue>({
    title,
    table,
    column,
    className,
    isResizable,
}: DataTableColumnHeaderProps<TData, TValue>) {
    return (
        <div
            className={cn(
                'flex w-fit items-center gap-x-2',
                className,
                isResizable && 'pr-1'
            )}
        >
            <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <Button
                        variant="ghost"
                        size="sm"
                        className="-ml-3 h-8 gap-x-2 data-[state=open]:bg-accent"
                    >
                        <span className="font-bold">{title}</span>
                        {column.getCanSort() && (
                            <>
                                {column.getIsSorted() === 'desc' ? (
                                    <ArrowUpIcon className="ml-2 size-3" />
                                ) : column.getIsSorted() === 'asc' ? (
                                    <ArrowDownIcon className="ml-2 size-3" />
                                ) : (
                                    <ArrowUpDownIcon className="ml-2 size-3" />
                                )}
                            </>
                        )}
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent
                    className="ecoop-scroll max-h-[60vh] min-w-40 overflow-y-scroll border-none shadow-xl"
                    align="start"
                >
                    {column.getCanPin() && (
                        <>
                            <DropdownMenuSeparator />
                            <DropdownMenuGroup>
                                <DropdownMenuLabel>
                                    Pin Column
                                </DropdownMenuLabel>
                                <DropdownMenuItem
                                    className="gap-x-2 px-2"
                                    disabled={column.getIsPinned() === 'left'}
                                    onClick={() => column.pin('left')}
                                >
                                    <PushPinIcon className="mr-2 size-3.5" />
                                    Pin Left
                                </DropdownMenuItem>
                                <DropdownMenuItem
                                    className="gap-x-2 px-2"
                                    disabled={column.getIsPinned() === 'right'}
                                    onClick={() => column.pin('right')}
                                >
                                    <PushPinIcon className="mr-2 size-3.5" />
                                    Pin Right
                                </DropdownMenuItem>
                                <DropdownMenuItem
                                    className="gap-x-2 px-2"
                                    onClick={() => column.pin(false)}
                                >
                                    <PushPinSlashIcon className="mr-2 size-3.5" />
                                    Unpin
                                </DropdownMenuItem>
                                <DropdownMenuItem
                                    className="gap-x-2 px-2"
                                    onClick={() => table.resetColumnPinning()}
                                >
                                    <PushPinSlashIcon className="mr-2 size-3.5" />
                                    Unpin All
                                </DropdownMenuItem>
                            </DropdownMenuGroup>
                        </>
                    )}
                    {column.getCanSort() && (
                        <>
                            <DropdownMenuSeparator />
                            <DropdownMenuGroup>
                                <DropdownMenuLabel>Sorting</DropdownMenuLabel>
                                <DropdownMenuItem
                                    className="gap-x-2 px-2"
                                    onClick={() => column.toggleSorting(false)}
                                >
                                    <ArrowUpIcon className="mr-2 size-3.5 text-muted-foreground/70" />
                                    Ascending
                                </DropdownMenuItem>
                                <DropdownMenuItem
                                    className="gap-x-2 px-2"
                                    onClick={() => column.toggleSorting(true)}
                                >
                                    <ArrowDownIcon className="mr-2 size-3.5 text-muted-foreground/70" />
                                    Descending
                                </DropdownMenuItem>
                                <DropdownMenuItem
                                    className="gap-x-2 px-2"
                                    onClick={() => column.clearSorting()}
                                >
                                    <ResetIcon className="mr-2 size-3.5 text-muted-foreground/70" />
                                    Reset
                                </DropdownMenuItem>
                            </DropdownMenuGroup>
                        </>
                    )}
                    {column.getCanHide() && (
                        <>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem
                                className="gap-x-2 px-2"
                                onClick={() => column.toggleVisibility(false)}
                            >
                                <EyeNoneIcon className="mr-2 h-3.5 w-3.5 text-muted-foreground/70" />
                                Hide
                            </DropdownMenuItem>
                        </>
                    )}
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    )
}
