import { toast } from 'sonner'
import { Navigate, useRouter } from '@tanstack/react-router'

import UserAvatar from '../user-avatar'
import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { isUserHasUnverified, isUserUnverified } from '@/helpers'
import useCurrentUser from '@/hooks/use-current-user'
import { IBaseComp, TAccountType, TPageType } from '@/types'
import { cn } from '@/lib'
import { PatchCheckIcon, PatchExclamationIcon, PatchMinusIcon } from '../icons'
import { ReactNode } from 'react'
import { UserData } from '@/horizon-corp/types'

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

        if (
            !currentUser.isSkipVerification &&
            currentUser.status !== 'Not Allowed' &&
            isUserHasUnverified(currentUser)
        ) {
            return (
                <BannerContainer>
                    <AccountInfoContent
                        currentUser={currentUser}
                        infoTitle="Verification Required"
                        infoDescription="It looks like you have unverified contacts. For security reasons, please verify your contact information to access all features."
                    />
                    <Button
                        className="rounded-full"
                        onClick={() => router.navigate({ to: '/auth/verify' })}
                    >
                        Verify Now
                    </Button>
                </BannerContainer>
            )
        }

        if (
            (currentUser &&
                !allowedAccountTypes.includes(currentUser?.accountType)) ||
            isUserUnverified(currentUser)
        )
            return (
                <BannerContainer>
                    <AccountInfoContent
                        infoTitle="Restricted"
                        currentUser={currentUser}
                        infoDescription="Sorry but your account type is restricted in accessing
                        this page"
                    />
                    <Button
                        variant="outline"
                        className="rounded-full"
                        onClick={() => router.history.back()}
                    >
                        Go Back
                    </Button>
                </BannerContainer>
            )
    }

    return <>{children}</>
}

const BannerContainer = ({ children }: { children?: ReactNode }) => {
    return (
        <div className="flex h-screen w-full flex-col items-center justify-center gap-y-4">
            {children}
        </div>
    )
}

const AccountInfoContent = ({
    infoTitle,
    infoDescription,
    currentUser,
}: {
    infoTitle: string
    infoDescription: string
    currentUser: UserData
}) => {
    return (
        <>
            <UserAvatar
                src={currentUser.media?.downloadURL ?? ''}
                fallback={currentUser.username.charAt(0) ?? '-'}
                className="size-36 border-4 text-2xl font-medium"
            />
            {currentUser.status === 'Pending' && (
                <PatchMinusIcon className="size-8 text-amber-500" />
            )}
            {currentUser.status === 'Verified' && (
                <PatchCheckIcon
                    className={cn(
                        'size-8 text-primary',
                        isUserHasUnverified(currentUser) && 'text-amber-500'
                    )}
                />
            )}
            {currentUser.status === 'Not Allowed' && (
                <PatchExclamationIcon className="size-8 text-rose-400" />
            )}
            <p className="text-xl font-medium">{infoTitle}</p>
            <p className="max-w-xl text-center text-foreground/80">
                {infoDescription}
            </p>
        </>
    )
}

export default AuthGuard
