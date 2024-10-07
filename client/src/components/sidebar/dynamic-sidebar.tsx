import Sidebar from '@/components/sidebar'
import { Button } from '@/components/ui/button'
import { SidebarIcon } from '@/components/icons'
import MobileSidebar from '@/components/sidebar/mobile-sidebar'

import { cn } from '@/lib/utils'
import { TSidebarItem } from '@/types/component/sidebar'

interface Props {
    className?: string
    sidebarItems: TSidebarItem[]
}

const DynamicSidebar = ({ className, sidebarItems }: Props) => {
    return (
        <>
            <div className="hidden sm:block">
                <Sidebar
                    enableCollapse
                    enableFocusBlur
                    items={sidebarItems}
                    className={cn('', className)}
                />
            </div>
            <div className="block sm:hidden">
                <MobileSidebar
                    items={sidebarItems}
                    triggerComponent={
                        <Button
                            size="icon"
                            variant="secondary"
                            className="fixed left-2 top-2 size-fit rounded-lg p-2"
                        >
                            <SidebarIcon className="size-4" />
                        </Button>
                    }
                    className={cn('', className)}
                />
            </div>
        </>
    )
}

export default DynamicSidebar
