import { Navigate } from '@tanstack/react-router'

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
    const { data: currentUser } = useCurrentUser()

    if (!allowAuthenticatedUser && currentUser) {
        if (currentUser.status !== 'Verified') {
            return <Navigate to="/auth/verify" />
        }

        const redirectUrl = getUsersAccountTypeRedirectPage(currentUser)
        return (
            <div className="flex h-full flex-col items-center justify-center text-center">
                <div className="flex items-center gap-x-4 rounded-xl bg-popover p-4">
                    <p className="">Redirecting...</p>
                    <Navigate to={redirectUrl} />
                </div>
            </div>
        )
    }

    return <>{children}</>
}

export default GuestGuard
