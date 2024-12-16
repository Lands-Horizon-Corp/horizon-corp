import {
    redirect,
    createRoute,
    lazyRouteComponent,
} from '@tanstack/react-router'

import CompanyErrorPage from './error'
import { companyIdPathSchema } from './route-schemas'
import { adminCompaniesManagement } from '../route'
import { CompanyResource } from '@/horizon-corp/types'
import { queryOptions } from '@tanstack/react-query'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'

// pre-loader of company desu
export const companyLoader = (companyId: number) =>
    queryOptions<CompanyResource>({
        queryKey: ['company', companyId],
        queryFn: async () => {
            const data = await CompanyService.getById(companyId)
            return data
        },
        retry: 0,
    })

const adminCompanyId = createRoute({
    getParentRoute: () => adminCompaniesManagement,
    path: '$companyId',
    params: {
        parse: companyIdPathSchema.parse,
    },
    errorComponent: CompanyErrorPage,
    loader: async ({ context, params: { companyId } }) => {
        await context.queryClient.ensureQueryData(companyLoader(companyId))
    },
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
    component: lazyRouteComponent(() => import('./view')),
})

const adminCompanyEditRoute = createRoute({
    getParentRoute: () => adminCompanyId,
    path: 'edit',
    component: lazyRouteComponent(() => import('./edit')),
})

const AdminCompanyIdRoute = adminCompanyId.addChildren([
    adminCompanyViewRoute,
    adminCompanyEditRoute,
    adminCompanyIndexRoute,
])

export default AdminCompanyIdRoute
