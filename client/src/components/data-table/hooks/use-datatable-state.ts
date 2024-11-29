import { useState } from 'react'
import { SortingState, PaginationState } from '@tanstack/react-table'

import {
    PAGINATION_INITIAL_INDEX,
    PAGINATION_INITIAL_PAGE_SIZE,
} from '@/constants'

export type TDataTableDisplayType = 'Default' | 'Full'

interface Props {
    pageIndex?: number
    pageSize?: number
    columnOrder?: string[]
}

const useDataTableState = (props?: Props) => {
    const [rowSelection, setRowSelection] = useState({})
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

    return {
        // states: {
            sorting,
            pagination,
            columnOrder,
            rowSelection,
            isScrollable,
            columnVisibility,
        // },
        // setters: {
            setSorting,
            setPagination,
            setColumnOrder,
            setRowSelection,
            setIsScrollable,
            setColumnVisibility,
        // },
    }
}

export default useDataTableState
