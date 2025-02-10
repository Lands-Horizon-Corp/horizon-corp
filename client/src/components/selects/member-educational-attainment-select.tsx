import {
    Select,
    SelectItem,
    SelectValue,
    SelectGroup,
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
import { IMemberEducationalAttainmentResource, TEntityId } from '@/server/types'
import { useFilteredPaginatedMemberEducationalAttainments } from '@/hooks/api-hooks/member/use-member-educational-attainment'

interface MemberEducationalAttainmentSelectProps extends IBaseCompNoChild {
    value?: TEntityId
    placeholder?: string
    disabled?: boolean
    filter?: TFilterObject
    onChange?: (selected: IMemberEducationalAttainmentResource) => void
}

const MemberEducationalAttainmentSelect: React.FC<
    MemberEducationalAttainmentSelectProps
> = ({
    value,
    placeholder = 'Select Educational Attainment',
    disabled = false,
    className,
    filter,
    onChange,
}) => {
    const { finalFilterPayload } = useFilterState({ defaultFilter: filter })

    const { data, isLoading } =
        useFilteredPaginatedMemberEducationalAttainments({
            filterPayload: finalFilterPayload,
            pagination: {
                pageIndex: 0,
                pageSize: PICKERS_SELECT_PAGE_SIZE,
            },
            enabled: !disabled,
        })

    return (
        <Select
            onValueChange={(selectedValue: TEntityId) => {
                const selectedOption = data.data.find(
                    (attainment: IMemberEducationalAttainmentResource) =>
                        attainment.id === selectedValue
                )
                if (selectedOption && onChange) {
                    onChange(selectedOption)
                }
            }}
            value={value}
            disabled={disabled || isLoading}
        >
            <SelectTrigger className={cn('', className)}>
                <SelectValue placeholder={placeholder} />
            </SelectTrigger>
            <SelectContent>
                <SelectGroup>
                    <SelectLabel>Educational Attainments</SelectLabel>
                    {isLoading && (
                        <SelectLabel className="text-sm font-normal text-foreground/70">
                            <LoadingSpinner className="inline size-3" /> Loading
                            Educational Attainments...
                        </SelectLabel>
                    )}
                    {data.data.map(
                        (attainment: IMemberEducationalAttainmentResource) => (
                            <SelectItem
                                key={attainment.id}
                                value={attainment.id}
                            >
                                {attainment.name}
                            </SelectItem>
                        )
                    )}
                </SelectGroup>
            </SelectContent>
        </Select>
    )
}

export default MemberEducationalAttainmentSelect
