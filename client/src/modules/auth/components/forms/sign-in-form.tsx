import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Link } from '@tanstack/react-router'

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
import LoadingCircle from '@/components/loader/loading-circle'
import FormErrorMessage from '@/modules/auth/components/form-error-message'

import { cn } from '@/lib/utils'
import { IAuthForm } from '@/types/auth/form-interface'
import { IBaseCompNoChild } from '@/types/component/base'
import { handleAxiosError } from '@/horizon-corp/helpers'
import UserService from '@/horizon-corp/server/auth/UserService'
import useLoadingErrorState from '@/hooks/use-loading-error-state'
import { signInFormSchema } from '@/modules/auth/validations/sign-in-form'

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
            email: defaultValues?.email ?? '',
            username: defaultValues?.username ?? '',
            mode: defaultValues?.mode ?? 'Member',
            password: '',
        },
    })

    const onFormSubmit = async (data: TSignIn) => {
        setError(null)
        setLoading(true)
        try {
            const parsedData = await signInFormSchema.parse(data) // parse form data
            // const response = await UserService.signIn(parsedData) // sign in and server return the user data along with auth cookie
            // onSuccess?.(response.data) // if onSuccess is given, trigger it and pass the data
        } catch (e) {
            const errorMessage = handleAxiosError(e)
            onError?.(e)
            setError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setLoading(false)
        }
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
                        name="email"
                        control={form.control}
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Email
                                    </FormLabel>
                                    <FormControl>
                                        <div className="flex-1 space-y-2">
                                            <Input
                                                {...field}
                                                id={field.name}
                                                autoComplete="on"
                                                placeholder="Email"
                                            />
                                        </div>
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="username"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Username
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            placeholder="Username"
                                            autoComplete="username"
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
                                            autoComplete="current-password"
                                            placeholder="Password"
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="mode"
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
                    <Button
                        type="submit"
                        disabled={loading || error !== null || readOnly}
                    >
                        {loading ? <LoadingCircle /> : 'Login'}
                    </Button>
                    <Link
                        className="text-sm text-primary"
                        to="/auth/forgot-password"
                        search={{
                            email: form.getValues('email'),
                            mode: form.getValues('mode'),
                        }}
                    >
                        Forgor Password
                    </Link>
                </div>
            </form>
        </Form>
    )
}

export default SignInForm
