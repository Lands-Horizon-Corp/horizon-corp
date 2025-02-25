import * as z from 'zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
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
import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { IPaymentsEntryRequest } from '@/server/types/transactions/payments-entry'
import { useCreatePaymentEntry } from '@/hooks/api-hooks/transactions/use-payments-entry'
import { IMemberProfileResource, IMemberResource, TEntityId } from '@/server'
import AccountsPicker from '@/components/pickers/accounts-picker'
import AccountsLedgerTable from '@/components/tables/transactions/accouting-ledger-table'
import PaymentsEntryProfile from '@/components/transaction-profile/payments-entry-profile'
import { Card, CardContent } from '@/components/ui/card'
import MemberPicker from '@/components/pickers/member-picker'
import Modal from '@/components/modals/modal'
import { Checkbox } from '@/components/ui/checkbox'
import { Textarea } from '@/components/ui/textarea'
import { toast } from 'sonner'
import { useFilteredPaginatedIAccountingLedger } from '@/hooks/api-hooks/transactions/use-accounting-ledger'
import CurrentPaymentAccountingTransactionLedger from '@/components/ledger/payments-entry/current-transaction-ledger'

export const paymentsEntrySchema = z.object({
    ORNumber: z.string().min(1, 'OR Number is required'),
    amount: z.coerce.number().positive('Amount must be greater than 0'),
    accountsId: z.string().min(1, 'Accounts is required'),
    isPrinted: z.boolean().optional(),
    notes: z.string().optional(),
})

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

type TPaymentFormValues = z.infer<typeof paymentsEntrySchema>

export interface IPaymentFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IPaymentsEntryRequest>, unknown, string> {
    open: boolean
    onOpenChange: (open: boolean) => void
    selectedMemberId?: TEntityId
}

export const PaymentsEntryForm = () => {
    const [selectedMember, setSelectedMember] =
        useState<IMemberResource | null>(null)
    const [openModal, setIsOpenModal] = useState(false)

    const {
        isPending,
        isRefetching,
        data: { data: sampleLedgerData },
        refetch,
    } = useFilteredPaginatedIAccountingLedger()

    const refetchAccountingLedger = () => {
        refetch()
    }

    const totalAmount = sampleLedgerData.reduce(
        (acc, ledger) => acc + ledger.credit,
        0
    )

    const isAllowedToCreatePayment = selectedMember !== null

    const handleOpenCreateModal = () => {
        if (isAllowedToCreatePayment) {
            setIsOpenModal(true)
        }
        toast.warning('Please select member first.')
    }

    return (
        <>
            <FormModal
                selectedMemberId={selectedMember?.id}
                open={openModal}
                onOpenChange={setIsOpenModal}
            />
            <div className="flex w-full flex-col gap-y-4 p-5">
                <div className="grid grid-cols-1 gap-x-4 gap-y-4 lg:grid-cols-4">
                    <legend className="text-lg font-semibold text-secondary-foreground">
                        Payment Transactions
                    </legend>
                    <div className="col-span-2 lg:col-span-4 lg:col-start-1">
                        <div className="flex space-x-2">
                            <MemberPicker
                                value={selectedMember?.id}
                                onSelect={(member) => {
                                    setSelectedMember(member)
                                }}
                                placeholder="Select Members"
                            />
                            <Button
                                onClick={(e) => {
                                    e.preventDefault()
                                    handleOpenCreateModal()
                                }}
                                disabled={!isAllowedToCreatePayment}
                            >
                                create
                            </Button>
                        </div>
                    </div>
                    <div className="col-span-2">
                        <PaymentsEntryProfile profile={selectedMember} />
                        <div className="max-h-96 space-y-4 overflow-auto py-4">
                            <CurrentPaymentAccountingTransactionLedger
                                isRefetching={isRefetching}
                                isPending={isPending}
                                data={sampleLedgerData}
                                onRefetch={refetchAccountingLedger}
                            />
                        </div>
                    </div>
                    <div className="col-span-1 md:col-span-2">
                        <AccountsLedgerTable className="h-full" />
                    </div>
                    <div className="col-span-2 w-full">
                        <Card>
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
                </div>
            </div>
        </>
    )
}

const FormModal = ({
    open,
    onOpenChange,
    onSuccess,
    onError,
    selectedMemberId,
}: IPaymentFormProps) => {
    const form = useForm<TPaymentFormValues>({
        resolver: zodResolver(paymentsEntrySchema),
        reValidateMode: 'onChange',
        mode: 'onSubmit',
        defaultValues: {
            ORNumber: '',
            amount: 0,
            accountsId: '',
            notes: '',
            isPrinted: false,
        },
    })

    const createMutation = useCreatePaymentEntry({
        onSuccess: (data) => {
            onSuccess?.(data)
            onOpenChange(false)
        },
        onError,
    })

    const onSubmit = form.handleSubmit((data) => {
        console.log('Form data:', data)
        createMutation.mutate({
            memberId: selectedMemberId ?? '',
            ...data,
            isPrinted: data.isPrinted ?? false,
            transactionType: '',
        })
    })

    return (
        <Modal
            open={open}
            onOpenChange={onOpenChange}
            title="Payment Entry Form"
            description="Fill in the details to create a payment entry."
            footer={
                <div className="flex justify-end gap-2 p-4">
                    <Button
                        variant="outline"
                        onClick={() => onOpenChange(false)}
                    >
                        Cancel
                    </Button>
                    <Button
                        onClick={onSubmit}
                        type="submit"
                        disabled={createMutation.isPending}
                    >
                        {createMutation.isPending ? (
                            <LoadingSpinner />
                        ) : (
                            'Submit'
                        )}
                    </Button>
                </div>
            }
        >
            <Form {...form}>
                <form onSubmit={onSubmit} className="space-y-4">
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
                    <FormField
                        control={form.control}
                        name="accountsId"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel htmlFor={field.name}>
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
                        )}
                    />
                    <FormField
                        name="notes"
                        control={form.control}
                        render={({ field }) => (
                            <FormItem className="col-span-1 space-y-1">
                                <FormLabel>Add some feedback/notes.</FormLabel>
                                <FormControl>
                                    <Textarea
                                        {...field}
                                        placeholder="Enter some notes."
                                        className="resize-none"
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="isPrinted"
                        render={({ field }) => (
                            <FormItem className="flex flex-row items-center space-x-2">
                                <FormLabel htmlFor={field.name}>
                                    Allow to make print
                                </FormLabel>
                                <FormControl>
                                    <Checkbox
                                        className="-translate-y-1"
                                        checked={field.value}
                                        onCheckedChange={field.onChange}
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormErrorMessage errorMessage={createMutation.error} />
                </form>
            </Form>
        </Modal>
    )
}
