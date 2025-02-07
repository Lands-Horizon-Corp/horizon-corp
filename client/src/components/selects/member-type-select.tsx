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
import { PICKERS_SELECT_PAGE_SIZE } from '@/constants'
import { useFilteredPaginatedMemberTypes } from '@/hooks/api-hooks/member/use-member-type'

import { IBaseCompNoChild } from '@/types'
import { IMemberTypeResource, TEntityId } from '@/server/types'

interface MemberTypeSelectProps extends IBaseCompNoChild {
    value?: TEntityId
    disabled?: boolean
    placeholder?: string
    onSelect?: (selectedMemberType: IMemberTypeResource) => void
}

const MemberTypeSelect = ({
    value,
    disabled,
    className,
    placeholder,
    onSelect,
}: MemberTypeSelectProps) => {
    const {
        data: { data: memberTypes },
        isLoading,
    } = useFilteredPaginatedMemberTypes({
        filterPayload: {},
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
                const selectedMemberType = memberTypes.find(
                    (mt) => mt.id === selectedValue
                )
                if (selectedMemberType && onSelect) {
                    onSelect(selectedMemberType)
                }
            }}
            disabled={disabled || isLoading}
        >
            <SelectTrigger className={cn('', className)}>
                <SelectValue
                    placeholder={placeholder || 'Select Member Type'}
                />
            </SelectTrigger>
            <SelectContent>
                <SelectGroup>
                    <SelectLabel>Member Types </SelectLabel>
                    {isLoading && (
                        <SelectLabel className="text-sm font-normal text-foreground/70">
                            <LoadingSpinner className="inline size-3" /> Loading
                            Member Types...
                        </SelectLabel>
                    )}
                    {memberTypes.map(({ id, name }) => (
                        <SelectItem key={id} value={id}>
                            {name}
                        </SelectItem>
                    ))}
                </SelectGroup>
            </SelectContent>
        </Select>
    )
}

export default MemberTypeSelect
