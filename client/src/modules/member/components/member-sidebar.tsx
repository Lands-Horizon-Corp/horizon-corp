import {
    HiOutlineCog,
    HiOutlineUser,
    HiOutlineSquaresPlus,
    HiOutlineBell,
    HiOutlineChatBubbleOvalLeft,
} from 'react-icons/hi2'

import { IBaseComp } from '@/types/component/base'
import { TSidebarItem } from '@/types/component/sidebar'
import DynamicSidebar from '@/components/sidebar/dynamic-sidebar'

const memberSidebarItems: TSidebarItem[] = [
    {
        text: 'Dashboard',
        url: '/member/dashboard',
        Icon: HiOutlineSquaresPlus,
    },
    {
        text: 'Profile',
        url: '/member/profile',
        Icon: HiOutlineUser,
    },
    {
        text: 'Notifications',
        url: '/member/notifications',
        Icon: HiOutlineBell,
    },
    {
        text: 'Messages',
        url: '/member/messages',
        Icon: HiOutlineChatBubbleOvalLeft,
    },
    {
        text: 'Settings',
        url: '/member/settings',
        Icon: HiOutlineCog,
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
