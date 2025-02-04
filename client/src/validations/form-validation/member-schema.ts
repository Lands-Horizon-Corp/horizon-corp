import z from 'zod'
import {
    emailSchema,
    lastNameSchema,
    passwordSchema,
    userNameSchema,
    birthDateSchema,
    firstNameSchema,
    middleNameSchema,
    contactNumberSchema,
    permanentAddressSchema,
    entityIdSchema,
} from '../common'

export const createMemberSchema = z.object({
    email: emailSchema,
    username: userNameSchema,
    firstName: firstNameSchema,
    middleName: middleNameSchema.optional(),
    lastName: lastNameSchema,
    birthDate: birthDateSchema,
    contactNumber: contactNumberSchema,
    permanentAddress: permanentAddressSchema,
    password: passwordSchema,
    confirmPassword: passwordSchema,
})

export const createMemberProfileSchema = z.object({
    oldReferenceID: z.string().optional(),
    passbookNumber: z.string().optional(),

    notes: z.string().min(1, 'Notes are required'),
    description: z.string().min(1, 'Description is required'),
    contactNumber: z.string().min(1, 'Contact number is required'),
    tinNumber: z.string().optional(),
    civilStatus: z
        .enum(['Married', 'Single', 'Widowed', 'Separated', 'N/A'])
        .default('Single'),
    occupation: z.string().optional(),
    sssNumber: z.string().optional(),
    businessAddress: z.string().optional(),
    businessContact: z.string().optional(),

    status: z.enum(['Pending', 'Verified', 'Not Allowed']).default('Pending'),
    isClosed: z.boolean(),

    pagibigNumber: z.string().optional(),
    philhealthNumber: z.string().optional(),
    isMutualFundMember: z.boolean().default(false),
    isMicroFinanceMember: z.boolean().default(false),

    mediaID: entityIdSchema.optional(),
    memberID: entityIdSchema.optional(),

    memberTypeID: entityIdSchema.optional(),
    memberClassificationID: entityIdSchema.optional(),
    memberGenderID: entityIdSchema.optional(),
    branchID: entityIdSchema.optional(),
    memberCenterID: entityIdSchema.optional(),
    memberEducationalAttainmentID: entityIdSchema.optional(),

    memberDescriptions: z
        .array(
            z.object({
                name: z.string().min(1, 'Name is required'),
                description: z.string().min(1, 'Description is required'),
            })
        )
        .optional(),

    memberAddress: z
        .array(
            z.object({
                postalCode: z.string().min(1, 'Postal Code is required'),
                province: z.string().min(1, 'Postal Code is required'),
                city: z.string().min(1, 'Postal Code is required'),
                barangay: z.string().min(1, 'Postal Code is required'),
                label: z.string().min(1, 'Postal Code is required'),
            })
        )
        .min(1, 'Must provide at least 1 address'),

    memberContactNumberReferences: z
        .array(
            z.object({
                name: z.string().min(1, 'Name is required'),
                description: z.string().min(1, 'Description is required'),
                contactNumber: z.string().min(1, 'Contact Number is required'),
            })
        )
        .min(1, 'Must provide at least 1 contact number'),

    memberIncome: z
        .array(
            z.object({
                name: z.string().min(1, 'Name is required'),
                amount: z.coerce.number().min(0, 'Amount must be non-negative'),
                date: z.string().min(1, 'Date is required'),
                description: z.string().min(1, 'Description is required'),
            })
        )
        .optional(),

    memberRelativeAccounts: z
        .array(
            z.object({
                membersProfileID: entityIdSchema,
                relativeProfileMemberID: entityIdSchema,
                familyRelationship: z
                    .string()
                    .min(1, 'Family relationship is required'),
                description: z.string().min(1, 'Description is required'),
            })
        )
        .optional(),

    memberAssets: z
        .array(
            z.object({
                entryDate: z.string().min(1, 'Entry Date is required'),
                description: z.string().min(1, 'Description is required'),
                name: z.string().min(1, 'Name is required'),
            })
        )
        .optional(),

    memberExpenses: z
        .array(
            z.object({
                name: z.string().min(1, 'Name is required'),
                date: z.string().min(1, 'Date is required'),
                amount: z.number().min(0, 'Amount must be non-negative'),
                description: z.string().min(1, 'Description is required'),
            })
        )
        .optional(),

    memberJointAccounts: z
        .array(
            z.object({
                description: z.string().min(1, 'Description is required'),
                firstName: z.string().min(1, 'First name is required'),
                lastName: z.string().min(1, 'Last name is required'),
                middleName: z.string().optional(),
                familyRelationship: z.string().optional(),
            })
        )
        .optional(),

    memberRecruits: z
        .array(
            z.object({
                membersProfileID: entityIdSchema,
                membersProfileRecruitedID: entityIdSchema,
                dateRecruited: z.string().min(1, 'Date recruited is required'),
                description: z.string().min(1, 'Description is required'),
                name: z.string().min(1, 'Name is required'),
            })
        )
        .optional(),

    memberGovernmentBenefits: z
        .array(
            z.object({
                country: z.string().min(1, 'Country is required'),
                name: z.string().min(1, 'Name is required'),
                description: z.string().min(1, 'Description is required'),
                value: z.number().min(0, 'Value must be non-negative'),
                frontMediaID: entityIdSchema.optional(),
                backMediaID: entityIdSchema.optional(),
            })
        )
        .optional(),
})
