import qs from 'query-string'

import {
    OwnerRequest,
    OwnerResource,
    OwnerPaginatedResource,
} from '@/horizon-corp/types/owner'
import { downloadFile } from '@/horizon-corp/helpers'
import UseServer from '@/horizon-corp/request/server'

/**
 * Service class to handle CRUD operations for owners.
 */
export default class OwnerService {
    private static readonly BASE_ENDPOINT = '/owner'

    /**
     * Retrieves an owner by their ID with optional preloads.
     *
     * @param {number} id - The ID of the owner to retrieve.
     * @param {string[]} [preloads] - Optional array of relations to preload.
     * @returns {Promise<OwnerResource>} - A promise that resolves to the owner resource.
     */
    public static async getById(
        id: number,
        preloads?: string[]
    ): Promise<OwnerResource> {
        const url = qs.stringifyUrl({
            url: `${OwnerService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        const response = await UseServer.get<OwnerResource>(url)
        return response.data
    }

    /**
     * Creates a new owner.
     *
     * @param {OwnerRequest} ownerData - The data for the new owner.
     * @returns {Promise<OwnerResource>} - A promise that resolves to the created owner resource.
     */
    public static async create(
        ownerData: OwnerRequest
    ): Promise<OwnerResource> {
        const response = await UseServer.post<OwnerRequest, OwnerResource>(
            OwnerService.BASE_ENDPOINT,
            ownerData
        )
        return response.data
    }

    /**
     * Updates an existing owner by their ID with optional preloads.
     *
     * @param {number} id - The ID of the owner to update.
     * @param {OwnerRequest} ownerData - The updated data for the owner.
     * @param {string[]} [preloads] - Optional array of relations to preload.
     * @returns {Promise<OwnerResource>} - A promise that resolves to the updated owner resource.
     */
    public static async update(
        id: number,
        ownerData: OwnerRequest,
        preloads?: string[]
    ): Promise<OwnerResource> {
        const url = qs.stringifyUrl({
            url: `${OwnerService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        const response = await UseServer.put<OwnerRequest, OwnerResource>(
            url,
            ownerData
        )
        return response.data
    }

    /**
     * Deletes an owner by their ID.
     *
     * @param {number} id - The ID of the owner to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async delete(id: number): Promise<void> {
        const endpoint = `${OwnerService.BASE_ENDPOINT}/${id}`
        await UseServer.delete<void>(endpoint)
    }

    /**
     * Filters owners based on provided options with optional pagination.
     *
     * @param {Object} [props] - The options for filtering owners.
     * @param {string} [props.filters] - The filters to apply for searching owners.
     * @param {string[]} [props.preloads] - Optional array of relations to preload.
     * @param {Object} [props.pagination] - Pagination parameters.
     * @param {number} props.pagination.pageIndex - The current page index.
     * @param {number} props.pagination.pageSize - The number of items per page.
     * @returns {Promise<OwnerPaginatedResource>} - A promise that resolves to paginated owner resources.
     */
    public static async getOwners(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<OwnerPaginatedResource> {
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

        const response = await UseServer.get<OwnerPaginatedResource>(url)
        return response.data
    }

    /**
     * Verifies an owner by their ID.
     *
     * @param {number} id - The ID of the owner to verify.
     * @returns {Promise<OwnerResource>} - A promise that resolves to the verified owner resource.
     */
    public static async verify(id: number): Promise<OwnerResource> {
        const endpoint = `${OwnerService.BASE_ENDPOINT}/verify/${id}`
        const response = await UseServer.post<void, OwnerResource>(endpoint)
        return response.data
    }

    /**
     * Exports all owners.
     *
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportAll(): Promise<void> {
        const url = `${OwnerService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_owners_export.csv')
    }

    /**
     * Exports selected owners.
     *
     * @param {number[]} ids - The IDs of the selected owners to export.
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
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

    /**
     * Deletes multiple owners by their IDs.
     *
     * @param {number[]} ids - The IDs of the owners to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async deleteMany(ids: number[]): Promise<void> {
        const url = `${OwnerService.BASE_ENDPOINT}/bulk-delete`
        const payload = { ids }

        await UseServer.delete<void>(url, payload)
    }
}
