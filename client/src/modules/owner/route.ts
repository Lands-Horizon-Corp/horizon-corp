import { createRoute } from '@tanstack/react-router'

import OwnerLandingPage from '@/modules/owner/pages'

import { rootRoute } from '@/root-route'
import OwnerLayout from '@/modules/owner/layout'

const ownerRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'owner',
    beforeLoad: () => {
        // TODO: Once middleware is implemented
    },
    component : OwnerLayout
})

const memberLandingRoute = createRoute({
    getParentRoute: () => ownerRoute,
    path: '/',
    component: OwnerLandingPage,
})


const OwnerRoute = ownerRoute.addChildren([
    memberLandingRoute,
])

export default OwnerRoute
