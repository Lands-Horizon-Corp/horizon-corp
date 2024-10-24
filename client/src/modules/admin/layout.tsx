import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import AdminSidebar from './components/admin-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'

const AdminLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Admin']}>
            <div className="grid min-h-[100dvh] grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
                <AdminSidebar />
                <main className="">
                    <UserNav />
                    <Outlet />
                </main>
            </div>
        </AuthGuard>
    )
}

export default AdminLayout
