import {
    BuildingBranchIcon,
    DashboardIcon,
    FeedbackIcon,
    FootstepsIcon,
    NotificationIcon,
    SettingsIcon,
    UserIcon,
    Users3Icon,
} from '@/components/icons'
import DynamicSidebar from '@/components/sidebar/dynamic-sidebar'

import { IBaseComp } from '@/types/component/base'
import { TSidebarItem } from '@/types/component/sidebar'

const adminSidebarItems: TSidebarItem[] = [
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
