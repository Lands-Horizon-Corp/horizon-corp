import { createRoute } from '@tanstack/react-router'

import AdminLandingPage from './pages'
import AdminDashboardPage from './pages/dashboard'
import AdminViewMembersPage from './pages/members/view-members'
import AdminMembersFeedbacksPage from './pages/members/feedbacks'
import AdminViewCompaniesPage from './pages/companies-management/view-companies'
import AdminCompaniesFeedbacksPage from './pages/companies-management/feedbacks'
import AdminFootstepsTrackingsPage from './pages/footsteps'
import AdminProfilePage from './pages/profile'
import AdminNotificationsPage from './pages/notifications'
import AdminSettingsPage from './pages/settings'

import AdminLayout from '@/modules/admin/layout'

import { rootRoute } from '@/root-route'

const adminRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'admin',
    component: AdminLayout,
})

const adminLandingRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: '/',
    component: AdminLandingPage,
})

const adminDashboardRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'dashboard',
    component: AdminDashboardPage,
})

const adminViewMembersRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'members-management/view-members',
    component: AdminViewMembersPage,
})

const adminMembersFeedbackRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'members-management/feedbacks',
    component: AdminMembersFeedbacksPage,
})

const adminViewCompaniesRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'companies-management/view-companies',
    component: AdminViewCompaniesPage,
})

const adminCompaniesFeedbackRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'companies-management/feedbacks',
    component: AdminCompaniesFeedbacksPage,
})

const adminFootstepTrackingRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'footstep-tracking',
    component: AdminFootstepsTrackingsPage,
})

const adminProfileRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'profile',
    component: AdminProfilePage,
})

const adminNotificationsRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'notifications',
    component: AdminNotificationsPage,
})

const adminSettingsRoute = createRoute({
    getParentRoute: () => adminRoute,
    path: 'settings',
    component: AdminSettingsPage,
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
