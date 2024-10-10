export interface ChangePasswordRequest {
  currentPassword?: string; // Optional if token exists
  newPassword: string;
  confirmPassword: string;
  token?: string; // Optional
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

export interface SignOutRequest {
  token: string;
}

export interface SignUpRequest {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  confirmPassword: string;
  birthdate: Date;
  contactNumber: string;
  permanentAddress: string;
}

export interface VerifyEmailRequest {
  email: string;
  token: string;
}

export interface VerifyOTPRequest {
  contactNumber: string;
  otp: string;
}
