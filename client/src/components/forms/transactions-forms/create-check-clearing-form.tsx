import { useForm, UseFormReturn } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'
import { cn } from '@/lib'

import { Form } from '@/components/ui/form'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Separator } from '@/components/ui/separator'

import FormErrorMessage from '@/components/ui/form-error-message'
import Modal, { IModalProps } from '@/components/modals/modal'
import { toast } from 'sonner'

import { TCheckClearingSchema } from '@/validations/check-bank/check-clearing-schema'
import BankPicker from '@/components/pickers/bank-picker'
import InputDatePicker from '@/components/date-time-pickers/input-date-picker'
import { useCreateCheckClearing } from '@/hooks/api-hooks/transactions/use-check-clearing'
import { ICheckClearingRequest } from '@/server/types/transactions/check-clearing'
import FileUploader from '@/components/ui/file-uploader'
import FormFieldWrapper from '@/components/ui/form-field-wrapper'
import { useShortcut } from '@/components/use-shorcuts'

import { useMemberPickerStore } from '@/store/use-member-picker-state-store'
import {
    usePaymentsDataStore,
    usePaymentsModalStore,
} from '@/store/transaction/payments-entry-store'

import { IPaymentsEntry } from '@/server/types/transactions/payments-entry'
import { TPaymentFormValues, TRANSACTION_TYPE } from './payments-entry-form'

import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableFooter,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/ui/table'
import { useState } from 'react'
import { useSinglePictureUpload } from '@/hooks/api-hooks/use-media-resource'
import { IBankResponse } from '@/server/types/bank'
import { ScrollArea } from '@/components/ui/scroll-area'
import {
    base64ImagetoFile,
    capitalize,
    commaSeparators,
    formatNumberOnBlur,
    isValidDecimalInput,
    sanitizeNumberInput,
} from '@/helpers'
import { TrashIcon } from '@/components/icons'
import { Progress } from '@/components/ui/progress'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import { TEntityId } from '@/server'
import useConfirmModalStore from '@/store/confirm-modal-store'

type TBankFormValues = z.infer<typeof TCheckClearingSchema>

type DepositCheckClearingListType = Omit<TBankFormValues, 'bankId'> & {
    bank?: IBankResponse
}

interface IClearingFormProps {
    readOnly?: boolean
    disabledFields?: (keyof TBankFormValues)[]
    onSuccess?: (data: ICheckClearingRequest[]) => void
    onError?: (error: string) => void
    resetPaymentsForm?: () => void
    values?: TPaymentFormValues
    amount?: number
    form: UseFormReturn<{
        checkNo: string
        bankId: TEntityId
        checkDate: Date
        picture?: string
        amount?: number
    }>
}

