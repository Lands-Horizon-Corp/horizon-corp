import { useEffect, useState } from 'react'
import {
    SortingState,
    PaginationState,
    RowSelectionState,
} from '@tanstack/react-table'

import {
    PAGINATION_INITIAL_INDEX,
    PAGINATION_INITIAL_PAGE_SIZE,
} from '@/constants'

export type TDataTableDisplayType = 'Default' | 'Full'

interface Props<TData> {
    pageSize?: number
    pageIndex?: number
    columnOrder?: string[]
    onSelectedData?: (data: TData[]) => void
}

const useDataTableState = <TData>(props?: Props<TData>) => {
    const [rowSelection, setRowSelection] = useState<RowSelectionState>({})
    const [selectedRowsData, setSelectedRowsData] = useState<
        Map<string | number, TData>
    >(new Map())

    const [sorting, setSorting] = useState<SortingState>([])
    const [columnOrder, setColumnOrder] = useState<string[]>(
        props?.columnOrder ?? []
    )
    const [columnVisibility, setColumnVisibility] = useState({})
    const [isScrollable, setIsScrollable] = useState<boolean>(true)
    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: props?.pageIndex ?? PAGINATION_INITIAL_INDEX,
        pageSize: props?.pageSize ?? PAGINATION_INITIAL_PAGE_SIZE,
    })

    useEffect(() => {
        props?.onSelectedData?.(
            Array.from<TData>(selectedRowsData.values()).map((val) => val)
        )
    }, [props, selectedRowsData])

    return {
        // states: {
        sorting,
        pagination,
        columnOrder,
        rowSelection,
        isScrollable,
        columnVisibility,
        selectedRowsData,
        // },
        // setters: {
        setSorting,
        setPagination,
        setColumnOrder,
        setRowSelection,
        setIsScrollable,
        setSelectedRowsData,
        setColumnVisibility,
        // },
    }
}

export default useDataTableState
