import z from 'zod'
import { PASSWORD_MIN_LENGTH } from '../constants'

export const emailSchema = z
    .string({ required_error: 'Email is required' })
    .email('Email must be valid')

export const memberTypeSchema = z.enum(
    ['Member', 'Owner', 'Admin', 'Employee'],
    {
        required_error: 'mode is required',
        message: 'Valid options are Member, Owner, Admin, Employee',
        invalid_type_error: 'Valid options are Member, Owner, Admin, Employee',
    }
)

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
