import {
    OnChangeFn,
    SortingState,
    PaginationState,
    RowSelectionState,
} from '@tanstack/react-table'
import { useCallback, useMemo, useState } from 'react'

import {
    PAGINATION_INITIAL_INDEX,
    PAGINATION_INITIAL_PAGE_SIZE,
} from '@/constants'
import { useSortingState } from '@/hooks/use-sorting-state'

export type TDataTableDisplayType = 'Default' | 'Full'

interface Props<TData> {
    pageSize?: number
    pageIndex?: number
    columnOrder?: string[]
    onSelectData?: (data: TData[]) => void
}

interface RowSelectionStateWithData<TData> {
    rowSelection: RowSelectionState
    selectedRowsData: Map<string | number, TData>
}

const useDataTableState = <TData extends { id: string | number }>(
    props?: Props<TData>
) => {
    const [rowSelectionState, setRowSelectionState] = useState<
        RowSelectionStateWithData<TData>
    >({
        rowSelection: {},
        selectedRowsData: new Map(),
    })
    const [columnOrder, setColumnOrder] = useState<string[]>(
        props?.columnOrder ?? []
    )
    const [columnVisibility, setColumnVisibility] = useState({})
    const [isScrollable, setIsScrollable] = useState<boolean>(true)
    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: props?.pageIndex ?? PAGINATION_INITIAL_INDEX,
        pageSize: props?.pageSize ?? PAGINATION_INITIAL_PAGE_SIZE,
    })
    const { sortingState, setSortingState } = useSortingState()

    const sorting: SortingState = useMemo(() => {
        return sortingState.map((sortItem) => ({
            id: sortItem.field,
            desc: sortItem.order === 'desc',
        }))
    }, [sortingState])

    const setSorting = (newSortingState: SortingState) => {
        setSortingState(
            newSortingState.map((sortItem) => ({
                field: sortItem.id,
                order: sortItem.desc ? 'desc' : 'asc',
            }))
        )
    }

    const getRowIdFn = useCallback((row: TData) => `${row.id}`, [])

    const createHandleRowSelectionChange = (
        data: TData[]
    ): OnChangeFn<RowSelectionState> => {
        return (updaterOrValue) => {
            setRowSelectionState((prev) => {
                const newRowSelection =
                    typeof updaterOrValue === 'function'
                        ? updaterOrValue(prev.rowSelection)
                        : updaterOrValue

                const newSelectedRowsData = new Map(prev.selectedRowsData)

                data.forEach((row) => {
                    if (newRowSelection[row.id]) {
                        newSelectedRowsData.set(row.id, row)
                    } else {
                        newSelectedRowsData.delete(row.id)
                    }
                })

                props?.onSelectData?.(Array.from(newSelectedRowsData.values()))

                return {
                    rowSelection: newRowSelection,
                    selectedRowsData: newSelectedRowsData,
                }
            })
        }
    }

    return {
        sortingState,
        getRowIdFn,
        sorting,
        setSorting,
        pagination,
        setPagination,
        columnOrder,
        setColumnOrder,
        isScrollable,
        setIsScrollable,
        columnVisibility,
        setColumnVisibility,
        rowSelectionState,
        setRowSelectionState,
        createHandleRowSelectionChange,
    }
}

export default useDataTableState
