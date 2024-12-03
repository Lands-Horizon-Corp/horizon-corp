import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import { KeyIcon } from '@/components/icons'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import FormErrorMessage from '@/components/ui/form-error-message'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib/utils'
import { withCatchAsync } from "@/utils"
import { serverRequestErrExtractor } from '@/helpers'
import { IAuthForm } from '@/types/auth/form-interface'
import { userAccountTypeSchema } from '@/validations/common'
import UserService from '@/horizon-corp/server/auth/UserService'
import useLoadingErrorState from '@/hooks/use-loading-error-state'

const forgotPasswordFormSchema = z.object({
    key: z.string().min(1, 'key is required'),
    accountType: userAccountTypeSchema,
})

export type TForgotPasswordEmail = z.infer<typeof forgotPasswordFormSchema>

const ForgotPasswordEmail = ({
    readOnly,
    className,
    defaultValues = { key: '', accountType: 'Member' },
    onError,
    onSuccess,
}: IAuthForm<TForgotPasswordEmail, TForgotPasswordEmail>) => {
    const { loading, setLoading, error, setError } = useLoadingErrorState()

    const form = useForm<TForgotPasswordEmail>({
        resolver: zodResolver(forgotPasswordFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues,
    })

    const onFormSubmit = async (data: TForgotPasswordEmail) => {
        setError(null)
        setLoading(true)

        const [error] = await withCatchAsync(UserService.ForgotPassword(data))

        setLoading(false)

        if (error) {
            const errorMessage = serverRequestErrExtractor({ error })
            onError?.(error)
            setError(errorMessage)
            toast.error(errorMessage)
            return
        }

        onSuccess?.(data)
        toast.success(`Password reset link was sent to ${data.key}`)
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
                        <KeyIcon className="size-[53px] text-[#FF7E47]" />
                        <div className="absolute inset-0 rounded-full bg-[#FF7E47]/20" />
                        <div className="absolute inset-5 rounded-full bg-[#FF7E47]/20" />
                    </div>
                    <p className="text-xl font-medium">Forgot Password?</p>
                    <p className="text-sm text-foreground/70">
                        Enter your registered email address to receive a link to
                        reset your password.
                    </p>
                </div>
                <fieldset disabled={loading || readOnly} className="space-y-4">
                    <FormField
                        name="key"
                        control={form.control}
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="w-full text-right font-medium"
                                >
                                    Key
                                </FormLabel>
                                <FormControl>
                                    <div className="flex-1 space-y-2">
                                        <Input
                                            {...field}
                                            id={field.name}
                                            autoComplete="off"
                                            placeholder="Enter email or contact"
                                        />
                                    </div>
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="accountType"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel
                                    htmlFor="account-type"
                                    className="w-full max-w-[90px] text-right font-medium"
                                >
                                    Account Type
                                </FormLabel>
                                <Select
                                    name={field.name}
                                    defaultValue={field.value}
                                    onValueChange={field.onChange}
                                >
                                    <FormControl>
                                        <SelectTrigger id="account-type">
                                            <SelectValue placeholder="Choose type" />
                                        </SelectTrigger>
                                    </FormControl>
                                    <SelectContent>
                                        <SelectItem value="Member">
                                            Member
                                        </SelectItem>
                                        <SelectItem value="Owner">
                                            Owner
                                        </SelectItem>
                                        <SelectItem value="Admin">
                                            Admin
                                        </SelectItem>
                                        <SelectItem value="Employee">
                                            Employee
                                        </SelectItem>
                                    </SelectContent>
                                </Select>
                            </FormItem>
                        )}
                    />
                </fieldset>

                <div className="mt-4 flex flex-col space-y-2">
                    <FormErrorMessage errorMessage={firstError || error} />
                    <Button type="submit" disabled={loading || readOnly}>
                        {loading ? <LoadingSpinner /> : 'Confirm Email'}
                    </Button>
                </div>
            </form>
        </Form>
    )
}

export default ForgotPasswordEmail
