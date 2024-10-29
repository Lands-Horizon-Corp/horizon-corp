import { MediaRequest } from '..'

export type AccountType = 'Member' | 'Employee' | 'Admin' | 'Owner'

export interface ChangePasswordRequest {
  otp?: string
  resetId?: string
  newPassword: string
  confirmPassword: string
}

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
  otp: string
}

export interface VerifyContactNumberRequest {
  otp: string
}




export interface NewPasswordRequest {
  PreviousPassword: string
  NewPassword: string
  ConfirmPassword: string
}

export interface AccountSettingRequest {
  birthdate: Date
  firstName: string
  lastName: string
  description: string
  permanentAddress: string
}

export interface ChangeEmailRequest {
  password: string
  email: string
}

export interface ChangeContactNumberRequest {
  password: string
  contactNumber: string
}
export interface ChangeUsernameRequest {
  password: string
  username: string
}