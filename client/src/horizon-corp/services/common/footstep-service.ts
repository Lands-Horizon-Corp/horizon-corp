import qs from 'query-string'

import { downloadFile } from '@/horizon-corp/helpers'
import UseServer from '@/horizon-corp/request/server'

import { FootstepPaginatedResource } from '@/horizon-corp/types'

/**
 * Service class to handle CRUD operations for footsteps.
 */
export default class FootstepService {
    private static readonly BASE_ENDPOINT = '/footstep'

    /**
     * Exports all footsteps.
     *
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportAll(): Promise<void> {
        const url = `${FootstepService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_footsteps_export.csv')
    }

    /**
     * Exports filtered footsteps.
     *
     * @param {string} [filters] - The filters to apply for exporting footsteps.
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportAllFiltered(filters?: string): Promise<void> {
        const filterQuery = filters
            ? `filter=${encodeURIComponent(filters)}`
            : ''
        const url = `${FootstepService.BASE_ENDPOINT}/export-search${filterQuery ? `?${filterQuery}` : ''}`
        await downloadFile(url, 'filtered_footsteps_export.csv')
    }

    /**
     * Exports selected footsteps.
     *
     * @param {string[]} ids - The IDs of the selected footsteps to export.
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportSelected(ids: string[]): Promise<void> {
        if (ids.length === 0) {
            throw new Error('No footstep IDs provided for export.')
        }

        const query = ids.map((id) => `ids=${encodeURIComponent(id)}`).join('&')
        const url = `${FootstepService.BASE_ENDPOINT}/export-selected?${query}`
        await downloadFile(url, 'selected_footsteps_export.csv')
    }

    /**
     * Retrieves a paginated list of footsteps based on provided options.
     *
     * @param {Object} [props] - The options for filtering footsteps.
     * @param {string} [props.filters] - The filters to apply for exporting footsteps.
     * @param {string[]} [props.preloads] - Optional array of relations to preload.
     * @param {Object} [props.pagination] - Pagination parameters.
     * @param {number} props.pagination.pageIndex - The current page index.
     * @param {number} props.pagination.pageSize - The number of items per page.
     * @returns {Promise<FootstepPaginatedResource>} - A promise that resolves to paginated footstep resources.
     */
    public static async getFootsteps(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<FootstepPaginatedResource> {
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

        const response = await UseServer.get<FootstepPaginatedResource>(url)
        return response.data
    }
}
