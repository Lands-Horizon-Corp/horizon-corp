import { createRoute } from '@tanstack/react-router'

import { rootRoute } from '@/root-route'
import MemberLandingPage from '@/modules/member/pages'

const memberRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'member',
    beforeLoad: () => {
        const user = false
        if (!user) alert('Bawal ka dito desyo')
    },
})

const signUpRoute = createRoute({
    getParentRoute: () => memberRoute,
    path: '/',
    component: MemberLandingPage,
})

const MemberRoute = memberRoute.addChildren([signUpRoute])

export default MemberRoute
