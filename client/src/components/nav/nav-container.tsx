import { cn } from '@/lib'
import { IBaseComp } from '@/types'

const NavContainer = ({ children, className }: IBaseComp) => {
    return (
        <div className={cn('flex items-center gap-x-2', className)}>
            {children}
        </div>
    )
}

export default NavContainer
