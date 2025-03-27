import { useCallback, useEffect, useState } from 'react'
import { useForm, UseFormReturn } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'

import { Input } from '@/components/ui/input'
import FormErrorMessage from '@/components/ui/form-error-message'
import FormFieldWrapper from '@/components/ui/form-field-wrapper'
import {
    Form,
    FormField,
    FormItem,
    FormControl,
    FormMessage,
} from '@/components/ui/form'
import { IForm } from '@/types/component/form'
import { Button } from '@/components/ui/button'

import Modal, { IModalProps } from '@/components/modals/modal'

import AccountsPicker from '@/components/pickers/accounts-picker'
import TransactionPaymentTypesPicker from '@/components/pickers/transaction-payment-types-picker'
import CheckClearingFormModal from './create-check-clearing-form'

import {
    DEFAULT_TRANSACTION_TYPE,
    paymentsEntrySchema,
} from '@/validations/transactions/payments-entry'

import {
    capitalize,
    commaSeparators,
    formatNumberOnBlur,
    isValidDecimalInput,
    sanitizeNumberInput,
} from '@/helpers'
import { cn } from '@/lib'
import { TEntityId } from '@/server'
import {
    IPaymentsEntry,
    IPaymentsEntryRequest,
} from '@/server/types/transactions/payments-entry'
import { IBaseCompNoChild } from '@/types'
import { Separator } from '@/components/ui/separator'
import { AccountsIcon, MoneyIcon, SavingsIcon } from '@/components/icons'
import useConfirmModalStore from '@/store/confirm-modal-store'
import { useShortcut } from '@/components/use-shorcuts'
import { useMemberPickerStore } from '@/store/use-member-picker-state-store'
import { Textarea } from '@/components/ui/textarea'
import { Checkbox } from '@/components/ui/checkbox'
import { Label } from '@/components/ui/label'
import {
    usePaymentsDataStore,
    usePaymentsModalStore,
} from '@/store/transaction/payments-entry-store'
import { toast } from 'sonner'
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group'
import DepositAccountPicker from '@/components/pickers/deposit-accounts-picker'
import { useCreatePaymentEntry } from '@/hooks/api-hooks/transactions/use-payments-entry'

const CHECK_BANK_LIST = ['trans-pay-002', 'trans-pay-003']

export enum TRANSACTION_TYPE {
    deposit = 'deposit',
    withdraw = 'withdraw',
    payment = 'payment',
}
export type TPaymentFormValues = z.infer<typeof paymentsEntrySchema>

export interface IPaymentFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IPaymentsEntryRequest>, unknown, string> {
    open: boolean
    onOpenChange: (open: boolean) => void
    selectedMemberId: TEntityId
    resetSelectedMember?: () => void
    form: UseFormReturn<{
        amount: number
        accountsId: string
        transactionType: string
        isPrinted?: boolean | undefined
        notes?: string | undefined
    }>
    resetForm: () => void
    error?: string
}

