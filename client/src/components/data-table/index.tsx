import {
    useSensor,
    useSensors,
    DndContext,
    MouseSensor,
    TouchSensor,
    DragEndEvent,
    closestCenter,
    KeyboardSensor,
} from '@dnd-kit/core'
import { Table } from '@tanstack/react-table'
import { arrayMove } from '@dnd-kit/sortable'
import { restrictToHorizontalAxis } from '@dnd-kit/modifiers'

import DataTableBody from './data-table-body'
import { Table as UITable } from '@/components/ui/table'
import DataTableHeaderGroup from './data-table-header-group'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'

interface Props<TData> extends IBaseCompNoChild {
    table: Table<TData>
    isStickyHeader?: boolean
    setColumnOrder?: React.Dispatch<React.SetStateAction<string[]>>
}

const DataTable = <TData,>({
    table,
    className,
    isStickyHeader,
    setColumnOrder,
}: Props<TData>) => {
    const handleDragEnd = (event: DragEndEvent) => {
        if (!setColumnOrder) return

        const { active, over } = event

        if (active && over && active.id !== over.id) {
            setColumnOrder((columnOrder) => {
                const oldIndex = columnOrder.indexOf(active.id as string)
                const newIndex = columnOrder.indexOf(over.id as string)
                return arrayMove(columnOrder, oldIndex, newIndex) //this is just a splice util
            })
        }
    }

    const sensors = useSensors(
        useSensor(MouseSensor, { }),
        useSensor(TouchSensor, {}),
        useSensor(KeyboardSensor, {})
    )

    return (
        <DndContext
            collisionDetection={closestCenter}
            modifiers={[restrictToHorizontalAxis]}
            onDragEnd={handleDragEnd}
            sensors={sensors}
        >
            <div
                className={cn(
                    'ecoop-scroll relative z-10 flex max-h-full overflow-y-scroll rounded-xl bg-secondary',
                    className
                )}
            >
                {table.getLeftHeaderGroups().length > 0 && (
                    <div className="ecoop-scroll sticky left-[0%] z-50 h-fit w-fit border-r border-popover shadow-xl">
                        <UITable
                            className="h-fit w-auto"
                            style={{
                                width:
                                    table.getLeftHeaderGroups().length > 0
                                        ? table.getLeftTotalSize()
                                        : 'fit-content',
                            }}
                        >
                            <DataTableHeaderGroup
                                isStickyHeader={isStickyHeader}
                                columnOrder={table.getState().columnOrder}
                                headerGroups={table.getLeftHeaderGroups()}
                            />
                            <DataTableBody
                                targetGroup="left"
                                rows={table.getRowModel().rows}
                            />
                        </UITable>
                    </div>
                )}
                {table.getCenterHeaderGroups().length > 0 && (
                    <div className="z-0 flex h-fit flex-1">
                        <UITable className="h-fit flex-1">
                            <DataTableHeaderGroup
                                isStickyHeader={isStickyHeader}
                                columnOrder={table.getState().columnOrder}
                                headerGroups={table.getCenterHeaderGroups()}
                            />
                            <DataTableBody rows={table.getRowModel().rows} />
                        </UITable>
                    </div>
                )}
                {table.getRightHeaderGroups().length > 0 && (
                    <div className="ecoop-scroll sticky right-0 z-10 h-fit w-fit border-l border-popover shadow-xl">
                        <UITable
                            className="w-auto"
                            style={{
                                width:
                                    table.getRightHeaderGroups().length > 0
                                        ? table.getRightTotalSize()
                                        : 'fit-content',
                            }}
                        >
                            <DataTableHeaderGroup
                                columnOrder={table.getState().columnOrder}
                                isStickyHeader={isStickyHeader}
                                headerGroups={table.getRightHeaderGroups()}
                            />
                            <DataTableBody
                                targetGroup="right"
                                rows={table.getRowModel().rows}
                            />
                        </UITable>
                    </div>
                )}
            </div>
        </DndContext>
    )
}

export default DataTable
