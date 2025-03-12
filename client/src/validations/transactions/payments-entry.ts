import { z } from 'zod'

export const DEFAULT_TRANSACTION_TYPE = 'trans-pay-001'

export const paymentsEntrySchema = z.object({
    ORNumber: z.string().min(1, 'OR Number is required'),
    amount: z.preprocess(
        (val) =>
            typeof val === 'string' ? parseFloat(val.replace(/,/g, '')) : val,
        z.number().min(0.01, 'Amount must be greater than 0')
    ),
    accountsId: z.string().min(1, 'Accounts is required'),
    isPrinted: z.boolean().optional(),
    notes: z.string().optional(),
    transactionType: z.string().default(DEFAULT_TRANSACTION_TYPE),
})
