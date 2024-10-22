import { Outlet } from '@tanstack/react-router'

import MemberSidebar from './components/member-sidebar'
import ProtectedRouteWrapper from '@/components/wrappers/protected-route-wrapper'

const MemberLayout = () => {
    return (
        <ProtectedRouteWrapper allowedAccountTypes={['Member']}>
            <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
                <MemberSidebar />
                <main className="">
                    <Outlet />
                </main>
            </div>
        </ProtectedRouteWrapper>
    )
}

export default MemberLayout
