import { createRoute, redirect } from '@tanstack/react-router'

import NotFoundPage from '@/modules/auth/not-found'
import SignUpPage from '@/modules/auth/pages/sign-up'
import SignInPage from '@/modules/auth/pages/sign-in'

import { rootRoute } from '@/root-route'
import AuthLayout from './layout'

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

const AuthRoute = authRoute.addChildren([signUpRoute, signInRoute])

export default AuthRoute
