import * as z from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import FormErrorMessage from '@/components/ui/form-error-message'
import FormFieldWrapper from '@/components/ui/form-field-wrapper'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from '@/components/ui/form'
import { cn } from '@/lib/utils'

import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { IPaymentsEntryRequest } from '@/server/types/transactions/payments-entry'
import { useCreatePaymentEntry } from '@/hooks/api-hooks/transactions/use-payments-entry'
import { IMemberProfileResource, IMemberResource, TEntityId } from '@/server'
import AccountsPicker from '@/components/pickers/accounts-picker'
import AccountsLedgerTable from '@/components/tables/transactions/accouting-ledger-table'
import PaymentsEntryProfile from '@/components/transaction-profile/payments-entry-profile'
import LedgerCard from '@/components/accounts-card'

import { sampleLedgerData } from './dummy-account-ledger'
import MembersPicker from '@/components/pickers/members-picker'
import { Card, CardContent } from '@/components/ui/card'

export type memberPassbookNumber = Pick<
    IMemberProfileResource,
    'passbookNumber'
>
export interface IMemberCardResource
    extends memberPassbookNumber,
        Pick<
            IMemberResource,
            | 'id'
            | 'fullName'
            | 'contactNumber'
            | 'email'
            | 'media'
            | 'permanentAddress'
        > {}

export const sampleMemberCardResource: IMemberCardResource = {
    id: 'member-001',
    fullName: 'Jane Smith',
    contactNumber: '+1-234-567-8901',
    email: 'jane.smith@example.com',
    passbookNumber: '1234567890',
    media: {
        id: '7f76efd0-940a-42f9-afa9-8644453e20aa',
        fileName: 'sample-image.png',
        fileSize: 5242880, // 5 MB in bytes
        fileType: 'image/jpeg',
        storageKey: 'media/resources/sample-image.jpg',
        url: 'https://images.pexels.com/photos/1462980/pexels-photo-1462980.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1', // This is viewable directly in the browser
        bucketName: 'example-media-bucket',
        createdAt: '2024-10-29T08:45:00Z',
        updatedAt: '2024-10-29T10:20:00Z',
        downloadURL: 'https://cdn.example.com/media/resources/sample-image.jpg',
    },
    permanentAddress: '123 Main Street, Los Angeles, CA 90001',
}

export const paymentsEntrySchema = z.object({
    memberId: z.string().min(5, { message: 'Invalid Member ID' }),
    ORNumber: z.string().min(1, 'OR Number is required'),
    amount: z.coerce.number().positive('Amount must be greater than 0'),
    accountsId: z.string().min(1, 'accounts is required'),
    transactionType: z
        .string()
        .uuid({ message: 'Invalid Transaction Type ID' }),
    isPrinted: z.boolean(),
    notes: z.string().optional(),
})
type TPaymentFormValues = z.infer<typeof paymentsEntrySchema>

export interface IPaymentFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IPaymentsEntryRequest>, unknown, string> {
    paymentId?: TEntityId
}

