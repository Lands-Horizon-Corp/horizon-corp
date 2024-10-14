import z from 'zod'
import { useState } from 'react'
import { useRouter, useSearch } from '@tanstack/react-router'

import SignInForm from '../components/forms/sign-in-form'
import AuthPageWrapper from '../components/auth-page-wrapper'
import AccountCancelled from '../components/account-cancelled'

import { UserData } from '@/horizon-corp/types'
import { emailSchema, userAccountTypeSchema } from '../validations'
import useCurrentUser from '@/hooks/use-current-user'
import { getUsersAccountTypeRedirectPage } from './helpers'
import { toast } from 'sonner'

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
    const { setCurrentUser } = useCurrentUser()
    const prefilledValues = useSearch({ from: '/auth/sign-in' })
    const [userData, setUserData] = useState<null | UserData>(null)

    const onSignInSuccess = (userData: UserData) => {
        const { isContactVerified, isEmailVerified, status } = userData
        setCurrentUser(userData)

        if (status === 'Not Allowed') {
            toast.error(
                'Your account has been canceled, and you can no longer log in.'
            )
            router.navigate({ to: '/auth/verify' })
        }

        if (status === 'Pending' || !isContactVerified || !isEmailVerified) {
            toast.warning('Your account is pending approval')
            router.navigate({ to: '/auth/verify' })
        }

        if (status === 'Verified') {
            const redirectUrl = getUsersAccountTypeRedirectPage(userData)
            toast.success("You're logged in!")
            router.navigate({ to: redirectUrl })
        }
    }

    const handleOnCancelBack = () => {
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
