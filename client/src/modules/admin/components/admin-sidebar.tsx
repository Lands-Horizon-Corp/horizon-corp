import { useMemo } from 'react'
import { Link, useRouter } from '@tanstack/react-router'

import {
    UserIcon,
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
import { INavGroupItem } from '@/components/app-sidebar/types'
import AppSidebarItem from '@/components/app-sidebar/app-sidebar-item'
import AppSidebarUser from '@/components/app-sidebar/app-sidebar-user'
import { flatSidebarGroupItem } from '@/components/app-sidebar/app-sidebar-utils'
import AppSidebarQruickNavigate from '@/components/app-sidebar/app-sidebar-quick-navigate'

import { IBaseComp } from '@/types/component'

export const adminSidebarGroupItems: INavGroupItem[] = [
    {
        title: 'home',
        navItems: [
            {
                title: 'Dashboard',
                icon: DashboardIcon,
                url: '/admin/dashboard',
                type: 'item',
            },
        ],
    },
    {
        title: 'Users',
        navItems: [
            {
                title: 'View Members',
                url: '/admin/members-management/view-members',
                icon: UserIcon,
                type: 'item',
            },
        ],
    },
    {
        title: 'Company Management',
        navItems: [
            {
                title: 'View Companies',
                url: '/admin/companies-management/view-companies',
                icon: BuildingBranchIcon,
                type: 'item',
            },
        ],
    },
    {
        title: 'Management',
        navItems: [
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
        ],
    },
]

const AdminSidebar = (props: IBaseComp) => {
    const router = useRouter()

    const item = useMemo(
        () =>
            flatSidebarGroupItem(adminSidebarGroupItems).map((item) => ({
                ...item,
                items: item.items.map((itm) => ({
                    ...itm,
                    onClick: (self: typeof itm) => {
                        router.navigate({ to: self.url })
                    },
                })),
            })),
        [router]
    )

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
                <AppSidebarQruickNavigate groups={item} />
            </SidebarHeader>
            <SidebarContent className="ecoop-scroll">
                {adminSidebarGroupItems.map((navGroupItem, i) => (
                    <SidebarGroup key={`${navGroupItem.title}-${i}`}>
                        <SidebarGroupLabel>
                            {navGroupItem.title}
                        </SidebarGroupLabel>
                        <SidebarGroupContent>
                            <SidebarMenu>
                                {navGroupItem.navItems.map((navItem, index) => (
                                    <AppSidebarItem
                                        key={index}
                                        navItem={{ ...navItem, depth: 1 }}
                                    />
                                ))}
                            </SidebarMenu>
                        </SidebarGroupContent>
                    </SidebarGroup>
                ))}
            </SidebarContent>
            <SidebarFooter>
                <AppSidebarUser />
            </SidebarFooter>
            <SidebarRail />
        </Sidebar>
    )
}

export default AdminSidebar
