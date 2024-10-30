import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { useMutation } from '@tanstack/react-query'
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
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib'
import { otpSchema } from '@/validations'
import UseCooldown from '@/hooks/use-cooldown'
import { UserData } from '@/horizon-corp/types'
import { serverRequestErrExtractor } from '@/helpers'
import UserService from '@/horizon-corp/server/auth/UserService'

type TVerifyMode = 'email' | 'mobile'

interface Props {
    verifyMode: TVerifyMode
    onSuccess?: (newUserData: UserData) => void
}

const VerifyContactBar = ({ verifyMode, onSuccess }: Props) => {
    const form = useForm<z.infer<typeof otpSchema>>({
        resolver: zodResolver(otpSchema),
        reValidateMode: 'onChange',
        defaultValues: {
            otp: '',
        },
    })

    const { mutate: handleVerify, isPending } = useMutation<
        UserData,
        string,
        z.infer<typeof otpSchema>
    >({
        mutationKey: ['verify', verifyMode],
        mutationFn: async (data) => {
            try {
                if (verifyMode === 'email') {
                    const response = await UserService.VerifyEmail(data)
                    onSuccess?.(response.data)
                    toast.success('Email verified')
                    return response.data
                }

                if (verifyMode === 'mobile') {
                    const response = await UserService.VerifyContactNumber(data)
                    onSuccess?.(response.data)
                    toast.success('Contact verified')
                    return response.data
                }

                throw 'Unknown verify mode'
            } catch (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }
        },
    })

    /*
Security
- username
- email
- password
- contact number
*/

    return (
        <div className="flex flex-col justify-between gap-y-4 rounded-xl border border-orange-400/20 bg-secondary/70 p-3 lg:flex-row">
            <div className="space-y-2 text-center text-xs sm:text-left sm:text-sm">
                <p className="text-sm font-medium capitalize">
                    Verify {verifyMode}
                </p>
                <p className="text-foreground/50 text-sm">
                    Please verify {verifyMode} to enable other features/actions.{' '}
                    <ResendCode
                        verifyMode={verifyMode}
                        duration={10}
                        interval={1_000}
                    />
                </p>
            </div>
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit((data) => handleVerify(data))}
                >
                    <fieldset
                        disabled={isPending}
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
                                                form.handleSubmit((data) =>
                                                    handleVerify(data)
                                                )()
                                            }
                                            pattern={
                                                REGEXP_ONLY_DIGITS_AND_CHARS
                                            }
                                            containerClassName="mx-auto capitalize w-fit"
                                            {...field}
                                        >
                                            <InputOTPGroup>
                                                <InputOTPSlot
                                                    index={0}
                                                    className="size-7 sm:size-8"
                                                />
                                                <InputOTPSlot
                                                    index={1}
                                                    className="size-7 sm:size-8"
                                                />
                                                <InputOTPSlot
                                                    index={2}
                                                    className="size-7 sm:size-8"
                                                />
                                            </InputOTPGroup>
                                            <InputOTPSeparator />
                                            <InputOTPGroup>
                                                <InputOTPSlot
                                                    index={3}
                                                    className="size-7 sm:size-8"
                                                />
                                                <InputOTPSlot
                                                    index={4}
                                                    className="size-7 sm:size-8"
                                                />
                                                <InputOTPSlot
                                                    index={5}
                                                    className="size-7 sm:size-8"
                                                />
                                            </InputOTPGroup>
                                        </InputOTP>
                                    </FormControl>
                                    <FormMessage className="text-xs text-rose-500" />
                                </FormItem>
                            )}
                        />
                    </fieldset>
                </form>
            </Form>
        </div>
    )
}

const ResendCode = ({
    disabled = false,
    verifyMode,
    duration,
    interval,
}: {
    disabled? : boolean
    verifyMode: TVerifyMode
    duration: number
    interval: number
}) => {
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
        <span
            className={cn(
                'cursor-pointer underline',
                cooldownCount > 0 && 'cursor-not-allowed no-underline',
                disabled && 'cursor-not-allowed opacity-30'
            )}
            onClick={(e) => {
                if(disabled) return;

                e.preventDefault()
                resendOtpVerification()
            }}
        >
            {isResending && <LoadingSpinner className="size-3 inline-block" />}
            {!isResending && cooldownCount <= 0 && "Didn't get the code?"}
            {cooldownCount > 0 && `Resend again in ${cooldownCount}s`}
        </span>
    )
}

export default VerifyContactBar
