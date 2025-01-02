import { Image2Icon } from '@/components/icons'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'

import { cn } from '@/lib'
import { IBaseComp } from '@/types'

interface Props extends IBaseComp {
    src?: string
    fallback?: string
    imageClassName?: string
    fallbackClassName?: string
}

const ImageDisplay = ({
    src,
    fallback,
    children,
    className,
    imageClassName,
    fallbackClassName,
}: Props) => {
    return (
        <Avatar
            className={cn('size-6 bg-secondary dark:bg-popover', className)}
        >
            <AvatarImage
                className={cn('object-cover', imageClassName)}
                src={src ?? '-'}
            />
            <AvatarFallback className={fallbackClassName}>
                {fallback ? fallback : <Image2Icon className='text-foreground/20 size-[50%]' />}
            </AvatarFallback>
            {children}
        </Avatar>
    )
}

export default ImageDisplay
