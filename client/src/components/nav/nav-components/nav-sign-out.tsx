import { toast } from 'sonner'
import { useRouter } from '@tanstack/react-router'
import { useMutation } from '@tanstack/react-query'

import { Button } from '@/components/ui/button'

import { withCatchAsync } from '@/lib'
import { serverRequestErrExtractor } from '@/helpers'
import { useUserAuthStore } from '@/store/user-auth-store'
import UserService from '@/horizon-corp/server/auth/UserService'
import useConfirmModalStore from '@/store/confirm-modal-store'

const NavSignOut = () => {
    const router = useRouter()
    const { onOpen } = useConfirmModalStore()
    const { authStatus, setCurrentUser } = useUserAuthStore()

    const { mutate: handleSignout, isPending: isSigningOut } = useMutation({
        mutationKey: ['sign-out'],
        mutationFn: async () => {
            const [error] = await withCatchAsync(UserService.SignOut())

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                return
            }

            setCurrentUser(null)
            router.navigate({ to: '/auth/sign-in' })
            toast.success('Signed out')
        },
    })

    if (authStatus === 'unauthorized') return null

    return (
        <Button
            variant="outline"
            disabled={isSigningOut}
            onClick={() =>
                onOpen({
                    title: 'Sign Out',
                    description: 'Are you sure you want to sign out?',
                    onConfirm: () => handleSignout(),
                })
            }
            className="scale-effects rounded-full"
        >
            Sign-Out
        </Button>
    )
}

export default NavSignOut
