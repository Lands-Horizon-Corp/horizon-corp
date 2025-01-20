// services/CompanyService.ts
import qs from 'query-string'

import {
    IMediaRequest,
    ICompanyRequest,
    ICompanyResource,
    ICompanyPaginatedResource,
} from '@/server/types'
import { downloadFile } from '@/horizon-corp/helpers'
import APIService from './api-service'

/**
 * Service class to handle CRUD operations for companies.
 */
export default class CompanyService {
    private static readonly BASE_ENDPOINT = '/company'

    public static async getById(
        id: number,
        preloads?: string[]
    ): Promise<ICompanyResource> {
        // Construct each preload as a separate 'preloads' query parameter
        const url = qs.stringifyUrl({
            url: `${CompanyService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        // Make the GET request with necessary headers
        const response = await APIService.get<ICompanyResource>(url, {
            headers: {
                Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
            },
        })

        return response.data
    }

    public static async create(
        companyData: ICompanyRequest
    ): Promise<ICompanyResource> {
        const response = await APIService.post<
            ICompanyRequest,
            ICompanyResource
        >(CompanyService.BASE_ENDPOINT, companyData)
        return response.data
    }

    public static async delete(id: number): Promise<void> {
        const endpoint = `${CompanyService.BASE_ENDPOINT}/${id}`
        await APIService.delete(endpoint)
    }

    public static async update(
        id: number,
        companyData: ICompanyRequest,
        preloads?: string[]
    ): Promise<ICompanyResource> {
        // Construct each preload as a separate 'preloads' query parameter
        const url = qs.stringifyUrl({
            url: `${CompanyService.BASE_ENDPOINT}/${id}`,
            query: { preloads },
        })

        // Make the PUT request with necessary headers
        const response = await APIService.put<
            ICompanyRequest,
            ICompanyResource
        >(url, companyData, {
            headers: {
                Authorization: `Bearer YOUR_TOKEN`, // Replace with actual token if needed
            },
        })
        return response.data
    }

    public static async getCompanies(props?: {
        sort?: string
        filters?: string
        preloads?: string[]
        pagination?: { pageIndex: number; pageSize: number }
    }): Promise<ICompanyPaginatedResource> {
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

        const response = await APIService.get<ICompanyPaginatedResource>(url)
        return response.data
    }

    public static async exportAll(): Promise<void> {
        const url = `${CompanyService.BASE_ENDPOINT}/export`
        await downloadFile(url, 'all_companies_export.csv')
    }

    public static async exportAllFiltered(filters?: string): Promise<void> {
        const filterQuery = filters
            ? `filter=${encodeURIComponent(filters)}`
            : ''
        const url = `${this.BASE_ENDPOINT}/export-search${filterQuery ? `?${filterQuery}` : ''}`
        await downloadFile(url, 'filtered_companies_export.csv')
    }

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

    public static async deleteMany(ids: number[]): Promise<void> {
        const endpoint = `${CompanyService.BASE_ENDPOINT}/bulk-delete`

        // Construct the request payload
        const payload = { ids }

        // Make the DELETE request with the payload
        await APIService.delete<void>(endpoint, payload)
    }

    public static async verify(id: number): Promise<ICompanyResource> {
        const endpoint = `${CompanyService.BASE_ENDPOINT}/verify/${id}`
        const response = await APIService.post<void, ICompanyResource>(endpoint)
        return response.data
    }

    public static async ProfilePicture(
        id: number,
        data: IMediaRequest,
        preloads: string[] = ['Media']
    ): Promise<ICompanyResource> {
        const preloadParams =
            preloads
                ?.map((preload) => `preloads=${encodeURIComponent(preload)}`)
                .join('&') || ''
        const separator = preloadParams ? '?' : ''
        const endpoint = `${CompanyService.BASE_ENDPOINT}/profile-picture/${id}${separator}${preloadParams}`
        const response = await APIService.post<IMediaRequest, ICompanyResource>(
            endpoint,
            data
        )
        return response.data
    }
}
