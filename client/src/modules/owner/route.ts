import { createRoute, lazyRouteComponent } from '@tanstack/react-router'

import { rootRoute } from '@/root-route'

const ownerRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'owner',
    beforeLoad: () => {},
    component: lazyRouteComponent(() => import('@/modules/owner/layout')),
})

const ownerLandingRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: '/',
    component: lazyRouteComponent(() => import('@/modules/owner/pages')),
})

const ownerDashboardRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'dashboard',
    component: lazyRouteComponent(() => import('./pages/dashboard')),
})

const ownerViewMembersRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: '/users/members/view-members',
    component: lazyRouteComponent(() => import('./pages/members/view-members')),
})

const ownerMembersActivityRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'users/members/members-activity',
    component: lazyRouteComponent(() => import('./pages/members/members-activity')),
})

const ownerViewEmployeesRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'users/employees/view-employees',
    component: lazyRouteComponent(() => import('./pages/employees/view-employees')),
})

const ownerEmployeeFootstepsRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'users/employees/employee-footsteps',
    component: lazyRouteComponent(() => import('./pages/employees/employee-footsteps')),
})

const ownerRolesManagementRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'roles-management',
    component: lazyRouteComponent(() => import('./pages/roles-management')),
})

const ownerCompanyProfileRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'company/profile',
    component: lazyRouteComponent(() => import('./pages/company/profile')),
})

const ownerCompanyBranchesRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'company/branches',
    component: lazyRouteComponent(() => import('./pages/company/branches')),
})

const ownerFootstepTrackingRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'footstep-tracking',
    component: lazyRouteComponent(() => import('./pages/footstep-tracking')),
})

const ownerReportsRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'reports',
    component: lazyRouteComponent(() => import('./pages/reports')),
})

const ownerNotificationsRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'notifications',
    component: lazyRouteComponent(() => import('./pages/notifications')),
})

const ownerProfileRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'profile',
    component: lazyRouteComponent(() => import('./pages/profile')),
})

const ownerSettingsRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'settings',
    component: lazyRouteComponent(() => import('./pages/settings')),
})

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