const PaymentsEntryFormModal = ({
    selectedMemberId,
    form,
    resetForm,
    error,
}: Omit<IPaymentFormProps, 'open' | 'onOpenChange'>) => {
    const [latestFormValues, setLatestFormValues] =
        useState<TPaymentFormValues>()
    const {
        openCheckClearingFormModal,
        setOpenCheckClearingFormModal,
        setOpenPaymentsEntryModal,
        transactionType,
    } = usePaymentsModalStore()

    const {
        selectedPayments,
        setSelectedPayments,
        ORNumber,
        selectedMember,
        setSelectedAccounts,
        selectedAccounts,
    } = usePaymentsDataStore()

    const { isOpen } = useMemberPickerStore()
    const { isOpen: isOpenConfirm } = useConfirmModalStore()
    const [pickerState, setPickerState] = useState(false)
    const {
        mutate: createPayment,
        isPending: isPendingCreatePayments,
        error: depositAndWithdrawError,
    } = useCreatePaymentEntry({
        onSuccess: () => {},
        onError: (error) => {
            toast.error(error)
        },
    })

    const errorMessage = depositAndWithdrawError || error

    const handleSubmit = useCallback(
        async (values: TPaymentFormValues) => {
            if (CHECK_BANK_LIST.includes(values.transactionType)) {
                setOpenCheckClearingFormModal(true)
                return
            }

            if (!selectedMember?.id) {
                toast.error('Please select a member')
                return
            }

            if (!selectedAccounts) {
                toast.error('Please select an account')
                return
            }

            const depositEntry: IPaymentsEntryRequest = {
                ...values,
                accountsId: values.accountsId ?? '',
                memberId: selectedMemberId,
                ORNumber: ORNumber,
                isPrinted: values.isPrinted ?? false,
            }

            const paymentEntry: IPaymentsEntry = {
                ...values,
                ORNumber,
                memberId: selectedMemberId,
                account: selectedAccounts,
                isPrinted: values.isPrinted ?? false,
            }

            if (transactionType === TRANSACTION_TYPE.payment) {
                setSelectedPayments([...selectedPayments, paymentEntry])
            }

            createPayment([depositEntry])
            toast.success('Added payments entry')
            resetForm()
            setOpenPaymentsEntryModal(false)
        },
        [
            selectedMember?.id,
            selectedAccounts,
            ORNumber,
            selectedMemberId,
            setSelectedPayments,
            resetForm,
            setOpenCheckClearingFormModal,
            setOpenPaymentsEntryModal,
        ]
    )

    useEffect(() => {
        if (openCheckClearingFormModal) {
            const currentValues = form.getValues()
            setLatestFormValues(currentValues)
        }
    }, [openCheckClearingFormModal, form])

    useShortcut('Enter', async (event) => {
        event.preventDefault()
        const activeElement = document.activeElement as HTMLElement
        if (
            activeElement.tagName === 'INPUT' ||
            activeElement.tagName === 'TEXTAREA'
        ) {
            return
        }
        const isValid = await form.trigger()
        if (isValid && !isOpen && !isOpenConfirm) {
            handleSubmit(form.getValues() as TPaymentFormValues)
        }
    })

    const depositOptions = [
        {
            value: 'Others',
            label: 'Others',
            disabled: false,
        },
        {
            value: 'Savings',
            label: 'Savings',
            disabled: false,
        },
        {
            value: 'Time',
            label: 'Time',
            disabled: true,
        },
    ]

    return (
        <>
            <CheckClearingFormModal
                formProps={{
                    resetPaymentsForm: resetForm,
                    values: latestFormValues,
                    amount: form.getValues('amount'),
                }}
                open={openCheckClearingFormModal}
                onOpenChange={setOpenCheckClearingFormModal}
            />

            <Form {...form}>
                <form onSubmit={form.handleSubmit(handleSubmit)}>
                    <div className="grid grid-cols-1 gap-3 md:grid-cols-2">
                        <FormFieldWrapper
                            className="md:col-span-2"
                            control={form.control}
                            name="amount"
                            label="Payment Amount"
                            render={({ field }) => (
                                <div className="relative w-full">
                                    <Input
                                        {...field}
                                        id={field.name}
                                        className={cn(
                                            'h-16 rounded-2xl pl-8 pr-10 text-lg font-bold text-primary placeholder:text-sm placeholder:font-normal placeholder:text-foreground/40'
                                        )}
                                        type="text"
                                        inputMode="decimal"
                                        placeholder="Enter the payment amount"
                                        autoComplete="off"
                                        value={
                                            field.value !== undefined &&
                                            field.value !== null
                                                ? commaSeparators(
                                                      field.value.toString()
                                                  )
                                                : ''
                                        }
                                        onChange={(e) => {
                                            const rawValue =
                                                sanitizeNumberInput(
                                                    e.target.value
                                                )
                                            if (isValidDecimalInput(rawValue)) {
                                                field.onChange(rawValue)
                                            }
                                        }}
                                        onBlur={(e) =>
                                            formatNumberOnBlur(
                                                e.target.value,
                                                field.onChange
                                            )
                                        }
                                    />
                                    <span className="absolute left-4 top-1/2 -translate-y-1/2 text-lg font-bold text-primary after:content-['₱']"></span>
                                </div>
                            )}
                        />
                        {transactionType === TRANSACTION_TYPE.deposit && (
                            <FormFieldWrapper
                                className="md:col-span-2"
                                control={form.control}
                                name="accountsId"
                                label=""
                                render={({ field }) => {
                                    return (
                                        <>
                                            <DepositAccountPicker
                                                value={field.value}
                                                onSelect={(account) => {
                                                    field.onChange(account.id)
                                                }}
                                                disabledButton
                                                setPickerState={setPickerState}
                                                pickerState={pickerState}
                                                placeholder="Deposit Account Name"
                                            />
                                            <RadioGroup
                                                onValueChange={(value) => {
                                                    if (value === 'Others') {
                                                        setPickerState(true)
                                                    }
                                                }}
                                                defaultValue={'Savings'}
                                                className="flex"
                                            >
                                                {depositOptions.map(
                                                    ({
                                                        value,
                                                        label,
                                                        disabled,
                                                    }) => (
                                                        <Label
                                                            key={value}
                                                            htmlFor={`deposit-type-${value}`}
                                                            className={`relative flex w-full cursor-pointer items-start gap-3 rounded-lg border border-input p-4 shadow-sm shadow-black/5 duration-300 ease-in-out hover:bg-secondary/60 has-[[data-state=checked]]:border-ring has-[[data-state=checked]]:bg-secondary/80 ${
                                                                disabled
                                                                    ? 'cursor-not-allowed opacity-50'
                                                                    : ''
                                                            }`}
                                                        >
                                                            <RadioGroupItem
                                                                value={value}
                                                                id={`deposit-type-${value}`}
                                                                aria-describedby={`deposit-type-${value}-description`}
                                                                className="order-1"
                                                                disabled={
                                                                    disabled
                                                                }
                                                            />
                                                            <div className="grow">
                                                                <div className="flex items-center gap-2">
                                                                    <span className="text-lg">
                                                                        <SavingsIcon />
                                                                    </span>
                                                                    {label}
                                                                </div>
                                                            </div>
                                                        </Label>
                                                    )
                                                )}
                                            </RadioGroup>
                                        </>
                                    )
                                }}
                            />
                        )}

                        <FormFieldWrapper
                            control={form.control}
                            name="transactionType"
                            label="Payment Method"
                            render={({ field }) => (
                                <div className="relative w-full">
                                    <TransactionPaymentTypesPicker
                                        value={field.value}
                                        onSelect={(type) =>
                                            field.onChange(type.id)
                                        }
                                        defaultValue={DEFAULT_TRANSACTION_TYPE}
                                        leftIcon={
                                            <MoneyIcon className="text-foreground" />
                                        }
                                        placeholder="Select a payment method"
                                    />
                                    <span className="absolute left-4 top-1/2 -translate-y-1/2 text-lg font-bold text-primary"></span>
                                </div>
                            )}
                        />
                        <FormFieldWrapper
                            control={form.control}
                            name="accountsId"
                            label="Select Account"
                            render={({ field }) => (
                                <AccountsPicker
                                    value={field.value}
                                    onSelect={(account) => {
                                        field.onChange(account.id)
                                        setSelectedAccounts({
                                            ...account,
                                            noGracePeriodDaily:
                                                account.noGracePeriodDaily ??
                                                false,
                                        })
                                    }}
                                    leftIcon={
                                        <AccountsIcon className="text-foreground" />
                                    }
                                    placeholder="Choose an account"
                                />
                            )}
                        />
                        <FormFieldWrapper
                            className="md:col-span-2"
                            control={form.control}
                            name="notes"
                            label="Additional Notes"
                            render={({ field }) => (
                                <Textarea
                                    {...field}
                                    placeholder="Add any extra details or comments (optional)"
                                    className="resize-none"
                                />
                            )}
                        />
                        {selectedAccounts?.print && (
                            <div className="md:col-span-2">
                                <FormField
                                    name="isPrinted"
                                    control={form.control}
                                    render={({ field }) => {
                                        const id = 'allowPrinting'
                                        return (
                                            <FormItem className="flex flex-col">
                                                <FormControl>
                                                    <div className="relative flex w-full items-start gap-2 rounded-lg border border-input shadow-sm shadow-black/5 duration-300 ease-in-out hover:bg-secondary/60 has-[[data-state=checked]]:border-ring has-[[data-state=checked]]:bg-secondary/80">
                                                        <Checkbox
                                                            id={id}
                                                            checked={
                                                                field.value
                                                            }
                                                            onCheckedChange={
                                                                field.onChange
                                                            }
                                                            className="absolute right-4 top-4"
                                                            aria-describedby={`${id}-description`}
                                                        />
                                                        <Label
                                                            htmlFor={id}
                                                            className="flex grow cursor-pointer items-center gap-3 p-4"
                                                        >
                                                            <div className="grid gap-2">
                                                                <p>
                                                                    Enable
                                                                    Printing{' '}
                                                                    <span className="text-xs font-normal leading-[inherit] text-muted-foreground">
                                                                        (Check
                                                                        this box
                                                                        if you
                                                                        need a
                                                                        printed
                                                                        copy)
                                                                    </span>
                                                                </p>
                                                            </div>
                                                        </Label>
                                                    </div>
                                                </FormControl>
                                                <FormMessage />
                                            </FormItem>
                                        )
                                    }}
                                />
                            </div>
                        )}
                    </div>
                    <FormErrorMessage errorMessage={errorMessage} />
                    <div className="col-span-2">
                        <Separator className="my-2 sm:my-4" />
                        <div className="flex items-center justify-end gap-x-2">
                            <Button
                                type="button"
                                variant="ghost"
                                onClick={() => resetForm()}
                                className="w-full self-end px-8 sm:w-fit"
                            >
                                Reset
                            </Button>
                            <Button
                                disabled={isPendingCreatePayments}
                                type="submit"
                                id="create-payment"
                                className="w-full self-end px-8 sm:w-fit"
                            >
                                create
                            </Button>
                        </div>
                    </div>
                </form>
            </Form>
        </>
    )
}

