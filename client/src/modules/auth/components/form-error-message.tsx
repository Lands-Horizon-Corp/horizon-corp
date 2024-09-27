import { MdError } from 'react-icons/md'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {
    errorMessage?: string
}

const FormErrorMessage = ({ className, errorMessage }: Props) => {
    if (!errorMessage) return null

    return (
        <span
            className={cn(
                'mt-2 flex items-center gap-x-2 rounded-md bg-destructive/10 p-2 py-2 text-sm text-destructive',
                className
            )}
        >
            <MdError className="size-4" />
            {errorMessage}
        </span>
    )
}

export default FormErrorMessage
