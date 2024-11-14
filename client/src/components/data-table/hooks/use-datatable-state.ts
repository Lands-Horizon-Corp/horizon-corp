import { useState } from 'react'
import {
    SortingState,
    PaginationState,
} from '@tanstack/react-table'

const useDataTableState = () => {
    const [rowSelection, setRowSelection] = useState({})
    const [sorting, setSorting] = useState<SortingState>([])
    const [columnVisibility, setColumnVisibility] = useState({})
    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: 0,
        pageSize: 10,
    })

    return {
        rowSelection,
        sorting,
        columnVisibility,
        pagination,
        setRowSelection,
        setSorting,
        setColumnVisibility,
        setPagination,
    }
}

export default useDataTableState
