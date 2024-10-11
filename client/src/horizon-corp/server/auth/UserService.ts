import type { ChangeContactNumberRequest, ChangeEmailRequest, ChangePasswordRequest, ForgotPasswordRequest, SendEmailVerificationRequest, SignInRequest, SignOutRequest, SignUpRequest, VerifyContactNumberRequest, VerifyEmailRequest, VerifyOTPRequest } from "@/horizon-corp/types";

export default class UserService {
  private static readonly BASE_ENDPOINT = '/auth'

  public static async SignUp(data: SignUpRequest): Promise<void> {
    console.log(data)
  }

  public static async SignIn(data: SignInRequest): Promise<void> {
    console.log(data)
  }

  public static async SignOut(data: SignOutRequest): Promise<void> {
    console.log(data)
  }

  public static async ForgotPassword(data: ForgotPasswordRequest): Promise<void> {
    console.log(data)
  }

  public static async ChangePassword(data: ChangePasswordRequest): Promise<void> {
    console.log(data)
  }

  public static async ChangeEmail(data: ChangeEmailRequest): Promise<void> {
    console.log(data)
  }

  public static async SendEmailVerification(data: SendEmailVerificationRequest): Promise<void> {
    console.log(data)
  }

  public static async VerifyEmail(data: VerifyEmailRequest): Promise<void> {
    console.log(data)
  }

  public static async ChangeContactNumber(data: ChangeContactNumberRequest): Promise<void> {
    console.log(data)
  }

  public static async SendContactNumberVerification(data: VerifyOTPRequest): Promise<void> {
    console.log(data)
  }

  public static async VerifyContactNumber(data: VerifyContactNumberRequest): Promise<void> {
    console.log(data)
  }
}