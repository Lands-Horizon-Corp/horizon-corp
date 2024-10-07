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
import Sidebar from '@/components/sidebar'

import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component/base'
import { TSidebarItem } from '@/types/component/sidebar'

export const employeeSidebarItems: TSidebarItem[] = [
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
        <div className="hidden sm:block">
            <Sidebar
                enableCollapse
                enableFocusBlur
                logoRedirectUrl='/employee'
                items={employeeSidebarItems}
                className={cn('', className)}
            />
        </div>
    )
}

export default EmployeeSidebar
