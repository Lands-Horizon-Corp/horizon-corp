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
    PiBriefcase,
    PiUsersThree,
    PiBuildingOffice,
} from 'react-icons/pi'
import {
    LiaUser,
    LiaUserClockSolid,
    LiaShieldAltSolid,
    LiaUserShieldSolid,
} from 'react-icons/lia'
import { CgDetailsMore } from 'react-icons/cg'
import { IoFootstepsOutline } from 'react-icons/io5'

import { IBaseComp } from '@/types/component/base'
import { TSidebarItem } from '@/types/component/sidebar'
import DynamicSidebar from '@/components/sidebar/dynamic-sidebar'

const ownerSidebarItem: TSidebarItem[] = [
    {
        text: 'Dashboard',
        url: '/owner/dashboard',
        Icon: HiOutlineSquaresPlus,
    },
    {
        text: 'Users',
        Icon: PiUsersThree,
        baseUrl: '/owner/users',
        subItems: [
            {
                text: 'Members',
                baseUrl: '/owner/users/members',
                Icon: PiUserList,
                subItems: [
                    {
                        text: 'View Members',
                        url: '/view-members',
                        Icon: LiaUser,
                    },
                    {
                        text: 'Members Activity',
                        url: '/members-activity',
                        Icon: LiaUserClockSolid,
                    },
                ],
            },
            {
                text: 'Employees',
                baseUrl: '/owner/users/employees',
                Icon: PiBriefcase,
                subItems: [
                    {
                        text: 'View employees',
                        url: '/view-employees',
                        Icon: LiaUserShieldSolid,
                    },
                    {
                        text: 'Employee Footsteps',
                        url: '/employee-footsteps',
                        Icon: IoFootstepsOutline,
                    },
                ],
            },
        ],
    },
    {
        text: 'Roles Management',
        Icon: LiaShieldAltSolid,
        url: '/owner/roles-management',
    },
    {
        text: 'Company',
        baseUrl: '/owner/company',
        Icon: PiBuilding,
        subItems: [
            {
                text: 'Profile',
                Icon: CgDetailsMore,
                url: '/profile',
            },
            {
                text: 'Branches',
                Icon: PiBuildingOffice,
                url: '/branches',
            },
        ],
    },
    {
        text: 'Footstep Tracking',
        Icon: IoFootstepsOutline,
        url: '/owner/footstep-tracking',
    },
    {
        text: 'Reports',
        Icon: PiNewspaper,
        url: '/owner/reports',
    },
    {
        Icon: HiOutlineBell,
        text: 'Notifications',
        url: '/owner/notifications',
    },
    {
        text: 'Profile',
        Icon: HiOutlineUser,
        url: '/owner/profile',
    },
    {
        text: 'Settings',
        Icon: HiOutlineCog,
        url: '/owner/settings',
    },
]

interface Props extends IBaseComp {}

const OwnerSidebar = ({ className }: Props) => {
    return (
        <DynamicSidebar sidebarItems={ownerSidebarItem} className={className} />
    )
}

export default OwnerSidebar
