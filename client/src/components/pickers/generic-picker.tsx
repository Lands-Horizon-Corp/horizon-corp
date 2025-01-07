import React from 'react'

import {
    Command,
    CommandItem,
    CommandList,
    CommandEmpty,
    CommandGroup,
    CommandInput,
} from '@/components/ui/command'
import Modal, { IModalProps } from '@/components/modals/modal'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib'
import { useInternalState } from '@/hooks/use-internal-state'

interface GenericPickerProps<T extends { id: number }> extends IModalProps {
    items: T[]
    isLoading?: boolean
    listHeading?: string
    commandClassName?: string
    searchPlaceHolder?: string
    onSelect?: (item: T) => void
    onSearchChange: (val: string) => void
    renderItem: (item: T) => React.ReactNode
}

const GenericPicker = <T extends { id: number }>({
    open,
    items,
    children,
    isLoading,
    listHeading,
    className,
    commandClassName,
    searchPlaceHolder,
    onSelect,
    renderItem,
    onOpenChange,
    onSearchChange,
    ...other
}: GenericPickerProps<T>) => {
    const [modalState, setModalState] = useInternalState(
        open,
        onOpenChange,
        false
    )

    return (
        <Modal
            {...other}
            open={modalState}
            onOpenChange={setModalState}
            titleClassName="!hidden"
            closeButtonClassName="hidden"
            descriptionClassName="!hidden"
            className={cn(
                '!h-fit max-w-[90vw] !gap-y-0 border p-0 shadow-none backdrop-blur-none sm:max-w-2xl',
                className
            )}
        >
            <Command
                shouldFilter={false}
                className={cn('bg-none', commandClassName)}
            >
                <CommandInput
                    onValueChange={onSearchChange}
                    placeholder={searchPlaceHolder ?? 'Search anything...'}
                />
                <CommandList className="ecoop-scroll max-h-[300px] min-h-[400px] px-1">
                    <CommandEmpty className="text-sm text-foreground/50">
                        {isLoading ? (
                            <LoadingSpinner className="inline" />
                        ) : (
                            'No Result'
                        )}
                    </CommandEmpty>
                    {items.length > 0 && (
                        <CommandGroup heading={listHeading}>
                            {items?.map((item) => (
                                <CommandItem
                                    key={item.id}
                                    onSelect={() => {
                                        onSelect?.(item)
                                        setModalState(false)
                                    }}
                                    className="cursor-pointer rounded-lg"
                                >
                                    {renderItem(item)}
                                </CommandItem>
                            ))}
                        </CommandGroup>
                    )}
                </CommandList>
            </Command>
            {children}
        </Modal>
    )
}

export default GenericPicker
