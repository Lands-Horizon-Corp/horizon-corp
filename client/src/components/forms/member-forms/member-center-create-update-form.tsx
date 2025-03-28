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
import { IMemberCenterRequest, TEntityId } from '@/server/types'
import {
    useCreateMemberCenter,
    useUpdateMemberCenter,
} from '@/hooks/api-hooks/member/use-member-center'
import { createMemberCenterSchema } from '@/validations/member/member-center-schema'

type TMemberCenterForm = z.infer<typeof createMemberCenterSchema>

export interface IMemberCenterCreateUpdateFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IMemberCenterRequest>, unknown, string> {
    memberCenterId?: TEntityId
}

const MemberCenterCreateUpdateForm = ({
    memberCenterId,
    readOnly,
    className,
    defaultValues,
    onError,
    onSuccess,
}: IMemberCenterCreateUpdateFormProps) => {
    const isUpdateMode = Boolean(memberCenterId)

    const form = useForm<TMemberCenterForm>({
        resolver: zodResolver(createMemberCenterSchema),
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
        mutate: createMemberCenter,
    } = useCreateMemberCenter({ onSuccess, onError })

    const {
        error: updateError,
        isPending: isUpdating,
        mutate: updateMemberCenter,
    } = useUpdateMemberCenter({ onSuccess, onError })

    const onSubmit = (formData: TMemberCenterForm) => {
        if (isUpdateMode && memberCenterId) {
            updateMemberCenter({ memberCenterId, data: formData })
        } else {
            createMemberCenter(formData)
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
                                    placeholder="Member Center Name"
                                    autoComplete="member-center-name"
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
                                    autoComplete="member-center-description"
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

export const MemberCenterCreateUpdateFormModal = ({
    title = 'Create Member Center',
    description = 'Fill out the form to add a new member center.',
    className,
    formProps,
    ...props
}: IModalProps & {
    formProps?: Omit<IMemberCenterCreateUpdateFormProps, 'className'>
}) => {
    return (
        <Modal
            title={title}
            description={description}
            className={cn('', className)}
            {...props}
        >
            <MemberCenterCreateUpdateForm {...formProps} />
        </Modal>
    )
}

export default MemberCenterCreateUpdateForm
