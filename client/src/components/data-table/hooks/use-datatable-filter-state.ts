import { useMemo, useState } from 'react'

import {
    TFinalFilter,
    TSearchFilter,
    TFilterObject,
    IDataTableFilterState,
    TFilterLogic,
} from '../data-table-filters/data-table-filter-context'
import useDebounce from '@/hooks/use-debounce'

const useDatableFilterState = (options?: {
    debounceFinalFilterMs?: number
    onFilterChange?: () => void
}): IDataTableFilterState => {
    const [filters, setFilters] = useState<TFilterObject>({})
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
        setFilters({})
    }

    const debouncedFilter = useDebounce(
        filters,
        options?.debounceFinalFilterMs ?? 800
    )

    const finalFilters: { filters: TFinalFilter[]; logic: TFilterLogic } =
        useMemo(() => {
            const filteredFilter: TFinalFilter[] = []

            Object.entries(debouncedFilter).forEach(([key, value]) => {
                if (!value || !value.value) return

                if (!value.mode || key === 'globalSearch') return

                if (
                    value.mode === 'range' &&
                    (value.from === undefined || value.to === undefined)
                ) {
                    return
                } else if (value.mode !== 'range' && value.value === undefined)
                    return

                filteredFilter.push({
                    field: key,
                    mode: value.mode,
                    dataType: value.dataType,
                    value:
                        value.mode === 'range'
                            ? { from: value.from, to: value.to }
                            : value.value,
                })
            })

            options?.onFilterChange?.()

            return { filters: filteredFilter, logic: filterLogic }
        }, [debouncedFilter, filterLogic])

    return {
        filters,
        filterLogic,
        finalFilters,
        setFilter,
        resetFilter,
        removeFilter,
        bulkSetFilter,
        setFilterLogic,
    }
}

export default useDatableFilterState
