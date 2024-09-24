import { useEffect, useState } from 'react'

import VerifyForm from '@/modules/auth/components/forms/verify-form'

import { UserBase, UserStatus } from '@/types'
import UserAvatar from '@/components/user-avatar'
import { useRouter } from '@tanstack/react-router'
import { HELP_CONTACT } from '../constants'
import { Separator } from '@/components/ui/separator'
import LoadingCircle from '@/components/loader/loading-circle'
import { Button } from '@/components/ui/button'

interface Props {}

type TSteps = 1 | 2 | 3 | 4

const Verify = ({}: Props) => {
    const [loading, setLoading] = useState(true)
    const [step, setStep] = useState<0 | TSteps>(0)
    const [userData, setUserData] = useState<UserBase | null>(null)

    const router = useRouter()

    useEffect(() => {
        if (!userData) return
        //  TODO: Check what user lacks in verification
        //  then set appropriate step
        //  1 - Verify Contact Number
        if (!userData.validContactNumber) return setStep(1)

        //  2 - Verify Email
        if (!userData.validEmail) return setStep(2)

        if (userData.status === UserStatus.Verified) {
            // TODO: Redirect depending on the account type
            // member - /member
            // owner - /company/$companyId
            // admin - /admin
        }

        //  3 - Indicates he/she is a member, show a join company form input

        setStep(4)
        //  4 - Show users account status
    }, [userData])

    useEffect(() => {
        // TODO: Fetch User Data
        // Remove code below
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

    const handleOnSuccess = (userData: UserBase, nextStep: TSteps) => {
        setUserData(userData)
        setStep(nextStep)
    }

    const handleComplete = () => {}

    return (
        <div className="flex flex-1 justify-center px-4 py-4">
            <div className="py-14">
                {loading && <LoadingCircle />}
                {userData && !loading && (
                    <>
                        {step === 1 && (
                            <VerifyForm
                                key="1"
                                id={userData.id}
                                verifyMode="mobile"
                                onSuccess={(data) => handleOnSuccess(data, 2)}
                            />
                        )}
                        {step === 2 && (
                            <VerifyForm
                                key="2"
                                id={userData.id}
                                verifyMode="email"
                                onSuccess={(data) => {
                                    const nextStep = 3
                                    // TODO : if user is member, show step 3
                                    // step 3 allows user to select or join a company
                                    handleOnSuccess(data, nextStep)
                                }}
                            />
                        )}
                        {
                            // TODO: Step 3
                        }
                        {step === 4 && (
                            <div className="flex max-w-sm flex-col items-center gap-y-4">
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
                                        router.navigate({ to: '/auth/sign-in' })
                                    }
                                    disabled={loading}
                                    className="mt-6 bg-[#34C759] w-full hover:bg-[#38b558]"
                                >
                                    Go to Login
                                </Button>
                            </div>
                        )}
                    </>
                )}
            </div>
        </div>
    )
}

export default Verify
