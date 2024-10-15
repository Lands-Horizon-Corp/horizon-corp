import { useRouter } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import useCurrentUser from '@/hooks/use-current-user'
import { IBaseComp, TAccountType, TPageType } from '@/types'

interface Props extends IBaseComp {
    pageType?: TPageType
    allowedAccountTypes: TAccountType[]
}

const ProtectedRouteWrapper = ({
    children,
    pageType = 'AUTHENTICATED',
    allowedAccountTypes,
}: Props) => {
    const router = useRouter()
    const { currentUser, loadingUser } = useCurrentUser({ loadOnMount: true })

    if (pageType === 'AUTHENTICATED') {
        if (loadingUser)
            return (
                <div className="relative flex h-screen w-full items-center justify-center">
                    <LoadingSpinner />
                </div>
            )

        if (
            currentUser &&
            !allowedAccountTypes.includes(currentUser?.accountType)
        )
            return (
                <div className="h-screen w-full flex flex-col items-center justify-center">
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

export default ProtectedRouteWrapper
