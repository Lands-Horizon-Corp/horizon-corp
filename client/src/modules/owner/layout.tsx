import { Outlet } from '@tanstack/react-router'

import OwnerSidebar from './components/owner-sidebar'
import ProtectedRouteWrapper from '@/components/wrappers/protected-route-wrapper'

const OwnerLayout = () => {
    return (
        <ProtectedRouteWrapper allowedAccountTypes={['Owner']}>
            <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
                <OwnerSidebar />
                <main className="">
                    <Outlet />
                </main>
            </div>
        </ProtectedRouteWrapper>
    )
}

export default OwnerLayout
