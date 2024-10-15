import { toast } from 'sonner'
import { useEffect, useState } from 'react'
import { useRouter } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import VerifyRoot from '@/modules/auth/components/verify-root'
import AuthPageWrapper from '@/modules/auth/components/auth-page-wrapper'
import AccountCancelled from '@/modules/auth/components/account-cancelled'

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

    const [display, setDisplay] = useState<
        null | 'verify' | 'verify-complete' | 'account-cancelled'
    >()

    const handleBackSignOut = async () => {
        if (loading || loadingUser) return

        setLoading(true)
        try {
            await UserService.SignOut()
            setCurrentUser(null)
            router.navigate({ to: '/auth' })
        } catch (e) {
            const errorMessage = handleAxiosError(e)
            toast.error(errorMessage)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        if (!currentUser) return

        if (currentUser.status === 'Verified') {
            const redirectPageUrl = getUsersAccountTypeRedirectPage(currentUser)
            router.navigate({ to: redirectPageUrl })
        }

        if (currentUser.status === 'Not Allowed')
            return setDisplay('account-cancelled')

        if (currentUser.isContactVerified && currentUser.isEmailVerified)
            setDisplay('verify-complete')

        if (!currentUser.isContactVerified || !currentUser.isEmailVerified)
            setDisplay('verify')
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
                                    setDisplay('verify-complete')
                                    setCurrentUser(newUserData)
                                }}
                            />
                        )}
                        {display === 'verify-complete' && (
                            <div className="flex max-w-sm flex-col items-center gap-y-4">
                                <p className="text-xl font-medium text-green-500">
                                    Account Verify Complete
                                </p>
                                <p className="text-center text-foreground/70">
                                    Your account email & phone number is
                                    verified.{' '}
                                    <span className="font-medium text-foreground/90">
                                        Thank you for joining with us!
                                    </span>
                                </p>
                                <UserAvatar
                                    className="my-8 size-28"
                                    src={currentUser?.media?.downloadURL ?? ''}
                                    fallback={
                                        currentUser?.username.charAt(0) ?? '-'
                                    }
                                />
                                <p className="text-center">
                                    {currentUser.status === 'Pending' && (
                                        <span>
                                            Please wait for 7 working days for
                                            validation before you can use your
                                            account, we will send an email once
                                            your account is activated.
                                            <br />
                                        </span>
                                    )}
                                </p>
                                <p className="px-4 text-center">
                                    Your ID is{' '}
                                    <span className="font-medium text-green-500">
                                        {currentUser.id}
                                    </span>
                                </p>
                                <Button
                                    disabled={loadingUser || loading}
                                    onClick={handleBackSignOut}
                                    className="mt-6 w-full bg-[#34C759] hover:bg-[#38b558]"
                                >
                                    Go Back to signin
                                </Button>
                            </div>
                        )}
                        {display === 'account-cancelled' && (
                            <AccountCancelled
                                loading={loading}
                                userData={currentUser}
                                onBack={handleBackSignOut}
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
