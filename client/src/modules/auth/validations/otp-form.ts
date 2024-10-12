import z from 'zod'

import { NUMBER_LETTER_REGEX } from '../constants'

export const otpFormSchema = z.object({
    otp : z
        .string()
        .min(6, 'OTP must be 6 alphanumeric characters')
        .max(6, 'OTP must be 6 alphanumeric characters')
        .regex(NUMBER_LETTER_REGEX, 'OTP must be valid alphanumeric characters'),
})
