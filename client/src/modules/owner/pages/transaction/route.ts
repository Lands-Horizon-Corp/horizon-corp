import { createRoute, redirect } from '@tanstack/react-router'
import { ownerRoute } from '../../route'
import { lazyRouteComponent } from '@tanstack/react-router'

export const ownerTransactionRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: '/transaction',
})

export const ownerTransactionIndexRoute = createRoute({
    getParentRoute: () => ownerTransactionRoute,
    path: '/',
    beforeLoad: () => {
        throw redirect({ to: '/owner/transaction/payments-entry' })
    },
})

export const ownerTransactionPaymentsEntry = createRoute({
    getParentRoute: () => ownerTransactionRoute,
    path: '/payments-entry',
    component: lazyRouteComponent(() => import('./payments-entry-page')),
})

export const ownerTransactionDepositEntry = createRoute({
    getParentRoute: () => ownerTransactionRoute,
    path: '/deposit-entry',
    component: lazyRouteComponent(() => import('./deposit-entry-page')),
})

export const ownerTransactionWithdrawal = createRoute({
    getParentRoute: () => ownerTransactionRoute,
    path: '/withdrawals',
    component: lazyRouteComponent(() => import('./withdrawals-page')),
})

export const ownerTransactionType = createRoute({
    getParentRoute: () => ownerTransactionRoute,
    path: '/transaction-type',
    component: lazyRouteComponent(() => import('./transaction-type-page')),
})

export const OwnerTransactionRoute = ownerTransactionRoute.addChildren([
    ownerTransactionType,
    ownerTransactionIndexRoute,
    ownerTransactionWithdrawal,
    ownerTransactionDepositEntry,
    ownerTransactionPaymentsEntry,
])
