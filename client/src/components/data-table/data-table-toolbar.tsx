import { ReactNode } from 'react'
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
import { IDataTableFilterLogicOptionProps } from './data-table-actions/data-table-options-menu/filter-logic-option'
import { type IDataTableDeleteSelectedProps } from '@/components/data-table/data-table-actions/data-table-delete-selected'
import { type IDataTableScrollableOptionProps } from '@/components/data-table/data-table-actions/data-table-options-menu/scroll-option'
import DataTableCreateAction, {
    IDataTableCreateActionProps,
} from './data-table-actions/data-table-create-action'

export interface IDataTableToolbarProps<TData = unknown>
    extends IBaseCompNoChild {
    table: Table<TData>
    refreshActionProps: IRefreshButtonProps
    globalSearchProps?: IGlobalSearchProps<TData>
    scrollableProps?: IDataTableScrollableOptionProps
    filterLogicProps?: IDataTableFilterLogicOptionProps
    exportActionProps?: Omit<IDataTableExportProps<TData>, 'table'>
    deleteActionProps?: Omit<IDataTableDeleteSelectedProps<TData>, 'table'>
    createActionProps?: IDataTableCreateActionProps
    otherActionRight?: ReactNode
    hideRefreshButton?: boolean
    hideDeleteButton?: boolean
    hideCreateButton?: boolean
    hideExportButton?: boolean
}

const DataTableToolbar = <TData,>({
    table,
    scrollableProps,
    hideCreateButton,
    hideDeleteButton,
    hideExportButton,
    hideRefreshButton,
    otherActionRight,
    filterLogicProps,
    globalSearchProps,
    deleteActionProps,
    exportActionProps,
    createActionProps,
    refreshActionProps,
}: IDataTableToolbarProps<TData>) => {
    return (
        <div className="ecoop-scroll flex w-full max-w-full items-center justify-between gap-x-2 overflow-x-auto">
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
                    {deleteActionProps && !hideDeleteButton && (
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
                    {!hideRefreshButton && (
                        <DataTableRefreshButton
                            {...{
                                ...refreshActionProps,
                                className: cn(
                                    'rounded-none border first:rounded-l-md last:rounded-r-md',
                                    refreshActionProps.className
                                ),
                            }}
                        />
                    )}
                    <DataTableOptionsMenu
                        table={table}
                        scrollOption={scrollableProps}
                        filterLogicOption={filterLogicProps}
                        className="rounded-none border first:rounded-l-md last:rounded-r-md"
                    />
                </div>

                {exportActionProps && !hideExportButton && (
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
                {createActionProps && !hideCreateButton && (
                    <DataTableCreateAction {...createActionProps} />
                )}
                {otherActionRight}
            </div>
        </div>
    )
}

export default DataTableToolbar
