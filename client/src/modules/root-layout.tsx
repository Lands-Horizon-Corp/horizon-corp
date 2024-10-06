import { Outlet } from '@tanstack/react-router'

import { Toaster } from '@/components/ui/sonner'
import ConnectionProvider from '@/components/providers/connection-provider'

const RootLayout = () => {
    return (
        <div>
            <Outlet />
            <Toaster richColors />
            <ConnectionProvider />
        </div>
    )
}

export default RootLayout
