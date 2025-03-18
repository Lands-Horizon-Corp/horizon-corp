import { useState } from 'react'
import { useForm, UseFormReturn } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'

import { Input } from '@/components/ui/input'
import { Checkbox } from '@/components/ui/checkbox'
import { Textarea } from '@/components/ui/textarea'
import FormErrorMessage from '@/components/ui/form-error-message'
import FormFieldWrapper from '@/components/ui/form-field-wrapper'
import {
    Form,
    FormField,
    FormItem,
    FormControl,
    FormMessage,
} from '@/components/ui/form'
import { Label } from '@/components/ui/label'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import { IForm } from '@/types/component/form'
import { Button } from '@/components/ui/button'

import Modal, { IModalProps } from '@/components/modals/modal'

import AccountsPicker from '@/components/pickers/accounts-picker'
import TransactionPaymentTypesPicker from '@/components/pickers/transaction-payment-types-picker'
import CheckClearingFormModal from '../check-bank-forms/create-check-clearing-form'

import { useCreatePaymentEntry } from '@/hooks/api-hooks/transactions/use-payments-entry'
import {
    DEFAULT_TRANSACTION_TYPE,
    paymentsEntrySchema,
} from '@/validations/transactions/payments-entry'

import {
    commaSeparators,
    formatNumberOnBlur,
    isValidDecimalInput,
    sanitizeNumberInput,
} from '@/helpers'
import { cn } from '@/lib'
import { TEntityId } from '@/server'
import { IPaymentsEntryRequest } from '@/server/types/transactions/payments-entry'
import { IBaseCompNoChild } from '@/types'
import { Separator } from '@/components/ui/separator'
import { AccountsIcon, MoneyIcon } from '@/components/icons'
import useConfirmModalStore from '@/store/confirm-modal-store'

const CHECK_BANK_LIST = ['trans-pay-002', 'trans-pay-003']

type TPaymentFormValues = z.infer<typeof paymentsEntrySchema>

export interface IPaymentFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IPaymentsEntryRequest>, unknown, string> {
    open: boolean
    onOpenChange: (open: boolean) => void
    selectedMemberId: TEntityId
    resetSelectedMember?: () => void
    form: UseFormReturn<
        {
            ORNumber: string
            amount: number
            accountsId: string
            transactionType: string
            isPrinted?: boolean | undefined
            notes?: string | undefined
        },
        any,
        undefined
    >
    resetForm: () => void
}

const PaymentsEntryFormModal = ({
    onSuccess,
    onError,
    selectedMemberId,
    resetSelectedMember,
    form,
    resetForm,
}: Omit<IPaymentFormProps, 'open' | 'onOpenChange'>) => {
    const [openCheckClearingFormModal, setOpenCheckClearingFormModal] =
        useState(false)

    const {
        mutateAsync: createPayment,
        isPending,
        error,
    } = useCreatePaymentEntry({
        onSuccess: (data) => {
            onSuccess?.(data)
        },
        onError,
    })

    const handleSubmit = async (values: TPaymentFormValues) => {
        if (CHECK_BANK_LIST.includes(values.transactionType)) {
            setOpenCheckClearingFormModal(true)
            return
        }
        await submitPayment(values)
    }

    const submitPayment = async (values: TPaymentFormValues) => {
        try {
            const payload: IPaymentsEntryRequest = {
                ...values,
                memberId: selectedMemberId,
                isPrinted: values.isPrinted ?? false,
            }
            await createPayment(payload)
            resetSelectedMember?.()
            resetForm()
        } catch (error) {
            console.error('Error submitting payment:', error)
        }
    }

    return (
        <>
            <CheckClearingFormModal
                formProps={{
                    onSubmitPaymentform: async () => {
                        await submitPayment({ ...form.getValues() })
                    },
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
                            name="ORNumber"
                            label="Official Receipt (OR) Number"
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    className={cn(
                                        'text-lg font-semibold placeholder:text-xs placeholder:font-normal placeholder:text-foreground/40'
                                    )}
                                    id={field.name}
                                    placeholder="Enter the OR Number"
                                    autoComplete="off"
                                />
                            )}
                        />
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
                                    onSelect={(account) =>
                                        field.onChange(account.id)
                                    }
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
                                                        checked={field.value}
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
                                                                Enable Printing{' '}
                                                                <span className="text-xs font-normal leading-[inherit] text-muted-foreground">
                                                                    (Check this
                                                                    box if you
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
                    </div>
                    <FormErrorMessage errorMessage={error} />
                    <div className="col-span-2">
                        <Separator className="my-2 sm:my-4" />
                        <div className="flex items-center justify-end gap-x-2">
                            <Button
                                type="button"
                                variant="ghost"
                                disabled={isPending}
                                onClick={() => resetForm()}
                                className="w-full self-end px-8 sm:w-fit"
                            >
                                Reset
                            </Button>
                            <Button
                                type="submit"
                                disabled={
                                    form.formState.isSubmitting || isPending
                                }
                                className="w-full self-end px-8 sm:w-fit"
                            >
                                {isPending ? <LoadingSpinner /> : 'Create'}
                            </Button>
                        </div>
                    </div>
                </form>
            </Form>
        </>
    )
}

export const TransactionPaymentEntryFormModal = ({
    title = 'Create Payment',
    description = 'Fill out the form to add a new transaction.',
    className,
    onOpenChange,
    formProps,
    ...props
}: IModalProps & {
    formProps?: Omit<IPaymentFormProps, 'className' | 'form' | 'resetForm'>
}) => {
    const { onOpen } = useConfirmModalStore()

    const form = useForm<TPaymentFormValues>({
        resolver: zodResolver(paymentsEntrySchema),
        mode: 'onSubmit',
        defaultValues: {
            ORNumber: '',
            amount: 0,
            transactionType: DEFAULT_TRANSACTION_TYPE,
            accountsId: '',
            notes: '',
            isPrinted: false,
        },
    })

    const resetForm = () => {
        form.reset({
            ORNumber: '',
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
                    title: 'Discard Payment',
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
            title={title}
            description={description}
            className={cn('', className)}
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
