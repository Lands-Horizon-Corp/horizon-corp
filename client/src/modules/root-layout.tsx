import { Outlet } from '@tanstack/react-router'
import { Toaster } from '@/components/ui/sonner'

const RootLayout = () => {
    return (
        <div>
            <Outlet />
            <Toaster richColors />
        </div>
    )
}

export default RootLayout
