import { createRoute, lazyRouteComponent } from '@tanstack/react-router'

import { rootRoute } from '@/root-route'

const adminRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'admin',
    component: lazyRouteComponent(()=> import("./layout")),
})

const adminLandingRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: '/',
    component: lazyRouteComponent(()=> import("./pages/")),
})

const adminDashboardRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'dashboard',
    component: lazyRouteComponent(()=> import("./pages/dashboard")),
})

const adminViewMembersRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'members-management/view-members',
    component: lazyRouteComponent(() => import("./pages/members/view-members")),
})

const adminMembersFeedbackRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'members-management/feedbacks',
    component: lazyRouteComponent(() => import("./pages/members/feedbacks")),
})

const adminViewCompaniesRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'companies-management/view-companies',
    component: lazyRouteComponent(()=> import("./pages/companies-management/view-companies")),
})

const adminCompaniesFeedbackRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'companies-management/feedbacks',
    component: lazyRouteComponent(() => import("./pages/companies-management/feedbacks")),
})

const adminFootstepTrackingRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'footstep-tracking',
    component: lazyRouteComponent(()=> import("./pages/footsteps")),
})

const adminProfileRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'profile',
    component: lazyRouteComponent(()=> import("./pages/profile")),
})

const adminNotificationsRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'notifications',
    component: lazyRouteComponent(()=> import("./pages/notifications")),
})

const adminSettingsRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'settings',
    component: lazyRouteComponent(()=> import("./pages/settings")),
})

const AdminRoute = adminRoute.addChildren([
    adminLandingRoute,
    adminDashboardRoute,
    adminViewMembersRoute,
    adminMembersFeedbackRoute,
    adminViewCompaniesRoute,
    adminCompaniesFeedbackRoute,
    adminFootstepTrackingRoute,
    adminProfileRoute,
    adminNotificationsRoute,
    adminSettingsRoute,
])

export default AdminRoute
