import { useRouter } from '@tanstack/react-router'

import { IBaseComp } from '@/types'
import useCurrentUser from '@/hooks/use-current-user'
import { getUsersAccountTypeRedirectPage } from '@/helpers'

interface IGuestGuardProps extends Omit<IBaseComp, 'className'> {
    allowAuthenticatedUser?: false
}

const GuestGuard = ({
    allowAuthenticatedUser = false,
    children,
}: IGuestGuardProps) => {
    const router = useRouter()
    const { data: currentUser } = useCurrentUser()

    if (!allowAuthenticatedUser && currentUser) {
        if (currentUser.status !== 'Verified') {
            router.navigate({ to: '/auth/verify' })
            return <></>
        }

        const redirectUrl = getUsersAccountTypeRedirectPage(currentUser)
        router.navigate({ to: redirectUrl })
        return (
            <div className="flex h-full flex-col items-center justify-center text-center">
                <div className="flex items-center gap-x-4 rounded-xl bg-popover p-4">
                    <p className="">Redirecting...</p>
                </div>
            </div>
        )
    }

    return <>{children}</>
}

export default GuestGuard
