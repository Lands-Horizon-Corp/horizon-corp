import z from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { Form } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import Modal, { IModalProps } from '@/components/modals/modal'
import FormErrorMessage from '@/components/ui/form-error-message'
import FormFieldWrapper from '@/components/ui/form-field-wrapper'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { IMemberOccupationRequest, TEntityId } from '@/server/types'
import {
    useCreateMemberOccupation,
    useUpdateMemberOccupation,
} from '@/hooks/api-hooks/member/use-member-occupation'

export const createMemberOccupationSchema = z.object({
    name: z.string().min(1, { message: 'Name is required' }).trim(),
    description: z
        .string()
        .min(1, { message: 'Description is required' })
        .trim(),
})

type TMemberOccupationForm = z.infer<typeof createMemberOccupationSchema>

export interface IMemberOccupationCreateUpdateFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IMemberOccupationRequest>, unknown, string> {
    memberOccupationId?: TEntityId
}

const MemberOccupationCreateUpdateForm = ({
    memberOccupationId,
    readOnly,
    className,
    defaultValues,
    onError,
    onSuccess,
}: IMemberOccupationCreateUpdateFormProps) => {
    const isUpdateMode = Boolean(memberOccupationId)

    const form = useForm<TMemberOccupationForm>({
        resolver: zodResolver(createMemberOccupationSchema),
        reValidateMode: 'onChange',
        mode: 'onSubmit',
        defaultValues: {
            name: '',
            description: '',
            ...defaultValues,
        },
    })

    const {
        error: createError,
        isPending: isCreating,
        mutate: createMemberOccupation,
    } = useCreateMemberOccupation({ onSuccess, onError })

    const {
        error: updateError,
        isPending: isUpdating,
        mutate: updateMemberOccupation,
    } = useUpdateMemberOccupation({ onSuccess, onError })

    const onSubmit = (formData: TMemberOccupationForm) => {
        if (isUpdateMode && memberOccupationId) {
            updateMemberOccupation({
                occupationId: memberOccupationId,
                data: formData,
            })
        } else {
            createMemberOccupation(formData)
        }
    }

    const combinedError = createError || updateError

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit(onSubmit)}
                className={cn('flex w-full flex-col gap-y-4', className)}
            >
                <fieldset
                    disabled={isCreating || isUpdating || readOnly}
                    className="grid gap-x-6 gap-y-4 sm:gap-y-3"
                >
                    <fieldset className="space-y-3">
                        <FormFieldWrapper
                            control={form.control}
                            name="name"
                            label="Name"
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    id={field.name}
                                    placeholder="Member Occupation Name"
                                    autoComplete="member-occupation-name"
                                    disabled={
                                        isCreating || isUpdating || readOnly
                                    }
                                />
                            )}
                        />

                        <FormFieldWrapper
                            control={form.control}
                            name="description"
                            label="Description"
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    id={field.name}
                                    placeholder="Description"
                                    autoComplete="member-occupation-description"
                                    disabled={
                                        isCreating || isUpdating || readOnly
                                    }
                                />
                            )}
                        />
                    </fieldset>
                </fieldset>

                <FormErrorMessage errorMessage={combinedError} />

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
                            disabled={isCreating || isUpdating}
                            className="w-full self-end px-8 sm:w-fit"
                        >
                            {isCreating || isUpdating ? (
                                <LoadingSpinner />
                            ) : isUpdateMode ? (
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

export const MemberOccupationCreateUpdateFormModal = ({
    title = 'Create Member Occupation',
    description = 'Fill out the form to add a new member occupation.',
    className,
    formProps,
    ...props
}: IModalProps & {
    formProps?: Omit<IMemberOccupationCreateUpdateFormProps, 'className'>
}) => {
    return (
        <Modal
            title={title}
            description={description}
            className={cn('', className)}
            {...props}
        >
            <MemberOccupationCreateUpdateForm {...formProps} />
        </Modal>
    )
}

export default MemberOccupationCreateUpdateForm
