import { createRootRoute } from '@tanstack/react-router';

import AuthRoute from '@/modules/auth/route';
import LandingRoute from '@/modules/landing/route';
import MemberRoute from '@/modules/member/route';

export const rootRoute = createRootRoute();

export const routeTree = rootRoute.addChildren([
    AuthRoute,
    LandingRoute,
    MemberRoute,
]);
