import { QueryClient } from '@tanstack/react-query';
import { createRouter } from '@tanstack/react-router';

import AuthRoute from '@/modules/auth/route';
import AdminRoute from '@/modules/admin/route';
import OwnerRoute from '@/modules/owner/route';
import MemberRoute from '@/modules/member/route';
import LandingRoute from '@/modules/landing/route';
import EmployeeRoute from '@/modules/employee/route';
import TestRoute from '@/modules/test/route';

import RootLayout from '@/modules/root-layout';
import NotFoundPage from '@/components/not-found';
import practiceLandingRoute from './public/route';
import { createRootRouteWithContext } from '@tanstack/react-router';

export type TRouterContext = {
    queryClient: QueryClient;
};

export const rootRoute = createRootRouteWithContext<TRouterContext>()();

const routeTree = rootRoute.addChildren([
    AuthRoute,
    AdminRoute,
    OwnerRoute,
    MemberRoute,
    LandingRoute,
    EmployeeRoute,
    TestRoute,
    practiceLandingRoute,
]);

const router = (queryClient: QueryClient) =>
    createRouter({
        routeTree,
        context: {
            queryClient,
        },
        defaultComponent: RootLayout,
        defaultNotFoundComponent: NotFoundPage,
    });

export default router;
