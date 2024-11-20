import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import { Input } from '@/components/ui/input'
import { useColumnFilterState } from './column-filter-state-context'
import { filterModeMap, TFilterModes } from '../../data-table-filter-context'

const TextFilter = () => {
    const {
        filterState: { filterMode, value },
        setFilterMode,
        setValue,
    } = useColumnFilterState()

    const filterModeOptions = filterModeMap['text']

    return (
        <div onKeyDown={(e) => e.stopPropagation()} className="space-y-2 p-1">
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
            <Input
                type="text"
                value={value}
                className="w-full"
                onChange={(inpt) => setValue(inpt.target.value)}
                placeholder={`value`}
            />
        </div>
    )
}

export default TextFilter
