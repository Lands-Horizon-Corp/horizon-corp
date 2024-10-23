import { Outlet } from '@tanstack/react-router'

import { Toaster } from '@/components/ui/sonner'
import CookieConsent from '@/components/cookie-consent'
import ConfirmModal from '@/components/modals/confirm-modal'
import ConnectionProvider from '@/providers/connection-provider'
import useFetchCurrentUser from '@/hooks/use-fetch-current-user'

const RootLayout = () => {

    useFetchCurrentUser();

    return (
        <div className="relative">
            <Outlet />
            <Toaster richColors />
            <ConnectionProvider />
            <CookieConsent />
            <ConfirmModal />
        </div>
    )
}

export default RootLayout
