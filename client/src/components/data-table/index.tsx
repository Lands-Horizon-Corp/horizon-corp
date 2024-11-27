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
import { memo, useCallback, useEffect, useRef } from 'react'
import { restrictToHorizontalAxis } from '@dnd-kit/modifiers'

import DataTableBody from './data-table-body'
import { Table as UITable } from '@/components/ui/table'
import DataTableHeaderGroup from './data-table-header-group'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import DataTableFooter from './data-table-footer'

interface Props<TData> extends IBaseCompNoChild {
    table: Table<TData>
    rowClassName?: string
    isScrollable?: boolean
    isStickyHeader?: boolean
    isStickyFooter?: boolean
    setColumnOrder?: React.Dispatch<React.SetStateAction<string[]>>
}

const DataTable = <TData,>({
    table,
    className,
    rowClassName,
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
                return arrayMove(columnOrder, oldIndex, newIndex) //this is just a splice util
            })
        }
    }

    const leftTableRef = useRef<HTMLTableElement>(null)
    const centerTableRef = useRef<HTMLTableElement>(null)
    const rightTableRef = useRef<HTMLTableElement>(null)

    const syncRowHeights = useCallback(() => {
        const leftRows = leftTableRef.current?.querySelectorAll('[data-row-id]')
        const centerRows =
            centerTableRef.current?.querySelectorAll('[data-row-id]')
        const rightRows =
            rightTableRef.current?.querySelectorAll('[data-row-id]')

        const rowIds = new Set<string>()
        if (leftRows)
            leftRows.forEach((row) =>
                rowIds.add((row as HTMLElement).dataset.rowId!)
            )
        if (centerRows)
            centerRows.forEach((row) =>
                rowIds.add((row as HTMLElement).dataset.rowId!)
            )
        if (rightRows)
            rightRows.forEach((row) =>
                rowIds.add((row as HTMLElement).dataset.rowId!)
            )

        rowIds.forEach((rowId) => {
            const rows = [
                leftTableRef.current?.querySelector(`[data-row-id="${rowId}"]`),
                centerTableRef.current?.querySelector(
                    `[data-row-id="${rowId}"]`
                ),
                rightTableRef.current?.querySelector(
                    `[data-row-id="${rowId}"]`
                ),
            ].filter(Boolean) as HTMLElement[]

            let maxHeight = 0
            rows.forEach((row) => {
                row.style.height = 'auto'
                maxHeight = Math.max(maxHeight, row.offsetHeight)
            })

            rows.forEach((row) => {
                row.style.height = `${maxHeight}px`
            })
        })
    }, [])

    const syncFooterRowHeights = useCallback(() => {
        const leftFooterRows = leftTableRef.current?.querySelectorAll(
            '[data-footer-row-id]'
        )
        const centerFooterRows = centerTableRef.current?.querySelectorAll(
            '[data-footer-row-id]'
        )
        const rightFooterRows = rightTableRef.current?.querySelectorAll(
            '[data-footer-row-id]'
        )

        const footerRowIds = new Set<string>()
        if (leftFooterRows)
            leftFooterRows.forEach((row) =>
                footerRowIds.add((row as HTMLElement).dataset.footerRowId!)
            )
        if (centerFooterRows)
            centerFooterRows.forEach((row) =>
                footerRowIds.add((row as HTMLElement).dataset.footerRowId!)
            )
        if (rightFooterRows)
            rightFooterRows.forEach((row) =>
                footerRowIds.add((row as HTMLElement).dataset.footerRowId!)
            )

        footerRowIds.forEach((footerRowId) => {
            const footerRows = [
                leftTableRef.current?.querySelector(
                    `[data-footer-row-id="${footerRowId}"]`
                ),
                centerTableRef.current?.querySelector(
                    `[data-footer-row-id="${footerRowId}"]`
                ),
                rightTableRef.current?.querySelector(
                    `[data-footer-row-id="${footerRowId}"]`
                ),
            ].filter(Boolean) as HTMLElement[]

            let maxHeight = 0
            footerRows.forEach((row) => {
                row.style.height = 'auto'
                maxHeight = Math.max(maxHeight, row.offsetHeight)
            })

            footerRows.forEach((row) => {
                row.style.height = `${maxHeight}px`
            })
        })
    }, [])

    const syncHeights = useCallback(() => {
        syncRowHeights()
        syncFooterRowHeights()
    }, [syncRowHeights, syncFooterRowHeights])

    useEffect(() => {
        syncHeights()

        const handleResize = () => syncHeights()
        window.addEventListener('resize', handleResize)

        return () => {
            window.removeEventListener('resize', handleResize)
        }
    }, [syncHeights])

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
            <div
                className={cn(
                    'relative z-10 flex rounded-xl border dark:bg-secondary',
                    isScrollable
                        ? 'ecoop-scroll max-h-full overflow-y-scroll'
                        : 'h-fit max-h-none',
                    className
                )}
            >
                {table.getLeftHeaderGroups().length > 0 && (
                    <div className="ecoop-scroll sticky left-[0%] z-50 h-fit w-fit border-r shadow dark:border-popover/80">
                        <UITable
                            ref={leftTableRef}
                            className="h-fit w-auto"
                            style={{
                                width:
                                    table.getLeftHeaderGroups().length > 0
                                        ? table.getLeftTotalSize()
                                        : 'fit-content',
                            }}
                        >
                            <DataTableHeaderGroup
                                isStickyHeader={isStickyHeader && isScrollable}
                                columnOrder={table.getState().columnOrder}
                                headerGroups={table.getLeftHeaderGroups()}
                            />
                            <DataTableBody
                                targetGroup="left"
                                rowClassName={rowClassName}
                                rows={table.getRowModel().rows}
                            />
                            <DataTableFooter
                                table={table}
                                isStickyFooter={isStickyFooter && isScrollable}
                                targetGroup="left"
                            />
                        </UITable>
                    </div>
                )}
                {table.getCenterHeaderGroups().length > 0 && (
                    <div className="z-0 flex h-fit flex-1">
                        <UITable ref={centerTableRef} className="h-fit flex-1">
                            <DataTableHeaderGroup
                                isStickyHeader={isStickyHeader && isScrollable}
                                columnOrder={table.getState().columnOrder}
                                headerGroups={table.getCenterHeaderGroups()}
                            />
                            <DataTableBody
                                rowClassName={rowClassName}
                                rows={table.getRowModel().rows}
                            />
                            <DataTableFooter
                                table={table}
                                isStickyFooter={isStickyFooter && isScrollable}
                            />
                        </UITable>
                    </div>
                )}
                {table.getRightHeaderGroups().length > 0 && (
                    <div className="ecoop-scroll sticky right-0 z-10 h-fit w-fit border-l border-popover shadow-xl">
                        <UITable
                            ref={rightTableRef}
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
                                isStickyHeader={isStickyHeader && isScrollable}
                                headerGroups={table.getRightHeaderGroups()}
                            />
                            <DataTableBody
                                targetGroup="right"
                                rowClassName={rowClassName}
                                rows={table.getRowModel().rows}
                            />
                            <DataTableFooter
                                table={table}
                                targetGroup="right"
                                isStickyFooter={isStickyFooter && isScrollable}
                            />
                        </UITable>
                    </div>
                )}
            </div>
        </DndContext>
    )
}

export default DataTable
