import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import AdminSidebar from './components/admin-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'
import { SidebarInset, SidebarProvider } from '@/components/ui/sidebar'

const AdminLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Admin']}>
            <SidebarProvider>
                <AdminSidebar />
                <SidebarInset className="ecoop-scroll max-h-[98vh] w-full overflow-y-auto">
                    <UserNav className="sticky top-0 bg-background" />
                    <main className="relative">
                        <Outlet />
                    </main>
                </SidebarInset>
            </SidebarProvider>
        </AuthGuard>
    )
}

export default AdminLayout
