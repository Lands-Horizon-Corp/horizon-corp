import MobileSidebar from './mobile-sidebar'
import SidebarContent, { ISidebarProps } from './sidebar-content'

import { useSidebarContext } from './sidebar-context'

const Sidebar = (props: ISidebarProps) => {
    const { isMobile } = useSidebarContext()

    if (isMobile) return <MobileSidebar {...props} />

    return <SidebarContent {...props} />
}

export default Sidebar
