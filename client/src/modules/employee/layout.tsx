import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import AuthGuard from '@/components/wrappers/auth-guard'
import EmployeeSidebar from './components/employee-sidebar'
import { SidebarProvider } from '@/components/sidebar/sidebar-provider'

const OwnerLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Employee']}>
            <div className="grid min-h-[100dvh] w-full grid-cols-[1fr] md:grid-cols-[auto_1fr]">
                <SidebarProvider>
                    <EmployeeSidebar />
                    <main className="ecoop-scroll max-h-screen overflow-y-scroll">
                        <UserNav />
                        <Outlet />
                    </main>
                </SidebarProvider>
            </div>
        </AuthGuard>
    )
}

export default OwnerLayout
