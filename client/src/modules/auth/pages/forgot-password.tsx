import z from 'zod'
import { toast } from 'sonner'
import { useState } from 'react'
import { useRouter, useSearch } from '@tanstack/react-router'

import { GoArrowLeft } from 'react-icons/go'

import { Button } from '@/components/ui/button'
import ForgotPasswordEmail from '@/modules/auth/components/forms/forgot-password-email'

import { memberTypeSchema, emailSchema } from '../validations'

export const ForgotPasswordPageSearchSchema = z.object({
    email: z.string().optional().default('').or(emailSchema),
    mode: z
        .string()
        .optional()
        .default('Member')
        .pipe(memberTypeSchema)
        .catch('Member'),
})

interface Props {}

const ForgotPasswordPage = (_props: Props) => {
    const router = useRouter()
    const [sent, setSent] = useState(false)
    const preFilledValues = useSearch({ from: '/auth/forgot-password' })

    return (
        <div className="flex flex-1 justify-center px-4 py-4">
            {!sent ? (
                <ForgotPasswordEmail
                    onSuccess={() => {
                        toast.success('Reset link sent to your email')
                        setSent(true)
                    }}
                    defaultValues={preFilledValues}
                />
            ) : (
                <div className="flex w-[390px] flex-col gap-y-4">
                    <div className="flex flex-col items-center gap-y-4 py-4 text-center">
                        <img src="/e-coop-logo-1.png" className="size-24" />
                        <p className="text-xl font-medium">
                            Password Reset Link Sent
                        </p>
                        <p className="text-sm text-foreground/70">
                            We'be sent the reset link to your email address
                        </p>
                    </div>
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
        </div>
    )
}

export default ForgotPasswordPage
