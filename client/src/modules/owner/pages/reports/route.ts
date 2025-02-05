import {
    createRoute,
    lazyRouteComponent,
    redirect,
} from '@tanstack/react-router'
import { ownerRoute } from '../../route'

const ownerReportsRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: '/reports',
})

const ownerCompanyIndexRoute = createRoute({
    getParentRoute: () => ownerReportsRoute,
    path: '/',
    beforeLoad: () => redirect({ to: '/owner/reports/dashboard' }),
})

const ReportDocumentPage = createRoute({
    getParentRoute: () => ownerReportsRoute,
    path: 'document',
    component: lazyRouteComponent(() => import('../../../../components/document-builder/index')),
})

const reportDashBoard = createRoute({
    getParentRoute: () => ownerReportsRoute,
    path: 'dashboard',
    component: lazyRouteComponent(() => import('./dashboard/report-dashboard')),
})

const userTeller = createRoute({
    getParentRoute: () => ownerReportsRoute,
    path: 'user-teller',
    component: lazyRouteComponent(() => import('./user-teller/view-user-teller')),
})

const userTrialBalance = createRoute({
    getParentRoute: () => ownerReportsRoute,
    path: 'trial-balance',
    component: lazyRouteComponent(() => import('./trial-balance')),
})



const OwnerReportsRoute = ownerReportsRoute.addChildren([
    ReportDocumentPage,
    reportDashBoard,
    ownerCompanyIndexRoute,
    userTeller,
    userTrialBalance, 
])

export default OwnerReportsRoute

