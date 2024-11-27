import { useState } from 'react'

import { Button } from '@/components/ui/button'
import { MagnifyingGlassIcon, XIcon } from '@/components/icons'
import { DebouncedInput } from '@/components/ui/debounced-input'

import {
    TFilterModes,
    TSearchFilter,
    useDataTableFilter,
} from './data-table-filter-context'

type KeysOfType<T, ValueType> = {
    [K in keyof T]: T[K] extends ValueType ? K : never
}[keyof T]

interface Props<T> {
    keysToSearch: Array<KeysOfType<T, string>>
    defaultMode: TFilterModes
    placeHolder?: string
}

const accessorKey = 'globalSearch'

const DataTableGlobalSearch = <T,>({
    keysToSearch,
    defaultMode,
    ...otherProps
}: Props<T>) => {
    const [visible, setVisible] = useState(false)
    const { filters, setFilter, bulkSetFilter } = useDataTableFilter()

    const filterVal: TSearchFilter = filters[accessorKey as string] ?? {
        dataType: 'text',
        mode: defaultMode,
        value: '',
        from: undefined,
        to: undefined,
    }

    return (
        <div className="flex items-center gap-x-1">
            <Button
                variant="secondary"
                size="sm"
                onClick={() => setVisible((prev) => !prev)}
            >
                <MagnifyingGlassIcon />
            </Button>
            {visible && (
                <span className="relative">
                    <DebouncedInput
                        className="animate-in text-xs fade-in-75"
                        value={filterVal.value}
                        placeholder={
                            otherProps.placeHolder ??
                            'Global Search (Text Only)'
                        }
                        onChange={(val) => {
                            setFilter(accessorKey, { ...filterVal, value: val })
                            bulkSetFilter(keysToSearch as string[], {
                                ...filterVal,
                                value: val,
                            })
                        }}
                    />
                    <Button
                        variant="ghost"
                        className="p-.5 absolute right-2 top-1/2 size-fit -translate-y-1/2 rounded-full"
                        onClick={() => {
                            setFilter(accessorKey)
                            bulkSetFilter(keysToSearch as string[])
                        }}
                    >
                        <XIcon className="size-4" />
                    </Button>
                </span>
            )}
        </div>
    )
}

export default DataTableGlobalSearch
