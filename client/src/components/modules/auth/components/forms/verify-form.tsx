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
import { REGEXP_ONLY_DIGITS_AND_CHARS } from 'input-otp'
import { verifyFormSchema } from '../../validations/verify-form'
import { IAuthForm } from '@/interfaces/components/form-interface'

type TVerifyForm = z.infer<typeof verifyFormSchema>

interface Props extends IAuthForm<TVerifyForm> {
    id: string
    verifyMode: 'mobile' | 'email'
}

const VerifyForm = ({
    id,
    className,
    readOnly = false,
    verifyMode = 'mobile',
    defaultValues = { code: '' },
}: Props) => {
    const [loading, _setLoading] = useState(false)

    const form = useForm({
        resolver: zodResolver(verifyFormSchema),
        reValidateMode: 'onChange',
        defaultValues,
    })

    const handleSubmit = (data: TVerifyForm) => {
        const parsedData = verifyFormSchema.parse(data)
        console.log(parsedData)
        // TODO: Add functionality
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
                                        containerClassName="mx-auto capitalize w-fit"
                                        maxLength={6}
                                        pattern={REGEXP_ONLY_DIGITS_AND_CHARS}
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
                    <p>
                        Didn&apos;t receive the code?{' '}
                        <span className="cursor-pointer text-[#34C759] hover:text-[#38b558] hover:underline">
                            Resend Code
                        </span>
                    </p>
                    <Button
                        type="submit"
                        disabled={readOnly}
                        className="mt-6 bg-[#34C759] hover:bg-[#38b558]"
                    >
                        {loading ? <LoadingCircle /> : 'Submit'}
                    </Button>
                </fieldset>
            </form>
        </Form>
    )
}

export default VerifyForm
