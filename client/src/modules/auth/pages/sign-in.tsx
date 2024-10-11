import z from 'zod'
import { useState } from 'react'
import { useRouter, useSearch } from '@tanstack/react-router'

import SignInForm from '../components/forms/sign-in-form'
import AccountCancelled from '../components/account-cancelled'

import { UserBase, UserStatus } from '@/types'
import { emailSchema, userAccountTypeSchema } from '../validations'
import AuthPageWrapper from '../components/auth-page-wrapper'

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
