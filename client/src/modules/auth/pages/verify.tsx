import { toast } from 'sonner'
import { useEffect, useState } from 'react'
import { useRouter } from '@tanstack/react-router'
import { useMutation } from '@tanstack/react-query'

import VerifyRoot from '@/modules/auth/components/verify-root'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import AuthPageWrapper from '@/modules/auth/components/auth-page-wrapper'
import ShowAccountStatus from '../components/verify-root/show-account-status'

import { withCatchAsync } from '@/lib'
import useCurrentUser from '@/hooks/use-current-user'
import { serverRequestErrExtractor } from '@/helpers'
import { getUsersAccountTypeRedirectPage } from '@/helpers'
import UserService from '@/horizon-corp/server/auth/UserService'

const Verify = () => {
    const router = useRouter()
    const { data: currentUser, isFetching, setCurrentUser } = useCurrentUser({})

    const { mutate: onBackSignOut, isPending } = useMutation<void, string>({
        mutationKey: ['sign-out'],
        mutationFn: async () => {
            const [error] = await withCatchAsync(UserService.SignOut())

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            router.navigate({ to: '/auth' })

            setCurrentUser(null)
            return
        },
    })

    const [display, setDisplay] = useState<null | 'verify' | 'account-status'>(
        null
    )

    useEffect(() => {
        if (!currentUser || isFetching) return
        else if (currentUser.status === 'Verified') {
            const redirectUrl = getUsersAccountTypeRedirectPage(currentUser)
            router.navigate({ to: redirectUrl })
        } else if (
            !currentUser.isContactVerified ||
            !currentUser.isEmailVerified
        ) {
            setDisplay('verify')
        } else {
            setDisplay('account-status')
        }
    }, [currentUser, isFetching, router])

    return (
        <div className="flex min-h-full flex-col items-center justify-center">
            <AuthPageWrapper>
                {isFetching && (
                    <div className="flex flex-col items-center gap-y-2">
                        <LoadingSpinner className="block" />
                        <p className="text-center text-sm text-foreground/50">
                            loading user info
                        </p>
                    </div>
                )}
                {currentUser && !isFetching && (
                    <>
                        {display === 'verify' && (
                            <VerifyRoot
                                userData={currentUser}
                                onVerifyChange={(newUserData) =>
                                    setCurrentUser(newUserData)
                                }
                                onVerifyComplete={(newUserData) => {
                                    setDisplay('account-status')
                                    setCurrentUser(newUserData)
                                }}
                            />
                        )}
                        {display === 'account-status' && (
                            <ShowAccountStatus
                                loading={isFetching || isPending}
                                onBackSignOut={onBackSignOut}
                                userData={currentUser}
                            />
                        )}
                    </>
                )}
                {!currentUser && !isFetching && (
                    <p>Couldn&apos;t load your info, please please try again</p>
                )}
            </AuthPageWrapper>
        </div>
    )
}

export default Verify
