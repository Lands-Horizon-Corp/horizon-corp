import { useRouter } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import { WarningCircleIcon } from '@/components/icons'

import { HELP_CONTACT } from '../constants'
import { UserData } from '@/horizon-corp/types'
import { IBaseComp } from '@/types/component/base'

interface Props extends IBaseComp {
    userData: UserData // TODO: Change based on auth response resource fromn @/horizon-corp/types/auth/...
    onBack?: () => void
}

const AccountCancelled = ({ userData, onBack }: Props) => {
    const router = useRouter()

    return (
        <div className="flex max-w-sm flex-col items-center gap-y-4">
            <p className="text-xl font-medium text-[#ED6E6E]">
                Account Canceled
            </p>
            <div className="relative my-8 rounded-full border-4 border-[#ED6E6E]">
                <UserAvatar
                    className="size-28"
                    src={userData?.media?.downloadURL ?? ''}
                    fallback={userData?.username.charAt(0) ?? '-'}
                />
                <WarningCircleIcon className="absolute -bottom-1 -right-1 size-6 text-[#ED6E6E]" />
            </div>
            <p className="text-center text-foreground/80">
                Your account has been canceled. Please contact the number below
                for assistance.
            </p>

            <p className="text-center font-medium">{HELP_CONTACT}</p>
            <Button
                onClick={() => {
                    onBack ? onBack() : router.history.back()
                }}
                className="w-full bg-[#34C759] hover:bg-[#38b558]"
            >
                Back to Sign In
            </Button>
        </div>
    )
}

export default AccountCancelled
