import { z } from 'zod'

// Common Timestamp Schema
const TimeStampsSchema = z.object({
    createdAt: z.string(),
    updatedAt: z.string(),
    deletedAt: z.string().nullable().optional(),
})

// Base Schema for Request
export const AccountsComputationTypeRequestSchema = z.object({
    id: z.string(),
    companyId: z.string(),
    name: z.string().min(1, 'Name is required'),
    description: z.string().min(1, 'Description is required'),
})

// Resource Schema (Extends Request + Timestamps + Created/Updated By)
export const AccountsComputationTypeResourceSchema = AccountsComputationTypeRequestSchema.extend({
    createdBy: z.string(),
    updatedBy: z.string(),
}).merge(TimeStampsSchema)

// Paginated Response Schema
export const AccountsComputationTypePaginatedResourceSchema = z.object({
    data: z.array(AccountsComputationTypeResourceSchema),
    pageIndex: z.number(),
    totalPage: z.number(),
    pageSize: z.number(),
    totalSize: z.number(),
    pages: z.array(z.unknown()), // Using `z.unknown()` as a placeholder for `IPages` structure
})
