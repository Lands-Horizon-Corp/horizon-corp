import qs from 'query-string'

import {
    IOwnerRequest,
    IOwnerResource,
    IOwnerPaginatedResource,
} from '../types'
import APIService from './api-service'
import { downloadFile } from '../helpers'

/**
 * Service class to handle CRUD operations for owners.
 */
export default class OwnerService {
    private static readonly BASE_ENDPOINT = '/owner'

    public static async getById(
        id: number,
        preloads?: string[]
    ): Promise<IOwnerResource> {
        const url = qs.stringifyUrl({
            url: `${OwnerService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        const response = await APIService.get<IOwnerResource>(url)
        return response.data
    }

    public static async create(
        ownerData: IOwnerRequest
    ): Promise<IOwnerResource> {
        const response = await APIService.post<IOwnerRequest, IOwnerResource>(
            OwnerService.BASE_ENDPOINT,
            ownerData
        )
        return response.data
    }

    public static async update(
        id: number,
        ownerData: IOwnerRequest,
        preloads?: string[]
    ): Promise<IOwnerResource> {
        const url = qs.stringifyUrl({
            url: `${OwnerService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        const response = await APIService.put<IOwnerRequest, IOwnerResource>(
            url,
            ownerData
        )
        return response.data
    }

    public static async delete(id: number): Promise<void> {
        const endpoint = `${OwnerService.BASE_ENDPOINT}/${id}`
        await APIService.delete<void>(endpoint)
    }

    public static async getOwners(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<IOwnerPaginatedResource> {
        const { filters, preloads, pagination, sort } = props || {}

        const url = qs.stringifyUrl(
            {
                url: `${OwnerService.BASE_ENDPOINT}`,
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

        const response = await APIService.get<IOwnerPaginatedResource>(url)
        return response.data
    }

    public static async verify(id: number): Promise<IOwnerResource> {
        const endpoint = `${OwnerService.BASE_ENDPOINT}/verify/${id}`
        const response = await APIService.post<void, IOwnerResource>(endpoint)
        return response.data
    }

    public static async exportAll(): Promise<void> {
        const url = `${OwnerService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_owners_export.csv')
    }

    public static async exportSelected(ids: number[]): Promise<void> {
        if (ids.length === 0) {
            throw new Error('No owner IDs provided for export.')
        }

        const url = qs.stringifyUrl({
            url: `${OwnerService.BASE_ENDPOINT}/export-selected`,
            query: { ids },
        })

        await downloadFile(url, 'selected_owners_export.csv')
    }

    public static async deleteMany(ids: number[]): Promise<void> {
        const url = `${OwnerService.BASE_ENDPOINT}/bulk-delete`
        const payload = { ids }

        await APIService.delete<void>(url, payload)
    }
}
