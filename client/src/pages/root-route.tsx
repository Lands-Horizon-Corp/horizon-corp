import { createRootRoute, Outlet } from '@tanstack/react-router'

import publicRoutes from './(public)/public-route'

export const rootRoute = createRootRoute({
    component: () => (
        <>
            <Outlet />
        </>
    ),
})

export const routeTree = rootRoute.addChildren([...publicRoutes])
