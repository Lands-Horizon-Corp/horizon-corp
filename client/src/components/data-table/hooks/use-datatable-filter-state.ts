import { useMemo, useState } from 'react'

import {
    TFilterLogic,
    TFinalFilter,
    TSearchFilter,
    TFilterObject,
    IDataTableFilterState,
    TFilterPayload,
} from '../data-table-filters/data-table-filter-context'
import useDebounce from '@/hooks/use-debounce'

const useDatableFilterState = (options?: {
    onFilterChange?: () => void
    debounceFinalFilterMs?: number
    defaultFilter?: TFilterObject
}): IDataTableFilterState => {
    const [filters, setFilters] = useState<TFilterObject>(
        options?.defaultFilter ?? {}
    )
    const [filterLogic, setFilterLogic] = useState<TFilterLogic>('AND')

    const setFilter = (field: string, filter?: TSearchFilter) => {
        setFilters((prev) => ({ ...prev, [field]: filter }))
    }

    const removeFilter = (field: string) => {
        const targetFilter = filters[field]

        setFilters((prev) => {
            const newFilters = { ...prev }
            delete newFilters[field]
            return newFilters
        })

        return targetFilter
    }

    const bulkSetFilter = (
        targets: { field: string; displayText: string }[],
        filterValue?: TSearchFilter
    ) => {
        const constructedObject = {} as TFilterObject
        targets.forEach(({ field, displayText }) => {
            constructedObject[field] = {
                ...filterValue,
                displayText,
            } as TSearchFilter
        })
        setFilters((prev) => ({ ...prev, ...constructedObject }))
    }

    const resetFilter = () => {
        setFilters(options?.defaultFilter ?? {})
    }

    const debouncedFilter = useDebounce(
        filters,
        options?.debounceFinalFilterMs ?? 800
    )

    const finalFilterPayload: TFilterPayload = useMemo(() => {
        const filteredFilter: TFinalFilter[] = []

        Object.entries(debouncedFilter).forEach(([key, value]) => {
            if (!value || !value.value) {
                // logger.log('Value failed', value)
                return
            }

            if (!value.mode || key === 'globalSearch') {
                // logger.log( 'value mode, globalSearch failed', value.mode, key)
                return
            }

            if (
                value.mode === 'range' &&
                !Array.isArray(value.value) &&
                (value.from === undefined || value.to === undefined)
            ) {
                // logger.log( 'line 80 failed -> invalid range from to', value.from, value.to)
                return
            } else if (value.mode !== 'range' && value.value === undefined) {
                // logger.log('line 86 failed -> invalid value', value.value)
                return
            } else if (Array.isArray(value.value) && value.value.length === 0) {
                // logger.log( 'line 89 failed -> invalid multi select array', value.value)
                return
            }

            filteredFilter.push({
                field: key,
                mode: value.mode,
                dataType: value.dataType,
                isStaticFilter: value.isStaticFilter,
                value:
                    value.mode === 'range' && !Array.isArray(value.value)
                        ? { from: value.from, to: value.to }
                        : value.value,
            })
        })

        options?.onFilterChange?.()

        return { filters: filteredFilter, logic: filterLogic }

        // WARNING: don't worry about this, if you remove this and follow the suggestion, infinite loop will happen
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [debouncedFilter, filterLogic])

    return {
        filters,
        filterLogic,
        finalFilterPayload,
        setFilter,
        resetFilter,
        removeFilter,
        bulkSetFilter,
        setFilterLogic,
    }
}

export default useDatableFilterState
