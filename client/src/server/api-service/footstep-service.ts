import qs from 'query-string'

import APIService from './api-service'
import { downloadFile } from '../helpers'

import { IFootstepPaginatedResource } from '@/server/types'

/**
 * Service class to handle CRUD operations for footsteps.
 */
export default class FootstepService {
    private static readonly BASE_ENDPOINT = '/footstep'

    public static async exportAll(): Promise<void> {
        const url = `${FootstepService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_footsteps_export.csv')
    }

    public static async exportAllFiltered(filters?: string): Promise<void> {
        const filterQuery = filters
            ? `filter=${encodeURIComponent(filters)}`
            : ''
        const url = `${FootstepService.BASE_ENDPOINT}/export-search${filterQuery ? `?${filterQuery}` : ''}`
        await downloadFile(url, 'filtered_footsteps_export.csv')
    }

    public static async exportSelected(ids: string[]): Promise<void> {
        if (ids.length === 0) {
            throw new Error('No footstep IDs provided for export.')
        }

        const query = ids.map((id) => `ids=${encodeURIComponent(id)}`).join('&')
        const url = `${FootstepService.BASE_ENDPOINT}/export-selected?${query}`
        await downloadFile(url, 'selected_footsteps_export.csv')
    }

    public static async getFootsteps(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<IFootstepPaginatedResource> {
        const { filters, preloads, pagination, sort } = props || {}

        const url = qs.stringifyUrl(
            {
                url: `${FootstepService.BASE_ENDPOINT}`,
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

        const response = await APIService.get<IFootstepPaginatedResource>(url)
        return response.data
    }
}
