import React from 'react';
import { RouterProvider } from '@tanstack/react-router';

import router from '@/root-route';
import { ThemeProvider } from '@/providers/theme-provider';
import useCurrentUser from './hooks/use-current-user';

const TanStackRouterDevtoolsPanel =
    process.env.NODE_ENV === 'production'
        ? () => null
        : React.lazy(() =>
              import('@tanstack/router-devtools').then((res) => ({
                  default: res.TanStackRouterDevtools,
              }))
          );

declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router;
    }
}

const App = () => {
    useCurrentUser({ loadOnMount: true });

    return (
        <ThemeProvider>
            <RouterProvider router={router} />
            <TanStackRouterDevtoolsPanel
                position="bottom-left"
                router={router}
            />
        </ThemeProvider>
    );
};

export default App;