export const TransactionPaymentEntryFormModal = ({
    description = 'Fill out the form to add a new transaction.',
    className,
    onOpenChange,
    formProps,
    ...props
}: IModalProps & {
    formProps?: Omit<IPaymentFormProps, 'className' | 'form' | 'resetForm'>
}) => {
    const { onOpen } = useConfirmModalStore()
    const { transactionType } = usePaymentsModalStore()

    const form = useForm<TPaymentFormValues>({
        resolver: zodResolver(paymentsEntrySchema),
        mode: 'onSubmit',
        defaultValues: {
            amount: 0,
            transactionType: DEFAULT_TRANSACTION_TYPE,
            accountsId: '',
            notes: '',
            isPrinted: false,
        },
    })

    const resetForm = () => {
        form.reset({
            amount: 0,
            transactionType: DEFAULT_TRANSACTION_TYPE,
            accountsId: '',
            notes: '',
            isPrinted: false,
        })
    }

    const onPaymentsEntryModalClose = (isOpen: boolean) => {
        const { isDirty } = form.formState

        if (!isOpen) {
            if (!isDirty) {
                onOpenChange?.(false)
            } else {
                onOpen({
                    title: `Discard ${capitalize(transactionType ?? '')}`,
                    description:
                        'Are you sure you want to discard this transaction?',
                    onConfirm: () => {
                        resetForm()
                        onOpenChange?.(false)
                    },
                    confirmString: 'Discard',
                })
            }
        }
    }

    return (
        <Modal
            title={`Create ${capitalize(transactionType ?? '')}`}
            description={description}
            className={cn('w-[44rem]', className)}
            onOpenChange={onPaymentsEntryModalClose}
            {...props}
        >
            <PaymentsEntryFormModal
                {...formProps}
                form={form}
                resetForm={resetForm}
                selectedMemberId={formProps?.selectedMemberId ?? ''}
                onSuccess={(createdData) => {
                    formProps?.onSuccess?.(createdData)
                    onOpenChange?.(false)
                }}
            />
        </Modal>
    )
}

export default TransactionPaymentEntryFormModal
