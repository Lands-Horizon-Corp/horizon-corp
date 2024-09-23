import { createRoute, redirect } from '@tanstack/react-router'

import Verify from '@/modules/auth/pages/verify'
import NotFoundPage from '@/modules/auth/not-found'
import SignUpPage from '@/modules/auth/pages/sign-up'
import SignInPage from '@/modules/auth/pages/sign-in'

import AuthLayout from './layout'

import { rootRoute } from '@/root-route'

const authRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'auth',
    component: AuthLayout,
    beforeLoad: ({ location }) => {
        if (location.pathname === '/auth')
            throw redirect({
                to: '/auth/sign-in',
            })
    },
    notFoundComponent: NotFoundPage,
})

const signUpRoute = createRoute({
    getParentRoute: () => authRoute,
    path: 'sign-up',
    component: SignUpPage,
})

const signInRoute = createRoute({
    getParentRoute: () => authRoute,
    path: 'sign-in',
    component: SignInPage,
})

const verifyRoute = createRoute({
    getParentRoute : () => authRoute,
    path : "verify/$verify_id",
    component : Verify
})

const AuthRoute = authRoute.addChildren([signUpRoute, signInRoute, verifyRoute])

export default AuthRoute
