import { IRolesResource } from './role'
import { IGenderResource } from './gender'
import { IFootstepResource } from './footstep'
import { IPaginatedResult } from './paginated-result'
import { ITimeStamps, TAccountStatus } from './common'
import { IMediaRequest, IMediaResource } from './media'

export interface IMemberRequest {
    username: string
    firstName: string
    lastName: string
    middleName?: string
    birthDate: Date

    contactNumber: string
    permanentAddress: string

    email: string
    password: string
    confirmPassword: string
    media?: IMediaRequest

    emailTemplate?: string
    contactTemplate?: string

    companyId?: number
}

export interface IMemberResource extends ITimeStamps {
    id: string
    accountType: 'Member'

    username: string
    description?: string
    isEmailVerified: boolean
    isContactVerified: boolean
    isSkipVerification: boolean

    email: string
    contactNumber: string
    permanentAddress?: string
    role?: IRolesResource // Optional due to `omitempty`
    gender?: IGenderResource // Optional due to `omitempty`

    firstName: string
    lastName: string
    middleName?: string // Optional due to `omitempty`
    fullName: string
    birthDate: Date // `time.Time` maps to `Date` in TypeScript

    status: TAccountStatus // Assuming `providers.UserStatus` maps to a string; adjust if it's an enum

    media?: IMediaResource // Optional due to `omitempty`
    longitude?: number // `float64` maps to `number` in TypeScript
    latitude?: number // `float64` maps to `number` in TypeScript
    footsteps?: IFootstepResource[] // Array of `FootstepResource`
    // memberProfile?: MemberProfileResource; // Optional due to `omitempty`
}

export type IMemberPaginatedResource = IPaginatedResult<IMemberResource>
