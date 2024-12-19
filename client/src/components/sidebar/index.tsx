import MobileSidebar from './mobile-sidebar'
import SidebarContent, { ISidebarProps } from './sidebar-content'

import { useIsMobile } from '@/hooks/use-mobile'

const Sidebar = (props: ISidebarProps) => {
    const isMobile = useIsMobile()

    if (isMobile) return <MobileSidebar {...props} />

    return <SidebarContent {...props} />
}

export default Sidebar
