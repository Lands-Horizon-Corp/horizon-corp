import { IUserBase } from './common'
import { IRolesResource } from './role'
import { IMediaRequest } from './media'
import { IBranchResource } from './branch'
import { IFootstepResource } from './footstep'
import { IPaginatedResult } from './paginated-result'

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
}

export interface IMemberResource extends IUserBase {
    accountType: string
    longitude?: number
    latitude?: number
    branch?: IBranchResource
    role?: IRolesResource
    footsteps?: IFootstepResource[]
}

export type IMemberPaginatedResource = IPaginatedResult<IMemberResource>
