import z from 'zod'
import { toast } from 'sonner'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { REGEXP_ONLY_DIGITS_AND_CHARS } from 'input-otp'

import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormMessage,
} from '@/components/ui/form'
import {
    InputOTP,
    InputOTPGroup,
    InputOTPSeparator,
    InputOTPSlot,
} from '@/components/ui/input-otp'
import { Button } from '@/components/ui/button'
import FormErrorMessage from '../form-error-message'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib/utils'
import { UserData } from '@/horizon-corp/types'
import { IAuthForm } from '@/types/auth/form-interface'
import { handleAxiosError } from '@/horizon-corp/helpers'
import useCountDown from '@/modules/auth/hooks/use-count-down'
import UserService from '@/horizon-corp/server/auth/UserService'
import useLoadingErrorState from '@/hooks/use-loading-error-state'
import { otpFormSchema } from '@/modules/auth/validations/otp-form'

type TVerifyForm = z.infer<typeof otpFormSchema>

interface Props extends IAuthForm<TVerifyForm, UserData> {
    verifyMode: 'mobile' | 'email'
    onSkip?: () => void
}

const VerifyForm = ({
    className,
    readOnly = false,
    verifyMode = 'mobile',
    defaultValues = { otp: '' },
    onSuccess,
    onError,
    onSkip,
}: Props) => {
    const [resent, setResent] = useState(false)
    const { loading, setLoading, error, setError } = useLoadingErrorState()

    const form = useForm({
        resolver: zodResolver(otpFormSchema),
        reValidateMode: 'onChange',
        defaultValues,
    })

    const handleSubmit = async (data: TVerifyForm) => {
        setError(null)
        setLoading(true)
        try {
            const parsedData = await otpFormSchema.parseAsync(data)

            // Send verify otp code to mark verify current signed in users email address
            if (verifyMode === 'email') {
                const response = await UserService.VerifyEmail(parsedData)
                onSuccess?.(response.data)
                toast.success('Email verified')
            }

            // Send verify otp code to mark verify current signed in users contact number
            if (verifyMode === 'mobile') {
                const response =
                    await UserService.VerifyContactNumber(parsedData)
                onSuccess?.(response.data)
                toast.success('Contact verified')
            }
        } catch (e) {
            const errorMessage = handleAxiosError(e)
            onError?.(e)
            setError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setLoading(false)
        }
    }

    const handleSendOTPVerification = async () => {
        setError(null)
        setLoading(true)
        try {
            // send otp verification code to current logged in user's email
            if (verifyMode === 'email') {
                await UserService.SendEmailVerification()
                setResent(true)
            }

            // send otp verification code to current logged in user's contact number
            if (verifyMode === 'mobile') {
                await UserService.SendContactVerification()
                setResent(true)
            }
        } catch (e) {
            const errorMessage = handleAxiosError(e)
            onError?.(e)
            setError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setLoading(false)
        }
    }

    return (
        <>
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit(handleSubmit)}
                    className={cn(
                        'flex min-w-[380px] flex-col gap-y-4',
                        className
                    )}
                >
                    <div className="flex flex-col items-center justify-center gap-y-4 pt-4">
                        <p className="text-xl font-medium">
                            Verify your{' '}
                            {verifyMode === 'mobile'
                                ? 'OTP Account'
                                : 'Email Address'}
                        </p>
                        <p className="max-w-[320px] text-center text-foreground/80">
                            Enter the one time password sent to your{' '}
                            {verifyMode === 'mobile'
                                ? 'Mobile Number'
                                : 'Email'}
                        </p>
                    </div>
                    <fieldset
                        disabled={readOnly || loading}
                        className="flex flex-col gap-y-4"
                    >
                        <FormField
                            control={form.control}
                            name="otp"
                            render={({ field }) => (
                                <FormItem className="flex flex-col items-center">
                                    <FormControl>
                                        <InputOTP
                                            autoFocus
                                            maxLength={6}
                                            onComplete={() =>
                                                form.handleSubmit(
                                                    handleSubmit
                                                )()
                                            }
                                            pattern={
                                                REGEXP_ONLY_DIGITS_AND_CHARS
                                            }
                                            containerClassName="mx-auto capitalize w-fit"
                                            {...field}
                                        >
                                            <InputOTPGroup>
                                                <InputOTPSlot index={0} />
                                                <InputOTPSlot index={1} />
                                                <InputOTPSlot index={2} />
                                            </InputOTPGroup>
                                            <InputOTPSeparator />
                                            <InputOTPGroup>
                                                <InputOTPSlot index={3} />
                                                <InputOTPSlot index={4} />
                                                <InputOTPSlot index={5} />
                                            </InputOTPGroup>
                                        </InputOTP>
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormErrorMessage errorMessage={error} />
                        {!resent && !loading && (
                            <p className="text-center text-sm opacity-85">
                                Didn&apos;t receive the code?{' '}
                                <span
                                    onClick={handleSendOTPVerification}
                                    className="cursor-pointer text-primary hover:underline"
                                >
                                    Resend Code
                                </span>
                            </p>
                        )}
                        {resent && !loading && (
                            <ResendCountDown
                                duration={5}
                                trigger={resent}
                                className="text-center text-sm opacity-85"
                                onComplete={() => setResent(false)}
                            />
                        )}
                        <div className="flex flex-col gap-y-2">
                            {onSkip && (
                                <Button
                                    variant={'outline'}
                                    disabled={readOnly}
                                    onClick={(e) => {
                                        e.preventDefault()
                                        onSkip()
                                    }}
                                >
                                    Skip
                                </Button>
                            )}
                            <Button type="submit">
                                {loading ? <LoadingSpinner /> : 'Submit'}
                            </Button>
                        </div>
                    </fieldset>
                </form>
            </Form>
        </>
    )
}

const ResendCountDown = ({
    trigger,
    duration,
    className,
    onComplete,
}: {
    trigger: boolean
    duration: number
    className?: string
    onComplete: () => void
}) => {
    const countDown = useCountDown({
        trigger,
        duration,
        onComplete,
    })

    return (
        <p className={className}>
            Please wait for {countDown}s to resend again
        </p>
    )
}

export default VerifyForm
