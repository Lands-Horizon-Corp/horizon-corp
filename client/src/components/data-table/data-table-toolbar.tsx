import { Table } from '@tanstack/react-table'

import { Separator } from '@/components/ui/separator'
import DataTableRefreshButton, {
    IRefreshButtonProps,
} from '@/components/refresh-button'
import DataTableOptionsMenu from '@/components/data-table/data-table-actions/data-table-options-menu'
import DataTableExportButton from '@/components/data-table/data-table-actions/data-table-export-button'
import DataTableActiveFilters from '@/components/data-table/data-table-actions/data-table-active-filters'
import DataTableDeleteSelected from '@/components/data-table/data-table-actions/data-table-delete-selected'

import { IBaseCompNoChild } from '@/types'
import { IDataTableScrollableOptionProps } from './data-table-actions/data-table-options-menu/scroll-option'
import { type IDataTableExportProps } from '@/components/data-table/data-table-actions/data-table-export-button'
import { type IDataTableDeleteSelectedProps } from '@/components/data-table/data-table-actions/data-table-delete-selected'
import DataTableGlobalSearch, {
    IGlobalSearchProps,
} from './data-table-filters/data-table-global-search'

export interface IDataTableToolbarProps<TData = unknown>
    extends IBaseCompNoChild {
    table: Table<TData>
    refreshActionProps: IRefreshButtonProps
    globalSearchProps?: IGlobalSearchProps<TData>
    scrollableProps?: IDataTableScrollableOptionProps
    exportActionProps?: Omit<IDataTableExportProps<TData>, 'table'>
    deleteActionProps?: Omit<IDataTableDeleteSelectedProps<TData>, 'table'>
}

const DataTableToolbar = <TData,>({
    table,
    scrollableProps,
    globalSearchProps,
    deleteActionProps,
    exportActionProps,
    refreshActionProps,
}: IDataTableToolbarProps<TData>) => {
    return (
        <div className="flex w-full max-w-full items-center justify-between gap-x-2">
            <div className="flex items-center gap-x-2">
                {globalSearchProps ? (
                    <DataTableGlobalSearch {...globalSearchProps} />
                ) : null}
                <DataTableActiveFilters />
            </div>
            <div className="flex items-center gap-x-2">
                {deleteActionProps && (
                    <DataTableDeleteSelected
                        table={table}
                        {...deleteActionProps}
                    />
                )}
                <DataTableRefreshButton {...refreshActionProps} />
                {scrollableProps && (
                    <DataTableOptionsMenu
                        table={table}
                        scrollOption={scrollableProps}
                    />
                )}
                {exportActionProps && (
                    <>
                        <Separator
                            orientation="vertical"
                            className="h-full min-h-7"
                        />
                        <DataTableExportButton
                            table={table}
                            {...exportActionProps}
                        />
                    </>
                )}
            </div>
        </div>
    )
}

export default DataTableToolbar
