import { createRoute, lazyRouteComponent } from '@tanstack/react-router'

import { adminCompanyId } from '../route'

const adminCompanyBranchesRoute = createRoute({
    getParentRoute: () => adminCompanyId,
    path: 'branches',
})

const adminCompanyBranchesIndexRoute = createRoute({
    getParentRoute: () => adminCompanyBranchesRoute,
    path: '/',
    component: lazyRouteComponent(() => import('./view-branches')),
})

const AdminCompanyBranchRoute = adminCompanyBranchesRoute.addChildren([
    adminCompanyBranchesIndexRoute,
])

export default AdminCompanyBranchRoute
