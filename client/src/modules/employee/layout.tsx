import { Outlet } from '@tanstack/react-router'
import EmployeeSidebar from './components/employee-sidebar'

const OwnerLayout = () => {
    return (
        <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
            <EmployeeSidebar />
            <main className="">
                <Outlet />
            </main>
        </div>
    )
}

export default OwnerLayout
