import { AiOutlineLoading } from 'react-icons/ai'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {}

const LoadingCircle = ({ className }: Props) => {
    return (
        <AiOutlineLoading
            className={cn(
                'size-4 animate-spin [animation-duration:.5s]',
                className
            )}
        />
    )
}

export default LoadingCircle
