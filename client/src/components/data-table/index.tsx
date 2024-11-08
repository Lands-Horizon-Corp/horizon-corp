import { Table } from '@tanstack/react-table'

import DataTableBody from './data-table-body'
import { Table as UITable } from '@/components/ui/table'
import DataTableHeaderGroup from './data-table-header-group'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'

interface Props<TData> extends IBaseCompNoChild {
    table: Table<TData>
    isStickyHeader?: boolean
}

const DataTable = <TData,>({
    table,
    className,
    isStickyHeader,
}: Props<TData>) => {
    return (
        <div
            className={cn(
                'ecoop-scroll relative flex max-h-full overflow-y-scroll rounded-xl bg-secondary',
                className
            )}
        >
            {table.getLeftHeaderGroups().length > 0 && (
                <div className="ecoop-scroll sticky border-r h-fit border-popover shadow-xl left-[0%] z-50 w-fit">
                    <UITable className="h-fit !w-fit">
                        <DataTableHeaderGroup
                            isStickyHeader={isStickyHeader}
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
                            headerGroups={table.getCenterHeaderGroups()}
                        />
                        <DataTableBody rows={table.getRowModel().rows} />
                    </UITable>
                </div>
            )}
            {table.getRightHeaderGroups().length > 0 && (
                <div className="ecoop-scroll sticky right-0 h-fit z-10 w-fit border-l border-popover shadow-xl">
                    <UITable className="!w-fit">
                        <DataTableHeaderGroup
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
    )
}

export default DataTable
