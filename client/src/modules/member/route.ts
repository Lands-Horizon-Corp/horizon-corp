import { createRoute } from '@tanstack/react-router'

import { rootRoute } from '@/root-route'
import MemberLayout from '@/modules/member/layout'
import MemberLandingPage from '@/modules/member/pages'

const memberRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'member',
    beforeLoad: () => {
        // TODO: Once middleware is implemented
    },
    component: MemberLayout,
})

const signUpRoute = createRoute({
    getParentRoute: () => memberRoute,
    path: '/',
    component: MemberLandingPage,
})

const MemberRoute = memberRoute.addChildren([signUpRoute])

export default MemberRoute
