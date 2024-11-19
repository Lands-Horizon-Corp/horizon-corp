import { Table } from '@tanstack/react-table'

import { TrashIcon } from '@/components/icons'
import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import useConfirmModalStore from '@/store/confirm-modal-store'
import { Badge } from '@/components/ui/badge'

interface Props<T> {
    table: Table<T>
    isLoading?: boolean
    canDelete?: boolean
    onClick: (selectedRows: T[]) => void
}

const DataTableDeleteSelected = <T,>({
    table,
    isLoading,
    onClick,
    canDelete = false,
}: Props<T>) => {
    const { onOpen } = useConfirmModalStore()

    const selectedRows = table
        .getSelectedRowModel()
        .flatRows.map((row) => row.original)

    return (
        <Button
            disabled={isLoading || !canDelete || selectedRows.length === 0}
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
            className="relative"
        >
            {isLoading ? (
                <LoadingSpinner />
            ) : (
                <span className="inline-flex items-center gap-x-2">
                    <TrashIcon className="inline" />
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
