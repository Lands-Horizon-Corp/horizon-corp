import z from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Path, useFieldArray, useForm } from 'react-hook-form'

import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { PlusIcon, XIcon } from '@/components/icons'
import { Separator } from '@/components/ui/separator'
import { Form, FormControl } from '@/components/ui/form'
import FormFieldWrapper from '@/components/ui/form-field-wrapper'

import { cn } from '@/lib'
import { TEntityId } from '@/server'
import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { useCloseMemberProfile } from '@/hooks/api-hooks/member/use-member-profile'
import { memberCreateCloseRemarksSchema } from '@/validations/member/member-close-remark-schema'

type TMemberCloseForm = z.infer<typeof memberCreateCloseRemarksSchema>

interface IMemberProfileCloseFormProps
    extends IBaseCompNoChild,
        IForm<Partial<TMemberCloseForm>, unknown, string> {
    profileId: TEntityId
}

const MemberProfileCloseForm = ({
    readOnly,
    className,
    profileId,
    hiddenFields,
    defaultValues,
    disabledFields,
    onSuccess,
}: IMemberProfileCloseFormProps) => {
    const form = useForm<TMemberCloseForm>({
        resolver: zodResolver(memberCreateCloseRemarksSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues: {
            remarks: [],
            ...defaultValues,
        },
    })

    const { mutate, isPending: isLoading } = useCloseMemberProfile({
        onSuccess,
    })

    const isDisabled = (field: Path<TMemberCloseForm>) =>
        readOnly || disabledFields?.includes(field) || false

    const {
        fields: remarksFields,
        append: addRemark,
        remove: removeRemark,
    } = useFieldArray({
        control: form.control,
        name: 'remarks',
    })

    const handleSubmit = ({ remarks }: TMemberCloseForm) => {
        mutate({ profileId, data: remarks })
    }

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit(handleSubmit)}
                className={cn('flex w-full flex-col gap-y-4', className)}
            >
                <fieldset
                    className="min-h-[60vh] gap-x-4 gap-y-4 space-y-5"
                    disabled={isLoading || readOnly}
                >
                    <FormFieldWrapper
                        name="remarks"
                        label="Address"
                        control={form.control}
                        hiddenFields={hiddenFields}
                        disabled={isDisabled('remarks')}
                        render={() => (
                            <>
                                <Separator />
                                <fieldset className="grid gap-4">
                                    {remarksFields.map((remarkField, index) => (
                                        <div
                                            key={remarkField.id}
                                            className="relative flex w-full flex-col gap-4 md:flex-row"
                                        >
                                            <FormFieldWrapper
                                                name={`remarks.${index}.description`}
                                                control={form.control}
                                                label="Label *"
                                                hiddenFields={hiddenFields}
                                                render={({ field }) => (
                                                    <FormControl>
                                                        <Input
                                                            {...field}
                                                            id={field.name}
                                                            placeholder="Enter close description/reason..."
                                                        />
                                                    </FormControl>
                                                )}
                                            />
                                            <Button
                                                size="icon"
                                                type="button"
                                                variant="secondary"
                                                onClick={() =>
                                                    removeRemark(index)
                                                }
                                                className="absolute right-0 top-0 size-fit self-center rounded-full p-1"
                                            >
                                                <XIcon className="size-4" />
                                            </Button>
                                        </div>
                                    ))}
                                </fieldset>
                                <Button
                                    size="sm"
                                    type="button"
                                    variant="secondary"
                                    className="!mt-2"
                                    onClick={() =>
                                        addRemark({
                                            membersProfileId: profileId,
                                            description: '',
                                        })
                                    }
                                >
                                    <PlusIcon className="mr-2" /> Add Close Reason
                                </Button>
                            </>
                        )}
                    />
                </fieldset>
            </form>
        </Form>
    )
}

export default MemberProfileCloseForm
