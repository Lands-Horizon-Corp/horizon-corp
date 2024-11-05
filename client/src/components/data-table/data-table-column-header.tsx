import { Column } from '@tanstack/react-table'

import {
    ArrowDownIcon,
    ArrowUpDownIcon,
    ArrowUpIcon,
    EyeNoneIcon,
} from '@/components/icons'
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'

import { cn } from '@/lib/utils'

interface DataTableColumnHeaderProps<TData, TValue>
    extends React.HTMLAttributes<HTMLDivElement> {
    column: Column<TData, TValue>
    title: string
}

export function DataTableColumnHeader<TData, TValue>({
    title,
    column,
    className,
}: DataTableColumnHeaderProps<TData, TValue>) {
    if (!column.getCanSort()) {
        return <div className={cn(className)}>{title}</div>
    }

    return (
        <div className={cn('flex items-center space-x-2', className)}>
            <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <Button
                        variant="ghost"
                        size="sm"
                        className="-ml-3 h-8 data-[state=open]:bg-accent"
                    >
                        <span className="font-bold">{title}</span>
                        {column.getIsSorted() === 'desc' ? (
                            <ArrowUpIcon className="ml-2 size-4" />
                        ) : column.getIsSorted() === 'asc' ? (
                            <ArrowDownIcon className="ml-2 size-4" />
                        ) : (
                            <ArrowUpDownIcon className="ml-2 size-4" />
                        )}
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent
                    className="border-none shadow-xl"
                    align="start"
                >
                    <DropdownMenuItem
                        className="gap-x-2 px-2"
                        onClick={() => column.toggleSorting(false)}
                    >
                        <ArrowUpIcon className="mr-2 size-3.5 text-muted-foreground/70" />
                        Asc
                    </DropdownMenuItem>
                    <DropdownMenuItem
                        className="gap-x-2 px-2"
                        onClick={() => column.toggleSorting(true)}
                    >
                        <ArrowDownIcon className="mr-2 size-3.5 text-muted-foreground/70" />
                        Desc
                    </DropdownMenuItem>
                    <DropdownMenuSeparator />
                    <DropdownMenuItem
                        className="gap-x-2 px-2"
                        onClick={() => column.toggleVisibility(false)}
                    >
                        <EyeNoneIcon className="mr-2 h-3.5 w-3.5 text-muted-foreground/70" />
                        Hide
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    )
}
