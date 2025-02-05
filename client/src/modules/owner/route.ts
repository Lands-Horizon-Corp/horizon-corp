import { createRoute, lazyRouteComponent } from '@tanstack/react-router'

import { rootRoute } from '@/root-route'
import OwnerUserRoute from './pages/users/route'
import { redirect } from '@tanstack/react-router'
import OwnerCompanyRoute from './pages/company/route'
import ownerReportsRoute from './pages/reports/route'

export const ownerRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'owner',
    beforeLoad: () => {},
    component: lazyRouteComponent(() => import('@/modules/owner/layout')),
})

const ownerLandingRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: '/',
    beforeLoad: () => redirect({ to: '/owner/dashboard' }),
})

const ownerDashboardRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'dashboard',
    component: lazyRouteComponent(() => import('./pages/dashboard')),
})

const ownerRolesManagementRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'roles-management',
    component: lazyRouteComponent(() => import('./pages/roles-management')),
})

const ownerFootstepTrackingRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'footstep-tracking',
    component: lazyRouteComponent(() => import('./pages/footstep-tracking')),
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
    ownerRolesManagementRoute,
    OwnerUserRoute,
    OwnerCompanyRoute,
    ownerFootstepTrackingRoute,
    ownerReportsRoute,
    ownerNotificationsRoute,
    ownerProfileRoute,
    ownerSettingsRoute,
])

export default OwnerRoute
