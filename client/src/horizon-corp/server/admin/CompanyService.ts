// services/CompanyService.ts
import { downloadFile } from '@/horizon-corp/helpers'
import UseServer from '../../request/server'
import {
  CompanyResource,
  CompanyRequest,
  CompanyPaginatedResource,
} from '../../types'

/**
 * Service class to handle CRUD operations for companies.
 */
export default class CompanyService {
  private static readonly BASE_ENDPOINT = '/company'

  /**
   * Retrieves all companies with optional preloads.
   *
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<CompanyResource[]>} - A promise that resolves to an array of company resources.
   */
  public static async getAll(preloads?: string[]): Promise<CompanyResource[]> {
    const query = preloads ? `?preloads=${preloads.join(',')}` : ''
    const endpoint = `${CompanyService.BASE_ENDPOINT}${query}`
    const response = await UseServer.get<CompanyResource[]>(endpoint)
    return response.data
  }

  /**
   * Retrieves a company by its ID with optional preloads.
   *
   * @param {number} id - The ID of the company to retrieve.
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<CompanyResource>} - A promise that resolves to the company resource.
   */
  public static async getById(id: number, preloads?: string[]): Promise<CompanyResource> {
    const query = preloads ? `?preloads=${preloads.join(',')}` : ''
    const endpoint = `${CompanyService.BASE_ENDPOINT}/${id}${query}`
    const response = await UseServer.get<CompanyResource>(endpoint)
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
    const query = preloads ? `?preloads=${preloads.join(',')}` : ''
    const endpoint = `${CompanyService.BASE_ENDPOINT}/${id}${query}`
    const response = await UseServer.put<CompanyRequest, CompanyResource>(
      endpoint,
      companyData
    )
    return response.data
  }

  /**
   * Filters companies based on provided filters with optional preloads.
   *
   * @param {string} [filters] - The filters to apply for exporting companies.
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<CompanyPaginatedResource>} - A promise that resolves to paginated company resources.
   */
  public static async filter(
    filters?: string,
    preloads?: string[]
  ): Promise<CompanyPaginatedResource> {
    const queryFilters = filters ? `filter=${encodeURIComponent(filters)}` : ''
    const queryPreloads = preloads ? `preloads=${preloads.join(',')}` : ''
    const query = [queryFilters, queryPreloads].filter(Boolean).join('&')
    const url = `${CompanyService.BASE_ENDPOINT}/search?${query}`
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
    const url = `${CompanyService.BASE_ENDPOINT}/export-search?filter=${encodeURIComponent(filters || '')}`
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
}
