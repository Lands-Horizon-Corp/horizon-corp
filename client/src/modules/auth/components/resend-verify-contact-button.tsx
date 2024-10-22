import { toast } from 'sonner'
import { useMutation } from '@tanstack/react-query'

import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib'
import UseCooldown from '@/hooks/use-cooldown'
import { serverRequestErrExtractor } from '@/helpers'
import UserService from '@/horizon-corp/server/auth/UserService'

interface Props {
    verifyMode: 'email' | 'mobile'
    duration: number
    interval: number
}

const ResendVerifyContactButton = ({
    verifyMode,
    duration,
    interval,
}: Props) => {
    const { cooldownCount, startCooldown } = UseCooldown({
        cooldownDuration: duration,
        counterInterval: interval,
    })

    const { mutate: resendOtpVerification, isPending: isResending } =
        useMutation<void, string>({
            mutationKey: ['verify-resend', verifyMode],
            mutationFn: async () => {
                try {
                    if (verifyMode === 'email') {
                        await UserService.SendEmailVerification()
                        startCooldown()
                        return
                    }

                    if (verifyMode === 'mobile') {
                        await UserService.SendContactVerification()
                        startCooldown()
                        return
                    }

                    throw 'Unkown verify mode'
                } catch (error) {
                    const errorMessage = serverRequestErrExtractor({ error })
                    toast.error(errorMessage)
                    throw errorMessage
                }
            },
        })
    return (
        <Button
            size="sm"
            variant="ghost"
            className={cn('underline', cooldownCount > 0 && 'no-underline')}
            onClick={(e) => {
                e.preventDefault()
                resendOtpVerification()
            }}
            disabled={isResending || cooldownCount > 0}
        >
            {isResending && <LoadingSpinner />}
            {!isResending && cooldownCount <= 0 && 'Resend Code'}
            {cooldownCount > 0 && `Resend again in ${cooldownCount}s`}
        </Button>
    )
}

export default ResendVerifyContactButton
