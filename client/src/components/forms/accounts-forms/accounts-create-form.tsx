import Modal, { IModalProps } from '@/components/modals/modal'
import { cn } from '@/lib'
import { IAccountRequest } from '@/server/types/accounts/accounts'
import { IForm } from '@/types/component/form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import {
    AccountRequestSchema,
    TAccountingAccountsEnum,
    TEarnedUnearnedInterestEnum,
    TOtherAccountInformationEnum,
} from '@/validations/form-validation/accounts-schema'
import { Checkbox } from '@/components/ui/checkbox'
import { Textarea } from '@/components/ui/textarea'
import {
    Form,
    FormField,
    FormItem,
    FormLabel,
    FormControl,
    FormMessage,
} from '@/components/ui/form'
import { useForm } from 'react-hook-form'
import { Separator } from '@/components/ui/separator'
import { IBaseCompNoChild } from '@/types'

type TAccountsCreateForm = z.infer<typeof AccountRequestSchema>

interface IAccountsCreateFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IAccountRequest>, unknown, string> {}

const AccountsCreateForm = ({
    readOnly,
    className,
    defaultValues,
    onError,
    onSuccess,
}: IAccountsCreateFormProps) => {
    const form = useForm<TAccountsCreateForm>({
        resolver: zodResolver(AccountRequestSchema),
        reValidateMode: 'onChange',
        mode: 'onSubmit',
        defaultValues: {
            id: '',
            companyId: '',
            accountCode: '',
            description: '',
            altDescription: '',
            type: 'Deposit',
            maxAmount: 0,
            minAmount: 0,
            computationType: '',
            headerRow: 0,
            centerRow: 0,
            totalRow: 0,
            print: false,
            addOn: false,
            allowRebate: false,
            taxable: false,
            finesAmort: 0,
            finesMaturity: 0,
            interestStandard: 0,
            interestSecured: 0,
            schemeNo: 0,
            altCode: 0,
            glCode: 0,
            finesGpAmort: 0,
            addtlGp: '',
            noGracePeriodDaily: false,
            finesGpMaturity: 0,
            earnedUnearnedInterest: 'None',
            otherInformationOfAnAccount: 'None',
        },
    })

    return (
        <Form {...form}>
            <Separator />
            <form
                onSubmit={form.handleSubmit((formData) => {
                    console.log('formData', formData)
                    onSuccess?.(formData)
                })}
                className={cn('flex w-full flex-col gap-y-6', className)}
            >
                <fieldset
                    disabled={readOnly}
                    className="grid grid-cols-3 gap-x-6 gap-y-6 sm:grid-cols-3 sm:gap-y-4"
                >
                    {/* Basic Information */}

                    <fieldset className="space-y-3">
                        <legend className="text-lg font-semibold text-secondary-foreground">
                            Basic Information
                        </legend>

                        <FormField
                            name="accountCode"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel>Account Code</FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            placeholder="Enter Account Code"
                                        />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />

                        <FormField
                            name="description"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel>Description</FormLabel>
                                    <FormControl>
                                        <Textarea
                                            {...field}
                                            placeholder="Enter Description"
                                            className="resize-none"
                                        />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />

                        <FormField
                            name="altDescription"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel>
                                        Alternative Description
                                    </FormLabel>
                                    <FormControl>
                                        <Textarea
                                            {...field}
                                            placeholder="Enter Alternative Description"
                                            className="resize-none"
                                        />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    </fieldset>

                    {/* Account Type & Related Info */}
                    <fieldset className="space-y-3">
                        <legend className="text-lg font-semibold text-secondary-foreground">
                            Account Type & Information
                        </legend>

                        <FormField
                            name="type"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel>Account Type</FormLabel>
                                    <FormControl>
                                        <Select
                                            onValueChange={field.onChange}
                                            defaultValue={field.value}
                                        >
                                            <SelectTrigger>
                                                <SelectValue placeholder="Select Account Type" />
                                            </SelectTrigger>
                                            <SelectContent>
                                                {TAccountingAccountsEnum.options.map(
                                                    (option) => (
                                                        <SelectItem
                                                            key={option}
                                                            value={option}
                                                        >
                                                            {option}
                                                        </SelectItem>
                                                    )
                                                )}
                                            </SelectContent>
                                        </Select>
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />

                        <FormField
                            name="earnedUnearnedInterest"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel>
                                        Earned/Unearned Interest
                                    </FormLabel>
                                    <FormControl>
                                        <Select
                                            onValueChange={field.onChange}
                                            defaultValue={field.value}
                                        >
                                            <SelectTrigger>
                                                <SelectValue placeholder="Select Earned/Unearned Interest" />
                                            </SelectTrigger>
                                            <SelectContent>
                                                {TEarnedUnearnedInterestEnum.options.map(
                                                    (option) => (
                                                        <SelectItem
                                                            key={option}
                                                            value={option}
                                                        >
                                                            {option}
                                                        </SelectItem>
                                                    )
                                                )}
                                            </SelectContent>
                                        </Select>
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />

                        <FormField
                            name="otherInformationOfAnAccount"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel>
                                        Other Account Information
                                    </FormLabel>
                                    <FormControl>
                                        <Select
                                            onValueChange={field.onChange}
                                            defaultValue={field.value}
                                        >
                                            <SelectTrigger>
                                                <SelectValue placeholder="Select Other Account Information" />
                                            </SelectTrigger>
                                            <SelectContent>
                                                {TOtherAccountInformationEnum.options.map(
                                                    (option) => (
                                                        <SelectItem
                                                            key={option}
                                                            value={option}
                                                        >
                                                            {option}
                                                        </SelectItem>
                                                    )
                                                )}
                                            </SelectContent>
                                        </Select>
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    </fieldset>

                    {/* Boolean Settings */}
                    {/* <fieldset className="space-y-3">
                        <legend className="text-lg font-semibold text-secondary-foreground">
                            Settings
                        </legend>
                    </fieldset> */}

                    {/* Financial Details */}
                    <fieldset className="space-y-3">
                        <legend className="text-lg font-semibold text-secondary-foreground">
                            Financial Details
                        </legend>

                        {(
                            [
                                'finesAmort',
                                'finesMaturity',
                                'interestStandard',
                                'interestSecured',
                            ] as const
                        ).map((fieldName) => (
                            <FormField
                                key={fieldName}
                                name={fieldName}
                                control={form.control}
                                render={({ field }) => (
                                    <FormItem className="col-span-1 space-y-1">
                                        <FormLabel>
                                            {fieldName
                                                .replace(/([A-Z])/g, ' $1')
                                                .replace(/^./, (str) =>
                                                    str.toUpperCase()
                                                )}
                                        </FormLabel>
                                        <FormControl>
                                            <Input
                                                type="number"
                                                {...field}
                                                placeholder={`Enter ${fieldName}`}
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                        ))}
                    </fieldset>

                    {/* Additional Settings */}
                    <fieldset className=" flex col-span-3  justify-between items-center">
                        <legend className="text-lg font-semibold text-secondary-foreground">
                            Additional Settings
                        </legend>
                        <FormField
                            name="finesGpMaturity"
                            control={form.control}
                            render={({ field }) => (
                                <FormItem className="col-span-1 space-y-1">
                                    <FormLabel>Fines GP Maturity</FormLabel>
                                    <FormControl>
                                        <Input
                                            type="number"
                                            {...field}
                                            placeholder="Enter Fines GP Maturity"
                                        />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <div className='w-full flex p-5 items-center justify-start  space-x-5'>
                        {(
                            [
                                'print',
                                'addOn',
                                'allowRebate',
                                'taxable',
                                'noGracePeriodDaily'
                            ] as const
                        ).map((fieldName) => (
                            <FormField
                                key={fieldName}
                                name={fieldName}
                                control={form.control}
                                render={({ field }) => 
                                    <FormItem className="flex flex-row items-center space-x-2 ">
                                        <FormControl className=''>
                                            <Checkbox
                                                checked={field.value}
                                                onCheckedChange={field.onChange}
                                            />
                                        </FormControl>
                                        <FormLabel className="text-sm font-medium -translate-y-1 text-secondary-foreground">
                                            {fieldName.charAt(0).toUpperCase() +
                                                fieldName.slice(1)}
                                        </FormLabel>
                                    </FormItem>
                                }
                            />
                        ))}
                        </div>
                    </fieldset>
                </fieldset>
                <div>
                    {/* <Separator className="my-2 sm:my-4" /> */}
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
                            disabled={false}
                            className="w-full self-end px-8 sm:w-fit"
                        >
                            Create
                            {/* {isCreating ? <LoadingSpinner /> : 'Create'} */}
                        </Button>
                    </div>
                </div>
            </form>
        </Form>
    )
}

export const AccountsCreateFormModal = ({
    title = 'Create Accounts',
    description = 'Fill out the form to add a new accounts.',
    className,
    formProps,
    ...props
}: IModalProps & { formProps?: Partial<IAccountsCreateFormProps> }) => {
    return (
        <Modal
            title={title}
            description={description}
            className={cn('sm:max-w-full lg:max-w-7xl', className)}
            {...props}
        >
            <AccountsCreateForm {...(formProps as IAccountsCreateFormProps)} />
        </Modal>
    )
}

export default AccountsCreateForm
