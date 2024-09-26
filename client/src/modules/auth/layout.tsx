import { Outlet } from '@tanstack/react-router'
import AuthFooter from './components/auth-footer'

interface Props {}

const AuthLayout = (_props: Props) => {
    return (
        <div className="grid min-h-[100dvh] grid-rows-[1fr_auto]">
            <Outlet />
            <AuthFooter />
        </div>
    )
}

export default AuthLayout
