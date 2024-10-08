import z from 'zod'

export const contactFormSchema = z.object({
    firstName: z
        .string({ required_error: 'First name is required' })
        .min(1, 'Last name is required'),
    lastName: z
        .string({ required_error: 'Last name is required' })
        .min(1, 'Last name is required'),
    email: z
        .string({ required_error: 'Email is required' })
        .email('Email must be valid'),
    contactNumber: z
        .string({ required_error: 'Contact number is required' })
        .min(10, 'Contact number must be at least 10 digits long')
        .max(11, 'Contact number must not exceed 11 digits')
        .regex(/^\d+$/, 'Contact number must contain only numbers'),
    message: z
        .string({ required_error: 'Message is required' })
        .min(20, 'Message must be at least 20 characters long')
        .max(100, 'Message must not exceed 100 characters'),
})
