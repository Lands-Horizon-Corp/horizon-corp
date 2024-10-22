import {
    UserIcon,
    SettingsIcon,
    MessagesIcon,
    DashboardIcon,
    NotificationIcon,
} from '@/components/icons'
import DynamicSidebar from '@/components/sidebar/dynamic-sidebar'

import { IBaseComp } from '@/types/component'
import { TSidebarItem } from '@/types/component/sidebar'

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

interface Props extends IBaseComp {}

const MemberSidebar = ({ className }: Props) => {
    return (
        <DynamicSidebar
            sidebarItems={memberSidebarItems}
            className={className}
        />
    )
}

export default MemberSidebar
