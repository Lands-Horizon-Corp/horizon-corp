import Sidebar from '.'
import SidebarMobileWrapper from './sidebar-mobile-wrapper'

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
                <SidebarMobileWrapper>
                    <Sidebar
                        enableFocusBlur
                        items={sidebarItems}
                        className={cn('', className)}
                    />
                </SidebarMobileWrapper>
            </div>
        </>
    )
}

export default DynamicSidebar
