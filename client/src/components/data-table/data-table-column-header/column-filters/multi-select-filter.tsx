import { useColumnFilterState } from './column-filter-state-context'

import MultiSelectFilter, {
    IMultiSelectOption,
} from '@/components/multi-select-filter'

const DataTableMultiSelectFilter = ({
    multiSelectOptions,
}: {
    multiSelectOptions: IMultiSelectOption[]
}) => {
    const {
        filterState: { value },
        setValue,
        clearFilter,
    } = useColumnFilterState()

    return (
        <MultiSelectFilter
            clearValues={clearFilter}
            value={value}
            setValues={setValue}
            multiSelectOptions={multiSelectOptions}
        />
    )
}

export default DataTableMultiSelectFilter
