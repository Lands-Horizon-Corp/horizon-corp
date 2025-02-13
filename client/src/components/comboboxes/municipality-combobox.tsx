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

import { getMunicipalitiesByProvince } from '@/helpers/address'

interface Props {
    id?: string
    name?: string
    value?: string
    disabled?: boolean
    className?: string
    placeholder?: string
    province?: string
    onChange?: (selected: string) => void
}

const MunicipalityCombobox = ({
    value,
    className,
    province,
    disabled = false,
    placeholder = 'Select Municipality...',
    onChange,
    ...other
}: Props) => {
    const [open, setOpen] = React.useState(false)

    const municipalities = React.useMemo(() => {
        return province ? getMunicipalitiesByProvince(province) : []
    }, [province])

    return (
        <Popover open={open} onOpenChange={setOpen}>
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
                    <span className="capitalize">
                        {value?.toLocaleLowerCase() || placeholder}
                    </span>
                    <ChevronDownIcon className="opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="max-h-[--radix-popover-content-available-height] w-[--radix-popover-trigger-width] p-0">
                <Command>
                    <CommandInput
                        placeholder="Search Municipality..."
                        className="h-9"
                    />
                    <CommandList className="ecoop-scroll">
                        <CommandEmpty>No Municipality found.</CommandEmpty>
                        <CommandGroup>
                            {municipalities.map((municipality) => (
                                <CommandItem
                                    key={municipality}
                                    value={municipality}
                                    onSelect={() => {
                                        setOpen(false)
                                        onChange?.(municipality)
                                    }}
                                >
                                    {municipality}
                                    <Check
                                        className={cn(
                                            'ml-auto',
                                            value === municipality
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

export default MunicipalityCombobox
