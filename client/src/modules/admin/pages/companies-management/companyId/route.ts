import {
  redirect,
  createRoute,
  lazyRouteComponent,
} from '@tanstack/react-router'
import { queryOptions } from '@tanstack/react-query'

import CompanyErrorPage from './error'
import { adminCompaniesManagement } from '../route'
import { companyIdPathSchema } from './route-schemas'
import { CompanyResource } from '@/horizon-corp/types'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'
import AdminCompanyBranchRoute from './branches/route'

// pre-loader of company desu
export const companyLoader = (companyId: number) =>
  queryOptions<CompanyResource>({
    queryKey: ['company', companyId],
    queryFn: async () => {
      const data = await CompanyService.getById(companyId, ["Owner"])
      return data
    },
    retry: 0,
  })

export const adminCompanyId = createRoute({
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
  AdminCompanyBranchRoute
])

export default AdminCompanyIdRoute
