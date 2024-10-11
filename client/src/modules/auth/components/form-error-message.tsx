import { MdError } from 'react-icons/md'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {
    errorMessage?: string | null
}

const FormErrorMessage = ({ className, errorMessage }: Props) => {
    if (!errorMessage || errorMessage === null) return null

    return (
        <span
            className={cn(
                'flex items-center gap-x-2 rounded-md bg-destructive/10 p-2 py-2 text-sm text-orange-600 dark:bg-destructive/40',
                className
            )}
        >
            <MdError className="size-4" />
            {errorMessage}
        </span>
    )
}

export default FormErrorMessage
