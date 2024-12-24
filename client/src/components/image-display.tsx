import { Avatar, AvatarFallback, AvatarImage } from './ui/avatar'

import { cn } from '@/lib'
import { IBaseComp } from '@/types'

interface Props extends IBaseComp {
    src?: string
    fallback?: string
    fallbackClassName?: string
}

const ImageDisplay = ({
    src,
    fallback,
    children,
    className,
    fallbackClassName,
}: Props) => {
    return (
        <Avatar className={cn('size-6', className)}>
            <AvatarImage src={src ?? '-'} />
            <AvatarFallback className={fallbackClassName}>
                {fallback}
            </AvatarFallback>
            {children}
        </Avatar>
    )
}

export default ImageDisplay
