import z from 'zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'
import { Button } from '@/components/ui/button'
import PasswordInput from '@/components/password-input'
import LoadingCircle from '@/components/loader/loading-circle'

import { cn } from '@/lib/utils'
import { IAuthForm } from '@/types/auth/form-interface'
import { PASSWORD_MIN_LENGTH } from '@/modules/auth/constants'
import FormErrorMessage from '../form-error-message'

const ResetPasswordFormSchema = z
    .object({
        password: z
            .string({ required_error: 'Password is required' })
            .min(
                PASSWORD_MIN_LENGTH,
                `Password must atleast ${PASSWORD_MIN_LENGTH} characters`
            ),
        confirmPassword: z
            .string({ required_error: 'Confirm password' })
            .min(PASSWORD_MIN_LENGTH, `Password doesn't match`),
    })
    .refine(({ password, confirmPassword }) => password === confirmPassword, {
        message: "Password doesn't match",
        path: ['confirm_password'],
    })

type TResetPasswordForm = z.infer<typeof ResetPasswordFormSchema>

interface Props extends IAuthForm<TResetPasswordForm> {}

const ResetPasswordForm = ({
    readOnly,
    className,
    defaultValues = { password: '', confirmPassword: '' },
    onSuccess,
}: Props) => {
    const [loading, setLoading] = useState(false)

    const form = useForm<TResetPasswordForm>({
        resolver: zodResolver(ResetPasswordFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues,
    })

    function onFormSubmit(data: TResetPasswordForm) {
        setLoading(true)
        const parsedData = ResetPasswordFormSchema.parse(data)
        // TODO: Logic to create a reset entry and will return
        // authService.resetViaEmail(email, accountType) // return uuid string
        // modify code bellow
        setInterval(() => {
            setLoading(false)
            onSuccess?.('')
        }, 1000)
    }

    const firstError = Object.values(form.formState.errors)[0]?.message

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit(onFormSubmit)}
                className={cn('flex w-[390px] flex-col gap-y-4', className)}
            >
                <div className="flex flex-col items-center gap-y-4 py-4 text-center">
                    <img src="/e-coop-logo-1.png" className="size-24" />
                    <p className="text-xl font-medium">Forgot Password?</p>
                    <p className="text-sm text-foreground/70">
                        Enter your registered email address to receive a link to
                        reset your password.
                    </p>
                </div>
                <fieldset disabled={loading || readOnly} className="space-y-4">
                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            placeholder="Password"
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="confirmPassword"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex w-full items-center justify-end gap-x-4">
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        Confirm Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            placeholder="Confirm Password"
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                </fieldset>

                <div className="mt-4 flex flex-col space-y-2">
                    <FormErrorMessage errorMessage={firstError} />
                    <Button
                        type="submit"
                        disabled={loading || readOnly}
                        className="bg-[#34C759] hover:bg-[#38b558]"
                    >
                        {loading ? <LoadingCircle /> : 'Save Password'}
                    </Button>
                </div>
            </form>
        </Form>
    )
}

export default ResetPasswordForm
