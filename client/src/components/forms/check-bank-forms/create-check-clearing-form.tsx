import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'

import { Form } from '@/components/ui/form'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Separator } from '@/components/ui/separator'

import FormErrorMessage from '@/components/ui/form-error-message'
import Modal, { IModalProps } from '@/components/modals/modal'

import { cn } from '@/lib'
import { TCheckClearingSchema } from '@/validations/check-bank/check-clearing-schema'
import BankPicker from '@/components/pickers/bank-picker'
import InputDatePicker from '@/components/date-time-pickers/input-date-picker'
import { useCreateCheckClearing } from '@/hooks/api-hooks/transactions/use-check-clearing'
import { ICheckClearingRequest } from '@/server/types/transactions/check-clearing'
import FileUploader from '@/components/ui/file-uploader'
import FormFieldWrapper from '@/components/ui/form-field-wrapper'

type TBankFormValues = z.infer<typeof TCheckClearingSchema>

interface IClearingFormProps {
    readOnly?: boolean
    disabledFields?: (keyof TBankFormValues)[]
    onSuccess?: (data: ICheckClearingRequest) => void
    onError?: (error: string) => void
    onSubmitPaymentform?: (e?: React.BaseSyntheticEvent) => Promise<void>
}

const CheckBankForm = ({
    onSuccess,
    onSubmitPaymentform,
}: IClearingFormProps) => {
    const form = useForm<TBankFormValues>({
        resolver: zodResolver(TCheckClearingSchema),
        defaultValues: {
            checkNo: '',
            bankId: '',
            checkDate: new Date(),
            picture: '',
        },
    })

    const { mutateAsync, error, isPending } = useCreateCheckClearing({
        onSuccess,
    })

    const handleSubmit = form.handleSubmit(async (data) => {
        try {
            await mutateAsync({
                ...data,
                checkDate: data.checkDate.toISOString(),
            })
            form.reset()
            if (onSubmitPaymentform) {
                await onSubmitPaymentform()
            }
        } catch (err) {
            console.error('Error submitting check clearing:', err)
        }
    })

    return (
        <Form {...form}>
            <form onSubmit={handleSubmit} className="w-full space-y-4">
                <div className="grid grid-cols-1 gap-3 md:grid-cols-2">
                    <FormFieldWrapper
                        className="md:col-span-2"
                        control={form.control}
                        name="checkNo"
                        label="Check Number"
                        render={({ field }) => (
                            <Input
                                {...field}
                                onChange={(e) => {
                                    field.onChange(e.target.value.toUpperCase())
                                }}
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
                                value={field.value}
                                onSelect={(bank) => field.onChange(bank.id)}
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
                                disabled={(date) => date > new Date()}
                            />
                        )}
                    />
                    <FormFieldWrapper
                        className="md:col-span-2"
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
                </div>
                {error && <FormErrorMessage errorMessage={error} />}
                <div>
                    <Separator className="my-2 sm:my-4" />
                    <div className="flex items-center justify-end gap-x-2">
                        <Button
                            type="button"
                            variant="ghost"
                            disabled={isPending}
                            onClick={() => form.reset()}
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
    ...props
}: IModalProps & {
    formProps?: Omit<IClearingFormProps, 'className'>
}) => {
    return (
        <Modal
            title={title}
            description={description}
            className={cn('', className)}
            {...props}
        >
            <CheckBankForm
                {...formProps}
                onSuccess={(createdData) => {
                    formProps?.onSuccess?.(createdData)
                    props.onOpenChange?.(false)
                }}
            />
        </Modal>
    )
}

export default TransactionCheckClearingFormModal
