import { Navigate, useRouter } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import useCurrentUser from '@/hooks/use-current-user'
import { IBaseComp, TAccountType, TPageType } from '@/types'

interface Props extends IBaseComp {
    pageType?: TPageType
    allowedAccountTypes?: TAccountType[]
}

const ProtectedRouteWrapper = ({
    children,
    allowedAccountTypes = [],
    pageType = 'AUTHENTICATED',
}: Props) => {
    const router = useRouter()

    const { data: currentUser, status } = useCurrentUser({})

    if (pageType === 'AUTHENTICATED') {
        if (status === 'pending')
            return (
                <div className="relative flex h-screen w-full items-center justify-center">
                    <LoadingSpinner />
                </div>
            )

        if (!currentUser) return <Navigate to="/auth/sign-in" />


        if (
            currentUser &&
            !allowedAccountTypes.includes(currentUser?.accountType) ||
            currentUser.status !== 'Verified'
        )
            return (
                <div className="flex h-screen w-full flex-col items-center justify-center gap-y-4">
                    <p className="text-xl font-medium">Restricted</p>
                    <p className="text-foreground/80">
                        Sorry but you&apos;r account type or account status is restricted in
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

export default ProtectedRouteWrapper
