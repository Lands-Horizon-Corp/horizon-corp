import { useState } from 'react'
import {
    FilterObject,
    TSearchFilter,
    IDataTableFilterState,
} from '../data-table-filters/data-table-filter-context'

const useDatableFilterState = (): IDataTableFilterState => {
    const [filters, setFilters] = useState<FilterObject>({})

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
        // eslint-disable-next-line prefer-const
        let constructedObject = {} as FilterObject
        fields.forEach((field) => (constructedObject[field] = filterValue))
        setFilters((prev) => ({ ...prev, ...constructedObject }))
    }

    const resetFilter = () => {
        setFilters({})
    }

    return { filters, bulkSetFilter, setFilter, removeFilter, resetFilter }
}

export default useDatableFilterState
