import { TSearchFilter, useDataTableFilter } from './data-table-filter-context'

import MultiSelectFilter, {
    IMultiSelectOption,
} from '@/components/multi-select-filter'

const DataTableMultiSelectFilter = <TData,>({
    multiSelectOptions,
    accessorKey,
}: {
    accessorKey: keyof TData
    multiSelectOptions: IMultiSelectOption[]
}) => {
    const { filters, setFilter } = useDataTableFilter()

    const filterVal: TSearchFilter = filters[accessorKey as string] ?? {
        dataType: 'date',
        mode: 'equal',
        value: undefined,
        from: undefined,
        to: undefined,
    }

    return (
        <MultiSelectFilter
            value={
                typeof filterVal.value === 'string'
                    ? [filterVal.value]
                    : filterVal.value
            }
            multiSelectOptions={multiSelectOptions}
            setValues={(selected) =>
                setFilter(accessorKey as string, {
                    ...filterVal,
                    value: selected,
                })
            }
            clearValues={() => setFilter(accessorKey as string)}
        />
    )
}

export default DataTableMultiSelectFilter
