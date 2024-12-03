import {
    IFilterComponentProps,
    TSearchFilter,
    useDataTableFilter,
} from './data-table-filter-context'

import MultiSelectFilter, {
    IMultiSelectOption,
} from '@/components/multi-select-filter'

interface IDatatableMultiFilter extends IFilterComponentProps {
    multiSelectOptions: IMultiSelectOption[]
}

const DataTableMultiSelectFilter = ({
    field,
    displayText,
    multiSelectOptions,
}: IDatatableMultiFilter) => {
    const { filters, setFilter } = useDataTableFilter<
        string,
        string,
        string[]
    >()

    const filterVal: TSearchFilter<string, string[]> = filters[field] ?? {
        displayText,
        mode: 'equal',
        to: undefined,
        from: undefined,
        dataType: 'enum',
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
