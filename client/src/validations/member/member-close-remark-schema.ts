import z from 'zod'
import { entityIdSchema } from '../common'

export const memberCloseRemarkSchema = z.object({
    id: entityIdSchema.optional(),
    membersProfileId: entityIdSchema,
    description: z.string().min(1, 'Description/Reason is required'),
    createdAt: z.string(),
})
