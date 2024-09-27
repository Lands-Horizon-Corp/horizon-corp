import z from 'zod'

import { memberTypeSchema } from '.'
import { PASSWORD_MIN_LENGTH, LETTERS_REGEX } from '../constants'

export const signUpFormSchema = z
    .object({
        email: z
            .string({ required_error: 'Email is required' })
            .email('Email must be valid'),
        username: z
            .string({ required_error: 'Username is required' })
            .min(1, 'User Name is required'),
        firstName: z
            .string({ required_error: 'First Name is required' })
            .min(1, 'First Name too short')
            .regex(LETTERS_REGEX, 'First Name must contain only letters'),
        middleName: z
            .string()
            .transform((val) => val || undefined)
            .optional()
            .refine(
                (val) => {
                    if (!val || val === '') return true
                    return LETTERS_REGEX.test(val)
                },
                {
                    message: 'First Name must contain only letters',
                }
            ),
        lastName: z
            .string({ required_error: 'Last Name is required' })
            .min(1, 'Last Name is required')
            .regex(LETTERS_REGEX, 'Last Name must contain only letters'),
        contactNumber: z
            .string()
            .min(1)
            .max(11)
            .regex(/^\d+$/, 'Contact number must contain only numbers'),
        password: z
            .string({ required_error: 'Password is required' })
            .min(
                PASSWORD_MIN_LENGTH,
                `Password must atleast ${PASSWORD_MIN_LENGTH} characters`
            ),
        confirmPassword: z
            .string({ required_error: 'Confirm password' })
            .min(PASSWORD_MIN_LENGTH, `Password doesn't match`),
        acceptTerms: z
            .boolean()
            .default(false)
            .refine(
                (val) => {
                    return val === true
                },
                {
                    message: 'You must accept the terms and conditions',
                }
            ),
        mode: memberTypeSchema,
    })
    .refine(({ password, confirmPassword }) => password === confirmPassword, {
        message: "Password doesn't match",
        path: ['confirm_password'],
    })
