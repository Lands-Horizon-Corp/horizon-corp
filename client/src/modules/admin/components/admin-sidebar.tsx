import {
    HiOutlineCog,
    HiOutlineUser,
    HiOutlineBell,
    HiOutlineSquaresPlus,
} from 'react-icons/hi2'
import { PiUserList, PiBuilding, PiUsersThree } from 'react-icons/pi'
import { IoFootstepsOutline } from 'react-icons/io5'

import { IBaseComp } from '@/types/component/base'
import { TSidebarItem } from '@/types/component/sidebar'
import DynamicSidebar from '@/components/sidebar/dynamic-sidebar'

const adminSidebarItems: TSidebarItem[] = [
    {
        text: 'Dashboard',
        url: '/admin/dashboard',
        Icon: HiOutlineSquaresPlus,
    },
    {
        text: 'Members Management',
        Icon: PiUsersThree,
        baseUrl: '/admin/members-management',
        subItems: [
            {
                text: 'View Members',
                url: '/view-members',
                Icon: PiUserList,
            },
            {
                text: 'Feedbacks',
                url: '/feedbacks',
                Icon: HiOutlineBell,
            },
        ],
    },
    {
        text: 'Companies Management',
        Icon: PiBuilding,
        baseUrl: '/admin/companies-management',
        subItems: [
            {
                text: 'View Companies',
                url: '/view-companies',
                Icon: PiBuilding,
            },
            {
                text: 'Feedbacks',
                url: '/feedbacks',
                Icon: HiOutlineBell,
            },
        ],
    },
    {
        text: 'Footstep Tracking',
        url: '/admin/footstep-tracking',
        Icon: IoFootstepsOutline,
    },
    {
        text: 'Profile',
        url: '/admin/profile',
        Icon: HiOutlineUser,
    },
    {
        text: 'Notifications',
        url: '/admin/notifications',
        Icon: HiOutlineBell,
    },
    {
        text: 'Settings',
        url: '/admin/settings',
        Icon: HiOutlineCog,
    },
]

interface Props extends IBaseComp {}

const AdminSidebar = ({ className }: Props) => {
    return (
        <DynamicSidebar
            sidebarItems={adminSidebarItems}
            className={className}
        />
    )
}

export default AdminSidebar
