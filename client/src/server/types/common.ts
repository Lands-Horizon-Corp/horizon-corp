import { IMediaResource } from './media'
import { IGenderResource } from './gender'

export type TAccountType = 'Member' | 'Employee' | 'Admin' | 'Owner'

export type TAccountStatus = 'Pending' | 'Verified' | 'Not Allowed'

export type TEntityId = string

export interface IUserBase extends ITimeStamps, ILongLat {
    id: TEntityId

    firstName: string
    lastName: string
    middleName?: string
    fullName: string // concat & returned by server
    birthDate: string

    email: string
    contactNumber: string
    permanentAddress?: string

    username: string
    description?: string

    mediaId?: number
    media?: IMediaResource

    genderId?: number
    gender?: IGenderResource

    status: TAccountStatus

    isEmailVerified: boolean
    isContactVerified: boolean
    isSkipVerification: boolean
}

export interface ILongLat {
    longitude?: number // `float64` maps to `number` in TypeScript
    latitude?: number // `float64` maps to `number` in TypeScript
}

export interface ITimeStamps {
    deletedAt?: string | null
    createdAt: string
    updatedAt?: string
}

export type TCivilStatus =
    | 'Married'
    | 'Single'
    | 'Widowed'
    | 'Separated'
    | 'N/A'

export const AccountClosureReasonTypes = [
    'Voluntary Withdrawal',
    'Inactive Membership',
    'Violation of Cooperative Rules',
    'Failure to Meet Financial Obligations',
    'Fraudulent Activities',
    'Transfer to Another Cooperative',
    'Death of Member',
    'Membership Revocation',
    'Request for Account Termination',
    'Other Personal Reasons',
    'Relocation Outside Service Area',
    'Better Opportunities Elsewhere',
    'Disagreement with Cooperative Direction',
    'Lack of Perceived Benefit',
    'Changes in Personal Circumstances (e.g., retirement, job change)',
    'Administrative Issues or Poor Service',
    'Merger or Dissolution of the Cooperative',
    'Financial Hardship',
    'Loss of Trust in Leadership',
    'Limited Engagement or Participation',
    'Becoming a Customer Only (if applicable)',
    'Desire for Simpler Financial Arrangements',
    'Conflict of Interest',
    'Feeling Undervalued or Unheard',
    'Changes in Cooperative Offerings',
    'Technological Disadvantages Compared to Other Options',
    'Success in Independent Ventures',
    'Inheritance or Windfall Allowing Independence',
    'Shifting Personal Values or Priorities',
    'Formation of a New, More Suitable Cooperative',
] as const

export type AccountClosureReasonType =
    (typeof AccountClosureReasonTypes)[number]
