import z from 'zod'
import { entityIdSchema } from '../common'

export const createMemberCenterSchema = z.object({
    id: entityIdSchema.optional(),
    name: z.string().min(1, 'Name is required'),
    description: z.string().min(1, 'Description is required'),
})
