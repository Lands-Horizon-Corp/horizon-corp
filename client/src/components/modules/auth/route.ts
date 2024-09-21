import { createRoute } from '@tanstack/react-router'

import SignUpPage from '@/modules/auth/pages/sign-up'
import SignInPage from '@/modules/auth/pages/sign-in'

import { rootRoute } from '@/root-route'

const authRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'auth',
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
