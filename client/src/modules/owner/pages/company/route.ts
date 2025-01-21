import {
    redirect,
    createRoute,
    lazyRouteComponent,
} from '@tanstack/react-router'

import { ownerRoute } from '../../route'

export const ownerCompanyRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: '/company',
})

const ownerCompanyIndexRoute = createRoute({
    getParentRoute: () => ownerCompanyRoute,
    path: '/',
    beforeLoad: () => redirect({ to: '/owner/company/profile' }),
})

const ownerCompanyProfileRoute = createRoute({
    getParentRoute: () => ownerCompanyRoute,
    path: 'profile',
    component: lazyRouteComponent(() => import('./profile')),
})

const ownerCompanyBranchesRoute = createRoute({
    getParentRoute: () => ownerCompanyRoute,
    path: 'branches',
    component: lazyRouteComponent(() => import('./branches')),
})

const OwnerCompanyRoute = ownerCompanyRoute.addChildren([
    ownerCompanyIndexRoute,
    ownerCompanyProfileRoute,
    ownerCompanyBranchesRoute,
])

export default OwnerCompanyRoute
