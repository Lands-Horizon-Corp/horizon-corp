import { Outlet } from '@tanstack/react-router'

import { Toaster } from '@/components/ui/sonner'
import ConnectionProvider from '@/components/providers/connection-provider'
import CookieConsentProvider from '@/components/providers/cookie-consent-provider'

const RootLayout = () => {
    return (
        <div>
            <CookieConsentProvider />
            <Outlet />
            <Toaster richColors />
            <ConnectionProvider />
        </div>
    )
}

export default RootLayout
