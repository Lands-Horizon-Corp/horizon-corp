import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import UserAvatar from '@/components/user-avatar'
import { UserBase } from '@/types'
import { IBaseComp } from '@/types/component/base'
import { useRouter } from '@tanstack/react-router'
import { HELP_CONTACT } from '../constants'

interface Props extends IBaseComp {
    userData: UserBase
    onBack?: () => void
}

const AccountCancelled = ({ userData, onBack }: Props) => {
    const router = useRouter()

    return (
        <div className="flex max-w-sm flex-col items-center gap-y-4">
            <UserAvatar
                className="size-28"
                src={userData?.profilePicture?.url ?? ''}
                fallback={userData?.username.charAt(0) ?? '-'}
            />
            <Separator className="w-full" />
            <p className="text-center">
                <span>
                    Your account has been cancelled. Please contact the number
                    <span className="font-medium text-green-500">
                        {HELP_CONTACT}
                    </span>{' '}
                    to verify/activate your account
                </span>
            </p>
            <Button
                onClick={() => {
                    onBack ? onBack() : router.history.back()
                }}
                className="mt-6 w-full bg-[#34C759] hover:bg-[#38b558]"
            >
                Back to Login
            </Button>
        </div>
    )
}

export default AccountCancelled
