import UseServer from '@/horizon-corp/request/server'
import type {
    ChangeContactNumberRequest,
    ChangeEmailRequest,
    ChangePasswordRequest,
    ForgotPasswordRequest,
    SendEmailVerificationRequest,
    SignInRequest,
    SignUpRequest,
    VerifyContactNumberRequest,
    VerifyEmailRequest,
    VerifyOTPRequest,
} from '@/horizon-corp/types'

export default class UserService {
    private static readonly BASE_ENDPOINT = '/auth'

    // GET - /auth/current-user
    public static async CurrentUser(): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/current-user`
        await UseServer.post<void, void>(endpoint)
    }

    /*
    Need : 🔥🔥
    Sa pag create pala man ng password reset entry/url sa db need ko lang 
    UUID as link tapos mode/accountType mastore sa db
    {
        id : myqlIndexId or pede din UUID
        userId : id ng user
        accountType : Yung type na ginawa mo na type sa horizon-corp/types/auth/request.ts @Line 3
    }

    Need : 🔥🔥
    Need ko din ng endpoint para macheck if yung UUID sa link below is valid
                                       👇
    /auth/password-reset/15de37fa-cb0f-4295-8fe9-63daeb944492

    need ko lamang mag return ka ng true or false or hahit response status 200 pag exist/valid or 404 if invalid

    Need: 🔥🔥
    Need ko din ng another endpoint for password-reset, ipapasa ko lang { resetId : UUID string, newPassword : string }
    - verify mo nlng sa db nandon din naman id, userId, accountType ng specific user na papalitan password

    response : wala kahit http status code
    */

    // POST - /auth/signup
    public static async SignUp(data: SignUpRequest): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/signup`
        await UseServer.post<SignUpRequest, void>(endpoint, data)
    }

    // POST - /auth/signin
    public static async SignIn(data: SignInRequest): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/signin`
        await UseServer.post<SignInRequest, void>(endpoint, data)
    }

    // POST - /auth/signout
    public static async SignOut(): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/signout`
        await UseServer.post(endpoint)
    }

    // POST - /auth/forgot-password
    public static async ForgotPassword(
        data: ForgotPasswordRequest
    ): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/forgot-password`
        await UseServer.post<ForgotPasswordRequest, void>(endpoint, data)
    }

    // POST - /auth/change-password
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
    public static async VerifyEmail(data: VerifyEmailRequest): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/verify-email`
        await UseServer.post<VerifyEmailRequest, void>(endpoint, data)
    }

    // POST - /auth/change-contact-number
    public static async ChangeContactNumber(
        data: ChangeContactNumberRequest
    ): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/change-contact-number`
        await UseServer.post<ChangeContactNumberRequest, void>(endpoint, data)
    }

    // POST - /auth/send-contact-number-verification
    public static async SendContactNumberVerification(
        data: VerifyOTPRequest
    ): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/send-contact-number-verification`
        await UseServer.post<ChangeContactNumberRequest, void>(endpoint, data)
    }

    // POST - /auth/verify-otp
    public static async VerifyContactNumber(
        data: VerifyContactNumberRequest
    ): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/auth/verify-otp`
        await UseServer.post<ChangeContactNumberRequest, void>(endpoint, data)
    }
}
