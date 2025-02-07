import DOMPurify from 'isomorphic-dompurify'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'

interface Props extends IBaseCompNoChild {
    description?: string
}

const BranchDescription = ({ className, description }: Props) => {
    return (
        <div
            dangerouslySetInnerHTML={{
                __html: DOMPurify.sanitize(
                    description && description.length > 0
                        ? description
                        : '<i>No Description</i>'
                ),
            }}
            className={cn(
                'prose dark:prose-invert w-full !max-w-full rounded-xl bg-secondary p-4 text-sm text-foreground/70 prose-h1: prose-p:text-foreground/80 prose-strong:text-foreground dark:bg-popover sm:text-sm',
                className
            )}
        />
    )
}

export default BranchDescription
