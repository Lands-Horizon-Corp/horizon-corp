import {
    TSearchFilter,
    useDataTableFilter,
    IFilterComponentProps,
    TColumnDataTypes,
} from './data-table-filter-context'

import MultiSelectFilter, {
    IMultiSelectOption,
} from '@/components/multi-select-filter'

interface IDatatableMultiFilter<T, TValue> extends IFilterComponentProps<T> {
    multiSelectOptions: IMultiSelectOption<TValue>[]
    dataType: TColumnDataTypes
}

const DataTableMultiSelectFilter = <TData, TValue>({
    field,
    dataType,
    displayText,
    multiSelectOptions,
}: IDatatableMultiFilter<TData, TValue>) => {
    const { filters, setFilter } = useDataTableFilter<
        string,
        typeof field,
        TValue[]
    >()

    const filterVal: TSearchFilter<string, TValue[]> = filters[field] ?? {
        displayText,
        mode: 'equal',
        to: undefined,
        from: undefined,
        dataType,
        value: undefined,
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
                setFilter(field, {
                    ...filterVal,
                    value: selected,
                })
            }
            clearValues={() => setFilter(field)}
        />
    )
}

export default DataTableMultiSelectFilter
