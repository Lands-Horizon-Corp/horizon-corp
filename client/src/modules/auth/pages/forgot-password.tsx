import z from 'zod'
import { useState } from 'react'
import { useRouter, useSearch } from '@tanstack/react-router'

import { GoArrowLeft } from 'react-icons/go'

import { Button } from '@/components/ui/button'
import { EmailCheckIcon } from '@/components/icons'
import ForgotPasswordEmail, {
    TForgotPasswordEmail,
} from '@/modules/auth/components/forms/forgot-password-email'
import FormErrorMessage from '../components/form-error-message'
import ResetPasswordButton from '../components/resent-password-button'
import AuthPageWrapper from '@/modules/auth/components/auth-page-wrapper'

import useLoadingErrorState from '@/hooks/use-loading-error-state'
import { userAccountTypeSchema, emailSchema } from '../validations/common'

export const ForgotPasswordPageSearchSchema = z.object({
    email: z.string().optional().default('').or(emailSchema),
    accountType: z
        .string()
        .optional()
        .default('Member')
        .pipe(userAccountTypeSchema)
        .catch('Member'),
})

interface Props {}

const ForgotPasswordPage = (_props: Props) => {
    const router = useRouter()
    const preFilledValues = useSearch({ from: '/auth/forgot-password' })
    const [sentTo, setSentTo] = useState<null | TForgotPasswordEmail>(null)

    const { error, setError } = useLoadingErrorState()

    return (
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
                                We've sent the reset link to your email address
                            </p>
                        </div>
                        <FormErrorMessage errorMessage={error} />
                        <p className="text-center text-sm text-foreground/80">
                            Didn&apos;t receive the email?
                        </p>
                        <ResetPasswordButton
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
                            <GoArrowLeft className="mr-2" /> Back to Login
                        </Button>
                    </div>
                )}
            </AuthPageWrapper>
        </div>
    )
}

export default ForgotPasswordPage
