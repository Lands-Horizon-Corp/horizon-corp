import { ReactNode } from '@tanstack/react-router'
import {
    Sheet,
    SheetTrigger,
    SheetContent,
} from '../ui/sheet'

interface Props { children? : ReactNode}

const SidebarMobileWrapper = ({ children }: Props) => {
    return (
        <Sheet>
            <SheetTrigger>Open</SheetTrigger>
            <SheetContent side="left" className="w-fit">
                {children}
            </SheetContent>
        </Sheet>
    )
}

export default SidebarMobileWrapper
