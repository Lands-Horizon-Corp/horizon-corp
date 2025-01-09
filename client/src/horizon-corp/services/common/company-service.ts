// services/CompanyService.ts
import qs from 'query-string'

import { downloadFile } from '@/horizon-corp/helpers'
import UseServer from '@/horizon-corp/request/server'
import {
    MediaRequest,
    CompanyRequest,
    CompanyResource,
    CompanyPaginatedResource,
} from '@/horizon-corp/types'

/**
 * Service class to handle CRUD operations for companies.
 */
export default class CompanyService {
    private static readonly BASE_ENDPOINT = '/company'

    /**
     * Retrieves a company by its ID with optional preloads.
     *
     * @param {number} id - The ID of the company to retrieve.
     * @param {string[]} [preloads] - Optional array of relations to preload.
     * @returns {Promise<CompanyResource>} - A promise that resolves to the company resource.
     */
    public static async getById(
        id: number,
        preloads?: string[]
    ): Promise<CompanyResource> {
        // Construct each preload as a separate 'preloads' query parameter
        const preloadParams =
            preloads
                ?.map((preload) => `preloads=${encodeURIComponent(preload)}`)
                .join('&') || ''
        const separator = preloadParams ? '?' : ''
        const endpoint = `${CompanyService.BASE_ENDPOINT}/${id}${separator}${preloadParams}`

        // Make the GET request with necessary headers
        const response = await UseServer.get<CompanyResource>(endpoint, {
            headers: {
                Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
            },
        })

        return response.data
    }

    /**
     * Creates a new company.
     *
     * @param {CompanyRequest} companyData - The data for the new company.
     * @returns {Promise<CompanyResource>} - A promise that resolves to the created company resource.
     */
    public static async create(
        companyData: CompanyRequest
    ): Promise<CompanyResource> {
        const response = await UseServer.post<CompanyRequest, CompanyResource>(
            CompanyService.BASE_ENDPOINT,
            companyData
        )
        return response.data
    }

    /**
     * Deletes a company by its ID.
     *
     * @param {number} id - The ID of the company to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async delete(id: number): Promise<void> {
        const endpoint = `${CompanyService.BASE_ENDPOINT}/${id}`
        await UseServer.delete<void>(endpoint)
    }

    /**
     * Updates an existing company by its ID with optional preloads.
     *
     * @param {number} id - The ID of the company to update.
     * @param {CompanyRequest} companyData - The updated data for the company.
     * @param {string[]} [preloads] - Optional array of relations to preload.
     * @returns {Promise<CompanyResource>} - A promise that resolves to the updated company resource.
     */
    public static async update(
        id: number,
        companyData: CompanyRequest,
        preloads?: string[]
    ): Promise<CompanyResource> {
        // Construct each preload as a separate 'preloads' query parameter
        const preloadParams =
            preloads
                ?.map((preload) => `preloads=${encodeURIComponent(preload)}`)
                .join('&') || ''
        const separator = preloadParams ? '?' : ''
        const endpoint = `${CompanyService.BASE_ENDPOINT}/${id}${separator}${preloadParams}`

        // Make the PUT request with necessary headers
        const response = await UseServer.put<CompanyRequest, CompanyResource>(
            endpoint,
            companyData,
            {
                headers: {
                    Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
                },
            }
        )
        return response.data
    }

    /**
     * Filters companies based on provided options with optional pagination.
     *
     * @param {Object} [props] - The options for filtering companies.
     * @param {string} [props.filters] - The filters to apply for exporting companies.
     * @param {string[]} [props.preloads] - Optional array of relations to preload.
     * @param {Object} [props.pagination] - Pagination parameters.
     * @param {number} props.pagination.pageIndex - The current page index.
     * @param {number} props.pagination.pageSize - The number of items per page.
     * @returns {Promise<CompanyPaginatedResource>} - A promise that resolves to paginated company resources.
     */
    public static async getCompanies(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<CompanyPaginatedResource> {
        const { filters, preloads, pagination, sort } = props || {}

        const url = qs.stringifyUrl(
            {
                url: `${CompanyService.BASE_ENDPOINT}`,
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

        // Output URL
        // /api/v1/company?filter=eyJmaWx0ZXJzIjpbXSwibG9naWMiOiJBTkQifQ%3D%3D&pageIndex=0&pageSize=10&preloads=Media&preloads=Owner

        const response = await UseServer.get<CompanyPaginatedResource>(url)
        return response.data
    }

    /**
     * Exports all companies.
     *
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportAll(): Promise<void> {
        const url = `${CompanyService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_companies_export.csv')
    }

    /**
     * Exports filtered companies.
     *
     * @param {string} [filters] - The filters to apply for exporting companies.
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportAllFiltered(filters?: string): Promise<void> {
        const filterQuery = filters
            ? `filter=${encodeURIComponent(filters)}`
            : ''
        const url = `${CompanyService.BASE_ENDPOINT}/export-search${filterQuery ? `?${filterQuery}` : ''}`
        await downloadFile(url, 'filtered_companies_export.csv')
    }

    /**
     * Exports selected companies.
     *
     * @param {number[]} ids - The IDs of the selected companies to export.
     * @returns {Promise<void>} - A promise that resolves when the export is complete.
     */
    public static async exportSelected(ids: number[]): Promise<void> {
        if (ids.length === 0) {
            throw new Error('No company IDs provided for export.')
        }

        // Construct each preload as a separate 'preloads' query parameter if needed
        // (Assuming export-selected might also accept preloads; if not, you can omit this)

        const query = ids.map((id) => `ids=${encodeURIComponent(id)}`).join('&')
        const url = `${CompanyService.BASE_ENDPOINT}/export-selected?${query}`
        await downloadFile(url, 'selected_companies_export.csv')
    }

    /**
     * Deletes multiple companies by their IDs.
     *
     * @param {number[]} ids - The IDs of the companies to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async deleteMany(ids: number[]): Promise<void> {
        const endpoint = `${CompanyService.BASE_ENDPOINT}/bulk-delete`

        // Construct the request payload
        const payload = { ids }

        // Make the DELETE request with the payload
        await UseServer.delete<void>(endpoint, payload)
    }

    /**
     * Verifies a branch by its ID.
     *
     * @param {number} id - The ID of the branch to verify.
     * @returns {Promise<BranchResource>} - A promise that resolves to the verified branch resource.
     */
    public static async verify(id: number): Promise<CompanyResource> {
        const endpoint = `${CompanyService.BASE_ENDPOINT}/verify/${id}`
        const response = await UseServer.post<void, CompanyResource>(endpoint)
        return response.data
    }

    /**
     * Uploads a profile picture for the company.
     *
     * @async
     * @function
     * @param {MediaRequest} data - The request payload containing the media information.
     * @param {string[]} [preloads=["Media"]] - Optional array of relations to preload.
     * @returns {Promise<CompanyResource>} - The response containing the updated company resource.
     */
    public static async ProfilePicture(
        id: number,
        data: MediaRequest,
        preloads: string[] = ['Media']
    ): Promise<CompanyResource> {
        const preloadParams =
            preloads
                ?.map((preload) => `preloads=${encodeURIComponent(preload)}`)
                .join('&') || ''
        const separator = preloadParams ? '?' : ''
        const endpoint = `${CompanyService.BASE_ENDPOINT}/profile-picture/${id}${separator}${preloadParams}`
        const response = await UseServer.post<MediaRequest, CompanyResource>(
            endpoint,
            data
        )
        return response.data
    }
}
