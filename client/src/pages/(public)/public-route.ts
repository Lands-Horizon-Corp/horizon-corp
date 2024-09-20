import { createRoute } from '@tanstack/react-router'

import LandingPage from '.'
import AboutPage from './about'
import ContactPage from './contact'
import authRoute from './auth/auth-route'
import { rootRoute } from '@/pages/root-route'

import PublicLayout from './layout'

export const publicRoute = createRoute({
    getParentRoute: () => rootRoute,
    id: 'public-layout',
    component: PublicLayout,
})

export const publicIndexRoute = createRoute({
    getParentRoute: () => publicRoute,
    path: '/',
    component: LandingPage,
})

const aboutRoute = createRoute({
    getParentRoute: () => publicRoute,
    path: 'about',
    component: AboutPage,
})

const contactRoute = createRoute({
    getParentRoute: () => publicRoute,
    path: 'contact',
    component: ContactPage,
})

const publicRoutes = [
    publicRoute.addChildren([publicIndexRoute, aboutRoute, contactRoute]),
    authRoute,
]

export default publicRoutes
