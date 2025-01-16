import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import AdminSidebar from './components/admin-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'
import { SidebarProvider } from '@/components/sidebar/sidebar-provider'

const AdminLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Admin']}>
            <div className="grid min-h-[100dvh] w-full grid-cols-[1fr] md:grid-cols-[auto_1fr]">
                <SidebarProvider>
                    <AdminSidebar />
                    <main className="ecoop-scroll relative max-h-screen overflow-y-scroll">
                        <UserNav homeUrl='/admin' />
                        <Outlet />
                    </main>
                </SidebarProvider>
            </div>
        </AuthGuard>
    )
}

export default AdminLayout
