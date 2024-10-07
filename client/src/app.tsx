import React from 'react';
import { RouterProvider } from '@tanstack/react-router';

import router from '@/root-route';
import { ThemeProvider } from '@/components/providers/theme-provider';

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
