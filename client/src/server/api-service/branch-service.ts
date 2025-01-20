// services/BranchService.ts
import qs from 'query-string'

import {
    IMediaRequest,
    IBranchRequest,
    IBranchResource,
    IBranchPaginatedResource,
} from '@/server/types'
import { downloadFile } from '@/server/helpers'
import UseServer from '@/horizon-corp/request/server'

/**
 * Service class to handle CRUD operations for branches.
 */
export default class BranchService {
    private static readonly BASE_ENDPOINT = '/branch'

    public static async getById(
        id: number,
        preloads?: string[]
    ): Promise<IBranchResource> {
        // Construct each preload as a separate 'preloads' query parameter
        const preloadParams =
            preloads
                ?.map((preload) => `preloads=${encodeURIComponent(preload)}`)
                .join('&') || ''
        const separator = preloadParams ? '?' : ''
        const endpoint = `${BranchService.BASE_ENDPOINT}/${id}${separator}${preloadParams}`

        // Make the GET request with necessary headers
        const response = await UseServer.get<IBranchResource>(endpoint, {
            headers: {
                Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
            },
        })

        return response.data
    }

  
    public static async create(
        branchData: IBranchRequest
    ): Promise<IBranchResource> {
        const response = await UseServer.post<IBranchRequest, IBranchResource>(
            BranchService.BASE_ENDPOINT,
            branchData
        )
        return response.data
    }

    public static async delete(id: number): Promise<void> {
        const endpoint = `${BranchService.BASE_ENDPOINT}/${id}`
        await UseServer.delete<void>(endpoint)
    }

    public static async update(
        id: number,
        branchData: IBranchRequest,
        preloads?: string[]
    ): Promise<IBranchResource> {
        const preloadParams =
            preloads
                ?.map((preload) => `preloads=${encodeURIComponent(preload)}`)
                .join('&') || ''
        const separator = preloadParams ? '?' : ''
        const endpoint = `${BranchService.BASE_ENDPOINT}/${id}${separator}${preloadParams}`

        const response = await UseServer.put<IBranchRequest, IBranchResource>(
            endpoint,
            branchData,
            {
                headers: {
                    Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
                },
            }
        )
        return response.data
    }

    public static async getBranches(props?: {
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
        sort?: string
    }): Promise<IBranchPaginatedResource> {
        const { filters, preloads, pagination, sort } = props || {}

        const url = qs.stringifyUrl(
            {
                url: `${BranchService.BASE_ENDPOINT}`,
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

        const response = await UseServer.get<IBranchPaginatedResource>(url)
        return response.data
    }

    public static async exportAll(): Promise<void> {
        const url = `${BranchService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_branches_export.csv')
    }

    public static async exportAllFiltered(filters?: string): Promise<void> {
        const filterQuery = filters
            ? `filter=${encodeURIComponent(filters)}`
            : ''
        const url = `${BranchService.BASE_ENDPOINT}/export-search${filterQuery ? `?${filterQuery}` : ''}`
        await downloadFile(url, 'filtered_branches_export.csv')
    }

    public static async exportSelected(ids: number[]): Promise<void> {
        if (ids.length === 0) {
            throw new Error('No branch IDs provided for export.')
        }

        // Construct each preload as a separate 'preloads' query parameter if needed
        // (Assuming export-selected might also accept preloads; if not, you can omit this)

        const query = ids.map((id) => `ids=${encodeURIComponent(id)}`).join('&')
        const url = `${BranchService.BASE_ENDPOINT}/export-selected?${query}`
        await downloadFile(url, 'selected_branches_export.csv')
    }

    public static async deleteMany(ids: number[]): Promise<void> {
        const endpoint = `${BranchService.BASE_ENDPOINT}/bulk-delete`

        // Construct the request payload
        const payload = { ids }

        // Make the DELETE request with the payload
        await UseServer.delete<void>(endpoint, payload)
    }

    public static async verify(id: number): Promise<IBranchResource> {
        const endpoint = `${BranchService.BASE_ENDPOINT}/verify/${id}`
        const response = await UseServer.post<void, IBranchResource>(endpoint)
        return response.data
    }

    public static async ProfilePicture(
        id: number,
        data: IMediaRequest,
        preloads: string[] = ['Media']
    ): Promise<IBranchResource> {
        const preloadParams =
            preloads
                ?.map((preload) => `preloads=${encodeURIComponent(preload)}`)
                .join('&') || ''
        const separator = preloadParams ? '?' : ''
        const endpoint = `${BranchService.BASE_ENDPOINT}/profile-picture/${id}${separator}${preloadParams}`
        const response = await UseServer.post<IMediaRequest, IBranchResource>(
            endpoint,
            data
        )
        return response.data
    }
}
