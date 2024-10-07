import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component/base'

interface Props extends IBaseComp {}

const AuthPageWrapper = ({ className, children }: Props) => {
    return (
        <div
            className={cn(
                'flex w-full justify-center p-6 sm:w-fit sm:rounded-2xl sm:border sm:bg-background/80 sm:shadow-xl sm:backdrop-blur-md',
                className
            )}
        >
            {children}
        </div>
    )
}

export default AuthPageWrapper
