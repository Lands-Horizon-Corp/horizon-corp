import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import AdminSidebar from './components/admin-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'

const AdminLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Admin']}>
            <div className="grid min-h-[100dvh] w-full grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
                <AdminSidebar />
                <main className="max-h-screen overflow-y-scroll">
                    <UserNav />
                    <Outlet />
                </main>
            </div>
        </AuthGuard>
    )
}

export default AdminLayout
