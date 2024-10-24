import { Link } from '@tanstack/react-router'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import EcoopLogo from '@/components/ecoop-logo'

interface Props extends IBaseCompNoChild {
    linkUrl?: string
}

const NavEcoopLogo = ({ linkUrl = '/', className }: Props) => {
    return (
        <Link to={linkUrl}>
            <EcoopLogo className={cn('size-[46px]', className)} />
        </Link>
    )
}

export default NavEcoopLogo
