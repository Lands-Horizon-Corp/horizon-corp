import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import AuthGuard from '@/components/wrappers/auth-guard'
import EmployeeSidebar from './components/employee-sidebar'

const OwnerLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Employee']}>
            <div className="grid min-h-[100dvh] w-full grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
                <EmployeeSidebar />
                <main className="overflow-x-hidden">
                    <UserNav />
                    <Outlet />
                </main>
            </div>
        </AuthGuard>
    )
}

export default OwnerLayout
