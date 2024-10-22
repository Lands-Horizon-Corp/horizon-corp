import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { cn } from '@/lib/utils'
import { IBaseComp } from '@/types/component'

interface Props extends IBaseComp {
    src: string
    fallback: string
}

const UserAvatar = ({ src, fallback, className, children }: Props) => {
    if (fallback.length === 0 || fallback.length > 2)
        throw new Error(
            'User avatar fallback must atleast contain 2 characters'
        )

    return (
        <Avatar className={cn('size-6', className)}>
            <AvatarImage src={src} />
            <AvatarFallback>{fallback}</AvatarFallback>
            {children}
        </Avatar>
    )
}

export default UserAvatar
