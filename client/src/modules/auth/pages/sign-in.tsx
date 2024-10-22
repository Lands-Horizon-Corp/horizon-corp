import z from 'zod'
import { toast } from 'sonner'
import { useCallback, useEffect } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import { useRouter, useSearch } from '@tanstack/react-router'

import SignInForm from '../components/forms/sign-in-form'
import GuestGuard from '@/components/wrappers/guest-guard'
import AuthPageWrapper from '../components/auth-page-wrapper'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { UserData } from '@/horizon-corp/types'
import useCurrentUser from '@/hooks/use-current-user'
import { userAccountTypeSchema } from '../validations/common'

export const SignInPageSearchSchema = z.object({
    key: z.string().optional(),
    accountType: z
        .string()
        .optional()
        .default('Member')
        .pipe(userAccountTypeSchema)
        .catch('Member'),
})

const SignInPage = () => {
    const router = useRouter()
    const queryClient = useQueryClient()
    const { data: currentUser, isFetched, isLoading } = useCurrentUser()
    const prefilledValues = useSearch({ from: '/auth/sign-in' })

    const onSignInSuccess = useCallback(
        (userData: UserData) => {
            const { isContactVerified, isEmailVerified, status } = userData

            queryClient.setQueryData(['current-user'], userData)

            if (status === 'Not Allowed') {
                toast.error(
                    'Your account has been canceled, and you can no longer log in.'
                )
                router.navigate({ to: '/auth/verify' })
            }

            if (
                status === 'Pending' ||
                !isContactVerified ||
                !isEmailVerified
            ) {
                toast.warning('Your account is pending approval')
                router.navigate({ to: '/auth/verify' })
            }
        },
        [queryClient, router]
    )

    useEffect(() => {
        if (!currentUser) return
        onSignInSuccess(currentUser)
    }, [currentUser, onSignInSuccess])

    return (
        <GuestGuard>
            <div className="flex min-h-full flex-col items-center justify-center">
                <AuthPageWrapper>
                    {!currentUser && isFetched && (
                        <SignInForm
                            defaultValues={prefilledValues}
                            onSuccess={onSignInSuccess}
                        />
                    )}
                    {currentUser && <LoadingSpinner />}
                    {isLoading && !isFetched && <LoadingSpinner />}
                </AuthPageWrapper>
            </div>
        </GuestGuard>
    )
}

export default SignInPage
