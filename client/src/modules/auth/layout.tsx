import { Outlet } from '@tanstack/react-router'
import AuthFooter from './components/auth-footer'
import AuthNavBar from './components/auth-navbar'

interface Props {}

const AuthLayout = (_props: Props) => {
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
