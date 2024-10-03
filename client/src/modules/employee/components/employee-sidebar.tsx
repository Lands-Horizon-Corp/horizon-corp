import {
    HiOutlineCog,
    HiOutlineUser,
    HiOutlineBell,
    HiOutlineSquaresPlus,
} from 'react-icons/hi2'
import {
    PiUserList,
    PiBuilding,
    PiNewspaper,
    PiUsersThree,
} from 'react-icons/pi'
import { IoFootstepsOutline } from 'react-icons/io5'

import { IBaseComp } from '@/types/component/base'
import { TSidebarItem } from '@/types/component/sidebar'
import DynamicSidebar from '@/components/sidebar/dynamic-sidebar'

const employeeSidebarItems: TSidebarItem[] = [
    {
        text: 'Dashboard',
        url: '/employee/dashboard',
        Icon: HiOutlineSquaresPlus,
    },
    {
        text: 'Users',
        Icon: PiUsersThree,
        baseUrl: '/employee/users',
        subItems: [
            {
                text: 'Members',
                url: '/members',
                Icon: PiUserList,
            },
        ],
    },
    {
        text: 'Branch',
        url: '/employee/branch',
        Icon: PiBuilding,
    },
    {
        text: 'Reports',
        url: '/employee/reports',
        Icon: PiNewspaper,
    },
    {
        text: 'Notifications',
        url: '/employee/notifications',
        Icon: HiOutlineBell,
    },
    {
        text: 'Footsteps',
        url: '/employee/footsteps',
        Icon: IoFootstepsOutline,
    },
    {
        text: 'Profile',
        url: '/employee/profile',
        Icon: HiOutlineUser,
    },
    {
        text: 'Settings',
        url: '/employee/settings',
        Icon: HiOutlineCog,
    },
]

interface Props extends IBaseComp {}

const EmployeeSidebar = ({ className }: Props) => {
    return (
        <DynamicSidebar
            sidebarItems={employeeSidebarItems}
            className={className}
        />
    )
}

export default EmployeeSidebar