export const PaymentsEntryForm = ({
    paymentId,
    // readOnly,
    className,
    defaultValues,
    onError,
    onSuccess,
}: IPaymentFormProps) => {
    const form = useForm<TPaymentFormValues>({
        resolver: zodResolver(paymentsEntrySchema),
        reValidateMode: 'onChange',
        mode: 'onSubmit',
        defaultValues: {
            memberId: '',
            ORNumber: '',
            amount: 0,
            accountsId: '',
            transactionType: '',
            isPrinted: false,
            notes: '',
            ...defaultValues,
        },
    })

    const createMutation = useCreatePaymentEntry({ onSuccess, onError })

    const onSubmit = form.handleSubmit((formData) => {
        createMutation.mutate(formData)
    })

    const { error, isPending } = createMutation

    const totalAmount = sampleLedgerData.reduce(
        (acc, ledger) => acc + ledger.credit,
        0
    )

    return (
        <Form {...form}>
            <form
                onSubmit={onSubmit}
                className={cn('flex w-full flex-col gap-y-4 p-5', className)}
            >
                <div
                    // disabled={isPending || readOnly}
                    className="grid grid-cols-1 gap-x-4 gap-y-4 lg:grid-cols-4"
                >
                    <legend className="text-lg font-semibold text-secondary-foreground">
                        Payment Transactions
                    </legend>
                    <div className="col-span-2 grid items-center gap-5 lg:col-span-4 lg:grid-cols-4">
                        <FormField
                            control={form.control}
                            name="memberId"
                            render={({ field }) => {
                                return (
                                    <FormItem className="col-span-2 space-y-1 sm:col-span-1">
                                        <FormLabel
                                            htmlFor={field.name}
                                            className="text-right text-sm font-normal text-foreground/60"
                                        >
                                            Members
                                        </FormLabel>
                                        <FormControl>
                                            <MembersPicker
                                                value={field.value}
                                                onSelect={(member) =>
                                                    field.onChange(member.id)
                                                }
                                                placeholder="Select Members"
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )
                            }}
                        />
                        <FormFieldWrapper
                            control={form.control}
                            name="memberId"
                            label="Member ID"
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    id={field.name}
                                    placeholder="Member ID"
                                    autoComplete="off"
                                />
                            )}
                        />
                        <FormFieldWrapper
                            control={form.control}
                            name="ORNumber"
                            label="OR Number"
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    id={field.name}
                                    placeholder="OR Number"
                                    autoComplete="off"
                                />
                            )}
                        />
                        <FormFieldWrapper
                            control={form.control}
                            name="amount"
                            label="Amount"
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    id={field.name}
                                    type="number"
                                    step="0.01"
                                    placeholder="Amount"
                                    autoComplete="off"
                                />
                            )}
                        />
                    </div>
                    <div className="col-span-2">
                        <PaymentsEntryProfile {...sampleMemberCardResource} />
                        <div className="max-h-96 space-y-4 overflow-auto py-4">
                            {sampleLedgerData.map((ledger) => (
                                <LedgerCard key={ledger.id} ledger={ledger} />
                            ))}
                        </div>
                    </div>
                    <div className="col-span-1 md:col-span-2">
                        <AccountsLedgerTable className="h-full" />
                    </div>
                    <div className="col-span-2 w-full">
                        <Card className="">
                            <CardContent className="flex items-center justify-between gap-x-2 py-5">
                                <label className="font-bold uppercase">
                                    Total Amount
                                </label>
                                <div className="text-right">
                                    <p className="text-2xl font-bold text-green-600 dark:text-green-400">
                                        ₱ {totalAmount}
                                    </p>
                                </div>
                            </CardContent>
                        </Card>
                    </div>
                    <div className="col-span-1 w-full md:col-span-2">
                        <FormField
                            control={form.control}
                            name="accountsId"
                            render={({ field }) => {
                                return (
                                    <FormItem className="col-span-2 space-y-1 sm:col-span-1">
                                        <FormLabel
                                            htmlFor={field.name}
                                            className="text-right text-sm font-normal text-foreground/60"
                                        >
                                            Accounts
                                        </FormLabel>
                                        <FormControl>
                                            <AccountsPicker
                                                value={field.value}
                                                onSelect={(accounts) =>
                                                    field.onChange(accounts.id)
                                                }
                                                placeholder="Select Accounts"
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )
                            }}
                        />
                    </div>
                </div>
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
                            disabled={isPending}
                            className="w-full self-end px-8 sm:w-fit"
                        >
                            {isPending ? (
                                <LoadingSpinner />
                            ) : paymentId ? (
                                'Update'
                            ) : (
                                'Create'
                            )}
                        </Button>
                    </div>
                </div>
            </form>
        </Form>
    )
}
