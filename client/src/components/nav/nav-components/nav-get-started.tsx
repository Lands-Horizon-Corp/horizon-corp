import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'

import { getUsersAccountTypeRedirectPage } from '@/helpers'
import { useUserAuthStore } from '@/store/user-auth-store'
import { toast } from 'sonner'
import logger from '@/helpers/loggers/logger'
import { useRouter } from '@tanstack/react-router'

const NavGetStarted = () => {
    const { navigate } = useRouter()
    const { currentUser } = useUserAuthStore()

    if (!currentUser) return null

    const redirectUrl = getUsersAccountTypeRedirectPage(currentUser)

    return (
        <Button
            onClick={() => {
                if (redirectUrl === '/') {
                    toast.error(
                        "'Get Started' is not available for your account type. Contact your admin for help."
                    )
                    return
                }

                navigate({ to: redirectUrl })
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
    )
}

export default NavGetStarted
