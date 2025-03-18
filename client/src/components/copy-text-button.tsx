import { toast } from 'sonner'
import { useState } from 'react'
import { CheckIcon, CopyIcon } from './icons'

import { cn } from '@/lib'

interface Props<TErr = unknown> {
    textContent: string
    copyInterval?: number
    className?: string
    successClassName?: string
    onCopySuccess?: () => void
    onCopyError?: (error: TErr) => void
}

const CopyTextButton = <TErr = unknown,>({
    className,
    textContent,
    successClassName,
    copyInterval = 2500,
    onCopyError,
    onCopySuccess,
}: Props<TErr>) => {
    const [copied, setCopied] = useState(false)

    const handleCopy = () => {
        navigator.clipboard
            .writeText(textContent)
            .then(() => {
                setCopied(true)
                toast.success('Coppied')
                onCopySuccess?.()
                setTimeout(() => setCopied(false), copyInterval)
            })
            .catch((err: TErr) => {
                onCopyError?.(err)
                toast.error('Sorry, Failed to copy')
            })
    }

    if (copied)
        return (
            <CheckIcon
                className={cn(
                    'inline text-primary',
                    className,
                    successClassName
                )}
            />
        )

    return (
        <CopyIcon
            onClick={() => handleCopy()}
            className={cn(
                'inline cursor-pointer duration-150 ease-in-out hover:text-foreground',
                className
            )}
        />
    )
}

export default CopyTextButton
