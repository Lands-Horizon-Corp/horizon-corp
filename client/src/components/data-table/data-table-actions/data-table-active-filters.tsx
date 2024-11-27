import { useMemo } from 'react'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import FilterChip from '@/components/filter-chip'
import { TrashBinIcon } from '@/components/icons'
import ActionTooltip from '@/components/action-tooltip'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { useDataTableFilter } from '../data-table-filter-context'

const DataTableActiveFilters = ({ className }: IBaseCompNoChild) => {
    const { filters, resetFilter, removeFilter } = useDataTableFilter()

    const mappedFilters = useMemo(() => {
        return Object.entries(filters)
            .map(([key, value]) => ({
                field: key,
                ...value,
            }))
            .filter((filter) => filter.value || (filter.from && filter.to))
    }, [filters])

    if (mappedFilters.length <= 0) return null

    return (
        <div className={cn('flex max-w-full items-center gap-x-2', className)}>
            <span className="inline-flex items-center">
                Filters
                <Badge variant="secondary" className="ml-1 size-fit text-xs">
                    {mappedFilters.length}
                </Badge>
            </span>
            <div className="ecoop-scroll flex flex-wrap items-center gap-x-2 gap-y-1">
                <ActionTooltip tooltipContent="Remove All Filters">
                    <Button
                        size="sm"
                        variant="secondary"
                        className="size-fit p-1 text-xs text-foreground/60"
                        onClick={() => resetFilter()}
                    >
                        <TrashBinIcon />
                    </Button>
                </ActionTooltip>
                {mappedFilters.map((filter) => (
                    <FilterChip
                        key={filter.field}
                        label={filter.field}
                        tooltipDescription="Remove Filter"
                        onClick={() => removeFilter(filter.field)}
                    />
                ))}
            </div>
        </div>
    )
}
export default DataTableActiveFilters
