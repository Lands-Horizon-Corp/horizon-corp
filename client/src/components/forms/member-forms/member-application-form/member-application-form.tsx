import z from 'zod'
import { useFieldArray, useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import {
    Form,
    FormItem,
    FormField,
    FormLabel,
    FormControl,
    FormMessage,
    FormHidableItem,
    FormDescription,
} from '@/components/ui/form'
import {
    Select,
    SelectItem,
    SelectValue,
    SelectTrigger,
    SelectContent,
} from '@/components/ui/select'
import { Input } from '@/components/ui/input'
import { IForm } from '@/types/component/form'
import { Button } from '@/components/ui/button'
import { Checkbox } from '@/components/ui/checkbox'
import { Separator } from '@/components/ui/separator'
import FormErrorMessage from '@/components/ui/form-error-message'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { useCreateMemberProfile } from '@/hooks/api-hooks/use-member-profile'
import { createMemberProfileSchema } from '@/validations/form-validation/member-schema'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import {
    PlusIcon,
    TrashIcon,
    VerifiedPatchIcon,
    XIcon,
} from '@/components/icons'
import TextEditor from '@/components/text-editor'
import { PhoneInput } from '@/components/contact-input/contact-input'
import { Textarea } from '@/components/ui/textarea'
import { FormLabeledInputField } from '@/components/ui/form-labeled-input-field'

type TMemberProfileForm = z.infer<typeof createMemberProfileSchema>

interface IMemberApplicationFormProps
    extends IBaseCompNoChild,
        IForm<Partial<TMemberProfileForm>, unknown, string> {}

const MemberApplicationForm = ({
    readOnly,
    className,
    defaultValues,
    onError,
    onSuccess,
    disabledFields,
    hiddenFields,
}: IMemberApplicationFormProps) => {
    const form = useForm<TMemberProfileForm>({
        resolver: zodResolver(createMemberProfileSchema),
        reValidateMode: 'onChange',
        mode: 'onSubmit',
        defaultValues: {
            notes: '',
            occupation: '',
            description: '',
            contactNumber: '',
            civilStatus: 'Single',
            status: 'Pending',
            isClosed: false,
            isMutualFundMember: false,
            isMicroFinanceMember: false,
            ...defaultValues,
        },
    })

    const isDisabled = (field: keyof TMemberProfileForm) =>
        readOnly || disabledFields?.includes(field)

    const { fields, append, remove } = useFieldArray({
        control: form.control,
        name: 'memberDescriptions',
    })

    const memberAddress = useFieldArray({
        control: form.control,
        name: 'memberAddress',
    })

    const memberContactNumberReferences = useFieldArray({
        control: form.control,
        name: 'memberContactNumberReferences',
    })

    const {
        fields: memberAssetsFields,
        append: addAsset,
        remove: removeAsset,
    } = useFieldArray({
        control: form.control,
        name: 'memberAssets',
    })

    const {
        fields: incomeFields,
        append: addIncome,
        remove: removeIncome,
    } = useFieldArray({
        control: form.control,
        name: 'memberIncome',
    })

    const {
        fields: relativeFields,
        append: addRelative,
        remove: removeRelative,
    } = useFieldArray({
        control: form.control,
        name: 'memberRelativeAccounts',
    })

    const {
        fields: memberExpensesFields,
        append: addExpense,
        remove: removeExpense,
    } = useFieldArray({
        control: form.control,
        name: 'memberExpenses',
    })

    const {
        fields: jointFields,
        append: addJoint,
        remove: removeJoint,
    } = useFieldArray({
        control: form.control,
        name: 'memberJointAccounts',
    })

    const {
        fields: recruitFields,
        append: addRecruit,
        remove: removeRecruit,
    } = useFieldArray({
        control: form.control,
        name: 'memberRecruits',
    })

    const {
        fields: govtFields,
        append: addGovtBenefit,
        remove: removeGovtBenefit,
    } = useFieldArray({
        control: form.control,
        name: 'memberGovernmentBenefits',
    })

    const {
        error,
        isPending: isCreating,
        mutate: createMemberProfile,
    } = useCreateMemberProfile({ onSuccess, onError })

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit((formData) =>
                    createMemberProfile(formData)
                )}
                className={cn('flex w-full flex-col gap-y-4', className)}
            >
                <fieldset
                    className="gap-x-4 gap-y-4 space-y-5"
                    disabled={isCreating || readOnly}
                >
                    <div className="grid gap-4 sm:grid-cols-3">
                        <div className="space-y-4">
                            <legend>Identification & Reference</legend>
                            <Separator />
                            <FormLabeledInputField
                                form={form}
                                showMessage={true}
                                name="passbookNumber"
                                label="Passbook Number"
                                hiddenFields={hiddenFields}
                                placeholder="Enter Passbook Number"
                                isDisabled={isDisabled('passbookNumber')}
                            />
                            <FormLabeledInputField
                                name="oldReferenceID"
                                form={form}
                                hiddenFields={hiddenFields}
                                isDisabled={isDisabled('oldReferenceID')}
                                label="Old Reference ID"
                                placeholder="Enter Old Passbook/ID"
                                showMessage={true}
                            />
                            <FormField
                                name="status"
                                control={form.control}
                                render={({ field }) => (
                                    <FormHidableItem
                                        field={field.name}
                                        hiddenFields={hiddenFields}
                                    >
                                        <FormItem className="col-span-1 space-y-1">
                                            <FormLabel htmlFor={field.name}>
                                                Status
                                            </FormLabel>
                                            <Select
                                                defaultValue={field.value}
                                                onValueChange={field.onChange}
                                            >
                                                <FormControl>
                                                    <SelectTrigger
                                                        disabled={isDisabled(
                                                            field.name
                                                        )}
                                                    >
                                                        <SelectValue
                                                            {...field}
                                                            id={field.name}
                                                            placeholder="Member Status"
                                                        />
                                                    </SelectTrigger>
                                                </FormControl>
                                                <SelectContent>
                                                    <SelectItem value="Pending">
                                                        Pending
                                                    </SelectItem>
                                                    <SelectItem value="Verified">
                                                        Verified
                                                    </SelectItem>
                                                    <SelectItem value="Not Allowed">
                                                        Not Allowed
                                                    </SelectItem>
                                                </SelectContent>
                                            </Select>
                                            <FormMessage className="text-xs" />
                                        </FormItem>
                                    </FormHidableItem>
                                )}
                            />
                            <FormField
                                name="isMutualFundMember"
                                control={form.control}
                                render={({ field }) => (
                                    <FormHidableItem
                                        field={field.name}
                                        hiddenFields={hiddenFields}
                                    >
                                        <FormItem className="col-span-1">
                                            <span className="inline-flex w-full items-center gap-x-2 rounded-xl bg-popover px-4 py-2">
                                                <FormControl>
                                                    <Checkbox
                                                        id={field.name}
                                                        disabled={isDisabled(
                                                            field.name
                                                        )}
                                                        checked={field.value}
                                                        onCheckedChange={(
                                                            checked
                                                        ) => {
                                                            field.onChange(
                                                                checked
                                                            )
                                                        }}
                                                    />
                                                </FormControl>
                                                <FormLabel
                                                    htmlFor={field.name}
                                                    className="ml-2 space-y-2"
                                                >
                                                    <span>
                                                        Mutualfund Member
                                                    </span>
                                                    <FormDescription className="font-normal">
                                                        Contributes to a pooled
                                                        investment (mutual
                                                        fund).
                                                    </FormDescription>
                                                </FormLabel>
                                            </span>
                                            <FormMessage className="text-xs" />
                                        </FormItem>
                                    </FormHidableItem>
                                )}
                            />
                            <FormField
                                name="isMicroFinanceMember"
                                control={form.control}
                                render={({ field }) => (
                                    <FormHidableItem
                                        field={field.name}
                                        hiddenFields={hiddenFields}
                                    >
                                        <FormItem className="col-span-1">
                                            <span className="inline-flex w-full items-center gap-x-2 rounded-xl bg-popover px-4 py-2">
                                                <FormControl>
                                                    <Checkbox
                                                        id={field.name}
                                                        disabled={isDisabled(
                                                            field.name
                                                        )}
                                                        checked={field.value}
                                                        onCheckedChange={(
                                                            checked
                                                        ) => {
                                                            field.onChange(
                                                                checked
                                                            )
                                                        }}
                                                    />
                                                </FormControl>
                                                <FormLabel
                                                    htmlFor={field.name}
                                                    className="ml-2 space-y-2"
                                                >
                                                    <span>
                                                        Micro Finance Member
                                                    </span>
                                                    <FormDescription className="font-normal">
                                                        Participates in
                                                        small-scale financial
                                                        services.
                                                    </FormDescription>
                                                </FormLabel>
                                            </span>
                                            <FormMessage className="text-xs" />
                                        </FormItem>
                                    </FormHidableItem>
                                )}
                            />
                        </div>
                        <div className="space-y-4">
                            <legend>Personal Information</legend>
                            <Separator />
                            <FormField
                                name="contactNumber"
                                control={form.control}
                                render={({
                                    field,
                                    fieldState: { invalid, error },
                                }) => (
                                    <FormHidableItem
                                        field={field.name}
                                        hiddenFields={hiddenFields}
                                    >
                                        <FormItem className="col-span-1 space-y-1">
                                            <FormLabel htmlFor={field.name}>
                                                Contact Number *
                                            </FormLabel>
                                            <FormControl>
                                                <div className="relative flex flex-1 items-center gap-x-2">
                                                    <VerifiedPatchIcon
                                                        className={cn(
                                                            'absolute right-2 top-1/2 z-20 size-4 -translate-y-1/2 text-primary delay-300 duration-300 ease-in-out',
                                                            (invalid ||
                                                                error) &&
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
                                    </FormHidableItem>
                                )}
                            />
                            <FormField
                                name="civilStatus"
                                control={form.control}
                                render={({ field }) => (
                                    <FormHidableItem
                                        field={field.name}
                                        hiddenFields={hiddenFields}
                                    >
                                        <FormItem className="col-span-1 space-y-1">
                                            <FormLabel htmlFor={field.name}>
                                                Civil Status
                                            </FormLabel>
                                            <Select
                                                onValueChange={field.onChange}
                                                defaultValue={field.value}
                                            >
                                                <FormControl>
                                                    <SelectTrigger
                                                        disabled={isDisabled(
                                                            field.name
                                                        )}
                                                    >
                                                        <SelectValue
                                                            {...field}
                                                            id={field.name}
                                                            placeholder="Select Civil Status"
                                                        />
                                                    </SelectTrigger>
                                                </FormControl>
                                                <SelectContent>
                                                    <SelectItem value="Married">
                                                        Married
                                                    </SelectItem>
                                                    <SelectItem value="Single">
                                                        Single
                                                    </SelectItem>
                                                    <SelectItem value="Widowed">
                                                        Widowed
                                                    </SelectItem>
                                                    <SelectItem value="Separated">
                                                        Separated
                                                    </SelectItem>
                                                    <SelectItem value="N/A">
                                                        N/A
                                                    </SelectItem>
                                                </SelectContent>
                                            </Select>
                                            <FormMessage className="text-xs" />
                                        </FormItem>
                                    </FormHidableItem>
                                )}
                            />
                            <FormLabeledInputField
                                name="occupation"
                                form={form}
                                hiddenFields={hiddenFields}
                                isDisabled={isDisabled('occupation')}
                                label="Occupation"
                                placeholder="Enter Occupation"
                                showMessage={true}
                            />
                            <FormField
                                control={form.control}
                                name="businessAddress"
                                render={({ field }) => (
                                    <FormHidableItem
                                        field={field.name}
                                        hiddenFields={hiddenFields}
                                    >
                                        <FormItem className="col-span-1 space-y-1">
                                            <FormLabel htmlFor={field.name}>
                                                Business Address
                                            </FormLabel>
                                            <FormControl>
                                                <Textarea
                                                    {...field}
                                                    id={field.name}
                                                    placeholder="Enter business address"
                                                    className="min-h-0"
                                                    disabled={isDisabled(
                                                        field.name
                                                    )}
                                                />
                                            </FormControl>
                                            <FormMessage className="text-xs" />
                                        </FormItem>
                                    </FormHidableItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="businessContact"
                                render={({
                                    field,
                                    fieldState: { invalid },
                                }) => (
                                    <FormHidableItem
                                        field={field.name}
                                        hiddenFields={hiddenFields}
                                    >
                                        <FormItem className="col-span-1 space-y-1">
                                            <FormLabel htmlFor={field.name}>
                                                Business Contact
                                            </FormLabel>
                                            <FormControl>
                                                <div className="relative flex flex-1 items-center gap-x-2">
                                                    <VerifiedPatchIcon
                                                        className={cn(
                                                            'absolute right-2 top-1/2 z-20 size-4 -translate-y-1/2 text-primary delay-300 duration-300 ease-in-out',
                                                            (invalid ||
                                                                error) &&
                                                                'text-destructive'
                                                        )}
                                                    />
                                                    <PhoneInput
                                                        {...field}
                                                        disabled={isDisabled(
                                                            field.name
                                                        )}
                                                        className="w-full"
                                                        defaultCountry="PH"
                                                    />
                                                </div>
                                            </FormControl>
                                            <FormMessage className="text-xs" />
                                        </FormItem>
                                    </FormHidableItem>
                                )}
                            />
                        </div>
                        <div className="space-y-4">
                            <legend>Government Related Info</legend>
                            <Separator />

                            <FormLabeledInputField
                                name="tinNumber"
                                form={form}
                                hiddenFields={hiddenFields}
                                isDisabled={isDisabled('tinNumber')}
                                label="TIN Number"
                                placeholder="Enter TIN Number"
                            />

                            <FormLabeledInputField
                                name="sssNumber"
                                form={form}
                                hiddenFields={hiddenFields}
                                isDisabled={isDisabled('sssNumber')}
                                label="SSS Number"
                                placeholder="Enter SSS Number"
                            />

                            <FormLabeledInputField
                                name="pagibigNumber"
                                form={form}
                                hiddenFields={hiddenFields}
                                isDisabled={isDisabled('pagibigNumber')}
                                label="PAGIBIG Number"
                                placeholder="Enter PAGIBIG Number"
                            />

                            <FormLabeledInputField
                                name="philhealthNumber"
                                form={form}
                                hiddenFields={hiddenFields}
                                isDisabled={isDisabled('philhealthNumber')}
                                label="PhilHealth Number"
                                placeholder="Enter PhilHealth Number"
                            />
                        </div>
                    </div>
                    <div className="col-span-1 space-y-4">
                        <legend>Notes & Description</legend>
                        <Separator />
                        <fieldset className="grid gap-x-2 sm:grid-cols-2">
                            <FormField
                                name="notes"
                                control={form.control}
                                render={({ field }) => (
                                    <FormHidableItem
                                        field={field.name}
                                        hiddenFields={hiddenFields}
                                    >
                                        <FormItem className="col-span-1 space-y-1">
                                            <FormLabel htmlFor={field.name}>
                                                Notes *
                                            </FormLabel>
                                            <FormControl>
                                                <TextEditor
                                                    {...field}
                                                    placeholder="Enter Notes"
                                                    className=""
                                                    textEditorClassName="!max-w-none"
                                                    disabled={isDisabled(
                                                        field.name
                                                    )}
                                                />
                                            </FormControl>
                                            <FormMessage className="text-xs" />
                                        </FormItem>
                                    </FormHidableItem>
                                )}
                            />
                            <FormField
                                name="description"
                                control={form.control}
                                render={({ field }) => (
                                    <FormHidableItem
                                        field={field.name}
                                        hiddenFields={hiddenFields}
                                    >
                                        <FormItem className="col-span-1 space-y-1">
                                            <FormLabel htmlFor={field.name}>
                                                Description *
                                            </FormLabel>
                                            <FormControl>
                                                <TextEditor
                                                    {...field}
                                                    placeholder="Enter Description"
                                                />
                                            </FormControl>
                                            <FormMessage className="text-xs" />
                                        </FormItem>
                                    </FormHidableItem>
                                )}
                            />
                        </fieldset>
                    </div>

                    <FormField
                        control={form.control}
                        name="memberAddress"
                        render={({ fieldState }) => (
                            <FormItem className="col-span-1 mt-8 space-y-2">
                                <FormLabel>Address</FormLabel>

                                {fieldState.error?.message && (
                                    <FormMessage className="text-xs text-red-500">
                                        {fieldState.error.message}
                                    </FormMessage>
                                )}
                                <Separator />
                                <fieldset className="grid gap-4">
                                    {memberAddress.fields.map(
                                        (addressField, index) => (
                                            <div
                                                key={addressField.id}
                                                className="flex w-full flex-col gap-4 md:flex-row"
                                            >
                                                <FormLabeledInputField
                                                    name={`memberAddress.${index}.label`}
                                                    form={form}
                                                    label="Label *"
                                                    placeholder="Label"
                                                    className="w-full"
                                                />

                                                <FormLabeledInputField
                                                    name={`memberAddress.${index}.barangay`}
                                                    form={form}
                                                    label="Barangay *"
                                                    placeholder="Barangay"
                                                    className="w-full"
                                                />

                                                <FormLabeledInputField
                                                    name={`memberAddress.${index}.city`}
                                                    form={form}
                                                    label="City *"
                                                    placeholder="City"
                                                    className="w-full"
                                                />

                                                <FormLabeledInputField
                                                    name={`memberAddress.${index}.province`}
                                                    form={form}
                                                    label="Province *"
                                                    placeholder="Province"
                                                    className="w-full"
                                                />
                                                <Button
                                                    size="icon"
                                                    type="button"
                                                    variant="secondary"
                                                    onClick={() =>
                                                        memberAddress.remove(
                                                            index
                                                        )
                                                    }
                                                    className="self-center rounded-full p-2"
                                                >
                                                    <XIcon className="size-4" />
                                                </Button>
                                            </div>
                                        )
                                    )}
                                </fieldset>
                                <Button
                                    size="sm"
                                    type="button"
                                    variant="secondary"
                                    onClick={() =>
                                        memberAddress.append({
                                            label: '',
                                            barangay: '',
                                            city: '',
                                            province: '',
                                            postalCode: '',
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" /> Add Address
                                </Button>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="memberContactNumberReferences"
                        render={({ fieldState }) => (
                            <FormItem className="col-span-1 mt-8 space-y-2">
                                <FormLabel>Contact Number References</FormLabel>
                                {fieldState.error?.message && (
                                    <FormMessage className="text-xs text-red-500">
                                        {fieldState.error.message}
                                    </FormMessage>
                                )}

                                <Separator />

                                <fieldset className="grid gap-4">
                                    {memberContactNumberReferences.fields.map(
                                        (contactField, index) => (
                                            <div
                                                key={contactField.id}
                                                className="flex w-full flex-col gap-4 md:flex-row"
                                            >
                                                <FormLabeledInputField
                                                    name={`memberContactNumberReferences.${index}.name`}
                                                    form={form}
                                                    label="Name *"
                                                    placeholder="Reference Name"
                                                    className="w-full"
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberContactNumberReferences.${index}.description`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Description *
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Textarea
                                                                    placeholder="Description Content"
                                                                    className="min-h-0"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberContactNumberReferences.${index}.contactNumber`}
                                                    render={({
                                                        field,
                                                        fieldState: {
                                                            invalid,
                                                            error,
                                                        },
                                                    }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel
                                                                htmlFor={
                                                                    field.name
                                                                }
                                                            >
                                                                Contact Number *
                                                            </FormLabel>
                                                            <FormControl>
                                                                <div className="relative flex flex-1 items-center gap-x-2">
                                                                    <VerifiedPatchIcon
                                                                        className={cn(
                                                                            'absolute right-2 top-1/2 z-20 size-4 -translate-y-1/2 text-primary delay-300 duration-300 ease-in-out',
                                                                            (invalid ||
                                                                                error) &&
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
                                                <Button
                                                    size="icon"
                                                    type="button"
                                                    variant="secondary"
                                                    onClick={() =>
                                                        memberContactNumberReferences.remove(
                                                            index
                                                        )
                                                    }
                                                    className="self-center rounded-full p-2"
                                                >
                                                    <XIcon className="size-4" />
                                                </Button>
                                            </div>
                                        )
                                    )}
                                </fieldset>

                                <Button
                                    size="sm"
                                    type="button"
                                    variant="secondary"
                                    onClick={() =>
                                        memberContactNumberReferences.append({
                                            name: '',
                                            description: '',
                                            contactNumber: '',
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" />
                                    Add Contact Reference
                                </Button>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="memberIncome"
                        render={({ fieldState }) => (
                            <FormItem className="col-span-1 mt-8 space-y-2">
                                <FormLabel>Member Income</FormLabel>
                                {fieldState.error?.message && (
                                    <FormMessage className="text-xs text-red-500">
                                        {fieldState.error.message}
                                    </FormMessage>
                                )}

                                <Separator />

                                <fieldset className="grid gap-4">
                                    {incomeFields.map((incomeField, index) => (
                                        <div
                                            key={incomeField.id}
                                            className="flex w-full flex-col gap-4 md:flex-row"
                                        >
                                            <FormLabeledInputField
                                                name={`memberIncome.${index}.name`}
                                                form={form}
                                                hiddenFields={hiddenFields}
                                                label="Name *"
                                                placeholder="Income Name"
                                                className="w-full"
                                            />

                                            <FormLabeledInputField
                                                name={`memberIncome.${index}.amount`}
                                                form={form}
                                                hiddenFields={hiddenFields}
                                                label="Amount *"
                                                placeholder="0"
                                                className="w-full"
                                                inputProps={{ type: 'number' }}
                                            />

                                            <FormLabeledInputField
                                                name={`memberIncome.${index}.date`}
                                                form={form}
                                                hiddenFields={hiddenFields}
                                                label="Date *"
                                                placeholder="YYYY-MM-DD"
                                                className="w-full"
                                                inputProps={{ type: 'date' }}
                                            />

                                            <FormField
                                                control={form.control}
                                                name={`memberIncome.${index}.description`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full space-y-1">
                                                        <FormLabel>
                                                            Description *
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Textarea
                                                                placeholder="Description"
                                                                className="min-h-0"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <Button
                                                size="icon"
                                                variant="secondary"
                                                type="button"
                                                onClick={() =>
                                                    removeIncome(index)
                                                }
                                                className="self-center rounded-full p-2"
                                            >
                                                <XIcon className="size-4" />
                                            </Button>
                                        </div>
                                    ))}
                                </fieldset>

                                <Button
                                    size="sm"
                                    type="button"
                                    variant="secondary"
                                    onClick={() =>
                                        addIncome({
                                            name: '',
                                            amount: 0,
                                            date: '',
                                            description: '',
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" />
                                    Add Income
                                </Button>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="memberRelativeAccounts"
                        render={({ fieldState }) => (
                            <FormItem className="col-span-1 mt-8 space-y-2">
                                <FormLabel>Relative Accounts</FormLabel>
                                {fieldState.error?.message && (
                                    <FormMessage className="text-xs text-red-500">
                                        {fieldState.error.message}
                                    </FormMessage>
                                )}

                                <Separator />

                                <fieldset className="grid gap-4">
                                    {relativeFields.map((relField, index) => (
                                        <div
                                            key={relField.id}
                                            className="flex w-full flex-col gap-4 md:flex-row"
                                        >
                                            <FormField
                                                control={form.control}
                                                name={`memberRelativeAccounts.${index}.membersProfileID`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Member's Profile ID
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                {...field}
                                                                placeholder="Member's Profile ID"
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberRelativeAccounts.${index}.relativeProfileMemberID`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Relative Profile
                                                            Member ID
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                {...field}
                                                                placeholder="Relative Profile Member ID"
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberRelativeAccounts.${index}.familyRelationship`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Family Relationship
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                {...field}
                                                                placeholder="e.g. Spouse, Sibling"
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberRelativeAccounts.${index}.description`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Description
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Textarea
                                                                {...field}
                                                                placeholder="Additional details"
                                                                className="min-h-0"
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <Button
                                                size="icon"
                                                type="button"
                                                variant="secondary"
                                                onClick={() =>
                                                    removeRelative(index)
                                                }
                                                className="self-center rounded-full p-2"
                                            >
                                                <XIcon className="size-4" />
                                            </Button>
                                        </div>
                                    ))}
                                </fieldset>
                                <Button
                                    size="sm"
                                    type="button"
                                    variant="secondary"
                                    onClick={() =>
                                        addRelative({
                                            membersProfileID: '',
                                            relativeProfileMemberID: '',
                                            familyRelationship: '',
                                            description: '',
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" />
                                    Add Relative
                                </Button>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="memberAssets"
                        render={({ fieldState }) => (
                            <FormItem className="col-span-1 mt-8 space-y-2">
                                <FormLabel>Assets</FormLabel>
                                {fieldState.error?.message && (
                                    <FormMessage className="text-xs text-red-500">
                                        {fieldState.error.message}
                                    </FormMessage>
                                )}
                                <Separator />
                                <fieldset className="grid gap-4">
                                    {memberAssetsFields.map(
                                        (assetField, index) => (
                                            <div
                                                key={assetField.id}
                                                className="flex w-full flex-col gap-4 md:flex-row"
                                            >
                                                <FormField
                                                    control={form.control}
                                                    name={`memberAssets.${index}.name`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Name
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Input
                                                                    placeholder="Name"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberAssets.${index}.entryDate`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Entry Date
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Input
                                                                    type="date"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberAssets.${index}.description`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Description
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Textarea
                                                                    className="min-h-0"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <Button
                                                    size="icon"
                                                    type="button"
                                                    variant="secondary"
                                                    onClick={() =>
                                                        removeAsset(index)
                                                    }
                                                    className="self-center rounded-full p-2"
                                                >
                                                    <XIcon className="size-4" />
                                                </Button>
                                            </div>
                                        )
                                    )}
                                </fieldset>
                                <Button
                                    size="sm"
                                    type="button"
                                    variant="secondary"
                                    onClick={() =>
                                        addAsset({
                                            entryDate: '',
                                            description: '',
                                            name: '',
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" />
                                    Add Asset
                                </Button>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="memberExpenses"
                        render={({ fieldState }) => (
                            <FormItem className="col-span-1 mt-8 space-y-2">
                                <FormLabel>Expenses</FormLabel>
                                {fieldState.error?.message && (
                                    <FormMessage className="text-xs text-red-500">
                                        {fieldState.error.message}
                                    </FormMessage>
                                )}
                                <Separator />
                                <fieldset className="grid gap-4">
                                    {memberExpensesFields.map(
                                        (expenseField, index) => (
                                            <div
                                                key={expenseField.id}
                                                className="flex w-full flex-col gap-4 md:flex-row"
                                            >
                                                <FormField
                                                    control={form.control}
                                                    name={`memberExpenses.${index}.name`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Name
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Input
                                                                    placeholder="Name"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberExpenses.${index}.date`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Date
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Input
                                                                    type="date"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberExpenses.${index}.amount`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Amount
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Input
                                                                    type="number"
                                                                    placeholder="0"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberExpenses.${index}.description`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Description
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Textarea
                                                                    placeholder="Description"
                                                                    className="min-h-0"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <Button
                                                    size="icon"
                                                    type="button"
                                                    variant="secondary"
                                                    onClick={() =>
                                                        removeExpense(index)
                                                    }
                                                    className="self-center rounded-full p-2"
                                                >
                                                    <XIcon className="size-4" />
                                                </Button>
                                            </div>
                                        )
                                    )}
                                </fieldset>
                                <Button
                                    size="sm"
                                    type="button"
                                    variant="secondary"
                                    onClick={() =>
                                        addExpense({
                                            name: '',
                                            date: '',
                                            amount: 0,
                                            description: '',
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" />
                                    Add Expense
                                </Button>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="memberJointAccounts"
                        render={({ fieldState }) => (
                            <FormItem className="col-span-1 mt-8 space-y-2">
                                <FormLabel>Joint Accounts</FormLabel>
                                {fieldState.error?.message && (
                                    <FormMessage className="text-xs text-red-500">
                                        {fieldState.error.message}
                                    </FormMessage>
                                )}
                                <Separator />
                                <fieldset className="grid gap-4">
                                    {jointFields.map((jointField, index) => (
                                        <div
                                            key={jointField.id}
                                            className="flex w-full flex-col gap-4 md:flex-row"
                                        >
                                            <FormField
                                                control={form.control}
                                                name={`memberJointAccounts.${index}.firstName`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            First Name
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                placeholder="e.g. John"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberJointAccounts.${index}.lastName`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Last Name
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                placeholder="e.g. Doe"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberJointAccounts.${index}.middleName`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Middle Name
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                placeholder="e.g. M."
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberJointAccounts.${index}.familyRelationship`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Family Relationship
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                placeholder="e.g. Cousin, Spouse"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberJointAccounts.${index}.description`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Description
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Textarea
                                                                className="min-h-0"
                                                                placeholder="Description"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <Button
                                                size="icon"
                                                type="button"
                                                variant="secondary"
                                                onClick={() =>
                                                    removeJoint(index)
                                                }
                                                className="self-center rounded-full p-2"
                                            >
                                                <XIcon className="size-4" />
                                            </Button>
                                        </div>
                                    ))}
                                </fieldset>
                                <Button
                                    size="sm"
                                    type="button"
                                    variant="secondary"
                                    onClick={() =>
                                        addJoint({
                                            description: '',
                                            firstName: '',
                                            lastName: '',
                                            middleName: '',
                                            familyRelationship: '',
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" />
                                    Add Joint Account
                                </Button>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="memberRecruits"
                        render={({ fieldState }) => (
                            <FormItem className="col-span-1 mt-8 space-y-2">
                                <FormLabel>Recruits</FormLabel>
                                {fieldState.error?.message && (
                                    <FormMessage className="text-xs text-red-500">
                                        {fieldState.error.message}
                                    </FormMessage>
                                )}
                                <Separator />
                                <fieldset className="grid gap-4">
                                    {recruitFields.map(
                                        (recruitField, index) => (
                                            <div
                                                key={recruitField.id}
                                                className="flex w-full flex-col gap-4 md:flex-row"
                                            >
                                                <FormField
                                                    control={form.control}
                                                    name={`memberRecruits.${index}.membersProfileID`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Member's Profile
                                                                ID
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Input
                                                                    placeholder="Member's Profile ID"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberRecruits.${index}.membersProfileRecruitedID`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Recruited
                                                                Member's Prof.
                                                                ID
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Input
                                                                    placeholder="Recruited Member's Profile"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberRecruits.${index}.dateRecruited`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Date Recruited
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Input
                                                                    type="date"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberRecruits.${index}.description`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Description
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Textarea
                                                                    className="min-h-0"
                                                                    placeholder="Description"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <FormField
                                                    control={form.control}
                                                    name={`memberRecruits.${index}.name`}
                                                    render={({ field }) => (
                                                        <FormItem className="w-full">
                                                            <FormLabel>
                                                                Name
                                                            </FormLabel>
                                                            <FormControl>
                                                                <Input
                                                                    placeholder="Name"
                                                                    {...field}
                                                                />
                                                            </FormControl>
                                                            <FormMessage />
                                                        </FormItem>
                                                    )}
                                                />
                                                <Button
                                                    size="icon"
                                                    variant="secondary"
                                                    type="button"
                                                    onClick={() =>
                                                        removeRecruit(index)
                                                    }
                                                    className="self-center rounded-full p-2"
                                                >
                                                    <XIcon className="size-4" />
                                                </Button>
                                            </div>
                                        )
                                    )}
                                </fieldset>
                                <Button
                                    size="sm"
                                    variant="secondary"
                                    type="button"
                                    onClick={() =>
                                        addRecruit({
                                            membersProfileID: '',
                                            membersProfileRecruitedID: '',
                                            dateRecruited: '',
                                            description: '',
                                            name: '',
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" />
                                    Add Recruit
                                </Button>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="memberGovernmentBenefits"
                        render={({ fieldState }) => (
                            <FormItem className="col-span-1 mt-8 space-y-2">
                                <FormLabel>Government Benefits</FormLabel>
                                {fieldState.error?.message && (
                                    <FormMessage className="text-xs text-red-500">
                                        {fieldState.error.message}
                                    </FormMessage>
                                )}
                                <Separator />
                                <fieldset className="grid gap-4">
                                    {govtFields.map((benefitField, index) => (
                                        <div
                                            key={benefitField.id}
                                            className="flex w-full flex-col gap-4 md:flex-row"
                                        >
                                            <FormField
                                                control={form.control}
                                                name={`memberGovernmentBenefits.${index}.country`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Country
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                placeholder="Country"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberGovernmentBenefits.${index}.name`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Name
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                placeholder="Name"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberGovernmentBenefits.${index}.description`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Description
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Textarea
                                                                className="min-h-0"
                                                                placeholder="Description"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />
                                            <FormField
                                                control={form.control}
                                                name={`memberGovernmentBenefits.${index}.value`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Value
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                type="number"
                                                                placeholder="0"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />

                                            <FormField
                                                control={form.control}
                                                name={`memberGovernmentBenefits.${index}.frontMediaID`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Front Media ID
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                placeholder="Front Media ID"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />

                                            <FormField
                                                control={form.control}
                                                name={`memberGovernmentBenefits.${index}.backMediaID`}
                                                render={({ field }) => (
                                                    <FormItem className="w-full">
                                                        <FormLabel>
                                                            Back Media ID
                                                        </FormLabel>
                                                        <FormControl>
                                                            <Input
                                                                placeholder="Back Media ID"
                                                                {...field}
                                                            />
                                                        </FormControl>
                                                        <FormMessage />
                                                    </FormItem>
                                                )}
                                            />

                                            <Button
                                                size="icon"
                                                variant="secondary"
                                                type="button"
                                                onClick={() =>
                                                    removeGovtBenefit(index)
                                                }
                                                className="self-center rounded-full p-2"
                                            >
                                                <XIcon className="size-4" />
                                            </Button>
                                        </div>
                                    ))}
                                </fieldset>
                                <Button
                                    size="sm"
                                    variant="secondary"
                                    type="button"
                                    onClick={() =>
                                        addGovtBenefit({
                                            country: '',
                                            name: '',
                                            description: '',
                                            value: 0,
                                            frontMediaID: undefined,
                                            backMediaID: undefined,
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" />
                                    Add Government Benefit
                                </Button>
                            </FormItem>
                        )}
                    />

                    <div className="col-span-1 mt-8 space-y-2">
                        <legend>Descriptions</legend>
                        <Separator />
                        <fieldset className="grid gap-4 sm:grid-cols-2">
                            {fields.map((field, index) => (
                                <div
                                    key={field.id}
                                    className="relative space-y-2 rounded-lg bg-popover p-4"
                                >
                                    <FormField
                                        control={form.control}
                                        name={`memberDescriptions.${index}.name`}
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>Name</FormLabel>
                                                <FormControl>
                                                    <Input
                                                        placeholder="Description Name"
                                                        {...field}
                                                    />
                                                </FormControl>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                    <FormField
                                        control={form.control}
                                        name={`memberDescriptions.${index}.description`}
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    Description
                                                </FormLabel>
                                                <FormControl>
                                                    <TextEditor
                                                        placeholder="Description Content"
                                                        {...field}
                                                    />
                                                </FormControl>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                    <Button
                                        size="icon"
                                        type="button"
                                        variant="secondary"
                                        onClick={() => remove(index)}
                                        className="absolute right-2 top-0 size-fit rounded-full p-2"
                                    >
                                        <TrashIcon className="size-4" />
                                    </Button>
                                </div>
                            ))}
                        </fieldset>
                        <Button
                            size="sm"
                            type="button"
                            variant="secondary"
                            className="items-center"
                            onClick={() =>
                                append({ name: '', description: '' })
                            }
                        >
                            <PlusIcon className="mr-2" /> Add Description
                        </Button>
                    </div>
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

export default MemberApplicationForm
