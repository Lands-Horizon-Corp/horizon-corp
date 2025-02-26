import { useState } from 'react'
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
} from '@/components/ui/command'
import { cn } from '@/lib/utils'
import { ChevronDownIcon, PlusIcon } from '../icons'
import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import {
    IMemberClassificationCreateUpdateFormProps,
    MemberClassificationCreateUpdateFormModal,
} from '../forms/member-forms/member-classification-create-update-form'

import { TEntityId, IMemberClassificationResource } from '@/server/types'
import { useFilteredPaginatedMemberClassifications } from '@/hooks/api-hooks/member/use-member-classification'

export interface IMemberClassificationComboboxCreateProps
    extends Pick<
        IMemberClassificationCreateUpdateFormProps,
        'defaultValues' | 'disabledFields' | 'hiddenFields'
    > {}

interface Props {
    value?: TEntityId
    disabled?: boolean
    className?: string
    placeholder?: string
    memberClassificationCreateProps?: IMemberClassificationComboboxCreateProps
    onChange?: (selected: IMemberClassificationResource) => void
}

const MemberClassificationCombobox = ({
    value,
    placeholder = 'Select Member Classification...',
    disabled = false,
    className,
    memberClassificationCreateProps,
    onChange,
}: Props) => {
    const [open, setOpen] = useState(false)
    const [createModal, setCreateModal] = useState(false)

    const {
        data: { data },
        isLoading,
    } = useFilteredPaginatedMemberClassifications({
        filterPayload: {},
        pagination: { pageIndex: 0, pageSize: 100 },
        enabled: !disabled,
        showMessage: false,
    })

    return (
        <>
            <MemberClassificationCreateUpdateFormModal
                open={createModal}
                onOpenChange={setCreateModal}
                formProps={{
                    ...memberClassificationCreateProps,
                    onSuccess: (data) => {
                        onChange?.(data)
                        setOpen(false)
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
                            placeholder="Search Member Classification..."
                            className="h-9"
                        />
                        {isLoading ? (
                            <CommandEmpty>
                                <LoadingSpinner className="mr-2 inline-block" />{' '}
                                Loading...
                            </CommandEmpty>
                        ) : (
                            <CommandList>
                                <CommandEmpty>
                                    No Member Classification found.
                                </CommandEmpty>
                                {memberClassificationCreateProps && (
                                    <CommandGroup>
                                        <CommandItem
                                            onClick={(e) => e.stopPropagation()}
                                            onSelect={() =>
                                                setCreateModal(true)
                                            }
                                        >
                                            <PlusIcon /> Create Member
                                            Classification
                                        </CommandItem>
                                    </CommandGroup>
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

export default MemberClassificationCombobox
