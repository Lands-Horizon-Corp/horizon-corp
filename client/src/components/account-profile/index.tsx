import AccountProfileBanner from './account-profile-banner'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { useUserAuthStore } from '@/store/user-auth-store'
import AccountProfileNavigation from './account-profile-navigation'
import AccountSettings from './account-settings'

const AccountProfile = ({ className }: IBaseCompNoChild) => {
    const { currentUser, setCurrentUser } = useUserAuthStore()

    if (!currentUser) return <span>Unauthorized</span>

    return (
        <div
            className={cn(
                'flex w-full max-w-full flex-1 flex-col gap-y-4 sm:max-w-4xl',
                className
            )}
        >
            <AccountProfileBanner
                currentUser={currentUser}
                updateUserData={setCurrentUser}
            />
            <AccountProfileNavigation />
            <AccountSettings />
        </div>
    )
}

export default AccountProfile
