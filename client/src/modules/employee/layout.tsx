import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import AuthGuard from '@/components/wrappers/auth-guard'
import EmployeeSidebar from './components/employee-sidebar'
import { SidebarProvider } from '@/components/sidebar/sidebar-provider'

const OwnerLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Employee']}>
            <div className="grid min-h-[100dvh] w-full grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
                <EmployeeSidebar />
                <main className="max-h-screen overflow-y-scroll ecoop-scroll">
                    <UserNav />
                    <Outlet />
                </main>
            </div>
        </AuthGuard>
    )
}

export default OwnerLayout
