import z from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { PASSWORD_MIN_LENGTH, LETTERS_REGEX } from '@/modules/auth/constants'
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'

interface Props {}

const signUpFormSchema = z.object({
    email: z.string().min(1, 'Email is required').email('Email must be valid'),
    username: z.string().min(1, 'User name is required'),
    first_name: z
        .string()
        .min(1, 'First Name is required')
        .regex(LETTERS_REGEX, 'First Name must contain only letters'),
    middle_name: z
        .string()
        .min(1, 'Middle Name is required')
        .regex(LETTERS_REGEX, 'First Name must contain only letters'),
    last_name: z
        .string()
        .min(1, 'Last Name is required')
        .regex(LETTERS_REGEX, 'Last Name must contain only letters'),
    contact_number: z.string().min(1).max(11),
    password: z
        .string()
        .min(
            PASSWORD_MIN_LENGTH,
            `Password must atleast ${PASSWORD_MIN_LENGTH}`
        ),
    confirm_password: z
        .string()
        .min(
            PASSWORD_MIN_LENGTH,
            `Password must atleast ${PASSWORD_MIN_LENGTH}`
        ),
})

const SignUpForm = ({}: Props) => {
    const form = useForm<z.infer<typeof signUpFormSchema>>({
        resolver: zodResolver(signUpFormSchema),
        defaultValues: {
            username: '',
        },
    })

    return (
        <Form {...form}>
            <form>
                <FormField
                    control={form.control}
                    name="username"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Username</FormLabel>
                            <FormControl>
                                <Input placeholder="shadcn" {...field} />
                            </FormControl>
                            <FormDescription>
                                This is your public display name.
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />
            </form>
        </Form>
    )
}

export default SignUpForm
