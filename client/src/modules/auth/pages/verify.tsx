import { toast } from 'sonner'
import { useEffect, useState } from 'react'
import { useRouter } from '@tanstack/react-router'

import VerifyRoot from '@/modules/auth/components/verify-root'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import AuthPageWrapper from '@/modules/auth/components/auth-page-wrapper'
import ShowAccountStatus from '../components/verify-root/show-account-status'

import { withCatchAsync } from '@/lib'
import useCurrentUser from '@/hooks/use-current-user'
import { serverRequestErrExtractor } from '@/helpers'
import { getUsersAccountTypeRedirectPage } from './helpers'
import UserService from '@/horizon-corp/server/auth/UserService'
import useLoadingErrorState from '@/hooks/use-loading-error-state'

const Verify = () => {
    const router = useRouter()
    const { loading, setLoading } = useLoadingErrorState()
    const { currentUser, setCurrentUser, authState } = useCurrentUser({
        loadOnMount: true,
    })

    const [display, setDisplay] = useState<null | 'verify' | 'account-status'>(
        null
    )

    const handleBackSignOut = async () => {
        if (loading || authState === 'Loading') return

        setLoading(true)

        const [error] = await withCatchAsync(UserService.SignOut())

        setLoading(false)

        if (error) {
            const errorMessage = serverRequestErrExtractor({ error })
            toast.error(errorMessage)
        }

        router.navigate({ to: '/auth' })
        setCurrentUser(null)
    }

    useEffect(() => {
        if (!currentUser) return
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
    }, [currentUser])

    return (
        <div className="flex min-h-full flex-col items-center justify-center">
            <AuthPageWrapper>
                {authState === 'Loading' && (
                    <div className="flex flex-col items-center gap-y-2">
                        <LoadingSpinner />
                        <p className="text-center text-sm text-foreground/50">
                            loading user info
                        </p>
                    </div>
                )}
                {currentUser && authState === 'Authenticated' && (
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
                                loading={loading}
                                onBackSignOut={handleBackSignOut}
                                userData={currentUser}
                            />
                        )}
                    </>
                )}
                {!currentUser && authState === 'UnAuthenticated' && (
                    <p>Couldn&apos;t load your info, please please try again</p>
                )}
            </AuthPageWrapper>
        </div>
    )
}

export default Verify
