import { IMediaRequest } from '../media'
import { IRolesResource } from '../role'
import { IGenderResource } from '../gender'
import { IFootstepResource } from '../footstep'
import { IPaginatedResult } from '../paginated-result'
import { ITimeStamps, IUserBase, TAccountStatus, TEntityId } from '../common'

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

    companyId?: TEntityId
}

export interface IMemberResource extends IUserBase, ITimeStamps {
    id: TEntityId
    accountType: 'Member'

    username: string
    description?: string
    isEmailVerified: boolean
    isContactVerified: boolean
    isSkipVerification: boolean

    role?: IRolesResource // Optional due to `omitempty`
    gender?: IGenderResource // Optional due to `omitempty`

    status: TAccountStatus // Assuming `providers.UserStatus` maps to a string; adjust if it's an enum

    footsteps?: IFootstepResource[] // Array of `FootstepResource`
    // memberProfile?: MemberProfileResource; // Optional due to `omitempty`
}

export type IMemberPaginatedResource = IPaginatedResult<IMemberResource>
