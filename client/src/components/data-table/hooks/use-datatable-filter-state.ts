import { useMemo, useState } from 'react'

import {
    TFinalFilter,
    TSearchFilter,
    TFilterObject,
    IDataTableFilterState,
} from '../data-table-filters/data-table-filter-context'

const useDatableFilterState = (): IDataTableFilterState => {
    const [filters, setFilters] = useState<TFilterObject>({})

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

    const bulkSetFilter = (fields: string[], filterValue?: TSearchFilter) => {
        const constructedObject = {} as TFilterObject
        fields.forEach((field) => (constructedObject[field] = filterValue))
        setFilters((prev) => ({ ...prev, ...constructedObject }))
    }

    const resetFilter = () => {
        setFilters({})
    }

    const finalFilters: TFinalFilter[] = useMemo(() => {
        const filteredFilter: TFinalFilter[] = []

        Object.entries(filters).forEach(([key, value]) => {
            if (!value) return

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
                preload: '',
                dataType: value.dataType,
                mode: value.mode,
                value:
                    value.mode === 'range'
                        ? { from: value.from, to: value.to }
                        : value.value,
            })
        })

        return filteredFilter
    }, [filters])

    return {
        filters,
        finalFilters,
        setFilter,
        removeFilter,
        resetFilter,
        bulkSetFilter,
    }
}

export default useDatableFilterState
