import z from 'zod'
import { toast } from 'sonner'
import { useCallback } from 'react'
import { useForm } from 'react-hook-form'
import { useMutation } from '@tanstack/react-query'

import { CheckIcon, CloseIcon } from '@/components/icons'
import {
    Form,
    FormItem,
    FormLabel,
    FormField,
    FormControl,
    FormDescription,
} from '@/components/ui/form'
import { Button } from '@/components/ui/button'
import PasswordInput from '@/components/ui/password-input'
import FormErrorMessage from '@/components/ui/form-error-message'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { withCatchAsync } from '@/lib'
import { serverRequestErrExtractor } from '@/helpers'
import { passwordSchema } from '@/validations/common'
import { zodResolver } from '@hookform/resolvers/zod'
import { NewPasswordRequest } from '@/horizon-corp/types'
import ProfileService from '@/horizon-corp/server/auth/ProfileService'

interface Props {
    onSave?: () => void
}

const contactOptionSchema = z
    .object({
        PreviousPassword: passwordSchema,
        NewPassword: passwordSchema,
        ConfirmPassword: passwordSchema,
    })
    .refine(
        ({ NewPassword, ConfirmPassword }) => NewPassword === ConfirmPassword,
        {
            message: "Password doesn't match",
            path: ['confirm_password'],
        }
    )

const PasswordOption = ({ onSave }: Props) => {
    const form = useForm<z.infer<typeof contactOptionSchema>>({
        resolver: zodResolver(contactOptionSchema),
        reValidateMode : "onChange",
        defaultValues: {
            ConfirmPassword: '',
            NewPassword: '',
            PreviousPassword: '',
        },
    })

    const { isPending, mutate: updatePassword } = useMutation<
        void,
        string,
        NewPasswordRequest
    >({
        mutationKey: ['account-security-contact'],
        mutationFn: async (data) => {
            const [error] = await withCatchAsync(
                ProfileService.NewPassword(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            toast.success('Contact number has been saved.')

            onSave?.()
        },
    })

    const hasChanges = form.formState.isDirty

    const firstError = Object.values(form.formState.errors)[0]?.message

    const handleReset = useCallback(() => {
        form.reset()
    }, [form])

    return (
        <>
            <Form {...form}>
                <form
                    className="w-full space-y-2"
                    onSubmit={form.handleSubmit((data) => updatePassword(data))}
                >
                    <div className="space-y-2">
                        <FormLabel>Password</FormLabel>
                        <FormDescription>
                            Update or change your password. Your password is at
                            least 8 characters
                        </FormDescription>
                    </div>
                    <fieldset
                        disabled={isPending}
                        className="grid grid-cols-3 gap-x-2 pt-4"
                    >
                        <FormField
                            control={form.control}
                            name="PreviousPassword"
                            render={({ field }) => (
                                <FormItem className="grid w-full gap-y-2">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="text-sm font-normal text-foreground/80"
                                    >
                                        Old Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            id={field.name}
                                            placeholder="Password"
                                            autoComplete="new-password"
                                        />
                                    </FormControl>
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="NewPassword"
                            render={({ field }) => (
                                <FormItem className="grid w-full gap-y-2">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="text-sm font-normal text-foreground/80"
                                    >
                                        New Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            id={field.name}
                                            placeholder="New Password"
                                            autoComplete="new-password"
                                        />
                                    </FormControl>
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="ConfirmPassword"
                            render={({ field }) => (
                                <FormItem className="grid w-full gap-y-2">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="text-sm font-normal text-foreground/80"
                                    >
                                        Confirm Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            id={field.name}
                                            placeholder="Confirm Password"
                                            autoComplete="new-password"
                                        />
                                    </FormControl>
                                </FormItem>
                            )}
                        />
                    </fieldset>
                    <FormErrorMessage errorMessage={firstError} />
                    <div className="flex items-center justify-end gap-x-1">
                        {hasChanges && !isPending && (
                            <>
                                <Button
                                    size="sm"
                                    type="reset"
                                    variant="secondary"
                                    onClick={(e) => {
                                        e.preventDefault()
                                        handleReset()
                                    }}
                                >
                                    <CloseIcon />
                                </Button>
                                <Button type="submit" size="sm">
                                    <CheckIcon />
                                </Button>
                            </>
                        )}
                        {isPending && (
                            <span className="text-xs text-foreground/70">
                                <LoadingSpinner className="mr-1 inline-block size-3" />{' '}
                                Saving...
                            </span>
                        )}
                    </div>
                </form>
            </Form>
        </>
    )
}

export default PasswordOption
