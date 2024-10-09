import { Outlet } from '@tanstack/react-router'
import EmployeeSidebar from './components/employee-sidebar'
import EmployeeNavbar from './components/employee-navbar'

const OwnerLayout = () => {
    return (
        <div className="grid min-h-[100dvh] grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
            <EmployeeSidebar />
            <main>
                <EmployeeNavbar />
                <Outlet />
            </main>
        </div>
    )
}

export default OwnerLayout
