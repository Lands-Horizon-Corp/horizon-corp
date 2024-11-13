import { createContext, useContext, useState, ReactNode } from 'react'

export type TColumnDataTypes = 'number' | 'text' | 'date'

export type TFilterModes =
    | 'equal'
    | 'nequal'
    | 'contains'
    | 'ncontains'
    | 'startswith'
    | 'endswith'
    | 'isempty'
    | 'isnotempty'
    | 'gt'
    | 'gte'
    | 'lt'
    | 'lte'
    | 'range'
    | 'before'
    | 'after'

export const filterModeMap: {
    [K in TColumnDataTypes]: { value: TFilterModes; label: string }[]
} = {
    text: [
        { value: 'equal', label: 'Equals' },
        { value: 'nequal', label: 'Does Not Equal' },
        { value: 'contains', label: 'Contains' },
        { value: 'ncontains', label: 'Does Not Contain' },
        { value: 'startswith', label: 'Starts With' },
        { value: 'endswith', label: 'Ends With' },
        { value: 'isempty', label: 'Is Empty' },
        { value: 'isnotempty', label: 'Is Not Empty' },
    ],
    number: [
        { value: 'equal', label: '(=) Equal' },
        { value: 'nequal', label: 'Not equal' },
        { value: 'gt', label: '(>) Greater Than' },
        { value: 'gte', label: '(>=) Greater Than or Equal' },
        { value: 'lt', label: '(<) Less than' },
        { value: 'lte', label: '(<=) Less than or Equal to' },
        { value: 'range', label: 'Range' },
    ],
    date: [
        { value: 'equal', label: 'On' },
        { value: 'nequal', label: 'Not On' },
        { value: 'before', label: 'Before' },
        { value: 'after', label: 'After' },
        { value: 'gte', label: 'On or After' },
        { value: 'lte', label: 'On or Before' },
        { value: 'range', label: 'Range' },
    ],
}

type TSearchFilter = {
    type: 'string' | 'number' | 'date'
    value: string | number | Date
    mode: TFilterModes // here will the filterModes comes
    // if mode is range, there will be fromValue and toValue
    from: any
    to: any
}

// Define a type for filters where any string key is allowed
type FilterObject = {
    [key: string]: TSearchFilter | undefined
}

interface DataTableFilterContextType {
    filters: FilterObject
    setFilter: (field: string, filter: TSearchFilter) => void
    removeFilter: (field: string) => void
    resetFilter: () => void
}

const DataTableFilterContext = createContext<
    DataTableFilterContextType | undefined
>(undefined)

export const useDataTableFilter = () => {
    const context = useContext(DataTableFilterContext)

    if (!context) {
        throw new Error(
            'useDataTableFilter must be used within a DataTableFilterProvider'
        )
    }

    return context
}

// Example Provider
export const DataTableFilterProvider = ({
    children,
}: {
    children: ReactNode
}) => {
    const [filters, setFilters] = useState<FilterObject>({})

    const setFilter = (field: string, filter: TSearchFilter) => {
        setFilters((prev) => ({ ...prev, [field]: filter }))
    }

    const removeFilter = (field: string) => {
        setFilters((prev) => {
            const newFilters = { ...prev }
            delete newFilters[field]
            return newFilters
        })
    }

    const resetFilter = () => {
        setFilters({})
    }

    return (
        <DataTableFilterContext.Provider
            value={{ filters, setFilter, removeFilter, resetFilter }}
        >
            {children}
        </DataTableFilterContext.Provider>
    )
}