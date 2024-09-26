import { createRoute, redirect } from '@tanstack/react-router'
import { zodSearchValidator } from '@tanstack/router-zod-adapter'

import Verify from '@/modules/auth/pages/verify'
import NotFoundPage from '@/modules/auth/not-found'
import SignUpPage from '@/modules/auth/pages/sign-up'
import SignInPage, {
    SignInPageSearchSchema,
} from '@/modules/auth/pages/sign-in'

import AuthLayout from './layout'

import { rootRoute } from '@/root-route'
import ForgotPasswordPage, {
    ForgotPasswordPageSearchSchema,
} from './pages/forgot-password'
import PasswordResetPage from './pages/password-reset'

const authRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'auth',
    component: AuthLayout,
    beforeLoad: ({ location }) => {
        if (location.pathname === '/auth' || location.pathname === '/auth/')
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
    validateSearch: zodSearchValidator(SignInPageSearchSchema),
})

const verifyRoute = createRoute({
    getParentRoute: () => authRoute,
    path: 'verify',
    component: Verify,
})

const forgotPasswordRoute = createRoute({
    getParentRoute: () => authRoute,
    path: 'forgot-password',
    component: ForgotPasswordPage,
    validateSearch: zodSearchValidator(ForgotPasswordPageSearchSchema),
})

const passwordResetRoute = createRoute({
    getParentRoute: () => authRoute,
    path: 'password-reset/$resetId',
    component: PasswordResetPage,
})

const AuthRoute = authRoute.addChildren([
    signUpRoute,
    signInRoute,
    verifyRoute,
    forgotPasswordRoute,
    passwordResetRoute,
])

export default AuthRoute
