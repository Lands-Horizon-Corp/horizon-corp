import { createRoute } from '@tanstack/react-router'

import LandingPage from '@/modules/landing/pages'
import AboutPage from '@/modules/landing/pages/about'
import ContactPage from '@/modules/landing/pages/contact'
import DevelopersPage from '@/modules/landing/pages/developers'
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

const developersRoute = createRoute({
    getParentRoute: () => landingRoute,
    path: 'developers',
    component: DevelopersPage,
})

const LandingRoute = landingRoute.addChildren([
    aboutRoute,
    contactRoute,
    developersRoute,
    landingIndexRoute,
])

export default LandingRoute
