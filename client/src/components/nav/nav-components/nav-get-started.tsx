import { Link } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'

import { getUsersAccountTypeRedirectPage } from '@/helpers'
import { useUserAuthStore } from '@/store/user-auth-store'

const NavGetStarted = () => {
    const { currentUser } = useUserAuthStore()

    if (!currentUser) return null

    const redirectUrl = getUsersAccountTypeRedirectPage(currentUser)

    return (
        <Link to={redirectUrl}>
            <Button className="scale-effects gap-x-2 rounded-full px-2">
                <UserAvatar
                    src={currentUser.media?.downloadURL ?? ''}
                    fallback={currentUser?.username.charAt(0) ?? '-'}
                    fallbackClassName="bg-secondary text-secondary-foreground"
                />
                <span className="mr-2">Get Started</span>
            </Button>
        </Link>
    )
}

export default NavGetStarted
