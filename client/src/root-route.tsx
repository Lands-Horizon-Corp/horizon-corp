import { createRootRoute, createRouter } from '@tanstack/react-router';

import AuthRoute from '@/modules/auth/route';
import AdminRoute from '@/modules/admin/route';
import OwnerRoute from '@/modules/owner/route';
import MemberRoute from '@/modules/member/route';
import LandingRoute from '@/modules/landing/route';
import EmployeeRoute from '@/modules/employee/route';

import RootLayout from '@/modules/root-layout';
import NotFoundPage from '@/components/not-found';

export const rootRoute = createRootRoute();

const routeTree = rootRoute.addChildren([
    AuthRoute,
    AdminRoute,
    OwnerRoute,
    MemberRoute,
    LandingRoute,
    EmployeeRoute,
]);

const router = createRouter({
    routeTree,
    defaultComponent: RootLayout,
    defaultNotFoundComponent: NotFoundPage,
});

export default router;
