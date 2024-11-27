import {
    createContext,
    useContext,
    useState,
    ReactNode,
    useEffect,
    useMemo,
} from 'react'

import {
    TColumnDataTypes,
    TFilterModes,
    useDataTableFilter,
} from '../../data-table-filter-context'
import useDebounce from '@/hooks/use-debounce'

interface ColumnFilterState {
    filterMode: TFilterModes | undefined
    value: any
    rangeValue: { from?: any; to?: any }
}

interface ColumnFilterContextType {
    fieldName: string
    dataType: TColumnDataTypes
    filterState: ColumnFilterState
    clearFilter: () => void
    setValue: (value: any) => void
    setFilterMode: (mode: TFilterModes | undefined) => void
    setRangeValue: (range: { from?: any; to?: any }) => void
}

export const ColumnFilterContext = createContext<
    ColumnFilterContextType | undefined
>(undefined)

export const ColumnFilterProvider = ({
    children,
    dataType,
    fieldName,
}: {
    children: ReactNode
    fieldName: string
    dataType: TColumnDataTypes
}) => {
    const { filters, setFilter } = useDataTableFilter()

    const columnFilterValue = filters[fieldName]

    const [value, setValue] = useState<any>('')
    const [filterMode, setFilterMode] = useState<TFilterModes | undefined>(
        'equal'
    )
    const [rangeValue, setRangeValue] = useState<{ from?: any; to?: any }>({
        from: undefined,
        to: undefined,
    })

    const clearFilter = () => {
        setValue('')
        setRangeValue({
            from: undefined,
            to: undefined,
        })
    }

    useEffect(() => {
        const colFilterVal = columnFilterValue
        if (!colFilterVal) return

        setValue(colFilterVal.value)
        setFilterMode(colFilterVal.mode)
        setRangeValue({
            from: colFilterVal.from,
            to: colFilterVal.to,
        })
    }, [setValue, setFilterMode, setRangeValue, fieldName])

    const memoizedFilter = useMemo(() => {
        return { mode: filterMode as any, value, dataType, ...rangeValue }
    }, [filterMode, value, dataType, rangeValue])

    const finalFilter = useDebounce(memoizedFilter, 500)

    useEffect(() => {
        setFilter(fieldName, finalFilter)
    }, [finalFilter, fieldName])

    useEffect(() => {
        const currentFilter = filters[fieldName]
        if (!currentFilter) return

        const {
            value: mainValue,
            from: mainFromValue,
            to: mainToValue,
        } = currentFilter

        if (value !== mainValue) setValue(mainValue)
        if (rangeValue.from !== mainFromValue || rangeValue.to !== mainToValue)
            setRangeValue((prev) => ({
                ...prev,
                from: mainFromValue,
                to: mainToValue,
            }))
    }, [fieldName, filters])

    return (
        <ColumnFilterContext.Provider
            value={{
                dataType,
                fieldName,
                filterState: { filterMode, value, rangeValue },
                setValue,
                clearFilter,
                setRangeValue,
                setFilterMode,
            }}
        >
            {children}
        </ColumnFilterContext.Provider>
    )
}

export const useColumnFilterState = () => {
    const context = useContext(ColumnFilterContext)
    if (!context) {
        throw new Error(
            'useColumnFilterState must be used within a ColumnFilterProvider'
        )
    }
    return context
}
