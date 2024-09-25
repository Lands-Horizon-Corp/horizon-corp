import { createRoute } from '@tanstack/react-router'

import LandingPage from '@/modules/landing/pages'
import AboutPage from '@/modules/landing/pages/about'
import ContactPage from '@/modules/landing/pages/contact'

import PublicLayout from '@/modules/landing/layout'

import { rootRoute } from '@/root-route'

const landingRoute = createRoute({
    getParentRoute: () => rootRoute,
    id: 'public',
    component: PublicLayout,
})

const landingIndexRoute = createRoute({
    getParentRoute: () => landingRoute,
    path: '/',
    component: LandingPage,
})

const aboutRoute = createRoute({
    getParentRoute: () => landingRoute,
    path: 'about',
    component: AboutPage,
})

const contactRoute = createRoute({
    getParentRoute: () => landingRoute,
    path: 'contact',
    component: ContactPage,
})

const LandingRoute = landingRoute.addChildren([
    aboutRoute,
    contactRoute,
    landingIndexRoute,
])

export default LandingRoute
