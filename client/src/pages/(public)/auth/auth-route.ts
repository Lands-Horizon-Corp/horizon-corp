import { createRoute } from '@tanstack/react-router'

import SignUpPage from './sign-up'
import SignInPage from './sign-in'
import { rootRoute } from '@/pages/root-route'

export const authRoute = createRoute({
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

const authRoutes = authRoute.addChildren([signUpRoute, signInRoute])

export default authRoutes
