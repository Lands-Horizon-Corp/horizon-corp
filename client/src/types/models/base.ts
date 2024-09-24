import { Media } from './media'
import { Role } from './role'

export interface BaseModel {
    createdAt: Date
    updatedAt: Date
    deletedAt?: Date
}

export interface UserBase extends BaseModel {
    id: string
    email: string
    username: string
    firstName: string
    middleName?: string
    lastName: string
    contactNumber: string
    permanentAddress: string
    description: string
    birthdate: Date
    validEmail: boolean
    validContactNumber: boolean
    mediaId?: string
    profilePicture?: Media
    roles: Role[]
    status: UserStatus
}

export enum UserStatus {
    'Pending' = 'Pending',
    'Verified' = 'Verified',
    'Not allowed' = 'Not allowed',
}
