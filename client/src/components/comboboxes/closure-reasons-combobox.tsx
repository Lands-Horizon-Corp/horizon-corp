import * as React from 'react'
import { Check } from 'lucide-react'

import {
    Command,
    CommandItem,
    CommandList,
    CommandEmpty,
    CommandGroup,
    CommandInput,
} from '@/components/ui/command'
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import { ChevronDownIcon } from '../icons'
import { Button } from '@/components/ui/button'

import { cn } from '@/lib/utils'
import { AccountClosureReasonType, AccountClosureReasonTypes } from '@/server'

interface Props {
    id?: string
    name?: string
    value?: AccountClosureReasonType
    disabled?: boolean
    className?: string
    placeholder?: string
    closureReasons?: string[]
    onChange?: (selected: AccountClosureReasonType) => void
}

const AccountClosureReasonCombobox = ({
    value,
    className,
    disabled = false,
    placeholder = 'Select Closure Reason...',
    closureReasons = AccountClosureReasonTypes as unknown as AccountClosureReasonType[],
    onChange,
    ...other
}: Props) => {
    const [open, setOpen] = React.useState(false)

    return (
        <Popover modal open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
                <Button
                    {...other}
                    type="button"
                    role="combobox"
                    variant="outline"
                    disabled={disabled}
                    aria-expanded={open}
                    className={cn('w-full justify-between px-3', className)}
                >
                    <span className="capitalize">{value || placeholder}</span>
                    <ChevronDownIcon className="opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="max-h-[--radix-popover-content-available-height] w-[--radix-popover-trigger-width] p-0">
                <Command>
                    <CommandInput
                        placeholder="Search Closure Reason..."
                        className="h-9"
                    />
                    <CommandList className="ecoop-scroll">
                        <CommandEmpty>No reason found.</CommandEmpty>
                        <CommandGroup>
                            {closureReasons.map((reason) => (
                                <CommandItem
                                    key={reason}
                                    value={reason}
                                    onSelect={() => {
                                        setOpen(false)
                                        onChange?.(
                                            reason as unknown as AccountClosureReasonType
                                        )
                                    }}
                                >
                                    <span className="capitalize">{reason}</span>
                                    <Check
                                        className={cn(
                                            'ml-auto',
                                            value === reason
                                                ? 'opacity-100'
                                                : 'opacity-0'
                                        )}
                                    />
                                </CommandItem>
                            ))}
                        </CommandGroup>
                    </CommandList>
                </Command>
            </PopoverContent>
        </Popover>
    )
}

export default AccountClosureReasonCombobox