const CheckBankForm = ({
    onSuccess,
    resetPaymentsForm,
    values,
    amount,
}: IClearingFormProps) => {
    const { isOpen } = useMemberPickerStore()
    const [depositCheckClearingList, setDepositCheckClearingList] = useState<
        DepositCheckClearingListType[]
    >([])
    const [newAmount, setnewAmount] = useState<number>(amount ?? 0)
    const [selectedBank, setSelectedBank] = useState<IBankResponse>()
    const [uploadMediaProgress, setUploadCheckMediaProgress] =
        useState<number>(0)

    const {
        selectedPayments,
        setSelectedPayments,
        ORNumber,
        selectedMember,
        selectedAccounts,
    } = usePaymentsDataStore()

    const {
        setOpenPaymentsEntryModal,
        transactionType,
        setOpenCheckClearingFormModal,
    } = usePaymentsModalStore()

    const isDeposit = transactionType === TRANSACTION_TYPE.deposit

    const {
        data: uploadedPhoto,
        isPending: isUploadingPhoto,
        mutateAsync: uploadPhoto,
    } = useSinglePictureUpload({
        onUploadProgressChange: (progress) =>
            setUploadCheckMediaProgress(progress),
    })

    const form = useForm<TBankFormValues>({
        resolver: zodResolver(TCheckClearingSchema),
        defaultValues: {
            checkNo: '',
            bankId: '',
            checkDate: new Date(),
            picture: '',
            amount: amount,
        },
    })

    const { mutateAsync, error, isPending } = useCreateCheckClearing({
        onSuccess,
    })

    const handleSubmit = form.handleSubmit(async (data) => {
        try {
            if (
                !values ||
                !values.amount ||
                !values.transactionType ||
                !selectedMember?.id ||
                !selectedAccounts
            ) {
                toast.warning('Incomplete payment data. Skipping submit.')
                return
            }

            const newEntry: IPaymentsEntry = {
                ...values,
                ORNumber,
                memberId: selectedMember.id,
                account: selectedAccounts,
                isPrinted: values.isPrinted ?? false,
            }

            //deposit thing
            if (isDeposit && data.amount) {
                const formattedData = {
                    ...data,
                    checkDate: data.checkDate.toISOString(),
                }

                if (data.amount > newAmount) {
                    toast.error(
                        'Invalid Input: The input amount must be less than the deposited amount!'
                    )
                    return
                }
                const updatedDepositList = [
                    ...depositCheckClearingList,
                    {
                        ...formattedData,
                        checkDate: new Date(formattedData.checkDate),
                        bank: selectedBank,
                    },
                ]

                setDepositCheckClearingList(updatedDepositList)

                const remainingAmount = Math.max(0, newAmount - data.amount)

                const sanitizedAmount =
                    parseFloat(
                        sanitizeNumberInput(remainingAmount.toString())
                    ) || 0

                setnewAmount(sanitizedAmount)

                form.setValue('amount', parseFloat(sanitizedAmount.toFixed(2)))

                if (sanitizedAmount === 0) {
                    const photoFile = base64ImagetoFile(
                        data.picture ?? '',
                        'logo.jpg'
                    ) as File

                    await uploadPhoto(photoFile)

                    const requestData = {
                        ...formattedData,
                        picture: uploadedPhoto?.id ?? '',
                    }

                    await mutateAsync([requestData])
                    setOpenCheckClearingFormModal(false)
                    form.reset()
                }
            } else {
                setSelectedPayments([...selectedPayments, newEntry])
                toast.success('Added payments entry')

                await mutateAsync([
                    {
                        ...data,
                        checkDate: data.checkDate.toISOString(),
                    },
                ])

                resetPaymentsForm?.()
                form.reset()
                setOpenPaymentsEntryModal(false)
            }
        } catch (err) {
            console.error('Error submitting check clearing:', err)
        }
    })

    const handleDeleteCheckItem = (amount: number, itemToDelete: number) => {
        setDepositCheckClearingList((prev) =>
            prev.filter((_, index) => index !== itemToDelete)
        )
        const newAmountValue = form.getValues('amount') ?? 0 + amount
        form.setValue('amount', newAmountValue)
        setnewAmount(newAmountValue)
    }

    const totalAmount = depositCheckClearingList.reduce(
        (acc, entry) => acc + (entry.amount ?? 0),
        0
    )

    useShortcut('Enter', async (event) => {
        event.preventDefault()
        const activeElement = document.activeElement as HTMLElement
        console.log()
        if (
            activeElement.tagName === 'INPUT' ||
            activeElement.tagName === 'TEXTAREA'
        ) {
            return
        }

        const isValid = await form.trigger()
        if (isValid && !isOpen) {
            await handleSubmit()
            resetPaymentsForm?.()
        }
    })

    const hasDepositClearingList = depositCheckClearingList.length <= 0

    return (
        <Form {...form}>
            <form onSubmit={handleSubmit} className="w-full space-y-4">
                <div className="grid grid-cols-1 gap-3 md:grid-cols-4">
                    {isDeposit && (
                        <FormFieldWrapper
                            className=""
                            control={form.control}
                            name="amount"
                            label="Payment Amount"
                            render={({ field }) => (
                                <div className="relative w-full">
                                    <Input
                                        {...field}
                                        id={field.name}
                                        className={cn(
                                            'rounded-2xl !bg-background/70 pl-8 pr-10 text-lg font-bold text-primary placeholder:text-sm placeholder:font-normal placeholder:text-foreground/40'
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
                                    <span className="absolute left-4 top-1/2 -translate-y-1/2 text-lg font-bold text-primary after:content-['₱']" />
                                </div>
                            )}
                        />
                    )}
                    <FormFieldWrapper
                        className=""
                        control={form.control}
                        name="checkNo"
                        label="Check Number"
                        render={({ field }) => (
                            <Input
                                {...field}
                                onChange={(e) => {
                                    field.onChange(e.target.value.toUpperCase())
                                }}
                                className="!bg-background/70"
                                id={field.name}
                                placeholder="Enter the check number"
                                autoComplete="off"
                            />
                        )}
                    />
                    <FormFieldWrapper
                        control={form.control}
                        name="bankId"
                        label="Select Bank"
                        render={({ field }) => (
                            <BankPicker
                                ButtonClassName="!bg-background/70"
                                value={field.value}
                                onSelect={(bank) => {
                                    setSelectedBank(bank)
                                    field.onChange(bank.id)
                                }}
                                placeholder="Choose a bank"
                            />
                        )}
                    />
                    <FormFieldWrapper
                        control={form.control}
                        name="checkDate"
                        label="Check Issue Date"
                        render={({ field }) => (
                            <InputDatePicker
                                value={field.value}
                                onChange={field.onChange}
                                captionLayout="dropdown-buttons"
                                className="bg-background/70"
                                disabled={(date) => date > new Date()}
                            />
                        )}
                    />
                    <FormFieldWrapper
                        className="md:col-span-4"
                        control={form.control}
                        name="picture"
                        label="Upload Check Image"
                        render={({ field }) => (
                            <FileUploader
                                defaultPhotos={field.value ?? ''}
                                className="max-h-42 min-h-10 w-full"
                                maxFiles={1}
                                accept={{
                                    'image/png': ['.png'],
                                    'image/jpeg': ['.jpg', '.jpeg'],
                                }}
                                selectedPhotos={(selectedPhoto) => {
                                    field.onChange(selectedPhoto)
                                }}
                            />
                        )}
                    />
                    {isUploadingPhoto && (
                        <div className="md:col-span-4">
                            <Progress
                                value={uploadMediaProgress}
                                className="h-1"
                            />
                            <div className="flex items-center justify-center gap-x-1 text-center text-xs text-foreground/60">
                                <LoadingSpinner className="size-2" />
                                {isUploadingPhoto && 'uploading picture...'}
                            </div>
                        </div>
                    )}
                </div>
                {transactionType === TRANSACTION_TYPE.deposit && (
                    <ScrollArea className="max-h-[400px] rounded-md border md:col-span-4">
                        {!hasDepositClearingList && (
                            <div className="flex w-full justify-end">
                                <Button
                                    disabled={hasDepositClearingList}
                                    onClick={(e) => {
                                        e.preventDefault()
                                        setDepositCheckClearingList([])
                                        const sanitizedAmount =
                                            parseFloat(
                                                sanitizeNumberInput(
                                                    amount?.toString() ?? '0'
                                                )
                                            ) || 0
                                        form.setValue('amount', sanitizedAmount)
                                        setnewAmount(sanitizedAmount)
                                    }}
                                    variant={'ghost'}
                                    className={cn('hover:bg-background/20')}
                                >
                                    clear all
                                </Button>
                            </div>
                        )}
                        <Table className="overflow">
                            <TableCaption>
                                A list of submitted check deposits.
                            </TableCaption>
                            <TableHeader>
                                <TableRow>
                                    <TableHead>Check No.</TableHead>
                                    <TableHead>Bank</TableHead>
                                    <TableHead>Deposit Date</TableHead>
                                    <TableHead></TableHead>
                                    <TableHead className="text-right">
                                        Amount
                                    </TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {depositCheckClearingList.map((item, idx) => (
                                    <TableRow key={item.checkNo}>
                                        <TableCell className="font-medium">
                                            {item.checkNo}
                                        </TableCell>
                                        <TableCell>{item.bank?.name}</TableCell>
                                        <TableCell>
                                            {item.checkDate.toLocaleDateString(
                                                undefined,
                                                {
                                                    year: 'numeric',
                                                    month: 'short',
                                                    day: 'numeric',
                                                }
                                            )}
                                        </TableCell>
                                        <TableCell>
                                            <TrashIcon
                                                onClick={() => {
                                                    handleDeleteCheckItem(
                                                        item.amount ?? 0,
                                                        idx
                                                    )
                                                }}
                                                size={18}
                                                className="text-red-500"
                                            />
                                        </TableCell>
                                        <TableCell className="text-right font-bold text-green-600 dark:text-green-400">
                                            ₱{' '}
                                            {commaSeparators(
                                                (item.amount ?? 0).toFixed(2)
                                            )}
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                            <TableFooter>
                                <TableRow>
                                    <TableCell
                                        colSpan={4}
                                        className="font-bold"
                                    >
                                        Total
                                    </TableCell>
                                    <TableCell className="text-right font-bold text-green-600 dark:text-green-400">
                                        ₱{' '}
                                        {commaSeparators(
                                            totalAmount.toFixed(2)
                                        )}
                                    </TableCell>
                                </TableRow>
                            </TableFooter>
                        </Table>
                    </ScrollArea>
                )}
                {error && <FormErrorMessage errorMessage={error} />}
                <div>
                    <Separator className="my-2 sm:my-4" />
                    <div className="flex items-center justify-end gap-x-2">
                        <Button
                            type="button"
                            variant="ghost"
                            disabled={isPending}
                            onClick={() => {
                                form.reset()
                                setDepositCheckClearingList([])
                            }}
                            className="w-full sm:w-fit"
                        >
                            Reset
                        </Button>
                        <Button
                            type="submit"
                            disabled={isPending}
                            className="w-full sm:w-fit"
                        >
                            {isPending ? 'Processing...' : 'Save'}
                        </Button>
                    </div>
                </div>
            </form>
        </Form>
    )
}

export const TransactionCheckClearingFormModal = ({
    title = 'Check Clearing',
    description = 'Fill out the form to save Check.',
    className,
    formProps,
    onOpenChange,
    ...props
}: IModalProps & {
    formProps?: Omit<IClearingFormProps, 'className' | 'form'>
}) => {
    const { onOpen } = useConfirmModalStore()
    const { transactionType } = usePaymentsModalStore()

    const form = useForm<TBankFormValues>({
        resolver: zodResolver(TCheckClearingSchema),
        defaultValues: {
            checkNo: '',
            bankId: '',
            checkDate: new Date(),
            picture: '',
            amount: formProps?.amount,
        },
    })

    const resetForm = () => {
        form.reset({
            checkNo: '',
            bankId: '',
            checkDate: new Date(),
            picture: '',
            amount: formProps?.amount,
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
            title={title}
            description={description}
            onOpenChange={onPaymentsEntryModalClose}
            className={cn('max-w-[64rem]', className)}
            {...props}
        >
            <CheckBankForm
                amount={formProps?.amount}
                form={form}
                {...formProps}
                onSuccess={(createdData) => {
                    formProps?.onSuccess?.(createdData)
                    onOpenChange?.(false)
                }}
            />
        </Modal>
    )
}

export default TransactionCheckClearingFormModal
