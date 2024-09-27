import { Outlet } from '@tanstack/react-router'
import AuthFooter from './components/auth-footer'
import AuthNavBar from './components/auth-navbar'

interface Props {}

const AuthLayout = (_props: Props) => {
    return (
        <div className="grid min-h-[100dvh] grid-rows-[1fr_auto]">
            <AuthNavBar />
            <main className="mt-16">
                <Outlet />
            </main>
            <AuthFooter />
        </div>
    )
}

export default AuthLayout
