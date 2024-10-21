import { Outlet } from '@tanstack/react-router'

import EmployeeNavbar from './components/employee-navbar'
import EmployeeSidebar from './components/employee-sidebar'
import ProtectedRouteWrapper from '@/components/wrappers/protected-route-wrapper'

const OwnerLayout = () => {
    return (
        <ProtectedRouteWrapper allowedAccountTypes={['Member']}>
            <div className="grid min-h-[100dvh] grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
                <EmployeeSidebar />
                <main>
                    <EmployeeNavbar />
                    <Outlet />
                </main>
            </div>
        </ProtectedRouteWrapper>
    )
}

export default OwnerLayout
