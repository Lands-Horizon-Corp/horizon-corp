import qs from 'query-string'

import APIService from './api-service'
import { downloadFile } from '../helpers'

import {
    TEntityId,
    IGenderRequest,
    IGenderResource,
    IGenderPaginatedResource,
} from '../types'

/**
 * Service class to handle CRUD operations for genders.
 */
export default class GenderService {
    private static readonly BASE_ENDPOINT = '/gender'

    public static async getAll(): Promise<IGenderResource[]> {
        const response = await APIService.get<IGenderResource[]>(
            GenderService.BASE_ENDPOINT
        )
        return response.data
    }

    public static async create(
        genderData: IGenderRequest
    ): Promise<IGenderResource> {
        const response = await APIService.post<IGenderRequest, IGenderResource>(
            GenderService.BASE_ENDPOINT,
            genderData
        )
        return response.data
    }

    public static async delete(id: TEntityId): Promise<void> {
        const endpoint = `${GenderService.BASE_ENDPOINT}/${id}`
        await APIService.delete<void>(endpoint)
    }

    public static async update(
        id: TEntityId,
        genderData: IGenderRequest
    ): Promise<IGenderResource> {
        const endpoint = `${GenderService.BASE_ENDPOINT}/${id}`
        const response = await APIService.put<IGenderRequest, IGenderResource>(
            endpoint,
            genderData
        )
        return response.data
    }

    public static async getById(id: TEntityId): Promise<IGenderResource> {
        const url = qs.stringifyUrl(
            {
                url: `${GenderService.BASE_ENDPOINT}/${id}`,
            },
            { skipNull: true }
        )
        const response = await APIService.get<IGenderResource>(url)
        return response.data
    }

    public static async getGenders(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<IGenderPaginatedResource> {
        const { filters, preloads, pagination, sort } = props || {}

        const url = qs.stringifyUrl(
            {
                url: `${GenderService.BASE_ENDPOINT}/search`,
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

        const response = await APIService.get<IGenderPaginatedResource>(url)
        return response.data
    }

    public static async exportAll(): Promise<void> {
        const url = `${GenderService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_genders_export.xlsx')
    }

    public static async exportAllFiltered(filters?: string): Promise<void> {
        const url = `${GenderService.BASE_ENDPOINT}/export-search?filter=${filters || ''}`
        await downloadFile(url, 'filtered_genders_export.xlsx')
    }

    public static async exportSelected(ids: TEntityId[]): Promise<void> {
        const url = qs.stringifyUrl(
            {
                url: `${GenderService.BASE_ENDPOINT}/export-selected`,
                query: { ids },
            },
            { skipNull: true }
        )

        await downloadFile(url, 'selected_genders_export.xlsx')
    }

    public static async exportCurrentPage(page: number): Promise<void> {
        const url = `${GenderService.BASE_ENDPOINT}/export-current-page/${page}`
        await downloadFile(url, `current_page_genders_${page}_export.xlsx`)
    }
}
