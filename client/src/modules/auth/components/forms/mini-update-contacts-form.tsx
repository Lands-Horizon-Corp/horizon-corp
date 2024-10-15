import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import FormErrorMessage from '../form-error-message'
import LoadingCircle from '@/components/loader/loading-circle'
import { Form, FormControl, FormField, FormItem } from '@/components/ui/form'

import { cn } from '@/lib'
import { IAuthForm } from '@/types/auth/form-interface'
import { IBaseCompNoChild } from '@/types/component/base'
import { handleAxiosError } from '@/horizon-corp/helpers'
import useLoadingErrorState from '@/hooks/use-loading-error-state'
import { emailSchema, contactNumberSchema } from '../../validations/common'

const miniUpdateContactsSchema = z.object({
    email: emailSchema.optional(),
    contactNumber: contactNumberSchema.optional(),
})

type TMiniUpdateContactsForm = z.infer<typeof miniUpdateContactsSchema>

export interface IMiniUpdateContactsProps
    extends IBaseCompNoChild,
        IAuthForm<Partial<TMiniUpdateContactsForm>> {
    onCancel?: () => void
}

const MiniUpdateContactsForm = ({
    readOnly,
    className,
    defaultValues,
    onError,
    onCancel,
    // onSuccess,
}: IMiniUpdateContactsProps) => {
    const { loading, setLoading, error, setError } = useLoadingErrorState()

    const form = useForm({
        resolver: zodResolver(miniUpdateContactsSchema),
        reValidateMode: 'onChange',
        defaultValues,
    })

    const handleSubmit = async (_data: TMiniUpdateContactsForm) => {
        setError(null)
        setLoading(true)
        try {
            // const parsedData = await miniUpdateContactsSchema.parseAsync(data)
            // const response = await AccountServiceDesyo .update account(parsedData)
            // onSuccess?.(response.data)
            // toast.success("Contact saved!")
        } catch (e) {
            const errorMessage = handleAxiosError(e)
            onError?.(e)
            setError(errorMessage)
            toast.error(errorMessage)
        } finally {
            setLoading(false)
        }
    }

    const firstError = Object.values(form.formState.errors)[0]?.message

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit(handleSubmit)}
                className={cn('flex min-w-[380px] flex-col gap-y-4', className)}
            >
                <fieldset
                    disabled={readOnly || loading}
                    className="flex flex-col gap-y-4"
                >
                    <FormField
                        control={form.control}
                        name="email"
                        render={({ field }) => (
                            <FormItem className="flex flex-col items-center">
                                <FormControl>
                                    <Input
                                        {...field}
                                        id={field.name}
                                        placeholder="Email Address"
                                        autoComplete="off"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="contactNumber"
                        render={({ field }) => (
                            <FormItem className="flex flex-col items-center">
                                <FormControl>
                                    <Input
                                        {...field}
                                        id={field.name}
                                        placeholder="Contact Number"
                                        autoComplete="off"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />{' '}
                    <FormErrorMessage errorMessage={error} />
                </fieldset>
                <FormErrorMessage errorMessage={error || firstError} />
                <div className="flex justify-end gap-x-2">
                    <Button
                        onClick={(e) => {
                            e.preventDefault()
                            onCancel?.()
                        }}
                        variant="outline"
                        disabled={loading || readOnly}
                    >
                        Cancel
                    </Button>
                    <Button type="submit" disabled={loading || readOnly}>
                        {loading ? (
                            <LoadingCircle className="size-4" />
                        ) : (
                            'Save'
                        )}
                    </Button>
                </div>
            </form>
        </Form>
    )
}

export default MiniUpdateContactsForm
