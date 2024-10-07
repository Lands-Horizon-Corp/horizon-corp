import { ReactNode } from 'react'

import {
    Sheet,
    SheetTrigger,
    SheetContent,
    SheetHeader,
    SheetTitle,
    SheetDescription,
} from '@/components/ui/sheet'
import Sidebar, { ISidebarProps } from '@/components/sidebar'

interface Props extends ISidebarProps {
    triggerComponent: ReactNode
}

const MobileSidebar = ({ triggerComponent, ...props }: Props) => {
    return (
        <Sheet>
            <SheetTrigger asChild>{triggerComponent}</SheetTrigger>
            <SheetContent side="left" className="h-fit w-fit p-0">
                <SheetHeader className="hidden">
                    <SheetTitle>Sidebar</SheetTitle>
                    <SheetDescription>Mobile sidebar</SheetDescription>
                </SheetHeader>
                <Sidebar {...props} />
            </SheetContent>
        </Sheet>
    )
}

export default MobileSidebar
