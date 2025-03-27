import { Link } from '@tanstack/react-router'

import {
    BankIcon,
    UserIcon,
    BillIcon,
    WalletIcon,
    ShieldIcon,
    Users3Icon,
    UserTagIcon,
    UserCogIcon,
    GendersIcon,
    ReportsIcon,
    BuildingIcon,
    SettingsIcon,
    UserListIcon,
    HandCoinsIcon,
    DashboardIcon,
    BriefCaseIcon,
    FootstepsIcon,
    UserClockIcon,
    UserShieldIcon,
    HandDepositIcon,
    BuildingCogIcon,
    NotificationIcon,
    HandWithdrawIcon,
    GraduationCapIcon,
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

import { IBaseComp } from '@/types/component'
import AppSidebarUser from '@/components/app-sidebar/app-sidebar-user'
import AppSidebarItem from '@/components/app-sidebar/app-sidebar-item'
import { INavGroupItem, INavItem } from '@/components/app-sidebar/types'

const ownerSidebarGroupItems: INavGroupItem[] = [
    {
        title: 'Home',
        navItems: [
            {
                title: 'Dashboard',
                icon: DashboardIcon,
                url: '/owner/dashboard',
                type: 'item',
            },
        ],
    },
    {
        title: 'Transaction',
        navItems: [
            {
                title: 'Deposit Entry',
                icon: HandDepositIcon,
                type: 'item',
                url: '/owner/transaction/deposit-entry',
            },
            {
                icon: HandCoinsIcon,
                title: 'Payments Entry',
                type: 'item',
                url: '/owner/transaction/payments-entry',
            },
            {
                icon: HandWithdrawIcon,
                title: 'Withdrawals',
                type: 'item',
                url: '/owner/transaction/withdrawals',
            },
            {
                icon: BillIcon,
                title: 'Payment Types',
                type: 'item',
                url: '/owner/transaction/payment-types',
            },
        ],
    },
    {
        title: 'Accounting',
        navItems: [
            {
                title: 'Accounts',
                type: 'item',
                url: '/owner/accounting/accounts',
                icon: BankIcon,
            },
            {
                icon: BankIcon,
                type: 'item',
                title: 'Computation Type',
                url: '/owner/accounting/computation-type',
            },
        ],
    },
    {
        title: 'Users',
        navItems: [
            {
                title: 'Members',
                url: '/owner/users/members',
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
                        title: 'Member Types',
                        url: '/member-types',
                        type: 'item',
                        icon: UserCogIcon,
                    },
                    {
                        title: 'Member Center',
                        url: '/member-center',
                        type: 'item',
                        icon: UserCogIcon,
                    },
                    {
                        title: 'Members Activity',
                        url: '/members-activity',
                        type: 'item',
                        icon: UserClockIcon,
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
                url: '/owner/users/employees',
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
            {
                title: 'Roles Management',
                icon: ShieldIcon,
                type: 'item',
                url: '/owner/roles-management',
            },
        ],
    },
    {
        title: 'Company',
        navItems: [
            {
                title: 'Profile',
                icon: BuildingCogIcon,
                type: 'item',
                url: '/owner/profile',
            },
            {
                title: 'Branches',
                type: 'item',
                icon: BuildingBranchIcon,
                url: '/owner/branches',
            },
        ],
    },
    {
        title: 'Others',
        navItems: [
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
        ],
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
                {ownerSidebarGroupItems.map((navGroupItem, i) => (
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

export default OwnerSidebar
