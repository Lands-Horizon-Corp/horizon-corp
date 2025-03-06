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

import { arrayMove } from '@dnd-kit/sortable'
import { restrictToHorizontalAxis } from '@dnd-kit/modifiers'
import { Table as TableInstance } from '@tanstack/react-table'

import { Table } from '../ui/table'
import DataTableBody from './data-table-body'
import DataTableHeader from './data-table-header'
import DataTableFooter from './data-table-footer'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'

interface Props<TData> extends IBaseCompNoChild {
    table: TableInstance<TData>
    rowClassName?: string
    isScrollable?: boolean
    isStickyHeader?: boolean
    isStickyFooter?: boolean
    setColumnOrder?: React.Dispatch<React.SetStateAction<string[]>>
}

const DataTable = <TData,>({
    table,
    className,
    isScrollable,
    isStickyHeader,
    isStickyFooter,
    setColumnOrder,
}: Props<TData>) => {
    const handleDragEnd = (event: DragEndEvent) => {
        if (!setColumnOrder) return

        const { active, over } = event

        if (active && over && active.id !== over.id) {
            setColumnOrder((columnOrder) => {
                const oldIndex = columnOrder.indexOf(active.id as string)
                const newIndex = columnOrder.indexOf(over.id as string)
                return arrayMove(columnOrder, oldIndex, newIndex)
            })
        }
    }

    const sensors = useSensors(
        useSensor(MouseSensor, {}),
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
            <Table
                wrapperClassName={cn(
                    'ecoop-scroll rounded-lg border',
                    className,
                    !isScrollable ? 'h-fit max-h-none min-h-fit' : 'h-full'
                )}
                className="table-fixed border-separate border-spacing-0 [&_td]:border-border [&_tfoot_td]:border-t [&_th]:border-b [&_th]:border-border [&_tr:not(:last-child)_td]:border-b [&_tr]:border-none"
                style={{
                    width: table.getTotalSize(),
                }}
            >
                <DataTableHeader
                    isStickyHeader={isStickyHeader}
                    columnOrder={table.getState().columnOrder}
                    headerGroups={table.getHeaderGroups()}
                />
                <DataTableBody rows={table.getRowModel().rows} />
                <DataTableFooter
                    table={table}
                    isStickyFooter={isStickyFooter && isScrollable}
                />
            </Table>
        </DndContext>
    )
}

export default DataTable
