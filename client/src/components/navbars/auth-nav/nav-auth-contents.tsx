import { toast } from 'sonner'
import { Link, useLocation, useRouter } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'

import {
    getUsersAccountTypeRedirectPage,
    serverRequestErrExtractor,
} from '@/helpers'
import { cn, withCatchAsync } from '@/lib/utils'
import { useMutation } from '@tanstack/react-query'
import { IBaseCompNoChild } from '@/types/component'
import useConfirmModalStore from '@/store/confirm-modal-store'
import UserService from '@/horizon-corp/server/auth/UserService'
import { useUserAuthStore } from '@/store/user-auth-store'

const NavAuthContents = ({ className }: IBaseCompNoChild) => {
    const router = useRouter()
    const { onOpen } = useConfirmModalStore()
    const { currentUser, setCurrentUser } = useUserAuthStore()

    const pathname = useLocation({
        select: (location) => location.pathname,
    })

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

    const shouldShowAuthButtons = !pathname.startsWith('/auth/verify')

    return (
        <div className={cn('flex items-center gap-x-2', className)}>
            {currentUser && (
                <>
                    <Button
                        disabled={isSigningOut}
                        onClick={() => {
                            router.navigate({
                                to: getUsersAccountTypeRedirectPage(
                                    currentUser
                                ),
                            })
                        }}
                        className="scale-effects gap-x-2 rounded-full px-2"
                    >
                        <UserAvatar
                            src={currentUser.media?.downloadURL ?? ''}
                            fallback={currentUser?.username.charAt(0) ?? '-'}
                            fallbackClassName="bg-secondary text-secondary-foreground"
                        />
                        <span className="mr-2">Get Started</span>
                    </Button>
                    <Button
                        variant="outline"
                        disabled={isSigningOut}
                        onClick={() =>
                            onOpen({
                                title: 'Sign Out',
                                description:
                                    'Are you sure you want to sign out?',
                                onConfirm: () => handleSignout(),
                            })
                        }
                        className="scale-effects rounded-full"
                    >
                        Sign-Out
                    </Button>
                </>
            )}
            {shouldShowAuthButtons &&
                !currentUser &&
                pathname !== '/auth/sign-in' && (
                    <Link to="/auth/sign-in">
                        <Button className="scale-effects rounded-full bg-green-500 text-white hover:bg-green-500">
                            Sign-In
                        </Button>
                    </Link>
                )}
            {shouldShowAuthButtons &&
                !currentUser &&
                pathname !== '/auth/sign-up' && (
                    <Link to="/auth/sign-up">
                        <Button
                            variant="outline"
                            className="scale-effects rounded-full"
                        >
                            Sign-Up
                        </Button>
                    </Link>
                )}
            {/* TODO: Once auth service is implemeented, add/show user avatar here */}
        </div>
    )
}

export default NavAuthContents
