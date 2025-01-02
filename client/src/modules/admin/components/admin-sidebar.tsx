import {
    UserIcon,
    Users3Icon,
    SettingsIcon,
    FeedbackIcon,
    DashboardIcon,
    FootstepsIcon,
    NotificationIcon,
    BuildingBranchIcon,
} from '@/components/icons'
import Sidebar from '@/components/sidebar'

import { IBaseComp } from '@/types/component'
import { TSidebarItem } from '@/components/sidebar/sidebar-types'

export const adminSidebarItems: TSidebarItem[] = [
    {
        text: 'Dashboard',
        url: '/admin/dashboard',
        Icon: DashboardIcon,
    },
    {
        text: 'Members Management',
        Icon: Users3Icon,
        baseUrl: '/admin/members-management',
        subItems: [
            {
                text: 'View Members',
                url: '/view-members',
                Icon: UserIcon,
            },
            {
                text: 'Feedbacks',
                url: '/feedbacks',
                Icon: FeedbackIcon,
            },
        ],
    },
    {
        text: 'Companies Management',
        Icon: BuildingBranchIcon,
        baseUrl: '/admin/companies-management',
        subItems: [
            {
                text: 'View Companies',
                url: '/view-companies',
                Icon: BuildingBranchIcon,
            },
            {
                text: 'Feedbacks',
                url: '/feedbacks',
                Icon: FeedbackIcon,
            },
        ],
    },
    {
        text: 'Footstep Tracking',
        url: '/admin/footstep-tracking',
        Icon: FootstepsIcon,
    },
    {
        text: 'Profile',
        url: '/admin/profile',
        Icon: UserIcon,
    },
    {
        text: 'Notifications',
        url: '/admin/notifications',
        Icon: NotificationIcon,
    },
    {
        text: 'Settings',
        url: '/admin/settings',
        Icon: SettingsIcon,
    },
]

const AdminSidebar = ({ className }: IBaseComp) => {
    return (
        <Sidebar
            enableCollapse
            className={className}
            logoRedirectUrl="/admin"
            items={adminSidebarItems}
        />
    )
}

export default AdminSidebar
