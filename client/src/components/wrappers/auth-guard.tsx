import { toast } from 'sonner'
import { Navigate, useRouter } from '@tanstack/react-router'

import UserAvatar from '../user-avatar'
import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { isUserUnverified } from '@/helpers'
import useCurrentUser from '@/hooks/use-current-user'
import { IBaseComp, TAccountType, TPageType } from '@/types'

interface Props extends IBaseComp {
    pageType?: TPageType
    allowedAccountTypes?: TAccountType[]
}

const AuthGuard = ({
    children,
    allowedAccountTypes = [],
    pageType = 'AUTHENTICATED',
}: Props) => {
    const router = useRouter()
    const { data: currentUser, status } = useCurrentUser({
        onUnauthorized: () => toast.error('You are not logged in'),
    })

    if (pageType === 'AUTHENTICATED') {
        if (status === 'pending')
            return (
                <div className="relative flex h-screen w-full items-center justify-center">
                    <LoadingSpinner />
                </div>
            )

        if (!currentUser) return <Navigate to="/auth/sign-in" />

        if (currentUser.isSkipVerification && isUserUnverified(currentUser)) {
            return (
                <div className="flex h-screen w-full flex-col items-center justify-center gap-y-4">
                    <UserAvatar
                        src={currentUser.media?.downloadURL ?? ''}
                        fallback={currentUser.username.charAt(0) ?? '-'}
                        className="size-24"
                    />
                    <p className="text-xl font-medium">Verify Account</p>
                    <p className="text-foreground/80">
                        Sorry but you&apos;r account status is pending or you
                        have unverified contacts. You can skip this on verify
                        page.
                    </p>
                    <Button
                        variant="outline"
                        onClick={() => router.navigate({ to: '/auth/verify' })}
                    >
                        Go & Verify
                    </Button>
                </div>
            )
        }

        if (
            (currentUser &&
                !allowedAccountTypes.includes(currentUser?.accountType)) ||
            isUserUnverified(currentUser)
        )
            return (
                <div className="flex h-screen w-full flex-col items-center justify-center gap-y-4">
                    <p className="text-xl font-medium">Restricted</p>
                    <p className="text-foreground/80">
                        Sorry but you&apos;r account type is restricted in
                        accessing this page
                    </p>
                    <Button
                        variant="outline"
                        onClick={() => router.history.back()}
                    >
                        Go Back
                    </Button>
                </div>
            )
    }

    return <>{children}</>
}

export default AuthGuard
