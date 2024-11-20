import { CheckIcon } from 'lucide-react'; 

import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
    CommandSeparator,
} from '@/components/ui/command';
import { useColumnFilterState } from './column-filter-state-context';

import { cn } from '@/lib/utils'; 

interface IMultiSelectOption {
    label: string;
    value: string;
}

const MultiSelectFilter = ({
    multiSelectOptions,
}: {
    multiSelectOptions: IMultiSelectOption[];
}) => {
    const {
        filterState: { value },
        setValue,
    } = useColumnFilterState();

    const selectedValues = new Set(value);

    return (
        <div onKeyDown={(e) => e.stopPropagation()} className="space-y-2 p-1">
            <p className="text-sm">Filter</p>
            <Command className="w-fit bg-transparent">
                <CommandInput placeholder="Search filters..." />
                <CommandList>
                    <CommandEmpty className="px-8 py-6 text-xs text-center">
                        No results found.
                    </CommandEmpty>
                    <CommandGroup>
                        {multiSelectOptions.map((option) => {
                            const isSelected = selectedValues.has(option.value);
                            return (
                                <CommandItem
                                    key={option.value}
                                    onSelect={() => {
                                        if (isSelected) {
                                            selectedValues.delete(option.value);
                                        } else {
                                            selectedValues.add(option.value);
                                        }
                                        const filterValues = Array.from(
                                            selectedValues
                                        );
                                        setValue(
                                            filterValues.length
                                                ? filterValues
                                                : []
                                        );
                                    }}
                                >
                                    <div
                                        className={cn(
                                            'mr-2 flex h-4 w-4 items-center justify-center rounded-sm border border-primary',
                                            isSelected
                                                ? 'bg-primary text-primary-foreground'
                                                : 'opacity-50 [&_svg]:invisible'
                                        )}
                                    >
                                        <CheckIcon
                                            className={cn('h-4 w-4')}
                                        />
                                    </div>
                                    <span>{option.label}</span>
                                </CommandItem>
                            );
                        })}
                    </CommandGroup>
                    {selectedValues.size > 0 && (
                        <>
                            <CommandSeparator />
                            <CommandGroup>
                                <CommandItem
                                    onSelect={() => setValue([])}
                                    className="justify-center text-center"
                                >
                                    Clear filters
                                </CommandItem>
                            </CommandGroup>
                        </>
                    )}
                </CommandList>
            </Command>
        </div>
    );
};

export default MultiSelectFilter;

