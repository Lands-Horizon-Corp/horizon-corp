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
import Sidebar from '@/components/sidebar'

import { IBaseComp } from '@/types/component'
import type { TSidebarItem } from '@/components/sidebar/sidebar-types'

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
        baseUrl: '/owner/reports',
        subItems: [
            {
                text: 'Reports-Dashboard',
                Icon: DashboardIcon,
                url: '/dashboard',
            },
        ],
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

const OwnerSidebar = ({ className }: IBaseComp) => {
    return (
        <Sidebar
            enableCollapse
            className={className}
            items={ownerSidebarItem}
            logoRedirectUrl="/owner"
        />
    )
}

export default OwnerSidebar
