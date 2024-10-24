import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component'

const RootNav = ({ className, children }: IBaseComp) => {
    return (
        <nav
            className={cn(
                'z-10 flex min-h-[68px] items-center justify-between gap-x-2 px-4 lg:px-16',
                className
            )}
        >
            {children}
        </nav>
    )
}

export default RootNav
