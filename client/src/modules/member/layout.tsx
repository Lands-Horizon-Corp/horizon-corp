import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import MemberSidebar from './components/member-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'

const MemberLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Member']}>
            <div className="grid min-h-[100dvh] w-full grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
                <MemberSidebar />
                <main className="max-h-screen overflow-y-scroll">
                    <UserNav />
                    <Outlet />
                </main>
            </div>
        </AuthGuard>
    )
}

export default MemberLayout
