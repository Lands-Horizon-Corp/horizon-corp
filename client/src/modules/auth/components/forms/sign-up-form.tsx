import z from 'zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { BsPatchCheck } from 'react-icons/bs'

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
import PasswordInput from '@/components/ui/password-input'
import LoadingCircle from '@/components/loader/loading-circle'
import FormErrorMessage from '@/modules/auth/components/form-error-message'

import { cn } from '@/lib/utils'
import { IAuthForm } from '@/types/auth/form-interface'
import { signUpFormSchema } from '@/modules/auth/validations/sign-up-form'

type TSignUpForm = z.infer<typeof signUpFormSchema>

const defaultValue: TSignUpForm = {
    acceptTerms: false,
    confirmPassword: '',
    password: '',
    contactNumber: '',
    email: '',
    firstName: '',
    lastName: '',
    middleName: '',
    username: '',
    mode: 'Member',
}

interface Props extends IAuthForm<TSignUpForm> {}

const SignUpForm = ({
    className,
    readOnly,
    defaultValues = defaultValue,
}: Props) => {
    const [loading, setLoading] = useState(false)

    const form = useForm<TSignUpForm>({
        resolver: zodResolver(signUpFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues,
    })

    function onFormSubmit(data: TSignUpForm) {
        const parsedData = signUpFormSchema.parse(data)
        // TODO: Logic
        /*
            after account creation need ko lang sa backend is:

            1- after sign up automatically log in narin (may authorization something cookie?)
            2- after sign up, need ko din nung data ng user

            then: redirect ko na sa auth/verify page
        */
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
                        name="contactNumber"
                        render={({ field }) => (
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
                                            <BsPatchCheck className="size-8 text-primary" />
                                        </div>
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
                    <FormErrorMessage errorMessage={firstError} />
                    <Button type="submit" disabled={loading || readOnly}>
                        {loading ? <LoadingCircle /> : 'Submit'}
                    </Button>
                </div>
            </form>
        </Form>
    )
}

export default SignUpForm
