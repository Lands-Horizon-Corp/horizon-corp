import z from 'zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

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
import LoadingCircle from '@/components/loader/loading-circle'

import { cn } from '@/lib/utils'
import useCountDown from '../../hooks/use-count-down'
import { REGEXP_ONLY_DIGITS_AND_CHARS } from 'input-otp'
import { verifyFormSchema } from '../../validations/verify-form'

import { UserBase, UserStatus } from '@/types'
import { IAuthForm } from '@/types/auth/form-interface'

type TVerifyForm = z.infer<typeof verifyFormSchema>

interface Props extends IAuthForm<TVerifyForm> {
    id: string
    verifyMode: 'mobile' | 'email'
}

const ResendCountDown = ({
    trigger,
    duration,
    onComplete,
}: {
    trigger: boolean
    duration: number
    onComplete: () => void
}) => {
    const countDown = useCountDown({
        trigger,
        duration,
        onComplete,
    })

    return <p>Please wait for {countDown}s to resend again</p>
}

const VerifyForm = ({
    className,
    readOnly = false,
    verifyMode = 'mobile',
    defaultValues = { code: '' },
    onSuccess,
}: Props) => {
    const [loading, setLoading] = useState(false)
    const [resent, setResent] = useState(false)

    const form = useForm({
        resolver: zodResolver(verifyFormSchema),
        reValidateMode: 'onChange',
        defaultValues,
    })

    const handleSubmit = (data: TVerifyForm) => {
        const parsedData = verifyFormSchema.parse(data)
        setLoading(true)
        // TODO: Add functionality, delete the code below,
        // it is just for mocking ui flow
        setTimeout(() => {
            onSuccess?.(
                (verifyMode === 'mobile'
                    ? {
                          id: '215',
                          username: 'Jervx',
                          validEmail: false,
                          validContactNumber: true,
                          status: UserStatus['Pending'],
                          profilePicture: {
                              url: 'https://mrwallpaper.com/images/hd/suit-rick-and-morty-phone-5divv4gzo6gowk46.jpg',
                          },
                      }
                    : {
                          id: '215',
                          username: 'Jervx',
                          validEmail: true,
                          validContactNumber: true,
                          status: UserStatus['Pending'],
                          profilePicture: {
                              url: 'https://mrwallpaper.com/images/hd/suit-rick-and-morty-phone-5divv4gzo6gowk46.jpg',
                          },
                      }) as any as UserBase
            )
            setLoading(false)
        }, 1000)
    }

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit(handleSubmit)}
                className={cn('flex min-w-[390px] flex-col gap-y-4', className)}
            >
                <div className="flex flex-col items-center justify-center gap-y-4 pt-4">
                    <img src="/e-coop-logo-1.png" className="size-40" />
                    <p className="text-xl font-medium">
                        Verify your{' '}
                        {verifyMode === 'mobile'
                            ? 'OTP Account'
                            : 'Email Address'}
                    </p>
                    <p className="max-w-[320px] text-center text-foreground/80">
                        Enter the one time password sent to your mobile number{' '}
                        {verifyMode === 'mobile' ? 'Mobile Number' : 'Email'}
                    </p>
                </div>
                <fieldset
                    disabled={readOnly || loading}
                    className="flex flex-col gap-y-4"
                >
                    <FormField
                        control={form.control}
                        name="code"
                        render={({ field }) => (
                            <FormItem className="flex flex-col items-center">
                                <FormControl>
                                    <InputOTP
                                        autoFocus
                                        maxLength={6}
                                        onComplete={() =>
                                            form.handleSubmit(handleSubmit)()
                                        }
                                        pattern={REGEXP_ONLY_DIGITS_AND_CHARS}
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
                    {!resent && !loading && (
                        <p>
                            Didn&apos;t receive the code?{' '}
                            <span
                                onClick={() => setResent(true)}
                                className="cursor-pointer text-[#34C759] hover:text-[#38b558] hover:underline"
                            >
                                Resend Code
                            </span>
                        </p>
                    )}
                    {resent && !loading && (
                        <ResendCountDown
                            duration={5}
                            trigger={resent}
                            onComplete={() => setResent(false)}
                        />
                    )}
                    <div className="flex flex-col gap-y-2">
                        <Button
                            variant={'outline'}
                            disabled={readOnly}
                            onClick={(e) => {
                                e.preventDefault()
                            }}
                            className="border-[#34C759]"
                        >
                            Change {verifyMode === 'email' ? 'Email' : 'Mobile'}
                        </Button>
                        <Button
                            type="submit"
                            disabled={readOnly}
                            className="bg-[#34C759] hover:bg-[#38b558]"
                        >
                            {loading ? <LoadingCircle /> : 'Submit'}
                        </Button>
                    </div>
                </fieldset>
            </form>
        </Form>
    )
}

export default VerifyForm
