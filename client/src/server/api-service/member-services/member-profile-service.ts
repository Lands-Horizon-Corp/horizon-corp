import qs from 'query-string'

import APIService from '../api-service'

import {
    IMemberProfilePaginatedPicker,
    IMemberProfileRequest,
    IMemberProfileResource,
} from '../../types/member/member-profile'

export interface IMemberProfilePickerParams {
    sort?: string
    filters?: string
    pagination?: { pageIndex: number; pageSize: number }
}

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

    public static async getMemberProfilesForPicker(
        params: IMemberProfilePickerParams
    ): Promise<IMemberProfilePaginatedPicker> {
        const { filters, pagination, sort } = params

        const url = qs.stringifyUrl(
            {
                url: `${MemberProfileService.BASE_ENDPOINT}/picker`,
                query: {
                    sort,
                    filters,
                    pageIndex: pagination?.pageIndex,
                    pageSize: pagination?.pageSize,
                },
            },
            { skipNull: true }
        )

        const response =
            await APIService.get<IMemberProfilePaginatedPicker>(url)
        return response.data
    }
}
