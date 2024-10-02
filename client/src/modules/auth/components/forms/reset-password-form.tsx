import z from 'zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { AiOutlineKey } from 'react-icons/ai'

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
                className={cn(
                    'flex w-full flex-col gap-y-4 sm:w-[390px]',
                    className
                )}
            >
                <div className="flex flex-col items-center gap-y-4 py-4 text-center">
                    <div className="relative p-8">
                        <AiOutlineKey className="size-[53px] text-green-500" />
                        <div className="absolute inset-0 rounded-full bg-green-500/20"></div>
                        <div className="absolute inset-5 rounded-full bg-green-500/20"></div>
                    </div>
                    <p className="text-xl font-medium">Set new password</p>
                    <p className="px-10 text-sm text-foreground/70">
                        Set a new password for your account, make sure to use a
                        strong password.
                    </p>
                </div>
                <fieldset disabled={loading || readOnly} className="space-y-4">
                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <FormLabel className="font-medium">
                                    Password
                                </FormLabel>
                                <FormControl>
                                    <PasswordInput
                                        {...field}
                                        placeholder="Password"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="confirmPassword"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <FormLabel className="font-medium">
                                    Confirm Password
                                </FormLabel>
                                <FormControl>
                                    <PasswordInput
                                        {...field}
                                        placeholder="Confirm Password"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                </fieldset>

                <div className="mt-4 flex flex-col space-y-2">
                    <FormErrorMessage errorMessage={firstError} />
                    <Button type="submit" disabled={loading || readOnly}>
                        {loading ? <LoadingCircle /> : 'Save Password'}
                    </Button>
                </div>
            </form>
        </Form>
    )
}

export default ResetPasswordForm
