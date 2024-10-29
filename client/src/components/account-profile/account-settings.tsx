import z from 'zod'
import { toast } from 'sonner'
import { format } from 'date-fns'
import { useForm } from 'react-hook-form'
import { useMutation } from '@tanstack/react-query'
import { zodResolver } from '@hookform/resolvers/zod'
import { useRouterState } from '@tanstack/react-router'
import { useCallback, useEffect, useState } from 'react'
import { useUserAuthStore } from '@/store/user-auth-store'

import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Textarea } from '@/components/ui/textarea'
import { Separator } from '@/components/ui/separator'
import ActionTooltip from '@/components/action-tooltip'
import InputDatePicker from '@/components/input-date-picker'
import { CloseIcon, EditPencilIcon } from '@/components/icons'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import FormErrorMessage from '@/modules/auth/components/form-error-message'

import {
    birthDateSchema,
    firstNameSchema,
    lastNameSchema,
    middleNameSchema,
} from '@/validations/common'
import { cn, withCatchAsync } from '@/lib'
import { serverRequestErrExtractor } from '@/helpers'
import useConfirmModalStore from '@/store/confirm-modal-store'
import { AccountSettingRequest, UserData } from '@/horizon-corp/types'
import ProfileService from '@/horizon-corp/server/auth/ProfileService'

const AccountSettingsFormSchema = z.object({
    lastName: lastNameSchema,
    birthDate: birthDateSchema,
    firstName: firstNameSchema,
    middleName: middleNameSchema,
    description: z.string().optional(),
    permanentAddress: z
        .string({ required_error: 'Address is required' })
        .min(1, 'Address is required'),
})

type TAccountSettings = z.infer<typeof AccountSettingsFormSchema>

