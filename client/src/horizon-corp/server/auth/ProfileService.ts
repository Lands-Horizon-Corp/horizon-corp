import { AxiosResponse } from 'axios'
import UseServer from '@/horizon-corp/request/server'
import {
  UserData,
  AccountSettingRequest,
  MediaRequest,
} from '@/horizon-corp/types'

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
}
