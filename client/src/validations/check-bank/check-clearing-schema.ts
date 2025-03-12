import { z } from 'zod'

export const TCheckClearingSchema = z.object({
    checkNo: z.string().min(1, 'check no. is required'),
    bankId: z.string().min(1, 'Bank is required'),
    checkDate: z.preprocess(
        (val) => (typeof val === 'string' ? new Date(val) : val),
        z.date()
    ),
    picture: z.string().optional(),
})
