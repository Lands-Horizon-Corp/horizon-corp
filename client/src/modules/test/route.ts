import { createRoute, lazyRouteComponent } from '@tanstack/react-router'

import { rootRoute } from '@/root-route'
import React from 'react'

const testRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: 'test',
    component: React.lazy(() => import('./layout')),
})

const testLandingRoute = createRoute({
    getParentRoute: () => testRoute,
    path: '/upload',
    component: React.lazy(() => import('./pages/upload')),
})

const imageDetailRoute = createRoute({
    getParentRoute: () => testRoute, // Sets `landingRoute` as the parent route
    path: `image/$imageId`, // Dynamic path with `id` parameter
    component: lazyRouteComponent(
        () => import('@/components/image-preview/image-details')
    ),
});

const TestRoute = testRoute.addChildren([testLandingRoute, imageDetailRoute])

export default TestRoute
