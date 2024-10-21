import { Outlet } from '@tanstack/react-router'

import AuthNavBar from '@/components/navbars/auth-nav'
import AuthFooter from '@/components/footers/auth-footer'

const AuthLayout = () => {
    return (
        <div className="grid min-h-[100dvh] grid-rows-[1fr_auto] bg-cover bg-fixed sm:bg-[url('/auth-bg.webp')]">
            <AuthNavBar />
            <main className="mt-24">
                <Outlet />
            </main>
            <AuthFooter />
        </div>
    )
}

export default AuthLayout
