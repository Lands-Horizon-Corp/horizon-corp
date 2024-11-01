import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import OwnerSidebar from './components/owner-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'
import { SidebarProvider } from '@/components/sidebar/sidebar-provider'

const OwnerLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Owner']}>
            <div className="grid min-h-[100dvh] w-full grid-cols-[1fr] sm:grid-cols-[auto_1fr]">
                <SidebarProvider>
                    <OwnerSidebar />
                    <main className="ecoop-scroll max-h-screen overflow-y-scroll">
                        <UserNav />
                        <Outlet />
                    </main>
                </SidebarProvider>
            </div>
        </AuthGuard>
    )
}

export default OwnerLayout
