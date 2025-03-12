import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'
import { base64ImagetoFile } from '@/helpers'
import { cn } from '@/lib'

import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from '@/components/ui/form'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Separator } from '@/components/ui/separator'
import FormErrorMessage from '@/components/ui/form-error-message'
import UserAvatar from '@/components/user-avatar'
import ActionTooltip from '@/components/action-tooltip'
import { PlusIcon, ReplaceIcon } from '@/components/icons'
import Modal, { IModalProps } from '@/components/modals/modal'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import { Progress } from '@/components/ui/progress'

import { IBankResponse } from '@/server/types/bank'
import { useCreateBank, useUpdateBank } from '@/hooks/api-hooks/use-bank'
import { useSinglePictureUpload } from '@/hooks/api-hooks/use-media-resource'
import { SinglePictureUploadModal } from '@/components/single-image-uploader/single-picture-uploader'
import { TBankSchema } from '@/validations/check-bank/create-bank-schema'

type TBankFormValues = z.infer<typeof TBankSchema>

interface IBankFormProps {
    bankId?: string
    readOnly?: boolean
    disabledFields?: (keyof TBankFormValues)[]
    onSuccess?: (data: IBankResponse) => void
    onError?: (error: string) => void
}

const CheckBankFormModal = ({ bankId, onSuccess, onError }: IBankFormProps) => {
    const [uploadMediaProgress, setUploadMediaProgress] = useState<number>(0)

    const form = useForm<TBankFormValues>({
        resolver: zodResolver(TBankSchema),
        defaultValues: {
            mediaId: '',
            name: '',
            description: '',
        },
    })

    const [openImagePicker, setOpenImagePicker] = useState(false)
    const createMutation = useCreateBank({ onSuccess })
    const updateMutation = useUpdateBank({ onSuccess, onError })

    const {
        data: uploadedPhoto,
        isPending: isUploadingPhoto,
        mutateAsync: uploadPhoto,
    } = useSinglePictureUpload({
        onUploadProgressChange: (progress) => setUploadMediaProgress(progress),
    })

    const onSubmit = form.handleSubmit(async (data) => {
        console.log(data)
        try {
            await uploadPhoto(
                base64ImagetoFile(data.mediaId, `logo.jpg`) as File
            )
            const requestData = { ...data, mediaId: uploadedPhoto?.id ?? '' }

            if (bankId) {
                await updateMutation.mutateAsync({ bankId, data: requestData })
            } else {
                await createMutation.mutateAsync(requestData)
                form.reset()
            }
        } catch (error) {
            console.error('Upload or mutation failed:', error)
        }
    })

    const { error, isPending } = bankId ? updateMutation : createMutation

    return (
        <Form {...form}>
            <form onSubmit={onSubmit} className="w-full space-y-4">
                <div className="flex w-full gap-x-5">
                    <div className="grow">
                        <FormField
                            control={form.control}
                            name="name"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Name</FormLabel>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            placeholder="Enter bank name"
                                        />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="description"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Description</FormLabel>
                                    <FormControl>
                                        <Textarea {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    </div>
                    <div className="grow">
                        <FormField
                            control={form.control}
                            name="mediaId"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Logo</FormLabel>
                                    <FormControl>
                                        <div className="relative mx-auto size-fit">
                                            <SinglePictureUploadModal
                                                open={openImagePicker}
                                                onOpenChange={
                                                    setOpenImagePicker
                                                }
                                                onPhotoChoose={(newImage) => {
                                                    console.log(
                                                        'newImage',
                                                        newImage
                                                    )
                                                    field.onChange(newImage)
                                                }}
                                                defaultImage={field.value ?? ''}
                                            />
                                            <div></div>
                                            <UserAvatar
                                                src={field.value ?? ''}
                                                className="size-48"
                                            />
                                            <ActionTooltip
                                                tooltipContent={
                                                    field.value
                                                        ? 'Replace'
                                                        : 'Insert'
                                                }
                                                align="center"
                                                side="right"
                                            >
                                                <Button
                                                    variant="secondary"
                                                    onClick={(e) => {
                                                        e.preventDefault()
                                                        setOpenImagePicker(true)
                                                    }}
                                                    className="absolute bottom-2 right-2 size-fit w-fit rounded-full border border-transparent p-1"
                                                >
                                                    {field.value ? (
                                                        <ReplaceIcon />
                                                    ) : (
                                                        <PlusIcon />
                                                    )}
                                                </Button>
                                            </ActionTooltip>
                                        </div>
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    </div>
                </div>
                {isUploadingPhoto && (
                    <>
                        <Progress value={uploadMediaProgress} className="h-1" />
                        <div className="flex items-center justify-center gap-x-1 text-center text-xs text-foreground/60">
                            <LoadingSpinner className="size-2" />
                            {isUploadingPhoto && 'uploading picture...'}
                        </div>
                    </>
                )}
                <FormErrorMessage errorMessage={error} />
                <div>
                    <Separator className="my-2 sm:my-4" />
                    <div className="flex items-center justify-end gap-x-2">
                        <Button
                            type="button"
                            variant="ghost"
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
                            {isPending
                                ? 'Processing...'
                                : bankId
                                  ? 'Update'
                                  : 'Create'}
                        </Button>
                    </div>
                </div>
            </form>
        </Form>
    )
}

export const TransactionCheckBankCreateUpdateFormModal = ({
    title = 'Create Check Bank',
    description = 'Fill out the form to add a new Check bank.',
    className,
    formProps,
    ...props
}: IModalProps & {
    formProps?: Omit<IBankFormProps, 'className'>
}) => {
    return (
        <Modal
            title={title}
            description={description}
            className={cn('!min-w-[50rem]', className)}
            {...props}
        >
            <CheckBankFormModal
                {...formProps}
                onSuccess={(createdData) => {
                    formProps?.onSuccess?.(createdData)
                    props.onOpenChange?.(false)
                }}
            />
        </Modal>
    )
}

export default TransactionCheckBankCreateUpdateFormModal
