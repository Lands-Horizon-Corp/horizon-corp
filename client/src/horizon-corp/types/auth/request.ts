import { MediaRequest } from "..";
export type AccountType = 'Member' | 'Employee' | 'Admin' | 'Owner'
export interface ChangePasswordRequest {
  currentPassword?: string;
  newPassword: string;
  confirmPassword: string;
  token?: string;
}

export interface ForgotPasswordRequest {
  email: string;
}

export interface SendEmailVerificationRequest {
  email: string;
}

export interface SendOTPVerificationRequest {
  contactNumber: string;
}

export interface SignInRequest {
  email: string;
  password: string;
}


export interface SignUpRequest {
  mode: AccountType,
  firstName: string;
  lastName: string;
  middleName?: string;
  email: string;
  password: string;
  confirmPassword: string;
  birthdate: string;
  contactNumber: string;
  permanentAddress: string;
  media?: MediaRequest;
}

export interface VerifyEmailRequest {
  email: string;
  token: string;
}

export interface VerifyOTPRequest {
  contactNumber: string;
  otp: string;
}


export interface ChangeEmailRequest {
  email: string;
}

export interface ChangeContactNumberRequest {
  contactNumber: string
}

export interface VerifyContactNumberRequest {
  contactNumber: string;
  otp: string;
}

