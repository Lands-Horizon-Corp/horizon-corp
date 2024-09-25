import { useEffect, useState } from 'react'
import { useRouter } from '@tanstack/react-router'

import LoadingCircle from '@/components/loader/loading-circle'

import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import VerifyRoot from '../components/verify-root'
import { Separator } from '@/components/ui/separator'
import AccountCancelled from '../components/account-cancelled'

import { UserBase, UserStatus } from '@/types'

interface Props {}

const Verify = ({}: Props) => {
    const [loading, setLoading] = useState(true)
    const [display, setDisplay] = useState<
        null | 'verify' | 'verify-complete' | 'account-cancelled'
    >()
    const [userData, setUserData] = useState<UserBase | null>(null)

    const router = useRouter()

    useEffect(() => {
        if (!userData) return

        if (userData.status === UserStatus['Not allowed'])
            return setDisplay('account-cancelled')

        if (userData.status === UserStatus.Verified)
            return setDisplay('verify-complete')

        if (!userData.validContactNumber || !userData.validEmail)
            setDisplay('verify')
    }, [userData])

    useEffect(() => {
        // TODO: Fetch User Data
        // Remove code below it just for simulating user data fetching
        // and use our Horizon service.auth-service.ts
        setLoading(true)
        setTimeout(() => {
            setUserData({
                id: '215',
                username: 'Jervx',
                validEmail: false,
                validContactNumber: false,
                status: UserStatus['Pending'],
                profilePicture: {
                    url: 'https://mrwallpaper.com/images/hd/suit-rick-and-morty-phone-5divv4gzo6gowk46.jpg',
                },
            } as any as UserBase)
            setLoading(false)
        }, 1000)
    }, [])

    const autoRedirectAccount = (userData: UserBase) => {
        // TODO Redirect once verified
        // if (userData.status === UserStatus.Verified) {
        //     // TODO: Auto redirect to page the account belongs
        //     // admin/
        //     // owner/
        //     // member/
        //     // employee
        // }
        // router.navigate({ to : "..."})
    }

    return (
        <div className="flex flex-1 flex-col items-center px-4 py-4">
            <div className="min-h-[50vh] max-w-lg py-4">
                {loading && (
                    <div className="flex flex-col items-center gap-y-2">
                        <LoadingCircle />
                        <p className="text-center text-sm text-foreground/50">
                            please wait.. loading your info
                        </p>
                    </div>
                )}
                {userData && !loading && (
                    <>
                        {display === 'verify' && (
                            <VerifyRoot
                                userData={userData}
                                onVerifyChange={(newUserData) =>
                                    setUserData(newUserData)
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
                                <p className="text-center text-foreground/60">
                                    Your account email & phone number is
                                    verified. Thank you for joining with us!
                                </p>
                                <UserAvatar
                                    className="size-28"
                                    src={userData?.profilePicture?.url ?? ''}
                                    fallback={
                                        userData?.username.charAt(0) ?? '-'
                                    }
                                />
                                <Separator className="w-full" />
                                <p className="text-center">
                                    {userData.status === UserStatus.Pending && (
                                        <span>
                                            Please wait for 7 working days for
                                            validation. Your ID is{' '}
                                            <span className="font-medium text-green-500">
                                                {userData.id}
                                            </span>
                                        </span>
                                    )}
                                </p>
                                <Button
                                    onClick={() =>
                                        autoRedirectAccount(userData)
                                    }
                                    disabled={loading}
                                    className="mt-6 w-full bg-[#34C759] hover:bg-[#38b558]"
                                >
                                    Proceed
                                </Button>
                            </div>
                        )}
                        {display === 'account-cancelled' && (
                            <AccountCancelled userData={userData} />
                        )}
                    </>
                )}
            </div>
        </div>
    )
}

export default Verify
