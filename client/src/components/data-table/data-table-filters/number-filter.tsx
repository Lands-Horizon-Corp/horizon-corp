import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import { Button } from '@/components/ui/button'
import { DebouncedInput } from '@/components/ui/debounced-input'

import {
    filterModeMap,
    TFilterModes,
    TSearchFilter,
    useDataTableFilter,
} from './data-table-filter-context'
import NumberRange from './number-range'

const NumberFilter = <TData,>({
    accessorKey,
}: {
    accessorKey: keyof TData
}) => {
    const filterModeOptions = filterModeMap['number']
    const { filters, setFilter } = useDataTableFilter<number, keyof TData>()

    const filterVal: TSearchFilter<number> = filters[accessorKey as string] ?? {
        dataType: 'number',
        mode: filterModeOptions[0].value,
        value: undefined,
        from: undefined,
        to: undefined,
    }

    return (
        <div onKeyDown={(e) => e.stopPropagation()} className="space-y-2 p-1">
            <p className="text-sm">Filter</p>
            <Select
                value={filterVal.mode}
                onValueChange={(val) => {
                    const newFilterMode = val as TFilterModes
                    setFilter(accessorKey, {
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
                <DebouncedInput
                    type="number"
                    value={filterVal.value}
                    className="w-full"
                    onChange={(inpt) =>
                        setFilter(accessorKey, {
                            ...filterVal,
                            value: inpt as number,
                            from: undefined,
                            to: undefined,
                        })
                    }
                    placeholder={`value`}
                />
            ) : (
                <NumberRange
                    value={{ from: filterVal.from, to: filterVal.to }}
                    onChange={(val) =>
                        setFilter(accessorKey, {
                            ...filterVal,
                            ...val,
                            value: undefined,
                        })
                    }
                />
            )}
            <Button
                size="sm"
                className="w-full"
                variant="secondary"
                onClick={() => setFilter(accessorKey)}
            >
                Clear Filter
            </Button>
        </div>
    )
}

export default NumberFilter
