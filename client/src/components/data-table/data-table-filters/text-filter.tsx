import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import { Button } from '@/components/ui/button'

import {
    filterModeMap,
    TFilterModes,
    TSearchFilter,
    useDataTableFilter,
} from './data-table-filter-context'
import { DebouncedInput } from '@/components/ui/debounced-input'

const TextFilter = <TData,>({ accessorKey }: { accessorKey: keyof TData }) => {
    const { filters, setFilter } = useDataTableFilter<string, keyof TData>()

    const filterVal: TSearchFilter<string> = filters[accessorKey as string] ?? {
        dataType: 'text',
        mode: 'equal',
        value: '',
        from: undefined,
        to: undefined,
    }

    const filterModeOptions = filterModeMap['text']

    return (
        <div onKeyDown={(e) => e.stopPropagation()} className="space-y-2 p-1">
            <p className="text-sm">Filter</p>
            <Select
                value={filterVal?.mode}
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
            <DebouncedInput
                type="text"
                className="w-full"
                placeholder="value"
                debounceTime={500}
                value={filterVal.value || ''}
                onChange={(inpt: string) =>
                    setFilter(accessorKey, {
                        ...filterVal,
                        value: inpt,
                    })
                }
            />
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

export default TextFilter
