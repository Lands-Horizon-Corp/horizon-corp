import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import DateRange from './date-range'
import { Button } from '@/components/ui/button'
import InputDatePicker from '@/components/input-date-picker'

import { isDate } from '@/helpers'
import {
    filterModeMap,
    TFilterModes,
    TSearchFilter,
    useDataTableFilter,
} from './data-table-filter-context'

const DateFilter = <TData,>({ accessorKey }: { accessorKey: keyof TData }) => {
    const { filters, setFilter } = useDataTableFilter()

    const filterModeOptions = filterModeMap['date']

    const filterVal: TSearchFilter = filters[accessorKey as string] ?? {
        dataType: 'date',
        mode: filterModeOptions[0].value,
        value: undefined,
        from: undefined,
        to: undefined,
    }

    return (
        <div
            onKeyDown={(e) => e.stopPropagation()}
            className="min-w-48 space-y-2 p-1"
        >
            <p className="text-sm">Filter</p>
            <Select
                value={filterVal.mode}
                onValueChange={(val) => {
                    const newFilterMode = val as TFilterModes
                    setFilter(accessorKey as string, {
                        ...filterVal,
                        mode: newFilterMode,
                    })
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
            {filterVal.mode !== 'range' ? (
                <InputDatePicker
                    id="date-picker-input"
                    captionLayout="dropdown"
                    value={
                        isDate(filterVal.value) ? filterVal.value : undefined
                    }
                    onChange={(newDate) => {
                        if (!newDate) return
                        setFilter(accessorKey as string, {
                            ...filterVal,
                            value: newDate,
                        })
                    }}
                />
            ) : (
                <DateRange
                    value={
                        {
                            from: filterVal.from,
                            to: filterVal.to,
                        } as unknown as DateRange
                    }
                    onChange={(val) =>
                        setFilter(accessorKey as string, {
                            ...filterVal,
                            from: val.from,
                            to: val.to,
                        })
                    }
                />
            )}
            <Button
                size="sm"
                className="w-full"
                variant="secondary"
                onClick={() => setFilter(accessorKey as string)}
            >
                Clear Filter
            </Button>
        </div>
    )
}

export default DateFilter
