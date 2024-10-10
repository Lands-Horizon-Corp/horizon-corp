export interface ChangePasswordRequest {
  token: string
  newPassword: string
  confirmPassword: string
}

export interface ForgotPasswordRequest {
  email: string
}

export interface VerifyEmailRequest {
  token: string
  email: string
}

export interface SignOutRequest {
  token: string
}

export interface SignUpRequest {
  firstName: string
  lastName: string
  email: string
  password: string
  confirmPassword: string
  birthdate: Date
  contactNumber: string
  permanentAddress?: string
}

export interface VerifyContactRequest {
  token: string
  contactNumber: string
}

export interface VerifyEmailRequest {
  token: string
  email: string
}