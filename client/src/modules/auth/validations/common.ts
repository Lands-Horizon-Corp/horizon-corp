import z from 'zod'
import { PASSWORD_MIN_LENGTH } from '@/constants'

export const contactNumberSchema = z
    .string()
    .min(1, 'Invalid contact')
    .max(11)
    .regex(/^\d+$/, 'Contact number must contain only numbers')

export const passwordSetupSchema = z.object({
    password: z
        .string({ required_error: 'Password is required' })
        .min(
            PASSWORD_MIN_LENGTH,
            `Password must atleast ${PASSWORD_MIN_LENGTH}`
        ),
    confirmPassword: z
        .string({ required_error: 'Confirm password' })
        .min(PASSWORD_MIN_LENGTH, `Password doesn't match`),
})
