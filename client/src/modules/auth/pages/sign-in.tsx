import z from 'zod'
import { toast } from 'sonner'
import { useCallback, useEffect } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import { useRouter, useSearch } from '@tanstack/react-router'

import SignInForm from '../components/forms/sign-in-form'
import AuthPageWrapper from '../components/auth-page-wrapper'

import { UserData } from '@/horizon-corp/types'
import useCurrentUser from '@/hooks/use-current-user'
import { getUsersAccountTypeRedirectPage } from './helpers'
import { userAccountTypeSchema } from '../validations/common'
import LoadingSpinner from '@/components/spinners/loading-spinner'

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
    const { data: currentUser, isFetching } = useCurrentUser({ loadOnMount : true })
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

            if (status === 'Verified') {
                const redirectUrl = getUsersAccountTypeRedirectPage(userData)
                toast.success("You're logged in!")
                router.navigate({ to: redirectUrl })
            }
        },
        [queryClient, router]
    )

    useEffect(() => {
        if (!currentUser) return
        onSignInSuccess(currentUser)
    }, [currentUser, onSignInSuccess])

    return (
        <div className="flex min-h-full flex-col items-center justify-center">
            <AuthPageWrapper>
                {currentUser === null && !isFetching && (
                    <SignInForm
                        defaultValues={prefilledValues}
                        onSuccess={onSignInSuccess}
                    />
                )}
                {
                    isFetching && <LoadingSpinner />
                }
            </AuthPageWrapper>
        </div>
    )
}

export default SignInPage
