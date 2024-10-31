import {
    Sheet,
    SheetContent,
    SheetHeader,
    SheetTitle,
    SheetDescription,
} from '@/components/ui/sheet'
import Sidebar, { ISidebarProps } from '@/components/sidebar'

import { useSidebarContext } from './sidebar-context'

const MobileSidebar = (props: ISidebarProps) => {
    const { isMobileSidebarVisible, toggleMobileSidebar } = useSidebarContext()

    return (
        <Sheet
            open={isMobileSidebarVisible}
            onOpenChange={(newState) => toggleMobileSidebar(newState)}
        >
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
