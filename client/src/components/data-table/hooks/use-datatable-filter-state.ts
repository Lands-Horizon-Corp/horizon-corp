import { useState } from 'react'
import { FilterObject, TSearchFilter } from '../data-table-filter-context'

const useDatableFilterState = () => {
    const [filters, setFilters] = useState<FilterObject>({})

    const setFilter = (field: string, filter: TSearchFilter) => {
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

    const resetFilter = () => {
        setFilters({})
    }

    console.log(filters)

    return { filters, setFilter, removeFilter, resetFilter }
}

export default useDatableFilterState
