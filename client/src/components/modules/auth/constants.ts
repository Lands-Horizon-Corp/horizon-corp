export const PASSWORD_MIN_LENGTH = 8;

export const LETTERS_REGEX = /^[A-Za-z\s]+$/

import z from 'zod'

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