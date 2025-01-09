import { toast } from 'sonner'
import { useMutation } from '@tanstack/react-query'

import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { withCatchAsync } from '@/utils'
import UseCooldown from '@/hooks/use-cooldown'
import { AccountType } from '@/horizon-corp/types'
import { serverRequestErrExtractor } from '@/helpers'
import UserService from '@/horizon-corp/services/auth/UserService'

type TSentTo = { key: string; accountType: AccountType }

interface Props {
    duration: number
    interval: number
    onErrorMessage: (errorMessage: string) => void
    sentTo: TSentTo
}

const ResendPasswordResetLinkButton = ({
    sentTo,
    duration,
    interval,
    onErrorMessage,
}: Props) => {
    const { cooldownCount, startCooldown } = UseCooldown({
        cooldownDuration: duration,
        counterInterval: interval,
    })

    const { mutate: resendResetLink, isPending } = useMutation<
        void,
        string,
        TSentTo
    >({
        mutationKey: ['resend-reset-password-link'],
        mutationFn: async (sentTo) => {
            if (!sentTo) return

            const [error] = await withCatchAsync(
                UserService.ForgotPassword(sentTo)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                onErrorMessage(errorMessage)
                toast.error(errorMessage)
                return
            }

            startCooldown()
            toast.success(`Password Reset Link was resent to ${sentTo.key}`)
        },
    })

    return (
        <Button
            disabled={isPending || cooldownCount > 0}
            variant={cooldownCount || isPending ? 'secondary' : 'default'}
            onClick={() => resendResetLink(sentTo)}
        >
            {isPending && <LoadingSpinner />}
            {!isPending && cooldownCount <= 0 && 'Resend'}
            {cooldownCount > 0 && `Resend again in ${cooldownCount}s`}
        </Button>
    )
}

export default ResendPasswordResetLinkButton
