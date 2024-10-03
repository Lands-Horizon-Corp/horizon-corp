import { createRootRoute, createRouter } from '@tanstack/react-router';

import AuthRoute from '@/modules/auth/route';
import OwnerRoute from '@/modules/owner/route';
import MemberRoute from '@/modules/member/route';
import LandingRoute from '@/modules/landing/route';

import RootLayout from '@/modules/root-layout';
import NotFoundPage from '@/components/not-found';

export const rootRoute = createRootRoute();

const routeTree = rootRoute.addChildren([
    AuthRoute,
    LandingRoute,
    OwnerRoute,
    MemberRoute,
]);

const router = createRouter({
    routeTree,
    defaultComponent: RootLayout,
    defaultNotFoundComponent: NotFoundPage,
});

export default router;
