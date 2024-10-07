import { Outlet } from '@tanstack/react-router'

import { Toaster } from '@/components/ui/sonner'
import CookieConsent from '@/components/cookie-consent'
import ConnectionProvider from '@/components/providers/connection-provider'

const RootLayout = () => {
    return (
        <div>
            <Outlet />
            <Toaster richColors />
            <ConnectionProvider />
            <CookieConsent />
        </div>
    )
}

export default RootLayout
