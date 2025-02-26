import { z } from 'zod'

export const TransactionTypeRequestSchema = z.object({
    id: z.string(),
    name: z.string().min(1, 'Name is required'),
    description: z.string().min(1, 'Description is required'),
    cheque_id: z.string().min(1, 'Cheque ID is required'),
})

