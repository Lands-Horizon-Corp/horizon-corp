// services/MemberService.ts
import qs from 'query-string'

import APIService from './api-service'
import { downloadFile } from '../helpers'
import {
    IMediaRequest,
    IMemberPaginatedResource,
    IMemberRequest,
    IMemberResource,
} from '../types'

/**
 * Service class to handle CRUD operations for members.
 */
export default class MemberService {
    private static readonly BASE_ENDPOINT = '/member'

    public static async getById(
        id: number,
        preloads?: string[]
    ): Promise<IMemberResource> {
        const url = qs.stringifyUrl({
            url: `${MemberService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        const response = await APIService.get<IMemberResource>(url, {
            headers: {
                Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
            },
        })

        return response.data
    }

    public static async create(data: IMemberRequest, preloads?: string[]) {
        const url = qs.stringifyUrl(
            {
                url: `${MemberService.BASE_ENDPOINT}`,
                query: { preloads },
            },
            { skipNull: true }
        )

        return (
            await APIService.post<IMemberRequest, IMemberResource>(url, data)
        ).data
    }

    public static async delete(id: number): Promise<void> {
        const endpoint = `${MemberService.BASE_ENDPOINT}/${id}`
        await APIService.delete<void>(endpoint)
    }

    public static async update(
        id: number,
        memberData: IMemberRequest,
        preloads?: string[]
    ): Promise<IMemberResource> {
        const url = qs.stringifyUrl({
            url: `${MemberService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        const response = await APIService.put<IMemberRequest, IMemberResource>(
            url,
            memberData,
            {
                headers: {
                    Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
                },
            }
        )
        return response.data
    }

    public static async getMembers(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<IMemberPaginatedResource> {
        const { filters, preloads, pagination, sort } = props || {}

        const url = qs.stringifyUrl(
            {
                url: `${MemberService.BASE_ENDPOINT}`,
                query: {
                    sort,
                    preloads,
                    filter: filters,
                    pageIndex: pagination?.pageIndex,
                    pageSize: pagination?.pageSize,
                },
            },
            { skipNull: true }
        )

        const response = await APIService.get<IMemberPaginatedResource>(url)
        return response.data
    }

    public static async exportAll(): Promise<void> {
        const url = `${MemberService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_members_export.csv')
    }

    public static async exportAllFiltered(filters?: string): Promise<void> {
        const filterQuery = filters
            ? `filter=${encodeURIComponent(filters)}`
            : ''
        const url = `${MemberService.BASE_ENDPOINT}/export-search${filterQuery ? `?${filterQuery}` : ''}`
        await downloadFile(url, 'filtered_members_export.csv')
    }

    public static async exportSelected(ids: number[]): Promise<void> {
        if (ids.length === 0) {
            throw new Error('No member IDs provided for export.')
        }

        const query = ids.map((id) => `ids=${encodeURIComponent(id)}`).join('&')
        const url = `${MemberService.BASE_ENDPOINT}/export-selected?${query}`
        await downloadFile(url, 'selected_members_export.csv')
    }

    public static async deleteMany(ids: number[]): Promise<void> {
        const endpoint = `${MemberService.BASE_ENDPOINT}/bulk-delete`

        const payload = { ids }

        await APIService.delete<void>(endpoint, payload)
    }

    public static async profilePicture(
        id: number,
        data: IMediaRequest,
        preloads: string[] = ['Media']
    ): Promise<IMemberResource> {
        const preloadParams =
            preloads
                ?.map((preload) => `preloads=${encodeURIComponent(preload)}`)
                .join('&') || ''
        const separator = preloadParams ? '?' : ''
        const endpoint = `${MemberService.BASE_ENDPOINT}/profile-picture/${id}${separator}${preloadParams}`
        const response = await APIService.post<IMediaRequest, IMemberResource>(
            endpoint,
            data
        )
        return response.data
    }
}
