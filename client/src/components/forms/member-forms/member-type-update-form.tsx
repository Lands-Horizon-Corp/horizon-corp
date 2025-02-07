import z from 'zod'
import { Path, useForm } from 'react-hook-form'
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
import { IMemberTypeRequest, TEntityId } from '@/server/types'
import { useUpdateMemberType } from '@/hooks/api-hooks/member/use-member-type'
import { createMemberTypeSchema } from '@/validations/form-validation/member/member-type-schema'

type TMemberTypeForm = z.infer<typeof createMemberTypeSchema>

interface IMemberTypeFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IMemberTypeRequest>, unknown, unknown> {
    memberTypeId: TEntityId
}

const MemberTypeUpdateForm = ({
    memberTypeId,
    readOnly,
    className,
    hiddenFields,
    defaultValues,
    disabledFields,
    onError,
    onSuccess,
}: IMemberTypeFormProps) => {
    const form = useForm<TMemberTypeForm>({
        resolver: zodResolver(createMemberTypeSchema),
        reValidateMode: 'onChange',
        mode: 'onSubmit',
        defaultValues: {
            name: '',
            prefix: '',
            description: '',
            ...defaultValues,
        },
    })

    const {
        error,
        isPending: isUpdating,
        mutate: updateMemberType,
    } = useUpdateMemberType({ onSuccess, onError })

    const isDisabled = (field: Path<TMemberTypeForm>) =>
        readOnly || disabledFields?.includes(field) || false

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit((formData) =>
                    updateMemberType({ memberTypeId, data: formData })
                )}
                className={cn('flex w-full flex-col gap-y-4', className)}
            >
                <fieldset
                    disabled={isUpdating || readOnly}
                    className="grid gap-x-6 gap-y-4 sm:gap-y-3"
                >
                    <fieldset className="space-y-3">
                        <FormFieldWrapper
                            name="name"
                            label="Name"
                            control={form.control}
                            hiddenFields={hiddenFields}
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    id={field.name}
                                    placeholder="Member Type Name"
                                    autoComplete="member-type-name"
                                    disabled={isDisabled(field.name)}
                                />
                            )}
                        />

                        <FormFieldWrapper
                            name="prefix"
                            label="Prefix"
                            control={form.control}
                            hiddenFields={hiddenFields}
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    id={field.name}
                                    placeholder="Prefix"
                                    autoComplete="member-type-prefix"
                                    disabled={isDisabled(field.name)}
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
                                    autoComplete="member-type-description"
                                    disabled={isDisabled(field.name)}
                                />
                            )}
                        />
                    </fieldset>
                </fieldset>

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
                            disabled={isUpdating}
                            className="w-full self-end px-8 sm:w-fit"
                        >
                            {isUpdating ? <LoadingSpinner /> : 'Update'}
                        </Button>
                    </div>
                </div>
            </form>
        </Form>
    )
}

export const MemberTypeUpdateFormModal = ({
    title = 'Update Member Type',
    description = 'Fill out the form to update member type.',
    className,
    formProps,
    ...props
}: IModalProps & {
    formProps: Omit<IMemberTypeFormProps, 'className'>
}) => {
    return (
        <Modal
            title={title}
            description={description}
            className={cn('', className)}
            {...props}
        >
            <MemberTypeUpdateForm {...formProps} />
        </Modal>
    )
}

export default MemberTypeUpdateForm
