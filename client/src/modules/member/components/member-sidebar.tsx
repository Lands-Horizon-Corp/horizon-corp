import Sidebar from '@/components/sidebar'

import {
    HiOutlineCog,
    HiOutlineUser,
    HiOutlineSquaresPlus,
    HiOutlineBell,
    HiOutlineChatBubbleOvalLeft,
} from 'react-icons/hi2'

import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component/base'
import { ISidebarItem } from '@/types/component/sidebar'

const memberSidebarItems: ISidebarItem[] = [
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
        <Sidebar
            enableCollapse
            items={memberSidebarItems}
            className={cn('', className)}
        />
    )
}

export default MemberSidebar
