import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component'

interface Props extends IBaseComp {
    src: string
    fallback?: string
    avatarClassname?: string
    fallbackClassName?: string
}

const UserAvatar = ({
    src,
    children,
    className,
    fallback = '-',
    avatarClassname,
    fallbackClassName,
}: Props) => {
    if (fallback.length === 0 || fallback.length > 2)
        throw new Error(
            'User avatar fallback must atleast contain & not exceed 2 characters'
        )

    return (
        <Avatar className={cn('size-6', className)}>
            <AvatarImage
                src={src}
                className={cn('object-cover', avatarClassname)}
            />
            <AvatarFallback className={fallbackClassName}>
                {fallback}
            </AvatarFallback>
            {children}
        </Avatar>
    )
}

export default UserAvatar
