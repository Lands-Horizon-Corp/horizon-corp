import { toast } from 'sonner'
import { useRouter } from '@tanstack/react-router'

import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import { LogoutIcon, UserIcon } from '@/components/icons'

import { withCatchAsync } from '@/lib'
import {
    getUsersAccountTypeRedirectPage,
    serverRequestErrExtractor,
} from '@/helpers'
import { useUserAuthStore } from '@/store/user-auth-store'
import useConfirmModalStore from '@/store/confirm-modal-store'
import UserService from '@/horizon-corp/server/auth/UserService'

const NavProfileMenu = () => {
    const router = useRouter()
    const { onOpen } = useConfirmModalStore()
    const { currentUser, setCurrentUser } = useUserAuthStore()

    const handleSignout = async () => {
        const [error] = await withCatchAsync(UserService.SignOut())

        if (error) {
            const errorMessage = serverRequestErrExtractor({ error })
            toast.error(errorMessage)
            return
        }

        router.navigate({ to: '/auth/sign-in' })

        setCurrentUser(null)
        toast.success('Signed out')
    }

    if (!currentUser) return null

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button
                    size="icon"
                    variant="outline"
                    className="rounded-full p-0.5"
                >
                    <UserAvatar
                        className="size-full"
                        fallbackClassName="bg-transparent"
                        src={currentUser.media?.downloadURL ?? ''}
                        fallback={currentUser.username.charAt(0) ?? '-'}
                    />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
                <DropdownMenuLabel>{currentUser.username}</DropdownMenuLabel>
                <DropdownMenuSeparator />
                <DropdownMenuItem
                    onClick={() => {
                        const userProfileUrl =
                            getUsersAccountTypeRedirectPage(currentUser) +
                            '/profile'
                        router.navigate({
                            to: userProfileUrl,
                            hash: 'account-settings',
                        })
                    }}
                >
                    <UserIcon className="mr-2 size-4 duration-150 ease-in-out" />
                    Profile
                </DropdownMenuItem>
                <DropdownMenuItem
                    onClick={() =>
                        onOpen({
                            title: 'Sign Out',
                            description: 'Are you sure you want to sign out?',
                            onConfirm: () => handleSignout(),
                        })
                    }
                >
                    <LogoutIcon className="mr-2 size-4 duration-150 ease-in-out" />
                    <span>Sign Out</span>
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default NavProfileMenu
