import {
    createRoute,
    lazyRouteComponent,
    redirect,
} from '@tanstack/react-router'

import { ownerRoute } from '../../route'
import OwnerUserMemberRoute from './members/route'
import OwnerUserEmployeesRoute from './employees/route'

export const ownerUserRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: 'users',
})

export const ownerUserIndexRoute = createRoute({
    path: '/',
    getParentRoute: () => ownerUserRoute,
    beforeLoad: () => {
        throw redirect({
            to: '/owner/users/members/view-members',
        })
    },
})

export const ownerGenderRoute = createRoute({
    path: 'genders',
    getParentRoute: () => ownerUserRoute,
    component: lazyRouteComponent(() => import('./genders-page')),
})

const OwnerUserRoute = ownerUserRoute.addChildren([
    ownerGenderRoute,
    ownerUserIndexRoute,
    OwnerUserEmployeesRoute,
    OwnerUserMemberRoute,
])

export default OwnerUserRoute
