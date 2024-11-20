import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import { useColumnFilterState } from './column-filter-state-context'
import { filterModeMap, TFilterModes } from '../../data-table-filter-context'
import DateRange from './date-range'
import InputDatePicker from '@/components/input-date-picker'

const DateFilter = () => {
    const {
        filterState: { filterMode, value, rangeValue },
        setFilterMode,
        setValue,
        setRangeValue,
    } = useColumnFilterState()

    const filterModeOptions = filterModeMap['date']

    return (
        <div onKeyDown={(e) => e.stopPropagation()} className="space-y-2 min-w-48 p-1">
            <p className="text-sm">Filter</p>
            <Select
                value={filterMode}
                onValueChange={(val) => {
                    const newFilterMode = val as TFilterModes
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
            {filterMode !== 'range' ? (
                <InputDatePicker
                    id="date-picker-input"
                    captionLayout="dropdown"
                    value={value instanceof Date ? value : undefined}
                    onChange={(newDate) => {
                        if (!newDate) return
                        setValue(newDate)
                    }}
                />
            ) : (
                <DateRange
                    value={rangeValue as any}
                    onChange={(val) => setRangeValue(val)}
                />
            )}
        </div>
    )
}

export default DateFilter
