import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component/base'

interface Props extends IBaseComp {}

const AuthPageWrapper = ({ className, children }: Props) => {
    return (
        <div
            className={cn(
                'flex w-full justify-center p-6 sm:w-fit sm:rounded-xl sm:bg-background/80 sm:shadow-2xl sm:backdrop-blur-sm',
                className
            )}
        >
            {children}
        </div>
    )
}

export default AuthPageWrapper
