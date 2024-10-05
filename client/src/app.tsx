import React, { useEffect } from 'react';
import { RouterProvider } from '@tanstack/react-router';

import router from '@/root-route';
import { ThemeProvider } from '@/components/providers/theme-provider';
import useErrorDetailsState from './horizon-corp/states/error-details-state';

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
    const { connection } = useErrorDetailsState();
    const checkConnection = async () => {
        const isConnected = await connection();
        console.log('Connection status:', isConnected);
    };

    useEffect(() => {
        checkConnection();
    }, [connection]);
    return (
        <ThemeProvider>
            <RouterProvider router={router} />
            <TanStackRouterDevtoolsPanel
                position="bottom-right"
                router={router}
            />
        </ThemeProvider>
    );
};

export default App;
