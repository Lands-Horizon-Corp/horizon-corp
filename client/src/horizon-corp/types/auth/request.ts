import { MediaRequest } from '..'

export type AccountType = 'Member' | 'Employee' | 'Admin' | 'Owner'

export interface ChangePasswordRequest {
    currentPassword?: string
    newPassword: string
    confirmPassword: string
    token?: string
}

export interface ForgotPasswordRequest {
    email: string
    accountType: AccountType
}

export interface SendEmailVerificationRequest {
    email: string
}

export interface SendOTPVerificationRequest {
    contactNumber: string
}

export interface SignInRequest {
    email?: string
    username?: string
    password: string
    accountType: AccountType
}

export interface SignUpRequest {
    accountType: AccountType
    username: string
    firstName: string
    lastName: string
    middleName?: string
    email: string
    password: string
    confirmPassword: string
    birthdate: Date
    contactNumber: string
    permanentAddress: string
    media?: MediaRequest

    emailTemplate?: string
    contactTemplate?: string
}

export interface VerifyEmailRequest {
    otp: string // 6 digit string/number
}

export interface VerifyContactNumberRequest {
    otp: string // 6 digit string/number
}

export interface ChangeEmailRequest {
    email: string
}

export interface ChangeContactNumberRequest {
    contactNumber: string
}
