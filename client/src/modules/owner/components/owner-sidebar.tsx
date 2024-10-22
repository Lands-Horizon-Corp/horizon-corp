import {
    UserIcon,
    ShieldIcon,
    Users3Icon,
    ReportsIcon,
    BuildingIcon,
    SettingsIcon,
    UserListIcon,
    DashboardIcon,
    FootstepsIcon,
    UserClockIcon,
    UserShieldIcon,
    BuildingCogIcon,
    NotificationIcon,
    BuildingBranchIcon,
} from '@/components/icons'
import DynamicSidebar from '@/components/sidebar/dynamic-sidebar'

import { IBaseComp } from '@/types/component'
import { TSidebarItem } from '@/types/component/sidebar'

const ownerSidebarItem: TSidebarItem[] = [
    {
        text: 'Dashboard',
        url: '/owner/dashboard',
        Icon: DashboardIcon,
    },
    {
        text: 'Users',
        Icon: Users3Icon,
        baseUrl: '/owner/users',
        subItems: [
            {
                text: 'Members',
                baseUrl: '/owner/users/members',
                Icon: UserIcon,
                subItems: [
                    {
                        text: 'View Members',
                        url: '/view-members',
                        Icon: UserListIcon,
                    },
                    {
                        text: 'Members Activity',
                        url: '/members-activity',
                        Icon: UserClockIcon,
                    },
                ],
            },
            {
                text: 'Employees',
                baseUrl: '/owner/users/employees',
                Icon: UserShieldIcon,
                subItems: [
                    {
                        text: 'View employees',
                        url: '/view-employees',
                        Icon: UserListIcon,
                    },
                    {
                        text: 'Employee Footsteps',
                        url: '/employee-footsteps',
                        Icon: FootstepsIcon,
                    },
                ],
            },
        ],
    },
    {
        text: 'Roles Management',
        Icon: ShieldIcon,
        url: '/owner/roles-management',
    },
    {
        text: 'Company',
        baseUrl: '/owner/company',
        Icon: BuildingIcon,
        subItems: [
            {
                text: 'Profile',
                Icon: BuildingCogIcon,
                url: '/profile',
            },
            {
                text: 'Branches',
                Icon: BuildingBranchIcon,
                url: '/branches',
            },
        ],
    },
    {
        text: 'Footstep Tracking',
        Icon: FootstepsIcon,
        url: '/owner/footstep-tracking',
    },
    {
        text: 'Reports',
        Icon: ReportsIcon,
        url: '/owner/reports',
    },
    {
        Icon: NotificationIcon,
        text: 'Notifications',
        url: '/owner/notifications',
    },
    {
        text: 'Profile',
        Icon: UserIcon,
        url: '/owner/profile',
    },
    {
        text: 'Settings',
        Icon: SettingsIcon,
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
