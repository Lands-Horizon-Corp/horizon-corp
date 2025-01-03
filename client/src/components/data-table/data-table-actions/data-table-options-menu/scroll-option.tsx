import {
    DropdownMenuGroup,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuRadioGroup,
    DropdownMenuRadioItem,
} from '@/components/ui/dropdown-menu'

export interface DataTableScrollableOptionProps {
    isScrollable: boolean
    setIsScrollable: (val: boolean) => void
}

const DataTableScrollOption = ({
    isScrollable,
    setIsScrollable,
}: DataTableScrollableOptionProps) => {
    return (
        <DropdownMenuGroup>
            <DropdownMenuLabel>Table Display</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuRadioGroup
                value={isScrollable ? 'true' : 'false'}
                onValueChange={(newVal) =>
                    setIsScrollable(newVal === 'true' ? true : false)
                }
            >
                <DropdownMenuRadioItem value="true">
                    Default (Scroll)
                </DropdownMenuRadioItem>
                <DropdownMenuRadioItem value="false">
                    Full (No Scroll)
                </DropdownMenuRadioItem>
            </DropdownMenuRadioGroup>
        </DropdownMenuGroup>
    )
}

export default DataTableScrollOption
