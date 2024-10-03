import { createRoute } from '@tanstack/react-router'

import OwnerLandingPage from '@/modules/owner/pages'
import OwnerDashboardPage from './pages/dashboard'
import OwnerViewMembersPage from './pages/members/view-members'
import OwnerMembersActivityPage from './pages/members/members-activity'
import OwnerViewEmployeesPage from './pages/employees/view-employees'
import OwnerEmployeeFootstepsPage from './pages/employees/employee-footsteps'
import OwnerRolesManagementPage from './pages/roles-management'
import OwnerCompanyProfilePage from './pages/company/profile'
import OwnerCompanyBranchesPage from './pages/company/branches'
import OwnerFootstepTrackingPage from './pages/footstep-tracking'
import OwnerReportsPage from './pages/reports'
import OwnerNotificationsPage from './pages/notifications'
import OwnerProfilePage from './pages/profile'
import OwnerSettingsPage from './pages/settings'

import { rootRoute } from '@/root-route'
import OwnerLayout from '@/modules/owner/layout'

const ownerRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'owner',
    beforeLoad: () => {
        // TODO: Once middleware is implemented
    },
    component: OwnerLayout,
})

const ownerLandingRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: '/',
    component: OwnerLandingPage,
})

const ownerDashboardRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'dashboard',
    component: OwnerDashboardPage,
})

const ownerViewMembersRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: '/users/members/view-members',
    component: OwnerViewMembersPage,
})

const ownerMembersActivityRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'users/members/members-activity',
    component: OwnerMembersActivityPage,
})

const ownerViewEmployeesRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'users/employees/view-employees',
    component: OwnerViewEmployeesPage,
})

const ownerEmployeeFootstepsRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'users/employees/employee-footsteps',
    component: OwnerEmployeeFootstepsPage,
})

const ownerRolesManagementRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'roles-management',
    component: OwnerRolesManagementPage,
})

const ownerCompanyProfileRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'company/profile',
    component: OwnerCompanyProfilePage,
})

const ownerCompanyBranchesRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'company/branches',
    component: OwnerCompanyBranchesPage,
})

const ownerFootstepTrackingRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'footstep-tracking',
    component: OwnerFootstepTrackingPage,
})

const ownerReportsRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'reports',
    component: OwnerReportsPage,
})

const ownerNotificationsRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'notifications',
    component: OwnerNotificationsPage,
})

const ownerProfileRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'profile',
    component: OwnerProfilePage,
})

const ownerSettingsRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'settings',
    component: OwnerSettingsPage,
})

// Add all child routes under the main owner route
const OwnerRoute = ownerRoute.addChildren([
    ownerLandingRoute,
    ownerDashboardRoute,
    ownerViewMembersRoute,
    ownerMembersActivityRoute,
    ownerViewEmployeesRoute,
    ownerEmployeeFootstepsRoute,
    ownerRolesManagementRoute,
    ownerCompanyProfileRoute,
    ownerCompanyBranchesRoute,
    ownerFootstepTrackingRoute,
    ownerReportsRoute,
    ownerNotificationsRoute,
    ownerProfileRoute,
    ownerSettingsRoute,
])

export default OwnerRoute
