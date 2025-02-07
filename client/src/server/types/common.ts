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

export interface ILongLat {
    longitude?: number // `float64` maps to `number` in TypeScript
    latitude?: number // `float64` maps to `number` in TypeScript
}

export interface ITimeStamps {
    deletedAt?: string | null
    createdAt: string
    updatedAt: string
}

export type TCivilStatus =
    | 'Married'
    | 'Single'
    | 'Widowed'
    | 'Separated'
    | 'N/A'
