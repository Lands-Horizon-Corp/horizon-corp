import z from 'zod'

import { NUMBER_LETTER_REGEX } from '../constants'

export const otpFormSchema = z.object({
    code: z
        .string()
        .min(6, 'Code must be 6 alphanumeric characters')
        .max(6, 'Code must be 6 alphanumeric characters')
        .regex(NUMBER_LETTER_REGEX, 'Code must be a number'),
})