const AccountSettings = () => {
    const { onOpen } = useConfirmModalStore()
    const [isEditing, setIsEditing] = useState(false)
    const { currentUser, setCurrentUser } = useUserAuthStore()

    const hash = useRouterState({
        select: ({ location }) => location.hash,
    })

    const form = useForm<TAccountSettings>({
        resolver: zodResolver(AccountSettingsFormSchema),
        defaultValues: {
            birthDate: new Date('10-10-1900'),
            description: '',
            firstName: '',
            lastName: '',
            middleName: '',
            permanentAddress: '',
        },
    })

    const handleReset = useCallback(() => {
        form.reset()
        if (!currentUser) return
        form.setValue('lastName', currentUser.lastName)
        form.setValue('firstName', currentUser.firstName)
        form.setValue('middleName', currentUser.middleName)
        form.setValue('description', currentUser.description)
        form.setValue('birthDate', currentUser.birthDate)
        form.setValue('permanentAddress', currentUser.permanentAddress)
    }, [form, currentUser])

    useEffect(() => {
        handleReset()
    }, [handleReset])

    const {
        error,
        isPending,
        mutate: updateAccountSettings,
    } = useMutation<UserData, string, AccountSettingRequest>({
        mutationKey: ['account-settings'],
        mutationFn: async (data) => {
            const [error, response] = await withCatchAsync(
                ProfileService.AccountSetting(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            setIsEditing(false)
            setCurrentUser(response.data)
            toast.success('Changes saved')
            return response.data
        },
    })

    if (hash !== 'account-settings') return null

    const firstError = Object.values(form.formState.errors)[0]?.message

    return (
        <div className="">
            <div className="flex items-center justify-between">
                <p className="text-lg font-medium">Account Settings</p>
                <ActionTooltip tooltipContent={isEditing ? 'Close' : 'Edit'}>
                    <Button
                        size="icon"
                        className="size-fit rounded-full p-1"
                        onClick={() =>
                            isEditing
                                ? onOpen({
                                      title: 'Discard Changes',
                                      description:
                                          'Are you sure to discard changes? Any changes made will be discarded.',
                                      onConfirm: () => {
                                          setIsEditing(false)
                                          handleReset()
                                      },
                                  })
                                : setIsEditing((val) => !val)
                        }
                        variant={isEditing ? 'outline' : 'ghost'}
                    >
                        {isEditing ? <CloseIcon /> : <EditPencilIcon />}
                    </Button>
                </ActionTooltip>
            </div>
            <p className="mt-1 text-xs text-foreground/60">
                Update your account settings
            </p>
            <Separator className="my-4" />
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit((data) =>
                        updateAccountSettings(data)
                    )}
                    className={cn('flex w-full flex-col gap-y-4')}
                >
                    <fieldset
                        disabled={isPending}
                        className="grid gap-x-4 gap-y-6 sm:grid-cols-2"
                    >
                        <FormField
                            control={form.control}
                            name="firstName"
                            render={({ field }) => (
                                <FormItem className="space-y-0">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="text-right text-sm font-normal text-foreground/60"
                                    >
                                        Firstname
                                    </FormLabel>
                                    {isEditing ? (
                                        <FormControl>
                                            <Input
                                                {...field}
                                                id={field.name}
                                                placeholder="Firstname"
                                                autoComplete="off"
                                            />
                                        </FormControl>
                                    ) : (
                                        <p className="py-1">{field.value}</p>
                                    )}
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="middleName"
                            render={({ field }) => (
                                <FormItem className="space-y-0">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="text-right font-normal text-foreground/60"
                                    >
                                        Middlename
                                    </FormLabel>
                                    {isEditing ? (
                                        <FormControl>
                                            <Input
                                                {...field}
                                                id={field.name}
                                                placeholder="Middlename"
                                                autoComplete="off"
                                            />
                                        </FormControl>
                                    ) : (
                                        <p className="py-1">{field.value}</p>
                                    )}
                                </FormItem>
                            )}
                        />{' '}
                        <FormField
                            control={form.control}
                            name="lastName"
                            render={({ field }) => (
                                <FormItem className="space-y-0">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="text-right font-normal text-foreground/60"
                                    >
                                        Lastname
                                    </FormLabel>
                                    {isEditing ? (
                                        <FormControl>
                                            <Input
                                                {...field}
                                                id={field.name}
                                                placeholder="Lastname"
                                                autoComplete="off"
                                            />
                                        </FormControl>
                                    ) : (
                                        <p className="pt-1">{field.value}</p>
                                    )}
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="birthDate"
                            render={({ field }) => (
                                <FormItem className="space-y-0">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="text-right text-sm font-normal text-foreground/60"
                                    >
                                        Birthdate
                                    </FormLabel>
                                    {isEditing ? (
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
                                    ) : (
                                        <p className="pt-1">
                                            {format(field.value, 'MMM dd yyyy')}
                                        </p>
                                    )}
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="permanentAddress"
                            render={({ field }) => (
                                <FormItem className="space-y-0 sm:col-span-2">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="text-right text-sm font-normal text-foreground/60"
                                    >
                                        Permanent Address
                                    </FormLabel>
                                    {isEditing ? (
                                        <FormControl>
                                            <Input
                                                {...field}
                                                id={field.name}
                                                placeholder="Permanent Address"
                                                autoComplete="off"
                                            />
                                        </FormControl>
                                    ) : (
                                        <p className="pt-1">{field.value}</p>
                                    )}
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="description"
                            render={({ field }) => (
                                <FormItem className="space-y-0 sm:col-span-2">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="text-right text-sm font-normal text-foreground/60"
                                    >
                                        Description
                                    </FormLabel>
                                    {isEditing ? (
                                        <FormControl>
                                            <Textarea
                                                placeholder="Write a short description about you"
                                                className="resize-none"
                                                {...field}
                                            />
                                        </FormControl>
                                    ) : (
                                        <p className="pt-1">{field.value}</p>
                                    )}
                                </FormItem>
                            )}
                        />
                    </fieldset>
                    <FormErrorMessage errorMessage={firstError || error} />
                    {isEditing && (
                        <>
                            <Separator className="my-2 sm:my-4" />
                            <Button
                                type="submit"
                                disabled={isPending}
                                className="w-full self-end px-8 sm:w-fit"
                            >
                                {isPending ? <LoadingSpinner /> : 'Save'}
                            </Button>
                        </>
                    )}
                </form>
            </Form>
        </div>
    )
}

export default AccountSettings
