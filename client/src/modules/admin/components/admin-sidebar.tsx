import { Link } from '@tanstack/react-router'

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
import {
    Sidebar,
    SidebarRail,
    SidebarMenu,
    SidebarGroup,
    SidebarHeader,
    SidebarFooter,
    SidebarContent,
    SidebarMenuItem,
    SidebarMenuButton,
    SidebarGroupLabel,
    SidebarGroupContent,
} from '@/components/ui/sidebar'
import EcoopLogo from '@/components/ecoop-logo'
import AppSidebarItem from '@/components/app-sidebar/app-sidebar-item'

import { IBaseComp } from '@/types/component'
import { INavItem } from '@/components/app-sidebar/types'
import AppSidebarUser from '@/components/app-sidebar/app-sidebar-user'

export const adminSidebarItems: INavItem[] = [
    {
        title: 'Dashboard',
        icon: DashboardIcon,
        url: '/admin/dashboard',
        type: 'item',
    },
    {
        title: 'Members Management',
        icon: Users3Icon,
        type: 'dropdown',
        url: '/admin/members-management',
        items: [
            {
                title: 'View Members',
                url: '/view-members',
                icon: UserIcon,
                type: 'item',
            },
        ],
    },
    {
        title: 'Companies Management',
        icon: BuildingBranchIcon,
        type: 'dropdown',
        url: '/admin/companies-management',
        items: [
            {
                title: 'View Companies',
                url: '/view-companies',
                icon: BuildingBranchIcon,
                type: 'item',
            },
        ],
    },
    {
        title: 'Footstep Tracking',
        icon: FootstepsIcon,
        type: 'item',
        url: '/admin/footstep-tracking',
    },
    {
        title: 'Profile',
        icon: UserIcon,
        type: 'item',
        url: '/admin/profile',
    },
    {
        title: 'Notifications',
        icon: NotificationIcon,
        type: 'item',
        url: '/admin/notifications',
    },
    {
        title: 'Feedbacks',
        icon: FeedbackIcon,
        type: 'item',
        url: '/admin/feedbacks',
    },
    {
        title: 'Settings',
        icon: SettingsIcon,
        type: 'item',
        url: '/admin/settings',
    },
]

const AdminSidebar = (props: IBaseComp) => {
    return (
        <Sidebar variant="inset" {...props}>
            <SidebarHeader>
                <SidebarMenu>
                    <SidebarMenuItem>
                        <SidebarMenuButton size="lg" asChild>
                            <Link to="/admin">
                                <EcoopLogo className="size-9" />
                                <div className="grid flex-1 text-left text-sm leading-tight">
                                    <span className="truncate font-semibold">
                                        eCOOP
                                    </span>
                                    <span className="truncate text-xs text-muted-foreground/80">
                                        Admin
                                    </span>
                                </div>
                            </Link>
                        </SidebarMenuButton>
                    </SidebarMenuItem>
                </SidebarMenu>
            </SidebarHeader>
            <SidebarContent className="ecoop-scroll">
                <SidebarGroup>
                    <SidebarGroupLabel>Navigation</SidebarGroupLabel>
                    <SidebarGroupContent>
                        <SidebarMenu>
                            {adminSidebarItems.map((navItem, index) => (
                                <AppSidebarItem
                                    key={index}
                                    navItem={{ ...navItem, depth: 1 }}
                                />
                            ))}
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
            <SidebarFooter>
                <AppSidebarUser />
            </SidebarFooter>
            <SidebarRail />
        </Sidebar>
    )
}

export default AdminSidebar
