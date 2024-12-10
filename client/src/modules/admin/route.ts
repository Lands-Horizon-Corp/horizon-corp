import { createRoute, lazyRouteComponent, redirect } from '@tanstack/react-router'

import { rootRoute } from '@/root-route'

const adminRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'admin',
    component: lazyRouteComponent(() => import('./layout')),
})

const adminLandingRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: '/',
    component: lazyRouteComponent(() => import('./pages/')),
})

const adminDashboardRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'dashboard',
    component: lazyRouteComponent(() => import('./pages/dashboard')),
})

const adminViewMembersRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'members-management/view-members',
    component: lazyRouteComponent(
        () => import('./pages/members/view-members/index')
    ),
})

const adminMembersFeedbackRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'members-management/feedbacks',
    component: lazyRouteComponent(() => import('./pages/members/feedbacks')),
})

const adminCompaniesDefaultRoute = createRoute({
  getParentRoute: () => adminRoute,
  path: 'companies-management',
  loader: () => {
    throw redirect({ to: '/admin/companies-management/view-companies' });
  },
});


const adminViewCompaniesRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'companies-management/view-companies',
    component: lazyRouteComponent(
        () => import('./pages/companies-management/view-companies')
    ),
})

const adminViewCompanyBranchesRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'companies-management/view-companies/$companyId/view-branches',
    component: lazyRouteComponent(
        () =>
            import(
                './pages/companies-management/company-branches'
            )
    ),
})

const adminCompaniesFeedbackRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'companies-management/feedbacks',
    component: lazyRouteComponent(
        () => import('./pages/companies-management/feedbacks')
    ),
})

const adminFootstepTrackingRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'footstep-tracking',
    component: lazyRouteComponent(() => import('./pages/footsteps')),
})

const adminProfileRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'profile',
    component: lazyRouteComponent(() => import('./pages/profile')),
})

const adminNotificationsRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'notifications',
    component: lazyRouteComponent(() => import('./pages/notifications')),
})

const adminSettingsRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'settings',
    component: lazyRouteComponent(() => import('./pages/settings')),
})

const AdminRoute = adminRoute.addChildren([
    adminLandingRoute,
    adminDashboardRoute,
    adminViewMembersRoute,
    adminMembersFeedbackRoute,
    
    adminCompaniesDefaultRoute,
    adminViewCompanyBranchesRoute,
    adminViewCompaniesRoute,

    adminCompaniesFeedbackRoute,
    adminFootstepTrackingRoute,
    adminProfileRoute,
    adminNotificationsRoute,
    adminSettingsRoute,
])

export default AdminRoute
