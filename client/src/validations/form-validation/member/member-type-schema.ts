import z from 'zod'

export const createMemberTypeSchema = z.object({
    name: z.string().min(1, 'Name is required'),
    prefix: z.string().min(1, 'Prefix is required').max(2, 'Max 2 characters'),
    description: z.string().min(1, 'Description is required'),
})
