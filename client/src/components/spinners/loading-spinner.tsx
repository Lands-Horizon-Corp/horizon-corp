import { LoadingCircleIcon } from '@/components/icons'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component'

interface Props extends IBaseCompNoChild {}

const LoadingSpinner = ({ className }: Props) => {
    return (
        <LoadingCircleIcon
            className={cn(
                'size-4 animate-spin [animation-duration:1s]',
                className
            )}
        />
    )
}

export default LoadingSpinner
