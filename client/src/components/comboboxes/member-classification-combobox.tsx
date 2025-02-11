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
} from '@/components/ui/command'
import { cn } from '@/lib/utils'
import { ChevronDownIcon } from '../icons'
import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { TEntityId, IMemberClassificationResource } from '@/server/types'
import { useFilteredPaginatedMemberClassifications } from '@/hooks/api-hooks/member/use-member-classification'

interface Props {
    value?: TEntityId
    placeholder?: string
    disabled?: boolean
    onChange?: (selected: IMemberClassificationResource) => void
    className?: string
}

const MemberClassificationCombobox = ({
    value,
    placeholder = 'Select Member Classification...',
    disabled = false,
    onChange,
    className,
}: Props) => {
    const [open, setOpen] = React.useState(false)

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
    )
}

export default MemberClassificationCombobox
