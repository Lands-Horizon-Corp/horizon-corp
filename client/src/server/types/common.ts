import { IMediaResource } from './media'
import { IGenderResource } from './gender'

export type TAccountType = 'Member' | 'Employee' | 'Admin' | 'Owner'

export type TAccountStatus = 'Pending' | 'Verified' | 'Not Allowed'

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

    status: TAccountStatus

    isEmailVerified: boolean
    isContactVerified: boolean
    isSkipVerification: boolean
}

export interface ITimeStamps {
    deletedAt?: Date | null
    createdAt: Date
    updatedAt: Date
}
