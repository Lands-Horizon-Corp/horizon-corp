import {
    BankIcon,
    UserIcon,
    ShieldIcon,
    Users3Icon,
    UserTagIcon,
    UserCogIcon,
    GendersIcon,
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
    GraduationCapIcon,
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
                    {
                        text: 'Member Types',
                        url: '/member-types',
                        Icon: UserCogIcon,
                    },
                    {
                        text: 'Educational Attainments',
                        url: '/member-educational-attainments',
                        Icon: GraduationCapIcon,
                    },
                    {
                        text: 'Member Classification',
                        url: '/member-classification',
                        Icon: UserTagIcon,
                    },
                    {
                        text: 'Member Occupation',
                        url: '/member-occupation',
                        Icon: BriefCaseIcon,
                    },
                    {
                        text: 'Genders',
                        Icon: GendersIcon,
                        url: '/genders',
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
        text: 'Accounting',
        Icon: BankIcon,
        baseUrl: '/owner/accounting',
        subItems: [
            { text: 'Accounts', url: '/accounts', Icon: BankIcon },
            {
                Icon: BankIcon,
                text: 'Computation Type',
                url: '/computation-type',
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
