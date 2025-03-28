import DOMPurify from 'dompurify'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'

interface Props extends IBaseCompNoChild {
    expandedClassName?: string
    content: string
}

const RawDescription = ({ className, content }: Props) => {
    return (
        <div
            dangerouslySetInnerHTML={{
                __html: DOMPurify.sanitize(content),
            }}
            className={cn('', className)}
        ></div>
    )
}

export default RawDescription
