import { toast } from 'sonner'
import { useEffect, useState } from 'react'
import { useRouter } from '@tanstack/react-router'

import VerifyRoot from '@/modules/auth/components/verify-root'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import AuthPageWrapper from '@/modules/auth/components/auth-page-wrapper'
import ShowAccountStatus from '../components/verify-root/show-account-status'

import useCurrentUser from '@/hooks/use-current-user'
import { handleAxiosError } from '@/horizon-corp/helpers'
import { getUsersAccountTypeRedirectPage } from './helpers'
import UserService from '@/horizon-corp/server/auth/UserService'
import useLoadingErrorState from '@/hooks/use-loading-error-state'

interface Props {}

const Verify = ({}: Props) => {
    const router = useRouter()
    const { loading, setLoading } = useLoadingErrorState()
    const { currentUser, setCurrentUser, loadingUser } = useCurrentUser({
        loadOnMount: true,
    })

    const [display, setDisplay] = useState<null | 'verify' | 'account-status'>(
        null
    )

    const handleBackSignOut = async () => {
        if (loading || loadingUser) return

        setLoading(true)
        try {
            await UserService.SignOut()
            router.navigate({ to: '/auth' })
            setCurrentUser(null)
        } catch (e) {
            const errorMessage = handleAxiosError(e)
            toast.error(errorMessage)
        } finally {
            setLoading(false)
        }
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
                {loadingUser && (
                    <div className="flex flex-col items-center gap-y-2">
                        <LoadingSpinner />
                        <p className="text-center text-sm text-foreground/50">
                            loading user info
                        </p>
                    </div>
                )}
                {currentUser && !loadingUser && (
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
                {!currentUser && !loadingUser && (
                    <p>Couldn&apos;t load your info, please please try again</p>
                )}
            </AuthPageWrapper>
        </div>
    )
}

export default Verify
