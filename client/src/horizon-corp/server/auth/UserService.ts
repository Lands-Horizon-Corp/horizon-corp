import { AxiosResponse } from 'axios'
import UseServer from '@/horizon-corp/request/server'
import type {
    ChangeContactNumberRequest,
    ChangeEmailRequest,
    ChangePasswordRequest,
    CurrentUserResource,
    ForgotPasswordRequest,
    SendEmailVerificationRequest,
    SignInRequest,
    SignInResource,
    SignUpRequest,
    SignUpResource,
    VerifyContactNumberRequest,
    VerifyContactResource,
    VerifyEmailRequest,
} from '@/horizon-corp/types'
import { getSMSContent, getEmailContent } from '@/lib'

export default class UserService {
    private static readonly BASE_ENDPOINT = '/auth'

    // GET - /auth/current-user
    public static async CurrentUser(): Promise<
        AxiosResponse<CurrentUserResource>
    > {
        const endpoint = `${UserService.BASE_ENDPOINT}/current-user`
        return await UseServer.get<CurrentUserResource>(endpoint)
    }

    // POST - /auth/signup
    /*
      pagka create return mo sakin user data except password

      NOTE: when creating new user also do the following below
      - create/new verify email entry
          - via redis? db? idk
          yung structure nya sa db/redis { id, userId, accountType, otp, email }
          - send email generated otp
      - create/new verify contact entry
          - via redis? db? idk
          same lang din yung structure nya sa db/redis { id, userId, accountType, otp, contact }
          - send generated otp code to contact number via sms
  */
    public static async SignUp(
        data: SignUpRequest
    ): Promise<AxiosResponse<SignUpResource>> {
        const endpoint = `${UserService.BASE_ENDPOINT}/signup`
        data.birthdate = new Date()
        data.emailTemplate = getEmailContent('otp')
        data.contactTemplate = getSMSContent('ContactNumber')
        return await UseServer.post<SignUpRequest, SignUpResource>(
            endpoint,
            data
        )
    }

    // POST - /auth/signin
    public static async SignIn(
        data: SignInRequest
    ): Promise<AxiosResponse<SignInResource>> {
        const endpoint = `${UserService.BASE_ENDPOINT}/signin`
        return await UseServer.post<SignInRequest, SignInResource>(
            endpoint,
            data
        )
    }

    // POST - /auth/signout
    public static async SignOut(): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/signout`
        await UseServer.post(endpoint)
    }

    // POST - /auth/forgot-password
    /* 
  Need : 🔥
  Sa pag create pala man ng password reset entry/url sa db need ko lang 
  UUID as link tapos mode/accountType mastore sa db
  example reset entry : 
  {
      id : myqlIndexId
      resetId : UUID
      userId : id ng user
      code : string 6 Digit code, pag sa mobile(for future) eto gagamitin
      accountType : Yung type na ginawa mo na type sa horizon-corp/types/auth/request.ts @Line 3
  }

  Need : 🔥
  Need ko din pala ng endpoint para macheck if yung UUID sa link below is valid
                                     👇
  /auth/password-reset/15de37fa-cb0f-4295-8fe9-63daeb944492

  need ko lamang mag return ka ng true or false or hahit response status 200 pag exist/valid or 404 if invalid
  */
    public static async ForgotPassword(
        data: ForgotPasswordRequest
    ): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/forgot-password`
        await UseServer.post<ForgotPasswordRequest, void>(endpoint, data)
    }

    // POST - /auth/change-password
    /*
  Need: 🔥
  Need ko din ng another endpoint for password-reset eto ata yun man?

  mas goods ba if nasa url na yung id? or ipass ko nlng sa req body?
  url : /auth/password-reset/${UUID resetIdDItoDesyo} then body { newPassword : string }
  or
  via body : /auth/password-reset then body { resetId : UUID string, newPassword : string }

  ikaw bahala man
  
  - then verify mo nlng sa db kung ano napili paglagyan (redis/db) 
  nandon din naman { id, userId, accountType } ng specific user na papalitan password
  
  walang response kahit http status code
  */
    public static async ChangePassword(
        data: ChangePasswordRequest
    ): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/change-password`
        await UseServer.post<ChangePasswordRequest, void>(endpoint, data)
    }

    // POST - /auth/change-email
    public static async ChangeEmail(data: ChangeEmailRequest): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/change-email`
        await UseServer.post<ChangeEmailRequest, void>(endpoint, data)
    }

    // POST - /auth/send-email-verification
    public static async SendEmailVerification(
        data: SendEmailVerificationRequest
    ): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/send-email-verification`
        await UseServer.post<SendEmailVerificationRequest, void>(endpoint, data)
    }

    // POST - /auth/verify-email
    /*
   Need: 🔥🔥
   I'll pass 
   {
      otp : string
      email : string,
      userId : number
   }

   You return new data of user
  */
    public static async VerifyEmail(data: VerifyEmailRequest): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/verify-email`
        await UseServer.post<VerifyEmailRequest, void>(endpoint, data)
    }

    // POST - /auth/verify-otp
    public static async VerifyContactNumber(
        data: VerifyContactNumberRequest
    ): Promise<AxiosResponse<VerifyContactResource>> {
        const endpoint = `${UserService.BASE_ENDPOINT}/auth/verify-otp`
        return await UseServer.post<VerifyContactNumberRequest, VerifyContactResource>(endpoint, data)
    }

    // POST - /auth/change-contact-number
    public static async ChangeContactNumber(
        data: ChangeContactNumberRequest
    ): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/change-contact-number`
        await UseServer.post<ChangeContactNumberRequest, void>(endpoint, data)
    }
}
