import z from 'zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Link, useRouter } from '@tanstack/react-router'

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
import PasswordInput from '@/components/password-input'
import LoadingCircle from '@/components/loader/loading-circle'
import FormErrorMessage from '@/modules/auth/components/form-error-message'

import { cn } from '@/lib/utils'
import { UserStatus } from '@/types'
import { IAuthForm } from '@/types/auth/form-interface'
import { IBaseCompNoChild } from '@/types/component/base'
import { signInFormSchema } from '@/modules/auth/validations/sign-in-form'

type TSignIn = z.infer<typeof signInFormSchema>

interface Props extends IBaseCompNoChild, IAuthForm<Partial<TSignIn>> {}

const SignInForm = ({
    defaultValues,
    className,
    readOnly,
    onSuccess,
}: Props) => {
    const [loading, setLoading] = useState(false)

    const router = useRouter()

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

    function onFormSubmit(data: TSignIn) {
        const parsedData = signInFormSchema.parse(data)
        // TODO: Logic
        onSuccess?.({
            id: '215',
            username: 'Jervx',
            validEmail: false,
            validContactNumber: true,
            status: UserStatus['Pending'],
            profilePicture: {
                url: 'https://mrwallpaper.com/images/hd/suit-rick-and-morty-phone-5divv4gzo6gowk46.jpg',
            },
        } as any)
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
                    <p className="text-xl">Login to your account</p>
                </div>

                <fieldset disabled={loading || readOnly} className="space-y-4">
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
                </fieldset>
                <div className="mt-6 flex flex-col space-y-2">
                    <FormErrorMessage errorMessage={firstError} />
                    <Button
                        type="submit"
                        disabled={loading || readOnly}
                        className="bg-[#34C759] hover:bg-[#38b558]"
                    >
                        {loading ? <LoadingCircle /> : 'Login'}
                    </Button>
                    <Link
                        className="text-sm text-green-500"
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
