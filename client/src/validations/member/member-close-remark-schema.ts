import z from 'zod'
import { entityIdSchema } from '../common'

export const memberCloseRemarkSchema = z.object({
    id: entityIdSchema.optional(),
    membersProfileId: entityIdSchema,
    description: z.string().min(1, 'Description/Reason is required'),
    createdAt: z.string(),
})

export const memberCreateCloseRemarkSchema = z.object({
    membersProfileId: entityIdSchema,
    description: z.string().min(1, 'Description/Reason is required'),
})

export const memberCreateCloseRemarksSchema = z.object({
    remarks: z
        .array(memberCreateCloseRemarkSchema)
        .min(1, 'Atleast 1 close remark'),
})
