import {
    redirect,
    createRoute,
    lazyRouteComponent,
} from '@tanstack/react-router'

import { employeeRoute } from '../../route'

export const employeeTransactionRoute = createRoute({
    path: '/transaction',
    getParentRoute: () => employeeRoute,
})

export const employeeTransactionIndexRoute = createRoute({
    path: '/',
    getParentRoute: () => employeeTransactionRoute,
    beforeLoad: () => {
        throw redirect({ to: '/employee/transaction/fund-movement' })
    },
})

export const employeeTransactionPaymentsEntry = createRoute({
    path: '/fund-movement',
    getParentRoute: () => employeeTransactionRoute,
    component: lazyRouteComponent(
        () => import('./transaction-payments-entry-page')
    ),
})

export const employeeTransactionPaymentType = createRoute({
    path: '/payment-types',
    getParentRoute: () => employeeTransactionRoute,
    component: lazyRouteComponent(
        () => import('./transaction-payment-type-page')
    ),
})

export const employeeTransactionMaintenance = createRoute({
    getParentRoute: () => employeeTransactionRoute,
    path: '/maintenance',
    component: lazyRouteComponent(
        () => import('./transaction-maintenance-page')
    ),
})

const EmployeeTransactionRoute = employeeTransactionRoute.addChildren([
    employeeTransactionIndexRoute,
    employeeTransactionPaymentType,
    employeeTransactionPaymentsEntry,
    employeeTransactionMaintenance,
])

export default EmployeeTransactionRoute
