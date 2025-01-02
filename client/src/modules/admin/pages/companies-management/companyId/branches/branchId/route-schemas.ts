import z from 'zod'

export const branchIdPathSchema = z.object({
    branchId: z.coerce.number({
        required_error: 'Branch ID is required',
        invalid_type_error: 'Invalid Branch ID',
    }),
})
