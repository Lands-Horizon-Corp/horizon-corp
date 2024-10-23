import { AxiosResponse } from 'axios'
import UseServer from '@/horizon-corp/request/server'
import type {
    ChangeContactNumberRequest,
    ChangeEmailRequest,
    ChangePasswordRequest,
    UserData,
    ForgotPasswordRequest,
    SendEmailVerificationRequest,
    SendContactNumberVerificationRequest,
    SignInRequest,
    UserData,
    SignUpRequest,
    UserData,
    UserData,
    VerifyContactNumberRequest,
    UserData,
    VerifyEmailRequest,
    UserData,
    NewPasswordRequest,
} from '@/horizon-corp/types'
import { getSMSContent, getEmailContent } from '@/lib'

export default class UserService {
    private static readonly BASE_ENDPOINT = '/auth'

    // GET - /auth/current-user
    public static async CurrentUser(): Promise<
        AxiosResponse<UserData>
    > {
        const endpoint = `${UserService.BASE_ENDPOINT}/current-user`
        return await UseServer.get<UserData>(endpoint)
    }

    // POST - /auth/signup
    public static async SignUp(
        data: SignUpRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${UserService.BASE_ENDPOINT}/signup`
        data.birthdate = new Date()
        data.emailTemplate = getEmailContent('otp')
        data.contactTemplate = getSMSContent('contactNumber')
        return await UseServer.post<SignUpRequest, UserData>(
            endpoint,
            data
        )
    }

    // POST - /auth/signin
    public static async SignIn(
        data: SignInRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${UserService.BASE_ENDPOINT}/signin`
        return await UseServer.post<SignInRequest, UserData>(
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
    public static async ForgotPassword(
        data: ForgotPasswordRequest
    ): Promise<void> {
        data.emailTemplate = getEmailContent('changePassword')
        data.contactTemplate = getSMSContent('changePassword')
        const endpoint = `${UserService.BASE_ENDPOINT}/forgot-password`
        await UseServer.post<ForgotPasswordRequest, void>(endpoint, data)
    }

    // GET - /auth/verify-reset-link/${resetId} ex: /auth/verify-reset-link/ba88ffCdD
    // - check if the resetId is valid
    // just return 200 ok
    public static async CheckResetLink(resetId: string): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/verify-reset-link/${resetId}`
        await UseServer.get<void>(endpoint)
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
        data.emailTemplate = getEmailContent('otp')
        await UseServer.post<ChangeEmailRequest, void>(endpoint, data)
    }

    // POST - /auth/send-email-verification
    public static async SendEmailVerification(): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/send-email-verification`
        const data: SendEmailVerificationRequest = {
            emailTemplate: getEmailContent('otp'),
        }
        await UseServer.post<SendEmailVerificationRequest>(endpoint, data)
    }

    // POST - /auth/verify-email
    public static async VerifyEmail(
        data: VerifyEmailRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${UserService.BASE_ENDPOINT}/verify-email`
        return await UseServer.post<VerifyEmailRequest, UserData>(
            endpoint,
            data
        )
    }

    // POST - /auth/send-contact-number-verification
    public static async SendContactVerification(): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/send-contact-number-verification`
        const data: SendContactNumberVerificationRequest = {
            contactTemplate: getSMSContent('contactNumber'),
        }
        await UseServer.post<SendContactNumberVerificationRequest>(
            endpoint,
            data
        )
    }

    // POST - /auth/verify-contact
    public static async VerifyContactNumber(
        data: VerifyContactNumberRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${UserService.BASE_ENDPOINT}/verify-contact-number`
        return await UseServer.post<
            VerifyContactNumberRequest,
            UserData
        >(endpoint, data)

        // on backend read the current user cookie in backend, based on the information there, you'll know the contact
        // number to send the OTP to
    }

    // POST - /auth/change-contact-number
    public static async ChangeContactNumber(
        data: ChangeContactNumberRequest
    ): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/change-contact-number`
        data.contactTemplate = getSMSContent('contactNumber')
        await UseServer.post<ChangeContactNumberRequest, void>(endpoint, data)
    }

    // POST - /auth/skip-verification
    public static async SkipVerification(): Promise<
        AxiosResponse<UserData>
    > {
        const endpoint = `${UserService.BASE_ENDPOINT}/skip-verification`
        return await UseServer.post<void, UserData>(endpoint)
    }

    // POST - /auth/new-password
    public static async NewPassword(data: NewPasswordRequest): Promise<void> {
        const endpoint = `${UserService.BASE_ENDPOINT}/new-password`
        await UseServer.post<NewPasswordRequest>(endpoint, data)
    }
}
