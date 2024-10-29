import { useState } from 'react'
import { useRouter, useSearch } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'
import GuestGuard from '@/components/wrappers/guest-guard'
import AuthPageWrapper from '../components/auth-page-wrapper'
import ForgotPasswordEmail, {
    TForgotPasswordEmail,
} from '@/modules/auth/components/forms/forgot-password-email'
import FormErrorMessage from '../../../components/ui/form-error-message'
import { EmailCheckIcon, ArrowLeftIcon } from '@/components/icons'
import ResendPasswordResetLinkButton from '../components/resend-password-reset-link-button'

import useLoadingErrorState from '@/hooks/use-loading-error-state'

const ForgotPasswordPage = () => {
    const router = useRouter()
    const preFilledValues = useSearch({ from: '/auth/forgot-password' })
    const [sentTo, setSentTo] = useState<null | TForgotPasswordEmail>(null)

    const { error, setError } = useLoadingErrorState()

    return (
        <GuestGuard>
            <div className="flex min-h-full flex-col items-center justify-center">
                <AuthPageWrapper>
                    {!sentTo ? (
                        <ForgotPasswordEmail
                            onSuccess={(data) => {
                                setSentTo(data)
                            }}
                            defaultValues={preFilledValues}
                        />
                    ) : (
                        <div className="flex w-full flex-col gap-y-4 sm:w-[390px]">
                            <div className="flex flex-col items-center gap-y-4 py-4 text-center">
                                <div className="relative p-8">
                                    <EmailCheckIcon className="size-[53px] text-primary" />
                                    <div className="absolute inset-0 rounded-full bg-[#FF7E47]/20" />
                                    <div className="absolute inset-5 rounded-full bg-[#FF7E47]/20" />
                                </div>
                                <p className="text-xl font-medium">
                                    Password Reset Link Sent
                                </p>
                                <p className="text-sm text-foreground/70">
                                    We've sent the reset link to your email
                                    address
                                </p>
                            </div>
                            <FormErrorMessage errorMessage={error} />
                            <p className="text-center text-sm text-foreground/80">
                                Didn&apos;t receive the email?
                            </p>
                            <ResendPasswordResetLinkButton
                                duration={12}
                                interval={1000}
                                sentTo={sentTo}
                                onErrorMessage={(errorMessage) =>
                                    setError(errorMessage)
                                }
                            />
                            <Button
                                variant={'ghost'}
                                className="text-foreground/60"
                                onClick={() => {
                                    router.navigate({
                                        to: '/auth/sign-in',
                                        search: preFilledValues,
                                    })
                                }}
                            >
                                <ArrowLeftIcon className="mr-2" /> Back to Login
                            </Button>
                        </div>
                    )}
                </AuthPageWrapper>
            </div>
        </GuestGuard>
    )
}

export default ForgotPasswordPage
