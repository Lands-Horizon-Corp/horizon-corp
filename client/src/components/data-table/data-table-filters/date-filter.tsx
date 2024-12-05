import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import { Button } from '@/components/ui/button'
// import InputDatePicker from '@/components/input-date-picker'

import DateRange from './date-range'

import { isDate } from '@/helpers'
import {
    TFilterModes,
    filterModeMap,
    TSearchFilter,
    useDataTableFilter,
    IFilterComponentProps,
} from './data-table-filter-context'
import InputDatePicker from '@/components/date-time-pickers/input-date-picker'

const DateFilter = <T,>({ field, displayText }: IFilterComponentProps<T>) => {
    const { filters, setFilter } = useDataTableFilter<Date, typeof field>()

    const filterModeOptions = filterModeMap['date']

    const filterVal: TSearchFilter<Date> = filters[field] ?? {
        displayText,
        to: undefined,
        from: undefined,
        dataType: 'date',
        value: undefined,
        mode: filterModeOptions[0].value,
    }

    return (
        <div
            onKeyDown={(e) => e.stopPropagation()}
            className="flex min-w-72 flex-col space-y-2 p-1"
        >
            <p className="text-sm">Filter</p>
            <Select
                value={filterVal.mode}
                onValueChange={(val) => {
                    const newFilterMode = val as TFilterModes
                    setFilter(field, {
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
                    fromYear={1960}
                    captionLayout="dropdown"
                    value={
                        isDate(filterVal.value) ? filterVal.value : undefined
                    }
                    onChange={(newDate) => {
                        if (!newDate) return
                        setFilter(field, {
                            ...filterVal,
                            value: newDate,
                            from: undefined,
                            to: undefined,
                        })
                    }}
                />
            ) : (
                <DateRange
                    modal
                    withTimePick
                    fromYear={1960}
                    captionLayout="dropdown-buttons"
                    value={
                        {
                            from: filterVal.from,
                            to: filterVal.to,
                            value: undefined,
                        } as unknown as DateRange
                    }
                    onChange={(val) =>
                        setFilter(field, {
                            ...filterVal,
                            from: val.from,
                            to: val.to,
                            value: undefined,
                        })
                    }
                />
            )}
            <Button
                size="sm"
                className="w-full"
                variant="secondary"
                onClick={() => setFilter(field)}
            >
                Clear Filter
            </Button>
        </div>
    )
}

export default DateFilter
