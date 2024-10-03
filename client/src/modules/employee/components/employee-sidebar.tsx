import {
    UserIcon,
    Users3Icon,
    ReportsIcon,
    SettingsIcon,
    DashboardIcon,
    FootstepsIcon,
    NotificationIcon,
    BuildingBranchIcon,
} from '@/components/icons'

import { IBaseComp } from '@/types/component/base'
import { TSidebarItem } from '@/types/component/sidebar'
import DynamicSidebar from '@/components/sidebar/dynamic-sidebar'

const employeeSidebarItems: TSidebarItem[] = [
    {
        text: 'Dashboard',
        url: '/employee/dashboard',
        Icon: DashboardIcon,
    },
    {
        text: 'Users',
        Icon: Users3Icon,
        baseUrl: '/employee/users',
        subItems: [
            {
                text: 'Members',
                url: '/members',
                Icon: UserIcon,
            },
        ],
    },
    {
        text: 'Branch',
        url: '/employee/branch',
        Icon: BuildingBranchIcon,
    },
    {
        text: 'Reports',
        url: '/employee/reports',
        Icon: ReportsIcon,
    },
    {
        text: 'Notifications',
        url: '/employee/notifications',
        Icon: NotificationIcon,
    },
    {
        text: 'Footsteps',
        url: '/employee/footsteps',
        Icon: FootstepsIcon,
    },
    {
        text: 'Profile',
        url: '/employee/profile',
        Icon: UserIcon,
    },
    {
        text: 'Settings',
        url: '/employee/settings',
        Icon: SettingsIcon,
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
