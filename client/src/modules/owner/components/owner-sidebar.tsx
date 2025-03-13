import { Link } from '@tanstack/react-router'

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
    BriefCaseIcon,
    FootstepsIcon,
    UserClockIcon,
    UserShieldIcon,
    BuildingCogIcon,
    NotificationIcon,
    GraduationCapIcon,
    BuildingBranchIcon,
    WalletIcon,
    BillIcon,
    HandCoinsIcon,
    HandDepositIcon,
    HandWithdrawIcon,
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

import { IBaseComp } from '@/types/component'
import { INavItem } from '@/components/app-sidebar/types'
import AppSidebarUser from '@/components/app-sidebar/app-sidebar-user'
import AppSidebarItem from '@/components/app-sidebar/app-sidebar-item'

const ownerSidebarItem: INavItem[] = [
    {
        title: 'Dashboard',
        icon: DashboardIcon,
        url: '/owner/dashboard',
        type: 'item',
    },
    {
        title: 'Transactions',
        icon: WalletIcon,
        type: 'dropdown',
        url: '/owner/transaction',
        items: [
            {
                title: 'Deposit Entry',
                icon: HandDepositIcon,
                type: 'item',
                url: '/deposit-entry',
            },
            {
                icon: HandCoinsIcon,
                title: 'Payments Entry',
                type: 'item',
                url: '/payments-entry',
            },
            {
                icon: HandWithdrawIcon,
                title: 'Withdrawals',
                type: 'item',
                url: '/withdrawals',
            },
            {
                icon: BillIcon,
                title: 'Payment Types',
                type: 'item',
                url: '/payment-types',
            },
        ],
    },

    {
        title: 'Users',
        icon: Users3Icon,
        type: 'dropdown',
        url: '/owner/users',
        items: [
            {
                title: 'Members',
                url: '/members',
                icon: UserIcon,
                type: 'dropdown',
                items: [
                    {
                        title: 'View Members',
                        url: '/view-members',
                        type: 'item',
                        icon: UserListIcon,
                    },
                    {
                        title: 'Members Activity',
                        url: '/members-activity',
                        type: 'item',
                        icon: UserClockIcon,
                    },
                    {
                        title: 'Member Types',
                        url: '/member-types',
                        type: 'item',
                        icon: UserCogIcon,
                    },
                    {
                        title: 'Educational Attainments',
                        url: '/member-educational-attainments',
                        type: 'item',
                        icon: GraduationCapIcon,
                    },
                    {
                        title: 'Member Classification',
                        url: '/member-classification',
                        type: 'item',
                        icon: UserTagIcon,
                    },
                    {
                        title: 'Member Occupation',
                        url: '/member-occupation',
                        type: 'item',
                        icon: BriefCaseIcon,
                    },
                    {
                        title: 'Genders',
                        icon: GendersIcon,
                        type: 'item',
                        url: '/genders',
                    },
                ],
            },
            {
                title: 'Employees',
                url: '/employees',
                type: 'dropdown',
                icon: UserShieldIcon,
                items: [
                    {
                        title: 'View Employees',
                        url: '/view-employees',
                        icon: UserListIcon,
                        type: 'item',
                    },
                    {
                        title: 'Employee Footsteps',
                        url: '/employee-footsteps',
                        icon: FootstepsIcon,
                        type: 'item',
                    },
                ],
            },
        ],
    },
    {
        title: 'Roles Management',
        icon: ShieldIcon,
        type: 'item',
        url: '/owner/roles-management',
    },
    {
        title: 'Company',
        url: '/owner/company',
        icon: BuildingIcon,
        type: 'dropdown',
        items: [
            {
                title: 'Profile',
                icon: BuildingCogIcon,
                type: 'item',
                url: '/profile',
            },
            {
                title: 'Branches',
                type: 'item',
                icon: BuildingBranchIcon,
                url: '/branches',
            },
        ],
    },
    {
        title: 'Accounting',
        icon: BankIcon,
        url: '/owner/accounting',
        type: 'dropdown',
        items: [
            {
                title: 'Accounts',
                type: 'item',
                url: '/accounts',
                icon: BankIcon,
            },
            {
                icon: BankIcon,
                type: 'item',
                title: 'Computation Type',
                url: '/computation-type',
            },
        ],
    },
    {
        title: 'Footstep Tracking',
        icon: FootstepsIcon,
        type: 'item',
        url: '/owner/footstep-tracking',
    },
    {
        title: 'Reports',
        icon: ReportsIcon,
        type: 'item',
        url: '/owner/reports',
    },
    {
        icon: NotificationIcon,
        title: 'Notifications',
        type: 'item',
        url: '/owner/notifications',
    },
    {
        title: 'Profile',
        icon: UserIcon,
        type: 'item',
        url: '/owner/profile',
    },
    {
        title: 'Settings',
        type: 'item',
        icon: SettingsIcon,
        url: '/owner/settings',
    },
]

const OwnerSidebar = (props: IBaseComp) => {
    return (
        <Sidebar variant="inset" {...props}>
            <SidebarHeader>
                <SidebarMenu>
                    <SidebarMenuItem>
                        <SidebarMenuButton size="lg" asChild>
                            <Link to="/owner">
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
                            {ownerSidebarItem.map((navItem, index) => (
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

export default OwnerSidebar
