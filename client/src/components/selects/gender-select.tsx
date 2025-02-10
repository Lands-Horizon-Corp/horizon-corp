import {
    Select,
    SelectItem,
    SelectGroup,
    SelectValue,
    SelectLabel,
    SelectContent,
    SelectTrigger,
} from '@/components/ui/select'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import useFilterState from '@/hooks/use-filter-state'
import { PICKERS_SELECT_PAGE_SIZE } from '@/constants'
import { TFilterObject } from '@/contexts/filter-context'
import { IGenderResource, TEntityId } from '@/server/types'

import { useFilteredPaginatedGenders } from '@/hooks/api-hooks/use-gender'

interface GenderSelectProps extends IBaseCompNoChild {
    value?: TEntityId
    disabled?: boolean
    placeholder?: string
    filter?: TFilterObject
    onChange?: (selectedGender: IGenderResource) => void
}

const GenderSelect = ({
    value,
    filter,
    disabled,
    className,
    placeholder,
    onChange,
}: GenderSelectProps) => {
    const { finalFilterPayload } = useFilterState({ defaultFilter: filter })

    const {
        data: { data: genders },
        isLoading,
    } = useFilteredPaginatedGenders({
        filterPayload: finalFilterPayload,
        pagination: {
            pageIndex: 0,
            pageSize: PICKERS_SELECT_PAGE_SIZE,
        },
        enabled: !disabled,
    })

    return (
        <Select
            value={value}
            onValueChange={(selectedValue: TEntityId) => {
                const selectedGender = genders.find(
                    (g) => g.id === selectedValue
                )
                if (selectedGender && onChange) {
                    onChange(selectedGender)
                }
            }}
            disabled={disabled || isLoading}
        >
            <SelectTrigger className={cn('', className)}>
                <SelectValue placeholder={placeholder || 'Select Gender'} />
            </SelectTrigger>
            <SelectContent>
                <SelectGroup>
                    <SelectLabel>Genders</SelectLabel>
                    {isLoading && (
                        <SelectLabel className="text-sm font-normal text-foreground/70">
                            <LoadingSpinner className="inline size-3" /> Loading
                            Genders...
                        </SelectLabel>
                    )}
                    {genders.map(({ id, name }) => (
                        <SelectItem key={id} value={id}>
                            {name}
                        </SelectItem>
                    ))}
                </SelectGroup>
            </SelectContent>
        </Select>
    )
}

export default GenderSelect
