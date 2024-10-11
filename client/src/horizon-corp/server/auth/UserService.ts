import type { ChangeContactNumberRequest, ChangeEmailRequest, ChangePasswordRequest, ForgotPasswordRequest, SendEmailVerificationRequest, SignInRequest, SignOutRequest, SignUpRequest, VerifyContactNumberRequest, VerifyEmailRequest, VerifyOTPRequest } from "@/horizon-corp/types";

export default class UserService {
  private static readonly BASE_ENDPOINT = '/auth'


  // GET - /auth/current-user
  public static async CurrentUser(): Promise<void> {

  }

  // POST - /auth/signup
  public static async SignUp(data: SignUpRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/signin
  public static async SignIn(data: SignInRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/signout
  public static async SignOut(data: SignOutRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/forgot-password
  public static async ForgotPassword(data: ForgotPasswordRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/change-password
  public static async ChangePassword(data: ChangePasswordRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/change-email
  public static async ChangeEmail(data: ChangeEmailRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/send-email-verification
  public static async SendEmailVerification(data: SendEmailVerificationRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/verify-email
  public static async VerifyEmail(data: VerifyEmailRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/change-contact-number
  public static async ChangeContactNumber(data: ChangeContactNumberRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/send-otp-verification
  public static async SendContactNumberVerification(data: VerifyOTPRequest): Promise<void> {
    console.log(data)
  }

  // POST - /auth/verify-otp
  public static async VerifyContactNumber(data: VerifyContactNumberRequest): Promise<void> {
    console.log(data)
  }
}