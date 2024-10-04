import { Outlet } from '@tanstack/react-router'
import { Toaster } from '@/components/ui/sonner'
import { VersionAndFeedBack } from '@/components/version'

const RootLayout = () => {
    return (
        <div className="relative">
            <VersionAndFeedBack />
            <Outlet />
            <Toaster richColors />
        </div>
    )
}

export default RootLayout
