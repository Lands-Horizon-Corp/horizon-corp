import { Outlet } from '@tanstack/react-router'

import { Toaster } from '@/components/ui/sonner'
import CookieConsent from '@/components/cookie-consent'
import { VersionAndFeedBack } from '@/components/version'
import ConnectionProvider from '@/providers/connection-provider'
import ConfirmModal from '@/components/modals/confirm-modal'

const RootLayout = () => {
    return (
        <div className="relative">
            <VersionAndFeedBack />
            <Outlet />
            <Toaster richColors />
            <ConnectionProvider />
            <CookieConsent />
            <ConfirmModal />
        </div>
    )
}

export default RootLayout
