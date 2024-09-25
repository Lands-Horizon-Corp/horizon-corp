import { useState } from 'react'
import { useRouter } from '@tanstack/react-router'

import SignInForm from '../components/forms/sign-in-form'
import AccountCancelled from '../components/account-cancelled'

import { UserBase, UserStatus } from '@/types'

interface Props {}

const SignInPage = (_props: Props) => {
    const router = useRouter()
    const [userData, setUserData] = useState<null | UserBase>(null)

    const onSignInSuccess = (userData: UserBase) => {
        const { validContactNumber, validEmail, status } = userData
        if (status === UserStatus['Not allowed']) return setUserData(userData)

        if (
            status === UserStatus.Pending &&
            (!validContactNumber || !validEmail)
        ) {
            router.navigate({ to: '/auth/verify' })
        }

        if (status === UserStatus.Verified) {
            // TODO
            // Navigate to respective page
            // owner
            // member
            // admin
            // employee
        }
    }

    const handleOnCancelBack = () => {
        // TODO:
        // Sign Out user since he/she is restricted
        // unset userData
        setUserData(null)
    }

    return (
        <div className="flex flex-1 justify-center px-4 py-4">
            {!userData && <SignInForm onSuccess={onSignInSuccess} />}
            {userData && (
                <AccountCancelled
                    userData={userData}
                    onBack={handleOnCancelBack}
                />
            )}
        </div>
    )
}

export default SignInPage
