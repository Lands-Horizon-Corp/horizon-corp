import * as React from 'react'
import { Check } from 'lucide-react'

import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import {
    Command,
    CommandItem,
    CommandList,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandSeparator,
} from '@/components/ui/command'
import { cn } from '@/lib/utils'
import { ChevronDownIcon, PlusIcon } from '../icons'
import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { PICKERS_SELECT_PAGE_SIZE } from '@/constants'
import { TEntityId, IMemberOccupationResource } from '@/server/types'
import { useFilteredPaginatedMemberOccupations } from '@/hooks/api-hooks/member/use-member-occupation'
import {
    IMemberOccupationCreateUpdateFormProps,
    MemberOccupationCreateUpdateFormModal,
} from '../forms/member-forms/member-occupation-create-update-form'

export interface IMemberOccupationComboboxCreateProps
    extends Pick<
        IMemberOccupationCreateUpdateFormProps,
        'defaultValues' | 'disabledFields' | 'hiddenFields'
    > {}

interface Props {
    value?: TEntityId
    disabled?: boolean
    className?: string
    placeholder?: string
    memberOccupationComboboxCreateProps?: IMemberOccupationComboboxCreateProps
    onChange?: (selected: IMemberOccupationResource) => void
}

const MemberOccupationCombobox = ({
    value,
    className,
    disabled = false,
    memberOccupationComboboxCreateProps,
    placeholder = 'Select Member Occupation...',
    onChange,
}: Props) => {
    const [open, setOpen] = React.useState(false)
    const [createModal, setCreateModal] = React.useState(false)

    const {
        data: { data },
        isLoading,
    } = useFilteredPaginatedMemberOccupations({
        filterPayload: {},
        pagination: { pageIndex: 0, pageSize: PICKERS_SELECT_PAGE_SIZE },
        enabled: !disabled,
        showMessage: false,
    })

    return (
        <>
            <MemberOccupationCreateUpdateFormModal
                open={createModal}
                onOpenChange={setCreateModal}
                formProps={{
                    ...memberOccupationComboboxCreateProps,
                    onSuccess: (newOccupation) => {
                        onChange?.(newOccupation)
                        setCreateModal(false)
                    },
                }}
            />
            <Popover open={open} onOpenChange={setOpen}>
                <PopoverTrigger asChild>
                    <Button
                        variant="outline"
                        role="combobox"
                        aria-expanded={open}
                        className={cn('w-full justify-between px-3', className)}
                        disabled={disabled || isLoading}
                    >
                        {value
                            ? data.find((option) => option.id === value)?.name
                            : placeholder}
                        <ChevronDownIcon className="opacity-50" />
                    </Button>
                </PopoverTrigger>
                <PopoverContent className="max-h-[--radix-popover-content-available-height] w-[--radix-popover-trigger-width] p-0">
                    <Command>
                        <CommandInput
                            placeholder="Search Member Occupation..."
                            className="h-9"
                        />
                        {isLoading ? (
                            <CommandEmpty>
                                <LoadingSpinner className="mr-2 inline-block" />{' '}
                                Loading...
                            </CommandEmpty>
                        ) : (
                            <CommandList className="ecoop-scroll">
                                <CommandEmpty>
                                    No Member Occupation found.
                                </CommandEmpty>
                                {memberOccupationComboboxCreateProps && (
                                    <>
                                        <CommandGroup>
                                            <CommandItem
                                                onSelect={() => {
                                                    setCreateModal(true)
                                                }}
                                                onClick={() => {}}
                                            >
                                                <PlusIcon /> Create new
                                                occupation occupation
                                            </CommandItem>
                                        </CommandGroup>
                                        <CommandSeparator />
                                    </>
                                )}
                                <CommandGroup>
                                    {data.map((option) => (
                                        <CommandItem
                                            key={option.id}
                                            value={option.name}
                                            onSelect={() => {
                                                setOpen(false)
                                                onChange?.(option)
                                            }}
                                        >
                                            {option.name}
                                            <Check
                                                className={cn(
                                                    'ml-auto',
                                                    value === option.id
                                                        ? 'opacity-100'
                                                        : 'opacity-0'
                                                )}
                                            />
                                        </CommandItem>
                                    ))}
                                </CommandGroup>
                            </CommandList>
                        )}
                    </Command>
                </PopoverContent>
            </Popover>
        </>
    )
}

export default MemberOccupationCombobox
