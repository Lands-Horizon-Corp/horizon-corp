import {
    redirect,
    createRoute,
    lazyRouteComponent,
} from '@tanstack/react-router'

import CompanyErrorPage from './error'
import { companyIdPathSchema } from './route-schemas'
import { adminCompaniesManagement } from '../route'

const adminCompanyId = createRoute({
    getParentRoute: () => adminCompaniesManagement,
    path: '$companyId',
    params: {
        parse: companyIdPathSchema.parse,
    },
    errorComponent: CompanyErrorPage,
})

const adminCompanyIndexRoute = createRoute({
    getParentRoute: () => adminCompanyId,
    path: '/',
    beforeLoad: ({ params }) => {
        throw redirect({
            to: `/admin/companies-management/${params.companyId}/view`,
            params,
        })
    },
})

const adminCompanyViewRoute = createRoute({
    getParentRoute: () => adminCompanyId,
    path: 'view',
    component: lazyRouteComponent(
        () => import('./view')
    ),
})

const adminCompanyEditRoute = createRoute({
    getParentRoute: () => adminCompanyId,
    path: 'edit',
    component: lazyRouteComponent(
        () => import('./edit')
    ),
})

const AdminCompanyIdRoute = adminCompanyId.addChildren([
    adminCompanyViewRoute,
    adminCompanyEditRoute,
    adminCompanyIndexRoute,
])

export default AdminCompanyIdRoute