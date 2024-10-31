import { Button } from '@/components/ui/button'
import { SidebarLeftIcon } from '@/components/icons'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { useSidebarContext } from './sidebar-context'

const SidebarMobileToggle = ({ className }: IBaseCompNoChild) => {
    const { isExpanded, toggleMobileSidebar } = useSidebarContext()

    return (
        <Button
            variant="secondary"
            className={cn('size-fit p-2', className)}
            onClick={() => toggleMobileSidebar(!isExpanded)}
        >
            <SidebarLeftIcon />
        </Button>
    )
}

export default SidebarMobileToggle
