import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import MemberSidebar from './components/member-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'
import { SidebarProvider } from '@/components/sidebar/sidebar-provider'

const MemberLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Member']}>
            <div className="grid min-h-[100dvh] w-full grid-cols-[1fr] md:grid-cols-[auto_1fr]">
                <SidebarProvider>
                    <MemberSidebar />
                    <main className="ecoop-scroll max-h-screen overflow-y-scroll">
                        <UserNav />
                        <Outlet />
                    </main>
                </SidebarProvider>
            </div>
        </AuthGuard>
    )
}

export default MemberLayout
