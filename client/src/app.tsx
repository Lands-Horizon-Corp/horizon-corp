import { useState } from 'react';
import { RouterProvider } from '@tanstack/react-router';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import createRouter from '@/root-route';
import { ThemeProvider } from '@/providers/theme-provider';

declare module '@tanstack/react-router' {
    interface Register {
        router: ReturnType<typeof createRouter>;
    }
}

const App = () => {
    const [queryClient] = useState(new QueryClient());
    return (
        <QueryClientProvider client={queryClient}>
            <ThemeProvider>
                <RouterProvider router={createRouter(queryClient)} />
            </ThemeProvider>
            <ReactQueryDevtools initialIsOpen={false} />
        </QueryClientProvider>
    );
};

export default App;
