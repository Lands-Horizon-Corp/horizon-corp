import { useState } from 'react'
import {
    SortingState,
    PaginationState,
} from '@tanstack/react-table'

export type TDataTableDisplayType = "Default" | "Full"

const useDataTableState = () => {
    const [rowSelection, setRowSelection] = useState({})
    const [sorting, setSorting] = useState<SortingState>([])
    const [columnVisibility, setColumnVisibility] = useState({})
    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: 0,
        pageSize: 10,
    })
    const [isScrollable, setIsScrollable] = useState<boolean>(true);

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
