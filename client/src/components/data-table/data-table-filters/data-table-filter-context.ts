import { KeysOfOrString } from '@/types'
import { createContext, useContext } from 'react'

export type TFilterLogic = 'AND' | 'OR'

export type TColumnDataTypes = 'number' | 'text' | 'date' | 'enum'

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
    [K in Exclude<TColumnDataTypes, 'enum'>]: {
        value: TFilterModes
        label: string
    }[]
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

export type TSearchFilter<T = unknown, TValue = T> = {
    to?: TValue
    from?: TValue
    value?: TValue
    mode: TFilterModes
    displayText: string
    dataType: TColumnDataTypes
}

export type TFinalFilter<T = unknown, TValue = T> = {
    field: string | keyof T
    value: TValue | { from: TValue; to: TValue }
} & Omit<TSearchFilter, 'value' | 'from' | 'to' | 'displayText'>

export type TFilterObject<T = unknown, TValue = T> = {
    [key: string]: TSearchFilter<T, TValue> | undefined
}

export interface IDataTableFilterState<
    T = unknown,
    TField = string,
    TValue = T,
> {
    filters: TFilterObject<T, TValue>
    filterLogic: TFilterLogic
    finalFilters: TFinalFilter[]
    setFilterLogic: (newFilterLogic: TFilterLogic) => void
    setFilter: (field: TField, filter?: TSearchFilter<TValue, TValue>) => void
    bulkSetFilter: (
        field: { field: TField; displayText: string }[],
        filterValue?: TSearchFilter<TValue, TValue>
    ) => void
    resetFilter: () => void
    removeFilter: (field: TField) => void
}

export interface IFilterComponentProps<T> {
    field: KeysOfOrString<T>
    displayText: string
}

const DataTableFilterContext = createContext<
    IDataTableFilterState<unknown> | undefined
>(undefined)

export const useDataTableFilter = <
    T = unknown,
    TField = string,
    TValue = T,
>() => {
    const context = useContext(DataTableFilterContext) as IDataTableFilterState<
        T,
        TField,
        TValue
    >

    if (!context) {
        throw new Error(
            'useDataTableFilter must be used within a DataTableFilterProvider'
        )
    }

    return context
}

export default DataTableFilterContext
