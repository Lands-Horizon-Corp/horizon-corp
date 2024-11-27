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
}

const useDataTableState = (props?: Props) => {
    const [rowSelection, setRowSelection] = useState({})
    const [sorting, setSorting] = useState<SortingState>([])
    const [columnVisibility, setColumnVisibility] = useState({})
    const [isScrollable, setIsScrollable] = useState<boolean>(true)
    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: props?.pageIndex ?? PAGINATION_INITIAL_INDEX,
        pageSize: props?.pageSize ?? PAGINATION_INITIAL_PAGE_SIZE,
    })

    return {
        sorting,
        pagination,
        rowSelection,
        isScrollable,
        columnVisibility,
        setSorting,
        setPagination,
        setRowSelection,
        setIsScrollable,
        setColumnVisibility,
    }
}

export default useDataTableState
