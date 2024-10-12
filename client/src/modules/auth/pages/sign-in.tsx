import z from 'zod'
import { useState } from 'react'
import { useRouter, useSearch } from '@tanstack/react-router'

import SignInForm from '../components/forms/sign-in-form'
import AuthPageWrapper from '../components/auth-page-wrapper'
import AccountCancelled from '../components/account-cancelled'

import { UserData } from '@/horizon-corp/types'
import { emailSchema, userAccountTypeSchema } from '../validations'

export const SignInPageSearchSchema = z.object({
    email: z.string().optional().default('').or(emailSchema),
    accountType: z
        .string()
        .optional()
        .default('Member')
        .pipe(userAccountTypeSchema)
        .catch('Member'),
})

const SignInPage = () => {
    const router = useRouter()
    const prefilledValues = useSearch({ from: '/auth/sign-in' })
    const [userData, setUserData] = useState<null | UserData>(null)

    const onSignInSuccess = (userData: UserData) => {
        const { isContactVerified, isEmailVerified, status } = userData
        if (status === 'Not Allowed') return setUserData(userData)

        if (status === 'Pending' && (!isContactVerified || !isEmailVerified)) {
            router.navigate({ to: '/auth/verify' })
        }

        if (status === 'Verified') {
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
        <div className="flex min-h-full flex-col items-center justify-center">
            <AuthPageWrapper>
                {!userData && (
                    <SignInForm
                        defaultValues={prefilledValues}
                        onSuccess={onSignInSuccess}
                    />
                )}
                {userData && (
                    <AccountCancelled
                        userData={userData}
                        onBack={handleOnCancelBack}
                    />
                )}
            </AuthPageWrapper>
        </div>
    )
}

export default SignInPage
