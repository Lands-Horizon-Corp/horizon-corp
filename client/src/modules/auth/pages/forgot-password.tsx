import z from 'zod'
import { toast } from 'sonner'
import { useState } from 'react'
import { useRouter, useSearch } from '@tanstack/react-router'

import { GoArrowLeft } from 'react-icons/go'
import { AiOutlineKey } from 'react-icons/ai'
import { MdMarkEmailRead } from 'react-icons/md'

import { Button } from '@/components/ui/button'
import AuthPageWrapper from '@/modules/auth/components/auth-page-wrapper'
import ForgotPasswordEmail, {
    TForgotPasswordEmail,
} from '@/modules/auth/components/forms/forgot-password-email'

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
    const [sent, setSent] = useState<null | TForgotPasswordEmail>(null)
    const preFilledValues = useSearch({ from: '/auth/forgot-password' })

    return (
        <div className="flex min-h-full flex-col items-center justify-center">
            <AuthPageWrapper>
                {!sent ? (
                    <ForgotPasswordEmail
                        onSuccess={(data) => {
                            toast.success('Reset link sent to your email')
                            setSent(data)
                        }}
                        defaultValues={preFilledValues}
                    />
                ) : (
                    <div className="flex w-full flex-col gap-y-4 sm:w-[390px]">
                        <div className="flex flex-col items-center gap-y-4 py-4 text-center">
                            <div className="relative p-8">
                                <AiOutlineKey className="size-[53px] text-[#FF7E47]" />
                                <div className="absolute inset-0 rounded-full bg-[#FF7E47]/20"></div>
                                <div className="absolute inset-5 rounded-full bg-[#FF7E47]/20"></div>
                            </div>
                            <MdMarkEmailRead className="size-[44px] text-[#34C759]" />
                            <p className="text-xl font-medium">
                                Password Reset Link Sent
                            </p>
                            <p className="text-sm text-foreground/70">
                                We've sent the reset link to your email address
                            </p>
                        </div>
                        <p className="text-center text-sm">
                            Didn&apos;t receive the email?
                        </p>
                        <Button
                            className="bg-[#34C759] hover:bg-[#38b558]"
                            onClick={() => {}}
                        >
                            Resend
                        </Button>
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
