import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
    Form,
    FormField,
    FormItem,
    FormLabel,
    FormControl,
    FormMessage,
    FormDescription,
} from '@/components/ui/form' // Ensure Form is from ShadCN
import { AccountRequestSchema } from '@/validations/form-validation/accounts-schema'

const formSchema = z.object({
    username: z.string().min(2, {
        message: 'Username must be at least 2 characters.',
    }),
})
type TAccountsCreateForm = z.infer<typeof AccountRequestSchema>

const AccountsCreateForm = () => {
    const form = useForm<TAccountsCreateForm>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            accountCode: '',
        },
    })

    function onSubmit(values: TAccountsCreateForm) {
        console.log(values)
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                <fieldset className="rounded-md border border-gray-300 p-4">
                    <legend className="px-2 text-lg font-semibold text-secondary-foreground">
                        Account Information
                    </legend>

                    <FormField
                        control={form.control}
                        name="accountCode"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Account Code</FormLabel>
                                <FormControl>
                                    <Input placeholder="shadcn" {...field} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                </fieldset>

                <Button type="submit">Submit</Button>
            </form>
        </Form>
    )
}
export default AccountsCreateForm
