import { AxiosResponse } from 'axios'
import UseServer from '@/horizon-corp/request/server'
import {
    UserData,
    DescriptionRequest,
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

    public static async Description(
        data: DescriptionRequest
    ): Promise<AxiosResponse<UserData>> {
        const endpoint = `${ProfileService.BASE_ENDPOINT}/description`
        return await UseServer.post<DescriptionRequest, UserData>(
            endpoint,
            data
        )
    }
}
