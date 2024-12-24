import {
    UserIcon,
    SettingsIcon,
    MessagesIcon,
    DashboardIcon,
    NotificationIcon,
} from '@/components/icons'
import Sidebar from '@/components/sidebar'

import { IBaseComp } from '@/types/component'
import { TSidebarItem } from '@/components/sidebar/sidebar-types'

const memberSidebarItems: TSidebarItem[] = [
    {
        text: 'Dashboard',
        url: '/member/dashboard',
        Icon: DashboardIcon,
    },
    {
        text: 'Profile',
        url: '/member/profile',
        Icon: UserIcon,
    },
    {
        text: 'Notifications',
        url: '/member/notifications',
        Icon: NotificationIcon,
    },
    {
        text: 'Messages',
        url: '/member/messages',
        Icon: MessagesIcon,
    },
    {
        text: 'Settings',
        url: '/member/settings',
        Icon: SettingsIcon,
    },
]

const MemberSidebar = ({ className }: IBaseComp) => {
    return (
        <Sidebar
            enableCollapse
            className={className}
            items={memberSidebarItems}
        />
    )
}

export default MemberSidebar
