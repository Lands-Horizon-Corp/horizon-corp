import { Outlet } from '@tanstack/react-router'
import AdminSidebar from './components/admin-sidebar'

const AdminLayout = () => {
    return (
        <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
            <AdminSidebar />
            <main className="">
                <Outlet />
            </main>
        </div>
    )
}

export default AdminLayout
