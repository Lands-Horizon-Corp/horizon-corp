import { toast } from 'sonner'
import { useCallback } from 'react'

import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { withCatchAsync } from '@/lib'
import UseCooldown from '@/hooks/use-cooldown'
import { AccountType } from '@/horizon-corp/types'
import { serverRequestErrExtractor } from '@/helpers'
import UserService from '@/horizon-corp/server/auth/UserService'
import useLoadingErrorState from '@/hooks/use-loading-error-state'

interface Props {
    duration: number
    interval: number
    onErrorMessage: (errorMessage: string) => void
    sentTo: { email: string; accountType: AccountType }
}

const ResetPasswordButton = ({
    sentTo,
    duration,
    interval,
    onErrorMessage,
}: Props) => {
    const { loading, setLoading } = useLoadingErrorState()
    const { cooldownCount, startCooldown } = UseCooldown({
        cooldownDuration: duration,
        counterInterval: interval,
    })

    const resentCode = useCallback(async () => {
        if (!sentTo) return
        setLoading(true)

        const [error] = await withCatchAsync(UserService.ForgotPassword(sentTo))
        setLoading(false)

        if (error) {
            const errorMessage = serverRequestErrExtractor({ error })
            onErrorMessage(errorMessage)
            toast.error(errorMessage)
            return
        }

        startCooldown()
        toast.success(`Password Reset Link was resent to ${sentTo.email}`)
    }, [sentTo, startCooldown])

    return (
        <Button
            disabled={loading || cooldownCount > 0}
            variant={cooldownCount || loading ? 'secondary' : 'default'}
            onClick={resentCode}
        >
            {loading && <LoadingSpinner />}
            {!loading && cooldownCount <= 0 && 'Resend'}
            {cooldownCount > 0 && `Resent again in ${cooldownCount}s`}
        </Button>
    )
}

export default ResetPasswordButton
