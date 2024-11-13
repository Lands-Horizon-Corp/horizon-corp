import { useMemo, useState } from 'react'
import { Column } from '@tanstack/react-table'

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
import { FunnelOutlineIcon } from '@/components/icons'
import { filterModeMap, TColumnDataTypes } from '../../data-table-filter-context'


interface Props<TData, TValue> {
    dataType: TColumnDataTypes
    column: Column<TData, TValue>
}


const ColumnFilter = <TData, TValue>({
    column,
    dataType,
}: Props<TData, TValue>) => {
    const filterModeOptions = useMemo(() => filterModeMap[dataType], [dataType])
    const [filterMode, setFilterMode] = useState<string | undefined>('equal')

    const [value, setValue] = useState('')
    const [rangeValue, setRangeValue] = useState<{ from: any; to: any }>({
        from: undefined,
        to: undefined,
    })

    return (
        <Popover>
            <PopoverTrigger onClick={(e) => e.stopPropagation()}>
                <Button
                    variant="ghost"
                    size="sm"
                    className="group-hover size-fit gap-x-2 p-1 data-[state=open]:bg-accent"
                >
                    <FunnelOutlineIcon className="size-3 opacity-55 group-hover:opacity-100" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="space-y-2 rounded-2xl border bg-popover/85 shadow-lg backdrop-blur">
                <p className="text-sm font-medium opacity-70">Filter Mode</p>
                <Select
                    value={filterMode}
                    onValueChange={(val) => setFilterMode(val)}
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
                {filterMode !== 'range' && (
                    <Input
                        className="w-full"
                        value={value}
                        onChange={(inpt) => setValue(inpt.target.value)}
                        placeholder={`search ${column.getFlatColumns.name}`}
                    />
                )}
                {filterMode === 'range' && dataType === 'number' && (
                    <div className="flex items-center gap-x-1">
                        <Input
                            type={dataType}
                            className="w-full"
                            value={rangeValue.from}
                            onChange={(inpt) =>
                                setRangeValue((prev) => ({
                                    ...prev,
                                    from: inpt.target.value,
                                }))
                            }
                            placeholder={`Min ${column.getFlatColumns.name}`}
                        />
                        <Input
                            type={dataType}
                            className="w-full"
                            value={rangeValue.to}
                            onChange={(inpt) =>
                                setRangeValue((prev) => ({
                                    ...prev,
                                    to: inpt.target.value,
                                }))
                            }
                            placeholder={`Max`}
                        />
                    </div>
                )}
            </PopoverContent>
        </Popover>
    )
}

export default ColumnFilter
