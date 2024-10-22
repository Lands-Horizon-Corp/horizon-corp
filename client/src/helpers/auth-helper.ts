import { UserData } from '@/horizon-corp/types'

export const getUsersAccountTypeRedirectPage = (currentUser: UserData) => {
    const { accountType } = currentUser
    switch (accountType) {
        case 'Admin':
            return '/admin'
        case 'Member':
            return '/member'
        case 'Employee':
            return '/employee'
        case 'Owner':
            return '/owner'
        default:
            return '/'
    }
}

export const isUserUnverified = (userData: UserData) => {
    return userData.status !== 'Verified'
}

export const isUserHasUnverified = (userData: UserData) => {
    return (
        userData.status === 'Pending' ||
        !userData.isEmailVerified ||
        !userData.isContactVerified
    )
}
