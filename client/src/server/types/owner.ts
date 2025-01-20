import { IMediaRequest } from "./media"
import { IUserBase, TAccountType } from "./common"
import { IPaginatedResult } from "./paginated-result"

export interface IOwnerResource extends IUserBase {}

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

