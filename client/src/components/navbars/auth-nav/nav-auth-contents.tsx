import { Link, useLocation } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component'

interface Props extends IBaseCompNoChild {}

const NavAuthContents = ({ className }: Props) => {
    const pathname = useLocation({
        select: (location) => location.pathname,
    })

    const shouldShowAuthButtons = !pathname.startsWith('/auth/verify')

    return (
        <div className={cn('flex items-center gap-x-2', className)}>
            {shouldShowAuthButtons && pathname !== '/auth/sign-in' && (
                <Link to="/auth/sign-in">
                    <Button className="scale-effects rounded-full bg-green-500 text-white hover:bg-green-500">
                        Sign-In
                    </Button>
                </Link>
            )}
            {shouldShowAuthButtons && pathname !== '/auth/sign-up' && (
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
