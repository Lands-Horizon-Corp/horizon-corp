import z from 'zod'
import { memberTypeSchema, emailSchema } from '.'

export const signInFormSchema = z.object({
    email: emailSchema,
    username: z
        .string({ required_error: 'Username is required' })
        .min(1, 'User Name is required'),
    password: z
        .string({ required_error: 'Password is required' })
        .min(1, 'Password is empty'),
    mode: memberTypeSchema,
})
