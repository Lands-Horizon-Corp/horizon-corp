import { Table } from '@tanstack/react-table'

import { Button } from '@/components/ui/button'
import ActionTooltip from '@/components/action-tooltip'
import { DashSquareDottedIcon } from '@/components/icons'

import { IBaseCompNoChild } from '@/types'
import useConfirmModalStore from '@/store/confirm-modal-store'

export interface IDataTableDeleteSelectedProps<T> extends IBaseCompNoChild {
    table: Table<T>
}

const DataTableUnselect = <T,>({ table, className }: IDataTableDeleteSelectedProps<T>) => {
    const { onOpen } = useConfirmModalStore();

    const selectedRows = table
        .getSelectedRowModel()
        .flatRows.map((row) => row.original)

    return (
        <ActionTooltip
            tooltipContent={`Unselect ${selectedRows.length} row(s)`}
        >
            <Button
                size="icon"
                disabled={selectedRows.length === 0}
                variant="secondary"
                className={className}
                onClick={() =>
                    onOpen({
                        title: 'Unselect',
                        description: `You are about to unselect ${selectedRows.length} items, are you sure you want to proceed?`,
                        onConfirm: () => table.resetRowSelection(),
                        confirmString: 'Unselect',
                    })
                }
            >
                <DashSquareDottedIcon />
            </Button>
        </ActionTooltip>
    )
}

export default DataTableUnselect
