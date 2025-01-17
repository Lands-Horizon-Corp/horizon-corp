import { IMediaResource } from './media'
import { IGenderResource } from './gender'

export type IAccountType = 'Member' | 'Employee' | 'Admin' | 'Owner'

export type IAccountStatus = 'Pending' | 'Verified' | 'Not Allowed'

export interface IUserBase extends ITimeStamps {
    id: number

    firstName: string
    lastName: string
    middleName?: string
    birthDate: Date

    email: string
    contactNumber: string
    permanentAddress?: string

    username: string
    description?: string

    mediaId?: number
    media?: IMediaResource

    genderId?: number
    gender?: IGenderResource

    status: IAccountStatus

    isEmailVerified: boolean
    isContactVerified: boolean
    isSkipVerification: boolean
}

export interface ITimeStamps {
    deletedAt?: Date | null
    createdAt: Date
    updatedAt: Date
}
