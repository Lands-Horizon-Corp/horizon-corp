import { useEffect, useState } from 'react'

import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import LoadingCircle from '@/components/loader/loading-circle'
import VerifyRoot from '@/modules/auth/components/verify-root'
import AuthPageWrapper from '@/modules/auth/components/auth-page-wrapper'
import AccountCancelled from '@/modules/auth/components/account-cancelled'

import useCurrentUser from '@/hooks/use-current-user'
import { UserData } from '@/horizon-corp/types'

interface Props {}

const Verify = ({}: Props) => {
    const { currentUser, setCurrentUser, loadingUser } = useCurrentUser({
        loadOnMount: true,
    })

    const [display, setDisplay] = useState<
        null | 'verify' | 'verify-complete' | 'account-cancelled'
    >()

    // const router = useRouter()

    useEffect(() => {
        if (!currentUser) return

        if (currentUser.status === 'Not Allowed')
            return setDisplay('account-cancelled')

        if (currentUser.status === 'Verified')
            return setDisplay('verify-complete')

        if (!currentUser.isContactVerified || !currentUser.isEmailVerified)
            setDisplay('verify')
    }, [currentUser])

    const autoRedirectAccount = (_currentUser: UserData) => {
        // TODO Redirect once verified
        // if (currentUser.status === UserStatus.Verified) {
        //     // TODO: Auto redirect to page the account belongs
        //     // admin/
        //     // owner/
        //     // member/
        //     // employee
        // }
        // router.navigate({ to : "..."})
    }

    return (
        <div className="flex min-h-full flex-col items-center justify-center">
            <AuthPageWrapper>
                {loadingUser && (
                    <div className="flex flex-col items-center gap-y-2">
                        <LoadingCircle />
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
                                onVerifyComplete={() => {
                                    setDisplay('verify-complete')
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
                                    {currentUser.status === "Pending" && (
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
                                    onClick={() =>
                                        autoRedirectAccount(currentUser)
                                    }
                                    disabled={loadingUser}
                                    className="mt-6 w-full bg-[#34C759] hover:bg-[#38b558]"
                                >
                                    Proceed
                                </Button>
                            </div>
                        )}
                        {display === 'account-cancelled' && (
                            <AccountCancelled userData={currentUser} />
                        )}
                    </>
                )}
                {!currentUser && !loadingUser && (
                    <p>Couldn&apos;t load info, please reload</p>
                )}
            </AuthPageWrapper>
        </div>
    )
}

export default Verify
