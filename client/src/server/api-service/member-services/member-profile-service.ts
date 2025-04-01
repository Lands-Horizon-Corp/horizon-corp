import qs from 'query-string'

import APIService from '../api-service'

import {
    IMemberProfileCloseRequest,
    IMemberProfileRequest,
    IMemberProfileResource,
} from '../../types/member/member-profile'
import { TEntityId } from '../../types/common'
import { IMemberCloseRemarkRequest } from '@/server/types'

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

    public static async getById(
        id: TEntityId,
        preloads?: string[]
    ): Promise<IMemberProfileResource> {
        const url = qs.stringifyUrl({
            url: `${MemberProfileService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })
        const response = await APIService.get<IMemberProfileResource>(url)
        return response.data
    }

    public static async update(
        id: TEntityId,
        memberData: IMemberProfileRequest,
        preloads?: string[]
    ): Promise<IMemberProfileResource> {
        const url = qs.stringifyUrl({
            url: `${MemberProfileService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        const response = await APIService.put<
            IMemberProfileRequest,
            IMemberProfileResource
        >(url, memberData, {
            headers: {
                Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
            },
        })
        return response.data
    }

    public static async closeAccount(
        id: TEntityId,
        closeRemark: IMemberCloseRemarkRequest[],
        preloads?: string[]
    ) {
        const url = qs.stringifyUrl({
            url: `${MemberProfileService.BASE_ENDPOINT}/${id}/close-account`,
            query: { preloads },
        })

        const response = await APIService.put<
            IMemberProfileCloseRequest,
            IMemberProfileResource
        >(url, closeRemark, {
            headers: {
                Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
            },
        })
        return response.data
    }

    // public static async getMemberProfilesForPicker(
    //     params: IMemberProfilePickerParams
    // ): Promise<IMemberProfilePaginatedPicker> {
    //     const { filters, pagination, sort } = params

    //     const url = qs.stringifyUrl(
    //         {
    //             url: `${MemberProfileService.BASE_ENDPOINT}/picker`,
    //             query: {
    //                 sort,
    //                 filters,
    //                 pageIndex: pagination?.pageIndex,
    //                 pageSize: pagination?.pageSize,
    //             },
    //         },
    //         { skipNull: true }
    //     )

    //     const response =
    //         await APIService.get<IMemberProfilePaginatedPicker>(url)
    //     return response.data
    // }
}
