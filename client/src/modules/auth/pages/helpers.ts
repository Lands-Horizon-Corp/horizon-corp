import { AccountStatus, UserData } from '@/horizon-corp/types'
import { toast } from 'sonner'

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

export const toastAccountStatus = (accountStatus: AccountStatus) => {
    switch (accountStatus) {
        case 'Pending':
            return toast.warning('Account pending')
        case 'Verified':
            return toast.success('Account Verified')
        case 'Not Allowed':
            return toast.error(
                'Account canceled, You are not allowed to sign in.'
            )
        default:
            return toast.error('Could not identify accountr status')
    }
}
