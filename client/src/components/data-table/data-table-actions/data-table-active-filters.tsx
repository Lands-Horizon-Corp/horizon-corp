import { useMemo } from 'react'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import ActionTooltip from '@/components/action-tooltip'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { useDataTableFilter } from '../data-table-filter-context'
import { TrashIcon, XIcon } from '@/components/icons'

const DataTableActiveFilters = ({ className }: IBaseCompNoChild) => {
    const { filters, resetFilter, removeFilter } = useDataTableFilter()

    const mappedFilters = useMemo(() => {
        return Object.entries(filters).map(([key, value]) => ({
            field: key,
            ...value,
        }))
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
                        <TrashIcon />
                    </Button>
                </ActionTooltip>
                {mappedFilters.map((filter) => (
                    <FilterChip
                        key={filter.field}
                        filterKey={filter.field}
                        onClick={() => removeFilter(filter.field)}
                    />
                ))}
            </div>
        </div>
    )
}

const FilterChip = ({
    filterKey,
    onClick,
}: {
    filterKey: string
    onClick: () => void
}) => {
    return (
        <ActionTooltip tooltipContent="remove">
            <Badge
                onClick={onClick}
                variant="secondary"
                className="group relative cursor-pointer pr-6 text-xs font-normal"
            >
                {filterKey}
                <span className="absolute right-2 top-1/2 block -translate-y-1/2 rounded-full group-hover:bg-background">
                    <XIcon />
                </span>
            </Badge>
        </ActionTooltip>
    )
}

export default DataTableActiveFilters
