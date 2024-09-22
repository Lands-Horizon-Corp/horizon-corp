import React from 'react';
import { RouterProvider } from '@tanstack/react-router';

import router from '@/root-route';

const TanStackRouterDevtoolsPanel =
    process.env.NODE_ENV === 'production'
        ? () => null // Render nothing in production
        : React.lazy(() =>
              // Lazy load in development
              import('@tanstack/router-devtools').then((res) => ({
                  default: res.TanStackRouterDevtools,
                  // For Embedded Mode
                  // default: res.TanStackRouterDevtoolsPanel
              }))
          );

// for type safety hahaha
declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router;
    }
}

const App = () => {
    return (
        <>
            <RouterProvider router={router} />
            <TanStackRouterDevtoolsPanel router={router} />
        </>
    );
};

export default App;
