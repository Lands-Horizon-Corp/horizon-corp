import { Outlet } from '@tanstack/react-router'

import UserNav from '@/components/nav/navs/user-nav'
import AuthGuard from '@/components/wrappers/auth-guard'
import { SidebarInset, SidebarProvider } from '@/components/ui/sidebar'
import OwnerSidebar from './components/owner-sidebar'

const OwnerLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Owner']}>
            <SidebarProvider>
                <OwnerSidebar />
                <SidebarInset className="ecoop-scroll max-h-[98vh] w-full overflow-y-auto">
                    <UserNav className="sticky top-0 bg-background" />
                    <main className="">
                        <Outlet />
                    </main>
                </SidebarInset>
            </SidebarProvider>
        </AuthGuard>
    )
}

export default OwnerLayout
{
    /* <main className="ecoop-scroll max-h-screen overflow-y-scroll"> */
}
{
    /* 
                    </main> */
}
