import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { useEffect, useState } from 'react';
import { RouterProvider } from '@tanstack/react-router';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';

import router from '@/root-route';
import { ThemeProvider } from '@/providers/theme-provider';
import { EnvironmentManager } from './manager';

// const TanStackRouterDevtoolsPanel =
//     process.env.NODE_ENV === 'production'
//         ? () => null
//         : React.lazy(() =>
//               import('@tanstack/router-devtools').then((res) => ({
//                   default: res.TanStackRouterDevtools,
//               }))
//           );

declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router;
    }
}

const App = () => {
    const [queryClient] = useState(new QueryClient());
    useEffect(() => {
        const initializeEnv = async () => await EnvironmentManager.initialize();
        initializeEnv();
    }, []);
    return (
        <QueryClientProvider client={queryClient}>
            <ThemeProvider>
                <RouterProvider router={router} />
                {/* <TanStackRouterDevtoolsPanel
                    position="bottom-left"
                    router={router}
                /> */}
            </ThemeProvider>
            <ReactQueryDevtools initialIsOpen={false} />
        </QueryClientProvider>
    );
};

export default App;
