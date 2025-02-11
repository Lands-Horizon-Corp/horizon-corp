import {
    redirect,
    createRoute,
    lazyRouteComponent,
} from '@tanstack/react-router'

import { ownerUserRoute } from '../route'
import OwnerMemberIdRoute from './member-id/route'

export const ownerUserMemberRoute = createRoute({
    getParentRoute: () => ownerUserRoute,
    path: '/members',
})

const ownerUserMemberIndexRoute = createRoute({
    getParentRoute: () => ownerUserMemberRoute,
    path: '/',
    beforeLoad: () => redirect({ to: '/owner/users/members/view-members' }),
})

const ownerViewMembersRoute = createRoute({
    getParentRoute: () => ownerUserMemberRoute,
    path: 'view-members',
    component: lazyRouteComponent(() => import('./view-members-page')),
})

const ownerMembersActivityRoute = createRoute({
    getParentRoute: () => ownerUserMemberRoute,
    path: 'members-activity',
    component: lazyRouteComponent(() => import('./members-activity-page')),
})

const ownerMemberTypesRoute = createRoute({
    getParentRoute: () => ownerUserMemberRoute,
    path: 'member-types',
    component: lazyRouteComponent(() => import('./member-types-page')),
})

const OwnerUserMemberRoute = ownerUserMemberRoute.addChildren([
    ownerUserMemberIndexRoute,
    ownerViewMembersRoute,
    ownerMembersActivityRoute,
    ownerMemberTypesRoute,
    OwnerMemberIdRoute,
])

export default OwnerUserMemberRoute
