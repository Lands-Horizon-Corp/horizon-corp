import {
    redirect,
    createRoute,
    lazyRouteComponent,
} from '@tanstack/react-router'

import { ownerUserRoute } from '../route'

export const ownerUserEmployeesRoute = createRoute({
    getParentRoute: () => ownerUserRoute,
    path: 'employees',
})

export const ownerEmployeesIndexRoute = createRoute({
    path: '/',
    getParentRoute: () => ownerUserEmployeesRoute,
    beforeLoad: () => {
        throw redirect({
            to: '/owner/users/employees/view-employees',
        })
    },
})

export const ownerEmployeesView = createRoute({
    getParentRoute: () => ownerUserEmployeesRoute,
    path: 'view-employees',
    component: lazyRouteComponent(() => import('./view-employees')),
})

const ownerEmployeesFootstepsRoute = createRoute({
    getParentRoute: () => ownerUserEmployeesRoute,
    path: 'employee-footsteps',
    component: lazyRouteComponent(() => import('./employee-footsteps')),
})

const OwnerUserEmployeesRoute = ownerUserEmployeesRoute.addChildren([
    ownerEmployeesView,
    ownerEmployeesIndexRoute,
    ownerEmployeesFootstepsRoute,
])

export default OwnerUserEmployeesRoute
