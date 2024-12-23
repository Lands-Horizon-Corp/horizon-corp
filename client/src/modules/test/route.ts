import { createRoute } from '@tanstack/react-router'

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

const documentBuilderRoute = createRoute({
    getParentRoute: () => testRoute,
    path: '/document',
    component: React.lazy(() => import('./pages/document')),
})

// const TableRoute = createRoute({
//     getParentRoute: () => testRoute,
//     path: '/tbl',
//     component: React.lazy(() => import('./pages/tbl')),
// })

const signatureRoute = createRoute({
    getParentRoute: () => testRoute,
    path: '/signature',
    component: React.lazy(() => import('./pages/signature')),
})



const testPageRoute = createRoute({
    getParentRoute: () => testRoute,
    path: '/page',
    component: React.lazy(() => import('./pages/testPaging')),
})


const TestRoute = testRoute.addChildren([testLandingRoute, documentBuilderRoute, signatureRoute, testPageRoute])

export default TestRoute
