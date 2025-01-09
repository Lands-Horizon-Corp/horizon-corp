import { GenderResource } from './gender'
import { MediaResource } from './media'

export type AccountStatus = 'Pending' | 'Verified' | 'Not Allowed'

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
    media?: MediaResource

    genderId?: number
    gender?: GenderResource

    status: AccountStatus

    isEmailVerified: boolean
    isContactVerified: boolean
    isSkipVerification: boolean
}

export interface ITimeStamps {
    deletedAt?: Date | null
    createdAt: Date
    updatedAt: Date
}
