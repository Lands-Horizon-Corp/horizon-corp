import z from 'zod'
import { useEffect, useState } from 'react'
import { useParams } from '@tanstack/react-router'
import { useRouter } from '@tanstack/react-router'

import { GoArrowLeft } from 'react-icons/go'
import { AiOutlineKey } from 'react-icons/ai'

import { Button } from '@/components/ui/button'
import LoadingCircle from '@/components/loader/loading-circle'
import AuthPageWrapper from '@/modules/auth/components/auth-page-wrapper'
import ResetPasswordForm from '@/modules/auth/components/forms/reset-password-form'

export const PasswordResetPagePathSchema = z.object({
    resetId: z
        .string({ required_error: 'Missing Reset Link' })
        .uuid('Invalid Reset Link'),
})

interface Props {}

const PasswordResetPage = (_props: Props) => {
    const router = useRouter()
    const pathParams = useParams({ from: '/auth/password-reset/$resetId' })
    const [done, setDone] = useState(false)
    const [loading, setLoading] = useState(false)
    const [resetEntry, setResetEntry] = useState<{} | null>(null) // replace {} with type of reset entry

    useEffect(() => {
        const { success, data } =
            PasswordResetPagePathSchema.safeParse(pathParams)

        if (!success) return setResetEntry(null)

        setLoading(true)
        // todo fetch reset entry from db if it exist
        // authService.verifyResetEntry(data?.resetId)
        // then set the data to resetEntry state
        setTimeout(() => {
            setLoading(false)
            setResetEntry(Math.floor(Math.random() * 2) % 2 === 0 ? null : {})
        }, 1000)
    }, [])

    return (
        <div className="flex min-h-full flex-col items-center justify-center">
            <AuthPageWrapper>
                {!done && resetEntry && (
                    <ResetPasswordForm onSuccess={() => setDone(true)} />
                )}
                {!done && !loading && !resetEntry && (
                    <div className="flex w-full flex-col gap-y-4 sm:w-[390px]">
                        <div className="flex flex-col items-center gap-y-4 py-4 text-center">
                            <div className="relative p-8">
                                <AiOutlineKey className="size-[53px] text-[#FF7E47]" />
                                <div className="absolute inset-0 rounded-full bg-[#FF7E47]/20"></div>
                                <div className="absolute inset-5 rounded-full bg-[#FF7E47]/20"></div>
                            </div>
                            <p className="text-xl font-medium">
                                Invalid reset link
                            </p>
                            <p className="text-sm text-foreground/70">
                                Sorry, but the reset link you have is invalid
                            </p>
                        </div>
                        <Button
                            variant="ghost"
                            className="text-foreground/60"
                            onClick={() => {
                                router.navigate({
                                    to: '/auth/sign-in',
                                })
                            }}
                        >
                            <GoArrowLeft className="mr-2" /> Back to Login
                        </Button>
                    </div>
                )}
                {loading && (
                    <div className="flex flex-col items-center gap-y-2 py-16">
                        <LoadingCircle />
                        <p className="text-center text-sm text-foreground/50">
                            verifying reset password link
                        </p>
                    </div>
                )}
                {done && (
                    <div className="flex w-full flex-col gap-y-4 sm:w-[390px]">
                        <div className="flex flex-col items-center gap-y-4 py-4 text-center">
                            <div className="relative p-8">
                                <AiOutlineKey className="size-[53px] text-green-500" />
                                <div className="absolute inset-0 rounded-full bg-green-500/20"></div>
                                <div className="absolute inset-5 rounded-full bg-green-500/20"></div>
                            </div>
                            <p className="text-xl font-medium">Password Set</p>
                            <p className="text-sm text-foreground/70">
                                Your new password has been saved.
                            </p>
                        </div>
                        <Button
                            onClick={() => {
                                router.navigate({
                                    to: '/auth/sign-in',
                                })
                            }}
                        >
                            Sign In Now
                        </Button>
                    </div>
                )}
            </AuthPageWrapper>
        </div>
    )
}

export default PasswordResetPage
