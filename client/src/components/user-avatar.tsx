import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component'

interface Props extends IBaseComp {
    src: string
    fallback?: string
    fallbackClassName?: string
}

const UserAvatar = ({
    src,
    fallback = '-',
    className,
    fallbackClassName,
    children,
}: Props) => {
    if (fallback.length === 0 || fallback.length > 2)
        throw new Error(
            'User avatar fallback must atleast contain & not exceed 2 characters'
        )

    return (
        <Avatar className={cn('size-6', className)}>
            <AvatarImage src={src} />
            <AvatarFallback className={fallbackClassName}>
                {fallback}
            </AvatarFallback>
            {children}
        </Avatar>
    )
}

export default UserAvatar
