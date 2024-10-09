import { Outlet } from '@tanstack/react-router'
import AdminNavbar from './components/admin-navbar'
import AdminSidebar from './components/admin-sidebar'

const AdminLayout = () => {
    return (
        <div className="grid min-h-[100dvh] grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
            <AdminSidebar />
            <main className="">
                <AdminNavbar />
                <Outlet />
            </main>
        </div>
    )
}

export default AdminLayout
