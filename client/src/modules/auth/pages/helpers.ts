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
