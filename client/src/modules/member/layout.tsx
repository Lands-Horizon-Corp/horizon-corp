import { Outlet } from '@tanstack/react-router'

import MemberSidebar from './components/member-sidebar'
import AuthGuard from '@/components/wrappers/auth-guard'

const MemberLayout = () => {
    return (
        <AuthGuard allowedAccountTypes={['Member']}>
            <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
                <MemberSidebar />
                <main className="">
                    <Outlet />
                </main>
            </div>
        </AuthGuard>
    )
}

export default MemberLayout
