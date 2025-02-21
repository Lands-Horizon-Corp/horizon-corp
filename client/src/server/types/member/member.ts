import { IMediaRequest } from '../media'
import { IRolesResource } from '../role'
import { IGenderResource } from '../gender'
import { IFootstepResource } from '../footstep'
import { IPaginatedResult } from '../paginated-result'
import { ITimeStamps, IUserBase, TAccountStatus, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

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
    fullName: string
    description?: string
    isEmailVerified: boolean
    isContactVerified: boolean
    isSkipVerification: boolean

    role?: IRolesResource
    gender?: IGenderResource

    status: TAccountStatus

    footsteps?: IFootstepResource[]
    memberProfile?: IMemberProfileResource
}

export type IMemberPaginatedResource = IPaginatedResult<IMemberResource>
