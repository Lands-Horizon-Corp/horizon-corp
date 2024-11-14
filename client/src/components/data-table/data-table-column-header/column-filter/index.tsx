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
import DateRange from './date-range'
import NumberRange from './number-range'
import useDebounce from '@/hooks/use-debounce'
import InputDatePicker from '@/components/input-date-picker'

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

    const [value, setValue] = useState<any>('')
    const [rangeValue, setRangeValue] = useState<{ from?: any; to?: any }>({
        from: undefined,
        to: undefined,
    })

    const resetFilter = () => {
        setRangeValue({ from: undefined, to: undefined })
        setValue('')
        setFilterMode('equal')
    }

    const filterInput = useMemo(
        () => ({
            mode: filterMode,
            ...(filterMode === 'range' ? undefined : { value }),
            rangeValue:
                filterMode === 'range'
                    ? rangeValue
                    : { from: undefined, to: undefined },
        }),
        [filterMode, value, rangeValue]
    )

    const finalFilterValue = useDebounce(filterInput, 2000)

    useEffect(() => {
        if (
            finalFilterValue.mode === undefined ||
            (finalFilterValue.mode !== 'range' && !finalFilterValue.value) ||
            (finalFilterValue.mode === 'range' &&
                (!finalFilterValue.rangeValue.to ||
                    !finalFilterValue.rangeValue.from))
        ) {
            console.log('Removed', removeFilter(column.id))
            return
        }

        const newFilter = {
            mode: finalFilterValue.mode,
            dataType,
            value: finalFilterValue.value,
            ...finalFilterValue.rangeValue,
        }

        console.log('Added ', newFilter)

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
            <PopoverContent className="w-fit min-w-72 space-y-2 rounded-2xl border bg-popover/85 shadow-lg backdrop-blur">
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
                {filterMode !== 'range' && dataType === 'date' && (
                    <InputDatePicker
                        id="date-picker-input"
                        captionLayout="dropdown"
                        value={value instanceof Date ? value : undefined}
                        onChange={(newDate) => {
                            if (!newDate) return
                            setValue(newDate)
                        }}
                    />
                )}
                {filterMode === 'range' && dataType === 'date' && (
                    <DateRange
                        onChange={setRangeValue}
                        value={rangeValue as any}
                    />
                )}
                {filterMode === 'range' && dataType === 'number' && (
                    <NumberRange
                        value={rangeValue}
                        onChange={(val) => setRangeValue(val)}
                    />
                )}
                {((filterMode === 'range' &&
                    rangeValue.from &&
                    rangeValue.to) ||
                    (filterMode !== 'range' && value)) && (
                    <div className="flex items-center justify-end pt-1">
                        <Button
                            onClick={() => resetFilter()}
                            variant="secondary"
                        >
                            Clear
                        </Button>
                    </div>
                )}
            </PopoverContent>
        </Popover>
    )
}

export default ColumnFilter
