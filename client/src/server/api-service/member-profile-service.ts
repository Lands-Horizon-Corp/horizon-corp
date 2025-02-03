import qs from 'query-string'

import APIService from './api-service'

import {
    IMemberProfileRequest,
    IMemberProfileResource,
} from '../types/member-profile'

export default class MemberProfileService {
    private static readonly BASE_ENDPOINT = '/member-profile'

    public static async create(
        data: IMemberProfileRequest,
        preloads?: string[]
    ) {
        const url = qs.stringifyUrl(
            {
                url: `${MemberProfileService.BASE_ENDPOINT}`,
                query: { preloads },
            },
            { skipNull: true }
        )

        return (
            await APIService.post<
                IMemberProfileRequest,
                IMemberProfileResource
            >(url, data)
        ).data
    }
}
