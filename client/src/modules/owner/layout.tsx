import { Outlet } from '@tanstack/react-router'

import OwnerSidebar from './components/owner-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'

const OwnerLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Owner']}>
            <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
                <OwnerSidebar />
                <main className="">
                    <Outlet />
                </main>
            </div>
        </AuthGuard>
    )
}

export default OwnerLayout
