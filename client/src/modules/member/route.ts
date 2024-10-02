import { createRoute } from '@tanstack/react-router'

import { rootRoute } from '@/root-route'
import MemberProfilePage from './pages/profile'
import MemberMessagesPage from './pages/messages'
import MemberSettingsPage from './pages/settings'
import MemberLayout from '@/modules/member/layout'
import MemberDashboardPage from './pages/dashboard'
import MemberLandingPage from '@/modules/member/pages'
import MemberNotificationPage from './pages/notifications'

const memberRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'member',
    beforeLoad: () => {
        // TODO: Once middleware is implemented
    },
    component: MemberLayout,
})

const memberLandingRoute = createRoute({
    getParentRoute: () => memberRoute,
    path: '/',
    component: MemberLandingPage,
})

const memberDashboardRoute = createRoute({
    getParentRoute: () => memberRoute,
    path: '/dashboard',
    component: MemberDashboardPage,
})

const memberProfileRoute = createRoute({
    getParentRoute: () => memberRoute,
    path: '/profile',
    component: MemberProfilePage,
})

const memberNotificationRoute = createRoute({
    getParentRoute: () => memberRoute,
    path: '/notifications',
    component: MemberNotificationPage,
})

const memberMessagesRoute = createRoute({
    getParentRoute: () => memberRoute,
    path: '/messages',
    component: MemberMessagesPage,
})

const memberSettingsRoute = createRoute({
    getParentRoute: () => memberRoute,
    path: '/settings',
    component: MemberSettingsPage,
})

const MemberRoute = memberRoute.addChildren([
    memberLandingRoute,
    memberDashboardRoute,
    memberProfileRoute,
    memberNotificationRoute,
    memberMessagesRoute,
    memberSettingsRoute,
])

export default MemberRoute
