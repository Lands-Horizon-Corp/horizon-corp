import { IUserData } from '@/server/types'

export const getUsersAccountTypeRedirectPage = (currentUser: IUserData) => {
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

export const isUserUnverified = (userData: IUserData) => {
    return userData.status !== 'Verified'
}

export const isUserHasUnverified = (userData: IUserData) => {
    return (
        userData.status === 'Pending' ||
        !userData.isEmailVerified ||
        !userData.isContactVerified
    )
}
