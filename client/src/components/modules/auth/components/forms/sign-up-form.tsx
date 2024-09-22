import z from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { BsPatchCheck } from 'react-icons/bs'

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

import { cn } from '@/lib/utils'
import { PASSWORD_MIN_LENGTH, LETTERS_REGEX } from '@/modules/auth/constants'

const signUpFormSchema = z
    .object({
        email: z
            .string({ required_error: 'Email is required' })
            .email('Email must be valid'),
        username: z
            .string({ required_error: 'Username is required' })
            .min(1, 'User Name is required'),
        first_name: z
            .string({ required_error: 'First Name is required' })
            .min(1, 'First Name too short')
            .regex(LETTERS_REGEX, 'First Name must contain only letters'),
        middle_name: z
            .string({ required_error: 'Middle Name is required' })
            .min(1, 'Middle Name is required')
            .regex(LETTERS_REGEX, 'First Name must contain only letters'),
        last_name: z
            .string({ required_error: 'Last Name is required' })
            .min(1, 'Last Name is required')
            .regex(LETTERS_REGEX, 'Last Name must contain only letters'),
        contact_number: z
            .string()
            .min(1)
            .max(11)
            .regex(/^\d+$/, 'Contact number must contain only numbers'),
        password: z
            .string({ required_error: 'Password is required' })
            .min(
                PASSWORD_MIN_LENGTH,
                `Password must atleast ${PASSWORD_MIN_LENGTH}`
            ),
        confirm_password: z
            .string({ required_error: 'Confirm password' })
            .min(
                PASSWORD_MIN_LENGTH,
                `Password must atleast ${PASSWORD_MIN_LENGTH} characters`
            ),
        accept_terms: z
            .boolean()
            .default(false)
            .refine(
                (val) => {
                    return val === true
                },
                {
                    message: 'You must accept the terms and conditions',
                }
            ),
    })
    .refine(({ password, confirm_password }) => password === confirm_password, {
        message: "Passwords don't match",
        path: ['confirm_password'],
    })

type TSignUpForm = z.infer<typeof signUpFormSchema>

interface Props {
    className?: string
}

const SignUpForm = ({ className }: Props) => {
    const form = useForm<TSignUpForm>({
        resolver: zodResolver(signUpFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues: {
            accept_terms: false,
            confirm_password: '',
            password: '',
            contact_number: '',
            email: '',
            first_name: '',
            last_name: '',
            middle_name: '',
            username: '',
        },
    })

    function onSubmit(values: TSignUpForm) {
        console.log(values)
        // TODO: add functionality sign up
    }

    const firstError = Object.values(form.formState.errors)[0]?.message

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit(onSubmit)}
                className={cn('flex w-[390px] flex-col gap-y-4', className)}
            >
                <div className="flex items-center justify-center gap-x-2 py-4 font-medium">
                    <img src="/e-coop-logo-1.png" className="size-24" />
                    <p className="text-xl">Create your profile</p>
                </div>

                <div className="space-y-4">
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
                        name="first_name"
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
                        name="middle_name"
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
                        name="last_name"
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
                        name="contact_number"
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
                        name="confirm_password"
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
                        name="accept_terms"
                        render={({ field }) => (
                            <FormItem className="min-w-[277px]">
                                <div className="flex items-center gap-x-2">
                                    <FormControl>
                                        <Checkbox
                                            checked={field.value}
                                            onCheckedChange={field.onChange}
                                        />
                                    </FormControl>
                                    <FormLabel className="m-0 p-0 font-medium">
                                        Accept terms and condition
                                    </FormLabel>
                                </div>
                            </FormItem>
                        )}
                    />
                </div>

                {firstError && (
                    <span className="mt-2 rounded-md bg-destructive/10 py-2 text-center text-sm text-destructive">
                        {firstError}
                    </span>
                )}
                <Button
                    type="submit"
                    disabled={
                        firstError !== undefined ||
                        !form.formState.isDirty ||
                        !form.formState.isValid
                    }
                    className="mt-6 bg-[#34C759] hover:bg-[#38b558]"
                >
                    Submit
                </Button>
            </form>
        </Form>
    )
}

export default SignUpForm
