import RootNav from '@/components/navbars/root-nav'
import { ThemeToggleMenu } from '@/components/theme-toggle'
import NavAuthContents from '@/components/navbars/auth-nav/nav-auth-contents'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component'

interface Props extends IBaseCompNoChild {}

const AuthNavBar = ({ className }: Props) => {
    return (
        <RootNav
            className={cn('fixed w-full pointer-events-none', className)}
            rightContents={
                <>
                    <NavAuthContents />
                    <ThemeToggleMenu />
                </>
            }
        />
    )
}

export default AuthNavBar
