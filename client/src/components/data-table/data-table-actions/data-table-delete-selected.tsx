import { Table } from '@tanstack/react-table'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { TrashIcon } from '@/components/icons'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import useConfirmModalStore from '@/store/confirm-modal-store'

export interface IDataTableDeleteSelectedProps<T> extends IBaseCompNoChild {
    table: Table<T>
    disabled?: boolean
    isLoading?: boolean
    canDelete?: boolean
    onClick: (selectedRows: T[]) => void
}

const DataTableDeleteSelected = <T,>({
    table,
    disabled,
    isLoading,
    className,
    canDelete = true,
    onClick,
}: IDataTableDeleteSelectedProps<T>) => {
    const { onOpen } = useConfirmModalStore()

    const selectedRows = table
        .getSelectedRowModel()
        .flatRows.map((row) => row.original)

    const isDisabled = !canDelete || selectedRows.length === 0 || disabled

    return (
        <Button
            disabled={isLoading || isDisabled}
            onClick={() =>
                onOpen({
                    title: 'Delete Selected',
                    description: `You are about to delete ${selectedRows.length} items, are you sure you want to proceed?`,
                    onConfirm: () => onClick(selectedRows),
                    confirmString: 'Proceed',
                })
            }
            size="icon"
            variant="destructive"
            className={cn('relative', className)}
        >
            {isLoading ? (
                <LoadingSpinner />
            ) : (
                <span className="inline-flex items-center gap-x-2">
                    <TrashIcon className="inline" />
                    {selectedRows.length > 0 && (
                        <Badge
                            variant="secondary"
                            className="absolute -right-[0%] top-0 size-fit w-fit -translate-y-1/2 translate-x-1/2 px-1.5 text-xs animate-in fade-in dark:bg-popover/60"
                        >
                            {selectedRows.length}
                        </Badge>
                    )}
                </span>
            )}
        </Button>
    )
}

export default DataTableDeleteSelected
