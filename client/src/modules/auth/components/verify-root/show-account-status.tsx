import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import { WarningCircleIcon } from '@/components/icons'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { HELP_CONTACT } from '../../constants'
import { UserData } from '@/horizon-corp/types'

interface Props {
    loading : boolean,
    userData: UserData,
    onBackSignOut : () => void
}

const ShowAccountStatus = ({ loading, userData, onBackSignOut }: Props) => {

    if (userData.status === 'Pending')
        return (
            <div className="flex max-w-sm flex-col items-center gap-y-4">
                <p className="text-xl font-medium text-green-500">
                    Account Pending
                </p>
                <p className="text-center text-foreground/70">
                    <span className="font-medium text-foreground/90">
                        Thank you for joining with us!
                    </span>
                </p>
                <UserAvatar
                    className="my-8 size-28"
                    src={userData?.media?.downloadURL ?? ''}
                    fallback={userData?.username?.charAt(0) ?? '-'}
                />
                <p className="text-center">
                    <span>
                        Please wait for 7 working days for validation before you
                        can use your account, we will send an email once your
                        account is activated.
                        <br />
                    </span>
                </p>
                <p className="px-4 text-center">
                    Your ID is{' '}
                    <span className="font-medium text-green-500">
                        {userData.id}
                    </span>
                </p>
                <Button
                    disabled={loading}
                    variant="secondary"
                    onClick={onBackSignOut}
                    className="mt-6 w-full"
                >
                    {loading ? <LoadingSpinner /> : 'Back to Sign In'}
                </Button>
            </div>
        )

    if (userData.status === 'Not Allowed')
        return (
            <div className="flex max-w-sm flex-col items-center gap-y-4">
                <p className="text-xl font-medium text-[#ED6E6E]">
                    Account Canceled
                </p>
                <div className="relative my-8 rounded-full border-4 border-[#ED6E6E]">
                    <UserAvatar
                        className="size-28"
                        src={userData?.media?.downloadURL ?? ''}
                        fallback={userData?.username?.charAt(0) ?? '-'}
                    />
                    <WarningCircleIcon className="absolute -bottom-1 -right-1 size-6 text-[#ED6E6E]" />
                </div>
                <p className="text-center text-foreground/80">
                    Your account has been canceled. Please contact the number
                    below for assistance.
                </p>

                <p className="text-center font-medium">{HELP_CONTACT}</p>
                <Button
                    disabled={loading}
                    onClick={onBackSignOut}
                    className="w-full"
                >
                    {loading ? <LoadingSpinner /> : 'Back to Sign In'}
                </Button>
            </div>
        )

    return (
        <p className="text-center">
            Sorry, couldn&apos;t identify you account status
        </p>
    )
}

export default ShowAccountStatus
