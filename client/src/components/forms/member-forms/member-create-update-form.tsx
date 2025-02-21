import z from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import {
    Form,
    FormItem,
    FormLabel,
    FormField,
    FormControl,
    FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import {
    ChecklistTemplate,
    ValueChecklistMeter,
} from '@/components/value-checklist-indicator'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import { VerifiedPatchIcon } from '@/components/icons'
import PasswordInput from '@/components/ui/password-input'
import Modal, { IModalProps } from '@/components/modals/modal'
import FormErrorMessage from '@/components/ui/form-error-message'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import { PhoneInput } from '@/components/contact-input/contact-input'
import InputDatePicker from '@/components/date-time-pickers/input-date-picker'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import {
    useCreateMember,
    useUpdateMember,
} from '@/hooks/api-hooks/member/use-member'
import { createMemberSchema } from '@/validations/form-validation/member/member-schema'

type TMemberCreateUpdateForm = z.infer<typeof createMemberSchema>

interface IMemberCreateUpdateFormProps
    extends IBaseCompNoChild,
        IForm<Partial<TMemberCreateUpdateForm>, unknown, string> {}

const MemberCreateUpdateForm = ({
    readOnly,
    className,
    defaultValues,
    onError,
    onSuccess,
}: IMemberCreateUpdateFormProps) => {
    const form = useForm<TMemberCreateUpdateForm>({
        resolver: zodResolver(createMemberSchema),
        reValidateMode: 'onChange',
        mode: 'onSubmit',
        defaultValues: {
            username: '',
            firstName: '',
            lastName: '',
            middleName: '',
            birthDate: new Date(),
            contactNumber: '',
            permanentAddress: '',
            email: '',
            password: '',
            confirmPassword: '',
            ...defaultValues,
        },
    })

    const {
        error: createError,
        isPending: isCreating,
        mutate: createMember,
    } = useCreateMember({ onSuccess, onError })

    const {
        error: updateError,
        isPending: isUpdating,
        mutate: updateMember,
    } = useUpdateMember({ onSuccess, onError })

    const handleSubmit = (data: TMemberCreateUpdateForm) => {
        if (data.id) {
            updateMember({ id: data.id, data })
        } else {
            createMember(data)
        }
    }

    const error = createError || updateError
    const isLoading = isUpdating || isCreating

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit(handleSubmit)}
                className={cn('flex w-full flex-col gap-y-4', className)}
            >
                <fieldset
                    disabled={isLoading || readOnly}
                    className="grid grid-cols-1 gap-x-6 gap-y-4 sm:grid-cols-3 sm:gap-y-3"
                >
                    <fieldset className="space-y-3">
                        <legend>Account Information</legend>
                        <FormField
                            name="username"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
                                    >
                                        Username
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            autoComplete="username"
                                            placeholder="Username"
                                        />
                                    </FormControl>
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                        <FormField
                            name="password"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
                                    >
                                        Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            maxLength={50}
                                            id={field.name}
                                            placeholder="Password"
                                            autoComplete="new-password"
                                        />
                                    </FormControl>
                                    <ValueChecklistMeter
                                        value={field.value}
                                        checkList={ChecklistTemplate[
                                            'password-checklist'
                                        ].concat([
                                            {
                                                regex: /^.{0,50}$/,
                                                text: 'No more than 50 characters',
                                            },
                                        ])}
                                    />
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                        <FormField
                            name="confirmPassword"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
                                    >
                                        Confirm Password
                                    </FormLabel>
                                    <FormControl>
                                        <PasswordInput
                                            {...field}
                                            id={field.name}
                                            autoComplete="off"
                                            placeholder="Confirm Password"
                                        />
                                    </FormControl>
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                    </fieldset>

                    <fieldset className="space-y-3">
                        <legend>Personal Information</legend>
                        <FormField
                            control={form.control}
                            name="firstName"
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
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
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="middleName"
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
                                    >
                                        Middle Name
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            placeholder="Middle Name"
                                            autoComplete="additional-name"
                                        />
                                    </FormControl>
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="lastName"
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
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
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="birthDate"
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
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
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                    </fieldset>

                    <fieldset className="space-y-3">
                        <legend>Contact Information</legend>
                        <FormField
                            name="email"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
                                    >
                                        Email
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            autoComplete="email"
                                            placeholder="Email"
                                        />
                                    </FormControl>
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                        <FormField
                            name="contactNumber"
                            control={form.control}
                            render={({
                                field,
                                fieldState: { invalid, error },
                            }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
                                    >
                                        Contact Number
                                    </FormLabel>
                                    <FormControl>
                                        <div className="relative flex flex-1 items-center gap-x-2">
                                            <VerifiedPatchIcon
                                                className={cn(
                                                    'absolute right-2 top-1/2 z-20 size-4 -translate-y-1/2 text-primary delay-300 duration-300 ease-in-out',
                                                    (invalid || error) &&
                                                        'text-destructive'
                                                )}
                                            />
                                            <PhoneInput
                                                {...field}
                                                className="w-full"
                                                defaultCountry="PH"
                                            />
                                        </div>
                                    </FormControl>
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                        <FormField
                            name="permanentAddress"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="w-full text-right text-sm font-medium"
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
                                    <FormMessage className="text-xs" />
                                </FormItem>
                            )}
                        />
                    </fieldset>
                </fieldset>

                <FormErrorMessage errorMessage={error} />
                <div>
                    <Separator className="my-2 sm:my-4" />
                    <div className="flex items-center justify-end gap-x-2">
                        <Button
                            type="button"
                            variant="ghost"
                            onClick={() => form.reset()}
                            className="w-full self-end px-8 sm:w-fit"
                        >
                            Reset
                        </Button>
                        <Button
                            type="submit"
                            disabled={isCreating}
                            className="w-full self-end px-8 sm:w-fit"
                        >
                            {isCreating ? <LoadingSpinner /> : 'Create'}
                        </Button>
                    </div>
                </div>
            </form>
        </Form>
    )
}

export const MemberCreateUpdateFormModal = ({
    title = 'Create Member',
    description = 'Fill out the form to add a new member.',
    className,
    formProps,
    ...props
}: IModalProps & {
    formProps: Omit<IMemberCreateUpdateFormProps, 'className'>
}) => {
    return (
        <Modal
            title={title}
            description={description}
            className={cn('sm:max-w-full lg:max-w-7xl', className)}
            {...props}
        >
            <MemberCreateUpdateForm {...formProps} />
        </Modal>
    )
}

export default MemberCreateUpdateForm
