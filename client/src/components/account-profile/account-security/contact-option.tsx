import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { useCallback, useState } from 'react'
import { useMutation } from '@tanstack/react-query'

import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import PasswordInputModal from './password-input-modal'
import { CheckIcon, CloseIcon } from '@/components/icons'

import { withCatchAsync } from '@/lib'
import { serverRequestErrExtractor } from '@/helpers'
import { contactNumberSchema } from '@/validations/common'
import { zodResolver } from '@hookform/resolvers/zod'
import { ChangeContactNumberRequest, UserData } from '@/horizon-corp/types'
import ProfileService from '@/horizon-corp/server/auth/ProfileService'

interface Props {
    contact: string
    onSave: (newUserData: UserData) => void
}

const contactOptionSchema = z.object({
    contactNumber: contactNumberSchema,
})

const ContactOption = ({ contact, onSave }: Props) => {
    const [pwdModalState, setPwdModalState] = useState(false)
    const { isPending, mutate: saveContact } = useMutation<
        UserData,
        string,
        ChangeContactNumberRequest
    >({
        mutationKey: ['account-security-contact'],
        mutationFn: async (data) => {
            const [error, response] = await withCatchAsync(
                ProfileService.ChangeContactNumber(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            toast.success("Contact number has been saved.")

            onSave(response.data)
            return response.data
        },
    })

    const form = useForm<z.infer<typeof contactOptionSchema>>({
        resolver: zodResolver(contactOptionSchema),
        mode: 'onChange',
        defaultValues: { contactNumber : contact },
    })

    const hasChanges = form.watch('contactNumber') !== contact

    const handleReset = useCallback(() => {
        if (!contact) return

        form.reset()
        form.setValue('contactNumber', contact)
    }, [contact, form])

    return (
        <>
            <PasswordInputModal
                state={pwdModalState}
                payloadData={form.getValues()}
                onClose={(state) => setPwdModalState(state)}
                onSubmit={(newPayload) => {
                    setPwdModalState(false)
                    saveContact(newPayload)
                }}
            />
            <Form {...form}>
                <form
                    className="w-full"
                    onSubmit={form.handleSubmit(() => {
                        if (!hasChanges) return
                        setPwdModalState(true)
                    })}
                >
                    <fieldset disabled={isPending}>
                        <FormField
                            control={form.control}
                            name="contactNumber"
                            render={({ field }) => (
                                <FormItem className="grid w-full gap-y-2">
                                    <div>
                                        <FormLabel
                                            htmlFor={field.name}
                                            className="text-right text-sm font-normal text-foreground/80"
                                        >
                                            Contact Number
                                        </FormLabel>
                                        <FormDescription className="text-xs">
                                            Will be used to contact you
                                        </FormDescription>
                                    </div>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            placeholder="Contact Number"
                                            autoComplete="off"
                                        />
                                    </FormControl>
                                    {hasChanges && (
                                        <div className="flex items-center gap-x-2">
                                            <Button
                                                size="sm"
                                                type="reset"
                                                variant="secondary"
                                                onClick={(e) => {
                                                    e.preventDefault()
                                                    handleReset()
                                                }}
                                            >
                                                <CloseIcon />
                                            </Button>
                                            <Button type="submit" size="sm">
                                                <CheckIcon />
                                            </Button>
                                        </div>
                                    )}
                                </FormItem>
                            )}
                        />
                    </fieldset>
                </form>
            </Form>
            <Separator />
        </>
    )
}

export default ContactOption
