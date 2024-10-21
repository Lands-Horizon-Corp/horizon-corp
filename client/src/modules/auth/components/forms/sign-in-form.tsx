// dependencies
import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { Link } from '@tanstack/react-router'
import { zodResolver } from '@hookform/resolvers/zod'

// components
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import EcoopLogo from '@/components/ecoop-logo'
import { Button } from '@/components/ui/button'
import PasswordInput from '@/components/ui/password-input'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import FormErrorMessage from '@/modules/auth/components/form-error-message'

// functions/hooks
import { cn, withCatchAsync } from '@/lib/utils'
import { serverRequestErrExtractor } from '@/helpers'
import UserService from '@/horizon-corp/server/auth/UserService'
import useLoadingErrorState from '@/hooks/use-loading-error-state'
import { signInFormSchema } from '@/modules/auth/validations/sign-in-form'

// types/interfaces
import { IBaseCompNoChild } from '@/types/component'
import { IAuthForm } from '@/types/auth/form-interface'

type TSignIn = z.infer<typeof signInFormSchema>

interface Props extends IBaseCompNoChild, IAuthForm<Partial<TSignIn>> {}

const SignInForm = ({
    defaultValues,
    className,
    readOnly,
    onSuccess,
    onError,
}: Props) => {
    const { loading, error, setError, setLoading } = useLoadingErrorState()

    const form = useForm<TSignIn>({
        resolver: zodResolver(signInFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues: {
            key: defaultValues?.key ?? '',
            password: '',
            accountType: defaultValues?.accountType ?? 'Member',
        },
    })

    const onFormSubmit = async (data: TSignIn) => {
        setError(null)
        setLoading(true)

        const [error, response] = await withCatchAsync(UserService.SignIn(data))

        setLoading(false)

        if (error) {
            const errorMessage = serverRequestErrExtractor({ error })
            onError?.(error)
            setError(errorMessage)
            toast.error(errorMessage)
            return
        }

        onSuccess?.(response.data)
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
                <div className="flex items-center justify-center gap-x-2 py-4 font-medium">
                    <EcoopLogo className="size-24" />
                    <p className="text-xl">Login to your account</p>
                </div>

                <fieldset disabled={loading || readOnly} className="space-y-4">
                    <FormField
                        control={form.control}
                        name="key"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Key
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            placeholder="Email, contact, or username"
                                            autoComplete="off"
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor="password-field"
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            id="password-field"
                                            autoComplete="off"
                                            placeholder="Password"
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="accountType"
                        render={({ field }) => (
                            <FormItem>
                                <div className="flex w-full items-center justify-end gap-x-4">
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
                                        <FormControl id="account-type">
                                            <SelectTrigger>
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
                                </div>
                            </FormItem>
                        )}
                    />
                </fieldset>
                <div className="mt-6 flex flex-col space-y-2">
                    <FormErrorMessage errorMessage={firstError || error} />
                    <Button type="submit" disabled={loading || readOnly}>
                        {loading ? <LoadingSpinner /> : 'Login'}
                    </Button>
                    <Link
                        className="text-sm text-primary"
                        to="/auth/forgot-password"
                        search={{
                            accountType: form.getValues('accountType'),
                        }}
                    >
                        Forgot Password
                    </Link>
                </div>
            </form>
        </Form>
    )
}

export default SignInForm
