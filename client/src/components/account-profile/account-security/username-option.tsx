import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { useMutation } from '@tanstack/react-query'
import { useCallback, useEffect, useState } from 'react'

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
import { userNameSchema } from '@/validations/common'
import { zodResolver } from '@hookform/resolvers/zod'
import { ChangeUsernameRequest, UserData } from '@/horizon-corp/types'
import ProfileService from '@/horizon-corp/server/auth/ProfileService'

interface Props {
    username: string
    onSave: (newUserData: UserData) => void
}

const userNameOptionSchema = z.object({
    username: userNameSchema,
})

const UsernameOption = ({ username, onSave }: Props) => {
    const [pwdModalState, setPwdModalState] = useState(false)
    const { isPending, mutate: saveUsername } = useMutation<
        UserData,
        string,
        ChangeUsernameRequest
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

            toast.success("Username has been saved.")

            onSave(response.data)
            return response.data
        },
    })

    const form = useForm<z.infer<typeof userNameOptionSchema>>({
        resolver: zodResolver(userNameOptionSchema),
        mode: 'onChange',
        defaultValues: { username: '' },
    })

    const hasChanges = form.watch('username') !== username

    const handleReset = useCallback(() => {
        if (!username) return

        form.reset()
        form.setValue('username', username)
    }, [username, form])

    useEffect(() => {
        handleReset()
    }, [handleReset])

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
            <Separator className="" />
        </>
    )
}

export default UsernameOption