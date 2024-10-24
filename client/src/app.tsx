import { useEffect, useState } from 'react';
import { RouterProvider } from '@tanstack/react-router';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import router from '@/root-route';
import { EnvironmentManager } from './manager';
import { ThemeProvider } from '@/providers/theme-provider';

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
            </ThemeProvider>
            <ReactQueryDevtools initialIsOpen={false} />
        </QueryClientProvider>
    );
};

export default App;
