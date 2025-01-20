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
import { CheckIcon, CloseIcon } from '@/components/icons'
import FormErrorMessage from '@/components/ui/form-error-message'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import { userNameSchema } from '@/validations/common'
import { zodResolver } from '@hookform/resolvers/zod'
import { IChangeUsernameRequest, IUserData } from '@/horizon-corp/types'
import ProfileService from '@/horizon-corp/services/auth/ProfileService'

interface Props {
    username: string
    onSave: (newUserData: IUserData) => void
}

const userNameOptionSchema = z.object({
    username: userNameSchema,
})

const UsernameOption = ({ username, onSave }: Props) => {
    const [pwdModalState, setPwdModalState] = useState(false)

    const form = useForm<z.infer<typeof userNameOptionSchema>>({
        resolver: zodResolver(userNameOptionSchema),
        mode: 'onChange',
        values: { username },
    })

    const hasChanges = form.formState.isDirty

    const firstError = Object.values(form.formState.errors)[0]?.message

    const handleReset = useCallback(() => {
        if (!username) return

        form.reset()
        form.setValue('username', username)
    }, [username, form])

    const { isPending, mutate: saveUsername } = useMutation<
        IUserData,
        string,
        IChangeUsernameRequest
    >({
        mutationKey: ['account-security-username'],
        mutationFn: async (data) => {
            const [error, response] = await withCatchAsync(
                ProfileService.ChangeUsername(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            toast.success('Username has been saved.')

            onSave(response.data)

            handleReset()
            return response.data
        },
    })

    return (
        <>
            <PasswordInputModal
                state={pwdModalState}
                payloadData={form.getValues()}
                onClose={(state) => setPwdModalState(state)}
                onSubmit={(newPayload) => {
                    setPwdModalState(false)
                    saveUsername(newPayload)
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
                            name="username"
                            render={({ field }) => (
                                <FormItem className="grid w-full gap-y-2">
                                    <div>
                                        <FormLabel
                                            htmlFor={field.name}
                                            className="text-right text-sm font-normal text-foreground/80"
                                        >
                                            Username
                                        </FormLabel>
                                        <FormDescription className="text-xs">
                                            Will be displayed in your profile
                                        </FormDescription>
                                    </div>
                                    <FormControl>
                                        <Input
                                            {...field}
                                            id={field.name}
                                            placeholder="Username"
                                            autoComplete="off"
                                            hidden
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
        </>
    )
}

export default UsernameOption
