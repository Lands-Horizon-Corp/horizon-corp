import { Outlet } from '@tanstack/react-router'
import AuthFooter from './components/auth-footer'

interface Props {}

const AuthLayout = (_props: Props) => {
    return (
        <div>
            <Outlet />
            <AuthFooter />
        </div>
    )
}

export default AuthLayout
