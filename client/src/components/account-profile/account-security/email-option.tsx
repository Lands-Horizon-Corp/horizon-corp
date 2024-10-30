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
import PasswordInputModal from './password-input-modal'
import {
    CheckIcon,
    CloseIcon,
    PatchCheckIcon,
    PatchMinusIcon,
} from '@/components/icons'

import { withCatchAsync } from '@/lib'
import { serverRequestErrExtractor } from '@/helpers'
import { emailSchema } from '@/validations/common'
import { zodResolver } from '@hookform/resolvers/zod'
import { ChangeEmailRequest, UserData } from '@/horizon-corp/types'
import ProfileService from '@/horizon-corp/server/auth/ProfileService'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import { VerifiedIcon } from 'lucide-react'
import ActionTooltip from '@/components/action-tooltip'

interface Props {
    email: string
    verified: boolean
    onSave: (newUserData: UserData) => void
}

const emailOptionSchema = z.object({
    email: emailSchema,
})

const EmailOption = ({ email, verified, onSave }: Props) => {
    const [pwdModalState, setPwdModalState] = useState(false)
    const { isPending, mutate: saveEmail } = useMutation<
        UserData,
        string,
        ChangeEmailRequest
    >({
        mutationKey: ['account-security-email'],
        mutationFn: async (data) => {
            const [error, response] = await withCatchAsync(
                ProfileService.ChangeEmail(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            toast.success('Email has been saved.')

            onSave(response.data)
            return response.data
        },
    })

    const form = useForm<z.infer<typeof emailOptionSchema>>({
        resolver: zodResolver(emailOptionSchema),
        mode: 'onChange',
        defaultValues: { email },
    })

    const hasChanges = form.watch('email') !== email

    const handleReset = useCallback(() => {
        if (!email) return

        form.reset()
        form.setValue('email', email)
    }, [email, form])

    return (
        <>
            <PasswordInputModal
                state={pwdModalState}
                payloadData={form.getValues()}
                onClose={(state) => setPwdModalState(state)}
                onSubmit={(newPayload) => {
                    setPwdModalState(false)
                    saveEmail(newPayload)
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
                            name="email"
                            render={({ field }) => (
                                <FormItem className="grid w-full gap-y-2">
                                    <div>
                                        <FormLabel
                                            htmlFor={field.name}
                                            className="text-right text-sm font-normal text-foreground/80"
                                        >
                                            Email{' '}
                                        </FormLabel>
                                        <FormDescription className="text-xs">
                                            Will be used to contact you.
                                            Verification is needed after
                                            updating.
                                        </FormDescription>
                                    </div>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            placeholder="Email"
                                            autoComplete="off"
                                        />
                                    </FormControl>
                                    {!verified ? (
                                        <span className="text-xs text-foreground/60">
                                            <PatchMinusIcon className="inline-block text-orange-500" />{' '}
                                            email not verified
                                        </span>
                                    ) : (
                                        <span className="text-xs text-foreground/60">
                                            <PatchCheckIcon className="inline-block text-primary" />{' '}
                                            email verified
                                        </span>
                                    )}
                                    <div className="flex items-center justify-end gap-x-1">
                                        {hasChanges && !isPending && (
                                            <>
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
                                            </>
                                        )}
                                        {isPending && (
                                            <span className="text-xs text-foreground/70">
                                                <LoadingSpinner className="mr-1 inline-block size-3" />{' '}
                                                Saving...
                                            </span>
                                        )}
                                    </div>
                                </FormItem>
                            )}
                        />
                    </fieldset>
                </form>
            </Form>
        </>
    )
}

export default EmailOption
