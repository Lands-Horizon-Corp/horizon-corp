import { Outlet } from '@tanstack/react-router'

import EmployeeNavbar from './components/employee-navbar'
import EmployeeSidebar from './components/employee-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'

const OwnerLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Employee']}>
            <div className="grid min-h-[100dvh] grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
                <EmployeeSidebar />
                <main>
                    <EmployeeNavbar />
                    <Outlet />
                </main>
            </div>
        </AuthGuard>
    )
}

export default OwnerLayout
