import { Check } from 'lucide-react'
import { useMemo, useState } from 'react'

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

type TGovernmentBenefitItem = {
    text: string
    value: string
}

interface Props {
    id?: string
    name?: string
    value?: string
    disabled?: boolean
    className?: string
    placeholder?: string
    governmentBenefitsList: TGovernmentBenefitItem[]
    onChange?: (selected: string) => void
}

const GovernmentBenefitsCombobox = ({
    value,
    className,
    disabled = false,
    governmentBenefitsList,
    placeholder = 'Select Government Id...',
    onChange,
    ...other
}: Props) => {
    const [open, setOpen] = useState(false)

    const selectedItem = useMemo(
        () => governmentBenefitsList.find((govBen) => govBen.value === value),
        [value, governmentBenefitsList]
    )

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
                    <span className="capitalize">
                        {selectedItem?.text || placeholder}
                    </span>
                    <ChevronDownIcon className="opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="max-h-[--radix-popover-content-available-height] w-[--radix-popover-trigger-width] p-0">
                <Command>
                    <CommandInput
                        placeholder="Search Barangay..."
                        className="h-9"
                    />
                    <CommandList className="ecoop-scroll">
                        <CommandEmpty>No Barangay found.</CommandEmpty>
                        <CommandGroup>
                            {governmentBenefitsList.map((govBenefits) => (
                                <CommandItem
                                    key={govBenefits.value}
                                    value={govBenefits.value}
                                    onSelect={() => {
                                        setOpen(false)
                                        onChange?.(govBenefits.value)
                                    }}
                                >
                                    <span className="capitalize">
                                        {govBenefits.text}
                                    </span>
                                    <Check
                                        className={cn(
                                            'ml-auto',
                                            value === govBenefits.value
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

export default GovernmentBenefitsCombobox

// NOTE: THE CODE BELLOW COMPOSE OF INPUT & COMBOBOX BARANGAY
// This enables the user to input barangay if not in choices
// for now it not neede, incase comment or suggestion about this arises the code below is
// ideal
//
// import * as React from 'react'
// import { Check } from 'lucide-react'

// import {
//     Popover,
//     PopoverContent,
//     PopoverTrigger,
// } from '@/components/ui/popover'
// import {
//     Command,
//     CommandItem,
//     CommandList,
//     CommandEmpty,
//     CommandGroup,
//     CommandInput,
// } from '@/components/ui/command'
// import { cn } from '@/lib/utils'
// import { ChevronDownIcon } from '../icons'
// import { Input } from '@/components/ui/input'
// import { Button } from '@/components/ui/button'

// import { getBarangaysByMunicipality } from '@/helpers/address'

// interface Props {
//     id?: string
//     name?: string
//     value?: string
//     disabled?: boolean
//     className?: string
//     placeholder?: string
//     municipality?: string
//     onChange?: (selected: string) => void
// }

// const InputBarangayCombobox = ({
//     value,
//     className,
//     municipality,
//     disabled = false,
//     placeholder = 'Select Barangay...',
//     onChange,
//     ...other
// }: Props) => {
//     const [open, setOpen] = React.useState(false)

//     const barangays = React.useMemo(() => {
//         return municipality ? getBarangaysByMunicipality(municipality) : []
//     }, [municipality])

//     return (
//         <div className="relative">
//             <Input
//                 {...other}
//                 value={value || ''}
//                 placeholder={placeholder}
//                 className={cn('w-full pr-10', className)}
//                 onChange={(e) => onChange?.(e.target.value)}
//             />
//             <Popover open={open} onOpenChange={setOpen}>
//                 <PopoverTrigger asChild>
//                     <Button
//                         type="button"
//                         variant="ghost"
//                         disabled={disabled}
//                         className="absolute top-1/2 -translate-y-1/2 size-fit p-1 right-2 flex items-center"
//                     >
//                         <ChevronDownIcon className="opacity-50" />
//                     </Button>
//                 </PopoverTrigger>
//                 <PopoverContent className="p-0">
//                     <Command>
//                         <CommandInput
//                             placeholder="Search Barangay..."
//                             className="h-9"
//                         />
//                         <CommandList className="ecoop-scroll">
//                             <CommandEmpty>No Barangay found.</CommandEmpty>
//                             <CommandGroup>
//                                 {barangays.map((barangay) => (
//                                     <CommandItem
//                                         key={barangay}
//                                         value={barangay}
//                                         onSelect={() => {
//                                             setOpen(false)
//                                             onChange?.(barangay)
//                                         }}
//                                     >
//                                         {barangay}
//                                         <Check
//                                             className={cn(
//                                                 'ml-auto',
//                                                 value === barangay
//                                                     ? 'opacity-100'
//                                                     : 'opacity-0'
//                                             )}
//                                         />
//                                     </CommandItem>
//                                 ))}
//                             </CommandGroup>
//                         </CommandList>
//                     </Command>
//                 </PopoverContent>
//             </Popover>
//         </div>
//     )
// }

// export default InputBarangayCombobox
