import { AccountType, MediaResource } from '..'

export type AccountStatus = 'Pending' | 'Verified' | 'Not Allowed'

export interface UserData {
    id: number
    username: string
    firstName: string
    middleName?: string
    lastName: string
    permanentAddress: string
    description: string
    birthDate: Date
    createdAt: Date
    email: string
    accountType: AccountType
    status: AccountStatus
    isEmailVerified: boolean
    isContactVerified: boolean
    media? : MediaResource
}

export interface SignUpResource extends UserData {}

export interface VerifyEmailResource extends UserData {}

export interface VerifyContactResource extends UserData {}

export interface CurrentUserResource extends UserData {}

export interface SignInResource extends UserData {}

// auth/forgot-password -> just http status code 200
// auth/password-reset/${resetId} -> just http status code
