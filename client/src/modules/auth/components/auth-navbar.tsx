import { Link, useLocation } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {}

const AuthNavBar = (className: Props) => {
    const pathname = useLocation({
        select: (location) => location.pathname,
    })

    return (
        <div
            className={cn(
                'fixed left-0 top-0 z-50 flex w-full justify-between bg-background/80 px-6 py-4 backdrop-blur-sm',
                className
            )}
        >
            <Link to="/">
                <img src="/e-coop-logo-1.png" className="size-12" />
            </Link>

            <div className="flex gap-x-2">
                {!pathname.startsWith('/auth/verify') && (
                    <>
                        <Link to="/auth/sign-up">
                            <Button variant="outline" className="rounded-full">
                                Sign Up
                            </Button>
                        </Link>
                        <Link to="/auth/sign-in">
                            <Button
                                variant="default"
                                className="rounded-full bg-[#34C759] hover:bg-[#38b558]"
                            >
                                Sign In
                            </Button>
                        </Link>
                    </>
                )}
            </div>
        </div>
    )
}

export default AuthNavBar
