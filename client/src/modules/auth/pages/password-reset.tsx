import z from 'zod'
import { useEffect, useState } from 'react'
import { useParams } from '@tanstack/react-router'
import { useRouter } from '@tanstack/react-router'

import { GoArrowLeft } from 'react-icons/go'

import { Button } from '@/components/ui/button'
import ResetPasswordForm from '@/modules/auth/components/forms/reset-password-form'
import LoadingCircle from '@/components/loader/loading-circle'

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
        <div className="flex min-h-[50dvh] flex-1 justify-center px-4 py-4">
            {!done && resetEntry && (
                <ResetPasswordForm onSuccess={() => setDone(true)} />
            )}
            {!done && !loading && !resetEntry && (
                <div className="flex w-[390px] flex-col gap-y-4">
                    <div className="flex flex-col items-center gap-y-4 py-4 text-center">
                        <img src="/e-coop-logo-1.png" className="size-24" />
                        <p className="text-xl font-medium">
                            Password Reset Link Invalid
                        </p>
                        <p className="text-sm text-foreground/70">
                            Sorry, but the link is invalid
                        </p>
                    </div>
                    <Button
                        variant={'ghost'}
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
                <div className="flex w-[390px] flex-col gap-y-4">
                    <div className="flex flex-col items-center gap-y-4 py-4 text-center">
                        <img src="/e-coop-logo-1.png" className="size-24" />
                        <p className="text-xl font-medium">Password Saved</p>
                        <p className="text-sm text-foreground/70">
                            Your new password has been set.
                        </p>
                    </div>
                    <Button
                        variant={'ghost'}
                        className="text-foreground/60"
                        onClick={() => {
                            router.navigate({
                                to: '/auth/sign-in',
                            })
                        }}
                    >
                        <GoArrowLeft className="mr-2" /> Go to login
                    </Button>
                </div>
            )}
        </div>
    )
}

export default PasswordResetPage
