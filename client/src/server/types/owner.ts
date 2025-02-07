import { ICompanyResource } from './company'
import { IFootstepResource } from './footstep'
import { IUserBase, TAccountType } from './common'
import { IPaginatedResult } from './paginated-result'
import { IMediaRequest, IMediaResource } from './media'

export interface IOwnerResource extends IUserBase {
    accountType: string
    media?: IMediaResource
    companies?: ICompanyResource[]
    footsteps?: IFootstepResource[]
}

export type IOwnerPaginatedResource = IPaginatedResult<IOwnerResource>

export interface IOwnerRequest {
    accountType: TAccountType
    username: string
    firstName: string
    lastName: string
    middleName?: string
    email: string
    password: string
    confirmPassword: string
    birthDate: Date
    contactNumber: string
    permanentAddress: string
    media?: IMediaRequest

    emailTemplate?: string
    contactTemplate?: string
}
