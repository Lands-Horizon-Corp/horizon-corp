import z from 'zod'
import { AuthModeSchema } from '.'

export const signInFormSchema = z.object({
    email: z
        .string({ required_error: 'Email is required' })
        .email('Email must be valid'),
    username: z
        .string({ required_error: 'Username is required' })
        .min(1, 'User Name is required'),
    password: z.string({ required_error: 'Password is required' }),
    mode: AuthModeSchema,
})
