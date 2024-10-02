import { Outlet } from '@tanstack/react-router'
import MemberSidebar from './components/member-sidebar'

const MemberLayout = () => {
    return (
        <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
            <MemberSidebar />
            <main className="">
                <Outlet />
            </main>
        </div>
    )
}

export default MemberLayout
