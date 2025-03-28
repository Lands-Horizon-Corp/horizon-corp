import z from 'zod'

import { memberAssetsSchema } from './member-assets-schema'
import { memberIncomeSchema } from './member-income-schema'
import { memberAddressSchema } from './member-address-schema'
import { memberExpensesSchema } from './member-expenses-schema'
import { memberRecruitsSchema } from './member-recruits-schema'
import { entityIdSchema, mediaResourceSchema } from '../common'
import { memberDescriptionSchema } from './member-description-schema'
import { memberGovernmentBenefits } from './member-government-benefits'
import { memberJointAccountsSchema } from './member-joint-accounts-schema'
import { memberRelativeAccountsSchema } from './member-relative-accounts-schema'
import { memberContactReferencesSchema } from './member-contact-number-references-schema'
import { memberCenterSchema } from './member-center-schema'

export const createMemberProfileSchema = z.object({
    id: entityIdSchema.optional(),
    oldReferenceId: z.string().optional(),
    passbookNumber: z.string().optional(),

    middleName: z.string().optional(),
    lastName: z.string().min(1, 'Lastname is required'),
    firstName: z.string().min(1, 'Firstname is required'),
    suffix: z.string().max(15).optional(),

    notes: z.string().min(1, 'Notes are required'),
    description: z.string().min(1, 'Description is required'),
    contactNumber: z.string().min(1, 'Contact number is required'),
    civilStatus: z
        .enum(['Married', 'Single', 'Widowed', 'Separated', 'N/A'])
        .default('Single'),
    occupationId: entityIdSchema.optional(),
    businessAddress: z.string().optional(),
    businessContact: z.string().optional(),

    status: z.enum(['Pending', 'Verified', 'Not Allowed']).default('Pending'),
    isClosed: z.boolean(),

    isMutualFundMember: z.boolean().default(false),
    isMicroFinanceMember: z.boolean().default(false),

    memberCenterId: entityIdSchema.optional(),
    memberCenter: memberCenterSchema.optional(),

    mediaId: entityIdSchema.optional(),
    memberId: entityIdSchema.optional(),
    media: mediaResourceSchema.optional(),

    signatureMediaId: entityIdSchema.optional(),
    signatureMedia: mediaResourceSchema.optional(),

    memberTypeId: z
        .string()
        .min(1, 'Member Type is required')
        .uuid('Invalid member type'),
    branchId: entityIdSchema.optional(),
    memberGenderId: entityIdSchema.optional(),
    memberClassificationId: entityIdSchema.optional(),
    memberEducationalAttainmentId: entityIdSchema.optional(),

    memberAddresses: z
        .array(memberAddressSchema)
        .min(1, 'Must provide at least 1 address'),

    memberContactNumberReferences: z
        .array(memberContactReferencesSchema)
        .min(1, 'Must provide at least 1 contact number'),

    memberIncome: z.array(memberIncomeSchema).optional(),
    memberAssets: z.array(memberAssetsSchema).optional(),
    memberExpenses: z.array(memberExpensesSchema).optional(),
    memberRecruits: z.array(memberRecruitsSchema).optional(),
    memberDescriptions: z.array(memberDescriptionSchema).optional(),
    memberJointAccounts: z.array(memberJointAccountsSchema).optional(),
    memberGovernmentBenefits: z.array(memberGovernmentBenefits).optional(),
    memberRelativeAccounts: z.array(memberRelativeAccountsSchema).optional(),
})
