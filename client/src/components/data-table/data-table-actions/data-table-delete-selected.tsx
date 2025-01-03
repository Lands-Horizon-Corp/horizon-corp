import { Table } from '@tanstack/react-table'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { TrashBinIcon } from '@/components/icons'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib'
import useConfirmModalStore from '@/store/confirm-modal-store'

interface Props<T> {
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
    canDelete = false,
    onClick,
}: Props<T>) => {
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
            variant="secondary"
            className={cn(
                'relative',
                !(isLoading || isDisabled) &&
                    'text-rose-400/80 hover:text-rose-400'
            )}
        >
            {isLoading ? (
                <LoadingSpinner />
            ) : (
                <span className="inline-flex items-center gap-x-2">
                    <TrashBinIcon className="inline" />
                    {selectedRows.length > 0 && (
                        <Badge
                            variant="secondary"
                            className="absolute left-[-10%] top-0 w-fit -translate-x-1/2 -translate-y-1/2 text-xs animate-in fade-in dark:bg-popover/60"
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
