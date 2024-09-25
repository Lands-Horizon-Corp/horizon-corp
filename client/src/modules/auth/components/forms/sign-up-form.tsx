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
import { Button } from '@/components/ui/button'
import { Checkbox } from '@/components/ui/checkbox'
import LoadingCircle from '@/components/loader/loading-circle'

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
    onError,
    onLoading,
    onSuccess,
    defaultValues = defaultValue,
}: Props) => {
    const [loading, setLoading] = useState(false)
    const [success, setSuccess] = useState(false)

    const form = useForm<TSignUpForm>({
        resolver: zodResolver(signUpFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues,
    })

    function onFormSubmit(data: TSignUpForm) {
        const parsedData = signUpFormSchema.parse(data)
        console.log(parsedData)
        // TODO: Logic
        // remove code bellow
    }

    const firstError = Object.values(form.formState.errors)[0]?.message

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit(onFormSubmit)}
                className={cn('flex w-[390px] flex-col gap-y-4', className)}
            >
                <div className="flex items-center justify-center gap-x-2 py-4 font-medium">
                    <img src="/e-coop-logo-1.png" className="size-24" />
                    <p className="text-xl">Create your profile</p>
                </div>

                <fieldset disabled={loading} className="space-y-4">
                    <FormField
                        name="email"
                        control={form.control}
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center justify-end gap-x-4">
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        Email
                                    </FormLabel>
                                    <FormControl>
                                        <div className="flex-1 space-y-2">
                                            <Input
                                                placeholder="Email"
                                                {...field}
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
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        Username
                                    </FormLabel>
                                    <FormControl>
                                        <Input
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
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        First Name
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            placeholder="First Name"
                                            {...field}
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
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        Middle Name
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            className="w-auto flex-1"
                                            placeholder="First Name"
                                            {...field}
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
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        Last Name
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            placeholder="First Name"
                                            {...field}
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
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        Contact Number
                                    </FormLabel>
                                    <FormControl>
                                        <div className="flex flex-1 items-center gap-x-2">
                                            <Input
                                                placeholder="Contact Number"
                                                {...field}
                                            />
                                            <BsPatchCheck className="size-8 text-green-500" />
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
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        Password
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            type="password"
                                            placeholder="Password"
                                            {...field}
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
                                        <Input
                                            placeholder="Confirm Password"
                                            {...field}
                                            type="password"
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
                                    <FormLabel className="w-full max-w-[90px] text-right font-medium">
                                        Account Type
                                    </FormLabel>
                                    <Select
                                        onValueChange={field.onChange}
                                        defaultValue={field.value}
                                    >
                                        <FormControl>
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
                                            checked={field.value}
                                            onCheckedChange={field.onChange}
                                            className="border-[#34c759] hover:text-[#38b558] data-[state=checked]:bg-[#34C759]"
                                        />
                                    </FormControl>
                                    <FormLabel className="m-0 cursor-pointer p-0 font-medium">
                                        Accept terms and condition
                                    </FormLabel>
                                </div>
                            </FormItem>
                        )}
                    />
                </fieldset>

                {firstError && (
                    <span className="mt-2 rounded-md bg-destructive/10 py-2 text-center text-sm text-destructive">
                        {firstError}
                    </span>
                )}
                <Button
                    type="submit"
                    disabled={firstError !== undefined || readOnly}
                    className="mt-6 bg-[#34C759] hover:bg-[#38b558]"
                >
                    {loading ? <LoadingCircle /> : 'Submit'}
                </Button>
            </form>
        </Form>
    )
}

export default SignUpForm
