import { Column } from '@tanstack/react-table'
import { useEffect, useMemo, useState } from 'react'

import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { FunnelFilledIcon } from '@/components/icons'

import {
    filterModeMap,
    TColumnDataTypes,
    TFilterModes,
    useDataTableFilter,
} from '../../data-table-filter-context'
import NumberRange from './number-range'
import useDebounce from '@/hooks/use-debounce'

interface Props<TData, TValue> {
    dataType: TColumnDataTypes
    column: Column<TData, TValue>
}

const ColumnFilter = <TData, TValue>({
    column,
    dataType,
}: Props<TData, TValue>) => {
    const { setFilter, removeFilter } = useDataTableFilter()
    const filterModeOptions = useMemo(() => filterModeMap[dataType], [dataType])
    const [filterMode, setFilterMode] = useState<TFilterModes | undefined>(
        'equal'
    )

    const [value, setValue] = useState('')
    const [rangeValue, setRangeValue] = useState<{ from?: any; to?: any }>({
        from: undefined,
        to: undefined,
    })

    const filterInput = useMemo(
        () => ({
            mode: filterMode,
            value,
            rangeValue:
                filterMode === 'range'
                    ? rangeValue
                    : { from: undefined, to: undefined },
        }),
        [filterMode, value, rangeValue]
    )

    const finalFilterValue = useDebounce(filterInput, 2000)

    useEffect(() => {
        if (finalFilterValue.mode === undefined || !finalFilterValue.value) {
            removeFilter(column.id)
            return
        }

        const newFilter = {
            mode: finalFilterValue.mode,
            dataType,
            value,
            ...finalFilterValue.rangeValue,
        }

        setFilter(column.id, newFilter)
    }, [finalFilterValue])

    return (
        <Popover>
            <PopoverTrigger onClick={(e) => e.stopPropagation()}>
                <Button
                    variant="ghost"
                    size="sm"
                    className="group-hover size-fit gap-x-2 p-1 data-[state=open]:bg-accent"
                >
                    <FunnelFilledIcon className="size-3 opacity-55 group-hover:opacity-100" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="space-y-2 rounded-2xl border bg-popover/85 shadow-lg backdrop-blur">
                <p className="text-sm font-medium opacity-70">Filter Mode</p>
                <Select
                    value={filterMode}
                    onValueChange={(val) => {
                        const newFilterMode = val as TFilterModes
                        if (newFilterMode === 'range') setValue('')
                        setFilterMode(newFilterMode)
                    }}
                >
                    <SelectTrigger className="">
                        <SelectValue placeholder="Select Filter" />
                    </SelectTrigger>
                    <SelectContent
                        onClick={(e) => e.stopPropagation()}
                        className="ecoop-scroll max-h-[60vh] min-w-40 overflow-y-scroll shadow-md"
                    >
                        {filterModeOptions.map((mode, i) => (
                            <SelectItem key={i} value={mode.value}>
                                {mode.label}
                            </SelectItem>
                        ))}
                    </SelectContent>
                </Select>
                {filterMode !== 'range' && dataType !== 'date' && (
                    <Input
                        className="w-full"
                        value={value}
                        type={dataType}
                        onChange={(inpt) => setValue(inpt.target.value)}
                        placeholder={`value`}
                    />
                )}
                {filterMode === 'range' && dataType === 'number' && (
                    <NumberRange
                        value={rangeValue}
                        onChange={(val) => setRangeValue(val)}
                    />
                )}
            </PopoverContent>
        </Popover>
    )
}

export default ColumnFilter
