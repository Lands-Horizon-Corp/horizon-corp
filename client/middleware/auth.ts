// import { defineNuxtRouteMiddleware, navigateTo, useRoute } from '#app';
// import { useAuth } from '~/composables/useAuth';

// export default defineNuxtRouteMiddleware(() => {
//   const auth = useAuth();
//   const route = useRoute();

//   if (!auth.value.isAuthenticated) {
//     return navigateTo('/login');
//   }

//   if (route.path !== '/dashboard') {
//     return navigateTo('/dashboard');
//   }
// });

import { defineNuxtRouteMiddleware, navigateTo } from '#app';
import { useUserStore } from '@/store/';

export default defineNuxtRouteMiddleware(() => {
  const userStore = useUserStore();

  if (!userStore.isLoggedIn && ['/login', '/register'].includes(useRoute().path)) {
    return;
  }

  if (!userStore.isLoggedIn) {
    navigateTo('/login');
  } else if (['/login', '/register'].includes(useRoute().path)) {
    navigateTo('/');
  }
});
