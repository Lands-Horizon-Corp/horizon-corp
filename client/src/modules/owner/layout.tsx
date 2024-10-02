import { Outlet } from '@tanstack/react-router'
import OwnerSidebar from './components/owner-sidebar'

const OwnerLayout = () => {
    return (
        <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
            <OwnerSidebar />
            <main className="">
                <Outlet />
            </main>
        </div>
    )
}

export default OwnerLayout
