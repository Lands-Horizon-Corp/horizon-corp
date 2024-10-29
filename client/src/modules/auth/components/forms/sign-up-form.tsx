import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { useRouter } from '@tanstack/react-router'
import { zodResolver } from '@hookform/resolvers/zod'

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
import { Checkbox } from '@/components/ui/checkbox'
import { VerifiedPatchIcon } from '@/components/icons'
import PasswordInput from '@/components/ui/password-input'
import InputDatePicker from '@/components/input-date-picker'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import FormErrorMessage from '@/components/ui/form-error-message'

import { cn, withCatchAsync } from '@/lib/utils'
import { serverRequestErrExtractor } from '@/helpers'
import UserService from '@/horizon-corp/server/auth/UserService'
import useLoadingErrorState from '@/hooks/use-loading-error-state'
import { signUpFormSchema } from '@/modules/auth/validations/sign-up-form'

import { IAuthForm } from '@/types/auth/form-interface'
import { useUserAuthStore } from '@/store/user-auth-store'

type TSignUpForm = z.infer<typeof signUpFormSchema>

const defaultValue: TSignUpForm = {
    acceptTerms: false,
    confirmPassword: '',
    password: '',
    contactNumber: '',
    email: '',
    permanentAddress: '',
    birthdate: new Date(),
    firstName: '',
    lastName: '',
    middleName: '',
    username: '',
    accountType: 'Member',
}

const SignUpForm = ({
    className,
    readOnly,
    defaultValues = defaultValue,
    onError,
    onSuccess,
}: IAuthForm<TSignUpForm>) => {
    const router = useRouter()
    const { setCurrentUser } = useUserAuthStore()
    const { loading, setLoading, error, setError } = useLoadingErrorState()

    const form = useForm<TSignUpForm>({
        resolver: zodResolver(signUpFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues,
    })

    const onFormSubmit = async (data: TSignUpForm) => {
        setError(null)
        setLoading(true)

        const [error, response] = await withCatchAsync(UserService.SignUp(data))

        setLoading(false)

        if (error) {
            const errorMessage = serverRequestErrExtractor({ error })
            onError?.(error)
            setError(errorMessage)
            toast.error(errorMessage)
            return
        }

        setCurrentUser(response.data)
        onSuccess?.(response.data)
        router.navigate({ to: '/auth/verify' })
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
                    <p className="text-xl">Create your profile</p>
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
                                                autoComplete="email"
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
                                            id={field.name}
                                            autoComplete="username"
                                            placeholder="Username"
                                            {...field}
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="firstName"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        First Name
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            autoComplete="given-name"
                                            placeholder="First Name"
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="middleName"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Middle Name
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            className="w-auto flex-1"
                                            placeholder="Middle Name"
                                            autoComplete="additional-name"
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="lastName"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Last Name
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            placeholder="Last Name"
                                            autoComplete="family-name"
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="birthdate"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Birth Date
                                    </FormLabel>
                                    <FormControl>
                                        <InputDatePicker
                                            id={field.name}
                                            value={field.value}
                                            onChange={field.onChange}
                                            captionLayout="dropdown-buttons"
                                            disabled={(date) =>
                                                date > new Date()
                                            }
                                        />
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="contactNumber"
                        render={({ field, fieldState: { invalid, error } }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Contact Number
                                    </FormLabel>
                                    <FormControl>
                                        <div className="flex flex-1 items-center gap-x-2">
                                            <Input
                                                {...field}
                                                id={field.name}
                                                autoComplete="tel-country-code"
                                                placeholder="Contact Number"
                                            />
                                            <VerifiedPatchIcon
                                                className={cn(
                                                    'size-8 text-primary delay-300 duration-300 ease-in-out',
                                                    (invalid || error) &&
                                                        'text-destructive'
                                                )}
                                            />
                                        </div>
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="permanentAddress"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Permanent Address
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            autoComplete="street-address"
                                            placeholder="Permanent Address"
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
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            id={field.name}
                                            placeholder="Password"
                                            autoComplete="new-password"
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
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full max-w-[90px] text-right font-medium"
                                    >
                                        Confirm Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            id={field.name}
                                            autoComplete="no"
                                            placeholder="Confirm Password"
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
                    <FormField
                        control={form.control}
                        name="acceptTerms"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="mt-8 flex items-center gap-x-2">
                                    <FormControl>
                                        <Checkbox
                                            id={field.name}
                                            name={field.name}
                                            checked={field.value}
                                            onCheckedChange={field.onChange}
                                        />
                                    </FormControl>
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="m-0 cursor-pointer p-0 font-medium"
                                    >
                                        Accept terms and condition
                                    </FormLabel>
                                </div>
                            </FormItem>
                        )}
                    />
                </fieldset>
                <div className="mt-4 flex flex-col space-y-2">
                    <FormErrorMessage errorMessage={firstError || error} />
                    <Button type="submit" disabled={loading || readOnly}>
                        {loading ? <LoadingSpinner /> : 'Submit'}
                    </Button>
                </div>
            </form>
        </Form>
    )
}

export default SignUpForm
