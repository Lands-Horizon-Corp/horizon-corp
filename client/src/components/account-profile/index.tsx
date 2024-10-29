import AccountProfileBanner from './account-profile-header'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { useUserAuthStore } from '@/store/user-auth-store'
import AccountProfileNavigation from './account-profile-navigation'

const AccountProfile = ({ className }: IBaseCompNoChild) => {
    const { currentUser, setCurrentUser } = useUserAuthStore()

    if (!currentUser) return <span>Unauthorized</span>

    return (
        <div className={cn('w-full max-w-4xl space-y-4 flex-1', className)}>
            <AccountProfileBanner
                currentUser={currentUser}
                updateUserData={setCurrentUser}
            />
            <AccountProfileNavigation />
        </div>
    )
}

export default AccountProfile
