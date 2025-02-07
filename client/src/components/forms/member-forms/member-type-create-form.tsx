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
import { IMemberTypeRequest } from '@/server/types'
import { useCreateMemberType } from '@/hooks/api-hooks/member/use-member-type'
import { createMemberTypeSchema } from '@/validations/form-validation/member/member-type-schema'

type TMemberTypeCreateForm = z.infer<typeof createMemberTypeSchema>

interface IMemberTypeCreateFormProps
    extends IBaseCompNoChild,
        IForm<Partial<IMemberTypeRequest>, unknown, string> {}

const MemberTypeCreateForm = ({
    readOnly,
    className,
    defaultValues,
    onError,
    onSuccess,
}: IMemberTypeCreateFormProps) => {
    const form = useForm<TMemberTypeCreateForm>({
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
        isPending: isCreating,
        mutate: createMemberType,
    } = useCreateMemberType({ onSuccess, onError })

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit((formData) =>
                    createMemberType(formData)
                )}
                className={cn('flex w-full flex-col gap-y-4', className)}
            >
                <fieldset
                    disabled={isCreating || readOnly}
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
                                    placeholder="Member Type Name"
                                    autoComplete="member-type-name"
                                    disabled={isCreating || readOnly}
                                />
                            )}
                        />

                        <FormFieldWrapper
                            control={form.control}
                            name="prefix"
                            label="Prefix"
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    id={field.name}
                                    placeholder="Prefix"
                                    autoComplete="member-type-prefix"
                                    disabled={isCreating || readOnly}
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
                                    autoCapitalize="member-type-description"
                                    disabled={isCreating || readOnly}
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
                            disabled={isCreating}
                            className="w-full self-end px-8 sm:w-fit"
                        >
                            {isCreating ? <LoadingSpinner /> : 'Create'}
                        </Button>
                    </div>
                </div>
            </form>
        </Form>
    )
}

export const MemberTypeCreateFormModal = ({
    title = 'Create Member Type',
    description = 'Fill out the form to add a new member type.',
    className,
    formProps,
    ...props
}: IModalProps & {
    formProps?: Omit<IMemberTypeCreateFormProps, 'className'>
}) => {
    return (
        <Modal
            title={title}
            description={description}
            className={cn('', className)}
            {...props}
        >
            <MemberTypeCreateForm {...formProps} />
        </Modal>
    )
}

export default MemberTypeCreateForm
