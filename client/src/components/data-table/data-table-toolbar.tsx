import { Table } from '@tanstack/react-table'

import DataTableRefreshButton, {
    IRefreshButtonProps,
} from '@/components/refresh-button'
import DataTableGlobalSearch, {
    IGlobalSearchProps,
} from './data-table-filters/data-table-global-search'
import { Separator } from '@/components/ui/separator'
import DataTableUnselect from './data-table-actions/data-table-unselect'
import DataTableExportButton from '@/components/data-table/data-table-actions/data-table-export'
import DataTableOptionsMenu from '@/components/data-table/data-table-actions/data-table-options-menu'
import DataTableActiveFilters from '@/components/data-table/data-table-actions/data-table-active-filters'
import DataTableDeleteSelected from '@/components/data-table/data-table-actions/data-table-delete-selected'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { type IDataTableExportProps } from '@/components/data-table/data-table-actions/data-table-export'
import { type IDataTableScrollableOptionProps } from './data-table-actions/data-table-options-menu/scroll-option'
import { IDataTableFilterLogicOptionProps } from './data-table-actions/data-table-options-menu/filter-logic-option'
import { type IDataTableDeleteSelectedProps } from '@/components/data-table/data-table-actions/data-table-delete-selected'

export interface IDataTableToolbarProps<TData = unknown>
    extends IBaseCompNoChild {
    table: Table<TData>
    refreshActionProps: IRefreshButtonProps
    globalSearchProps?: IGlobalSearchProps<TData>
    scrollableProps?: IDataTableScrollableOptionProps
    filterLogicProps?: IDataTableFilterLogicOptionProps
    exportActionProps?: Omit<IDataTableExportProps<TData>, 'table'>
    deleteActionProps?: Omit<IDataTableDeleteSelectedProps<TData>, 'table'>
}

const DataTableToolbar = <TData,>({
    table,
    scrollableProps,
    filterLogicProps,
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
                <div className="flex items-center">
                    <DataTableUnselect
                        className="rounded-none border first:rounded-l-md last:rounded-r-md"
                        table={table}
                    />
                    {deleteActionProps && (
                        <DataTableDeleteSelected
                            table={table}
                            {...{
                                ...deleteActionProps,
                                className: cn(
                                    'rounded-none border first:rounded-l-md last:rounded-r-md',
                                    deleteActionProps.className
                                ),
                            }}
                        />
                    )}
                    <DataTableRefreshButton
                        {...{
                            ...refreshActionProps,
                            className: cn(
                                'rounded-none border first:rounded-l-md last:rounded-r-md',
                                refreshActionProps.className
                            ),
                        }}
                    />
                    <DataTableOptionsMenu
                        table={table}
                        scrollOption={scrollableProps}
                        filterLogicOption={filterLogicProps}
                        className="rounded-none border first:rounded-l-md last:rounded-r-md"
                    />
                </div>

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
