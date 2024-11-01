import RootNav from '@/components/nav/root-nav'
import NavContainer from '@/components/nav/nav-container'
import NavTimeInBar from '../nav-components/nav-time-in-bar'
import NavThemeToggle from '@/components/nav/nav-components/nav-theme-toggle'
import NavProfileMenu from '@/components/nav/nav-components/nav-profile-menu'

import { useUserAuthStore } from '@/store/user-auth-store'
import SidebarMobileToggle from '@/components/sidebar/sidebar-mobile-toggle'

const UserNav = () => {
    const { currentUser } = useUserAuthStore()
    return (
        <RootNav className="pointer-events-none relative justify-between lg:px-4">
            <NavContainer className="pointer-events-auto">
                <SidebarMobileToggle className="md:hidden" />
            </NavContainer>
            <NavContainer className="pointer-events-auto">
                {currentUser &&
                    ['Admin', 'Employee'].includes(currentUser.accountType) && (
                        <NavTimeInBar currentUser={currentUser} />
                    )}
                <NavProfileMenu />
                <NavThemeToggle />
            </NavContainer>
        </RootNav>
    )
}

export default UserNav
