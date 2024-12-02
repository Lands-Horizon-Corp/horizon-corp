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
    const { filters, setFilter } = useDataTableFilter<
        string,
        keyof TData,
        string[]
    >()

    const filterVal: TSearchFilter<string, string[]> = filters[
        accessorKey as string
    ] ?? {
        dataType: 'enum',
        mode: 'equal',
        value: undefined,
        from: undefined,
        to: undefined,
    }

    return (
        <MultiSelectFilter
            value={
                filterVal.value
                    ? typeof filterVal.value === 'string'
                        ? [filterVal.value]
                        : filterVal.value
                    : []
            }
            multiSelectOptions={multiSelectOptions}
            setValues={(selected) =>
                setFilter(accessorKey, {
                    ...filterVal,
                    value: selected,
                })
            }
            clearValues={() => setFilter(accessorKey)}
        />
    )
}

export default DataTableMultiSelectFilter
