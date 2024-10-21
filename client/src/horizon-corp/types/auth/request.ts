import { MediaRequest } from '..'

export type AccountType = 'Member' | 'Employee' | 'Admin' | 'Owner'

export interface ChangePasswordRequest {
    otp?: string // token sent to user (for future implementation sa mobile, otp code sya instead of link)
    resetId?: string // eto yung id ng reset link example: ecoop/auth/password-reset/88ff6b8aCc -> yung resetId ay 88ff6b8aCc
    newPassword: string // bagong password
    confirmPassword: string // idk if need pato since validated na sa frontend pero dapat equal to sa new password
} // note: pwede otp/reset id isend, pero bawal undefined both

export interface ForgotPasswordRequest {
    key: string
    accountType: string
    emailTemplate?: string
    contactTemplate?: string
}

export interface SendEmailVerificationRequest {
    emailTemplate?: string
}

export interface SendContactNumberVerificationRequest {
    contactTemplate?: string
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
