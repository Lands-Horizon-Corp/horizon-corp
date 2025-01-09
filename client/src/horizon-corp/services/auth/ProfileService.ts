import { AxiosResponse } from 'axios'
import UseServer from '@/horizon-corp/request/server'
import {
    UserData,
    AccountSettingRequest,
    MediaRequest,
    NewPasswordRequest,
    ChangeContactNumberRequest,
    ChangeEmailRequest,
    ChangeUsernameRequest,
} from '@/horizon-corp/types'
import UserService from './UserService'

export default class ProfileService {
    private static readonly BASE_ENDPOINT = '/profile'

    // POST - /profile/profile-picture
    public static async ProfilePicture(
        data: MediaRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${ProfileService.BASE_ENDPOINT}/profile-picture`
        return await UseServer.post<MediaRequest, UserData>(endpoint, data)
    }

    public static async AccountSetting(
        data: AccountSettingRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${ProfileService.BASE_ENDPOINT}/account-setting`
        return await UseServer.post<AccountSettingRequest, UserData>(
            endpoint,
            data
        )
    }

    public static async NewPassword(data: NewPasswordRequest): Promise<void> {
        return UserService.NewPassword(data)
    }

    public static async ChangeContactNumber(
        data: ChangeContactNumberRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${ProfileService.BASE_ENDPOINT}/change-contact-number`
        return await UseServer.post<ChangeContactNumberRequest, UserData>(
            endpoint,
            data
        )
    }

    public static async ChangeEmail(
        data: ChangeEmailRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${ProfileService.BASE_ENDPOINT}/change-email`
        return await UseServer.post<ChangeEmailRequest, UserData>(
            endpoint,
            data
        )
    }

    public static async ChangeUsername(
        data: ChangeUsernameRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${ProfileService.BASE_ENDPOINT}/change-username`
        return await UseServer.post<ChangeUsernameRequest, UserData>(
            endpoint,
            data
        )
    }
}
