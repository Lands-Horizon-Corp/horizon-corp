import z from 'zod'
import { PASSWORD_MIN_LENGTH } from '../constants'

export const emailSchema = z
    .string({ required_error: 'Email is required' })
    .email('Invalid email')

const AccountTypes = ['Member', 'Owner', 'Admin', 'Employee'] as const

export const userAccountTypeSchema = z.enum(AccountTypes, {
    required_error: 'Account type is required',
    message: `Valid options are ${AccountTypes.join(',')}`,
    invalid_type_error: `Valid options are ${AccountTypes.join(',')}`,
})

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
