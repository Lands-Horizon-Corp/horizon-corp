import { createContext, useContext } from 'react'

type TSearchFilter = {
    field: string
    type: 'string' | 'number' | 'date'
    value: string | number | Date
    mode: 'starts with' | 'ends with' | 'include' | 'not include' | 'specific'
}

type TRangeFilter = {
    field: string
    type: 'number' | 'date'
    from: number | Date
    to: number | Date
}

interface DataTableFilterContext {
    filters: {
        search: TSearchFilter[]
        range: TRangeFilter[]
    }
}

const DataTableFilterContext = createContext<
    DataTableFilterContext | undefined
>(undefined)

export const useDataTableFilter = () => {
    const filterContext = useContext(DataTableFilterContext)

    if (!filterContext)
        throw new Error('This hook must be used within the TableFilterContext')

    return filterContext
}

export default DataTableFilterContext
