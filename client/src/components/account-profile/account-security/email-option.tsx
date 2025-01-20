import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { useCallback, useState } from 'react'
import { useMutation } from '@tanstack/react-query'

import {
    CheckIcon,
    CloseIcon,
    BadgeCheckFillIcon,
    BadgeMinusFillIcon,
} from '@/components/icons'
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
import FormErrorMessage from '@/components/ui/form-error-message'
import VerifyContactBar from '../verify-notice/verify-contact-bar'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { cn } from '@/lib'
import { withCatchAsync } from '@/utils'
import { emailSchema } from '@/validations/common'
import { serverRequestErrExtractor } from '@/helpers'
import { zodResolver } from '@hookform/resolvers/zod'
import { IChangeEmailRequest, IUserData } from '@/server/types'
import ProfileService from '@/horizon-corp/services/auth/ProfileService'

interface Props {
    email: string
    verified: boolean
    onSave: (newUserData: IUserData) => void
}

const emailOptionSchema = z.object({
    email: emailSchema,
})

const EmailOption = ({ email, verified, onSave }: Props) => {
    const [pwdModalState, setPwdModalState] = useState(false)
    const { isPending, mutate: saveEmail } = useMutation<
        IUserData,
        string,
        IChangeEmailRequest
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
        reValidateMode: 'onChange',
        values: { email },
    })

    const hasChanges = form.formState.isDirty

    const firstError = Object.values(form.formState.errors)[0]?.message

    const handleReset = useCallback(() => {
        if (!email) return

        form.reset()
        form.setValue('email', email)
    }, [email, form])

    return (
        <div className={cn('space-y-0', hasChanges && 'space-y-2')}>
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
                                            <span className="inline-flex items-center text-xs text-foreground/50">
                                                {!verified ? (
                                                    <>
                                                        <BadgeMinusFillIcon className="mr-1 inline-block text-orange-500" />{' '}
                                                        Not verified
                                                    </>
                                                ) : (
                                                    <>
                                                        <BadgeCheckFillIcon className="mr-1 inline-block text-primary" />{' '}
                                                        Verified
                                                    </>
                                                )}
                                            </span>
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
                                    <FormErrorMessage
                                        errorMessage={firstError}
                                    />
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

            {!verified && (
                <VerifyContactBar
                    verifyMode="email"
                    key="verify-bar-email"
                    onSuccess={onSave}
                />
            )}
        </div>
    )
}

export default EmailOption
