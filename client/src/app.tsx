import { createRouter, RouterProvider } from '@tanstack/react-router';

import { routeTree } from '@/root-route';

const router = createRouter({ routeTree });

// Register the router instance for type safety
declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router;
    }
}

const App = () => {
    return <RouterProvider router={router} />;
};

export default App;
