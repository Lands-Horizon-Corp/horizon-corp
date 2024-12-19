import { Header, Table } from '@tanstack/react-table'

import { DragHandleIcon } from '@/components/icons'

import { cn } from '@/lib'

interface Props<TData, TValue> {
    table: Table<TData>
    header: Header<TData, TValue>
}

const ColumnResizeHandle = <TData, TValue>({
    table,
    header,
}: Props<TData, TValue>) => {
    if (!header.column.getIsPinned()) return null

    return (
        <div
            className={cn(
                'absolute right-2 top-0 flex h-full w-[1px] items-center justify-center duration-100 ease-in-out hover:bg-border/40 hover:bg-border/80',
                header.column.getIsResizing() && 'z-50 bg-foreground/80'
            )}
            style={{
                transform:
                    table.options.columnResizeMode === 'onEnd' &&
                    header.column.getIsResizing()
                        ? `translateX(${
                              (table.options.columnResizeDirection === 'rtl'
                                  ? -1
                                  : 1) *
                              (table.getState().columnSizingInfo.deltaOffset ??
                                  0)
                          }px)`
                        : '',
            }}
        >
            <span
                className={cn(
                    'p-.5 cursor-col-resize touch-none select-none rounded-sm bg-secondary/30 text-foreground/60 hover:bg-secondary/80',
                    header.column.getIsResizing() &&
                        'bg-background text-foreground/80 hover:bg-background'
                )}
                onMouseDown={header.getResizeHandler()}
                onTouchStart={header.getResizeHandler()}
                onDoubleClick={() => header.column.resetSize()}
            >
                <DragHandleIcon className="size-3" />
            </span>
        </div>
    )
}

export default ColumnResizeHandle
