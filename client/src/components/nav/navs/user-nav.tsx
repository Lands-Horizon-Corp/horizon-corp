import RootNav from '@/components/nav/root-nav'
import NavContainer from '@/components/nav/nav-container'
import NavTimeInBar from '../nav-components/nav-time-in-bar'
import NavThemeToggle from '@/components/nav/nav-components/nav-theme-toggle'
import NavProfileMenu from '@/components/nav/nav-components/nav-profile-menu'

import { useUserAuthStore } from '@/store/user-auth-store'

const UserNav = () => {
    const { currentUser } = useUserAuthStore()
    return (
        <RootNav className="pointer-events-none justify-end lg:px-4">
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
