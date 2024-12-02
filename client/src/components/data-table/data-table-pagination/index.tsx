import { Table } from '@tanstack/react-table'

import {
    Select,
    SelectItem,
    SelectValue,
    SelectContent,
    SelectTrigger,
} from '@/components/ui/select'
import { Button } from '@/components/ui/button'

import {
    ChevronLeftIcon,
    ChevronRightIcon,
    ChevronsLeftIcon,
    ChevronsRightIcon,
} from '@/components/icons'
import { cn } from '@/lib/utils'
import { PAGE_SIZES_DENSE } from '@/constants'
import logger from '@/helpers/loggers/logger'
import { useEffect } from 'react'

interface DataTablePaginationProps<TData> {
    table: Table<TData>
    className?: string
    pageSizes?: number[]
    totalSize: number
    hideSelectedIndicator?: boolean
}

const DataTablePagination = <TData,>({
    table,
    className,
    totalSize,
    pageSizes = PAGE_SIZES_DENSE,
    hideSelectedIndicator = false,
}: DataTablePaginationProps<TData>) => {
    const finalPageSizes = pageSizes.filter((size) => size < totalSize)

    useEffect(() => {
        if (
            table.getRowCount() < finalPageSizes[0] &&
            table.getState().pagination.pageSize !== finalPageSizes[0]
        )
            logger.log('Less')

        logger.log('Not less')
    }, [table])

    return (
        <div
            className={cn(
                'mt-1 flex flex-col items-center justify-between space-y-4 px-2 md:flex-row md:space-y-0',
                className,
                hideSelectedIndicator && 'justify-end'
            )}
        >
            {!hideSelectedIndicator && (
                <div className="flex-1 text-sm text-muted-foreground">
                    {table.getFilteredSelectedRowModel().rows.length} of{' '}
                    {table.getFilteredRowModel().rows.length} row(s) selected.
                </div>
            )}
            <div className="flex flex-col items-center space-y-4 md:flex-row md:space-x-6 md:space-y-0 lg:space-x-8">
                <div className="flex items-center space-x-2">
                    <p className="text-sm font-medium">Rows per page</p>
                    <Select
                        value={`${table.getState().pagination.pageSize}`}
                        onValueChange={(value) => {
                            table.setPageSize(Number(value))
                        }}
                    >
                        <SelectTrigger className="h-8 w-[70px]">
                            <SelectValue
                                placeholder={
                                    table.getState().pagination.pageSize
                                }
                            />
                        </SelectTrigger>
                        <SelectContent side="top">
                            {finalPageSizes.map((pageSize) => (
                                <SelectItem
                                    key={pageSize}
                                    value={`${pageSize}`}
                                >
                                    {pageSize}
                                </SelectItem>
                            ))}
                            <SelectItem value={`${totalSize}`}>All</SelectItem>
                        </SelectContent>
                    </Select>
                </div>
                <div className="flex w-[100px] items-center justify-center text-sm font-medium">
                    Page {table.getState().pagination.pageIndex + 1} of{' '}
                    {table.getPageCount()}
                </div>
                <div className="flex items-center space-x-2">
                    <Button
                        variant="outline"
                        className="hidden size-8 p-0 lg:flex"
                        onClick={() => table.setPageIndex(0)}
                        disabled={!table.getCanPreviousPage()}
                    >
                        <span className="sr-only">Go to first page</span>
                        <ChevronsLeftIcon className="size-4" />
                    </Button>
                    <Button
                        variant="outline"
                        className="size-8 p-0"
                        onClick={() => table.previousPage()}
                        disabled={!table.getCanPreviousPage()}
                    >
                        <span className="sr-only">Go to previous page</span>
                        <ChevronLeftIcon className="size-4" />
                    </Button>
                    <Button
                        variant="outline"
                        className="size-8 p-0"
                        onClick={() => table.nextPage()}
                        disabled={!table.getCanNextPage()}
                    >
                        <span className="sr-only">Go to next page</span>
                        <ChevronRightIcon className="size-4" />
                    </Button>
                    <Button
                        variant="outline"
                        className="hidden size-8 p-0 lg:flex"
                        onClick={() =>
                            table.setPageIndex(table.getPageCount() - 1)
                        }
                        disabled={!table.getCanNextPage()}
                    >
                        <span className="sr-only">Go to last page</span>
                        <ChevronsRightIcon className="size-4" />
                    </Button>
                </div>
            </div>
        </div>
    )
}

export default DataTablePagination
