import { useMemo } from 'react'
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
import { useDataTableFilter } from '../../data-table-filter-context'

export type ColumnDataTypes = 'number' | 'string' | 'Date'

interface Props<TData, TValue> {
    dataType: ColumnDataTypes
    column: Column<TData, TValue>
}

const ColumnFilter = <TData, TValue>({
    column,
    dataType,
}: Props<TData, TValue>) => {
    const { filters, setFilter } = useDataTableFilter()

    const filterMode = useMemo(() => {
        switch (dataType) {
            case 'string':
                return ['equals', 'include', 'starts with', 'ends with']
            case 'Date':
                return ['equals', 'range']
            case 'number':
                return ['equals', 'less than', 'greater than', 'range']
            default:
                return []
        }
    }, [dataType])

    console.log(filters)

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
            <PopoverContent className="space-y-2">
                <Select
                    value={filters[column.id]?.value}
                    defaultValue={
                        filterMode.length === 0 ? undefined : filterMode[0]
                    }
                    onValueChange={(val) => setFilter(column.id, { mode : val, type : dataType })}
                >
                    <SelectTrigger className="">
                        <SelectValue placeholder="Select Filter" />
                    </SelectTrigger>
                    <SelectContent
                        onClick={(e) => e.stopPropagation()}
                        className="ecoop-scroll max-h-[60vh] min-w-40 overflow-y-scroll shadow-md"
                    >
                        {filterMode.map((mode, i) => (
                            <SelectItem key={i} value={mode}>
                                {mode}
                            </SelectItem>
                        ))}
                    </SelectContent>
                </Select>
                <Input
                    className="w-full"
                    placeholder={`search ${column.getFlatColumns.name}`}
                />
            </PopoverContent>
        </Popover>
    )
}

export default ColumnFilter
