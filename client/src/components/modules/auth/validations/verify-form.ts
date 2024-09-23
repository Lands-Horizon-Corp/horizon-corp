import z from 'zod'
import { NUMBER_LETTER_REGEX } from '../constants'

export const verifyFormSchema = z.object({
    code: z
        .string()
        .min(6, 'Code must be 6 digit long')
        .max(6, 'Code must be 6 digit long')
        .regex(NUMBER_LETTER_REGEX, "Code must be a number"),
})
