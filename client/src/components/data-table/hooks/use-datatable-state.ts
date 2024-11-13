import { PaginationState, SortingState } from '@tanstack/react-table'
import { useState } from 'react'

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
