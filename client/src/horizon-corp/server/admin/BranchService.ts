import { downloadFile } from '@/horizon-corp/helpers'
import UseServer from '../../request/server'
import {
  BranchResource,
  BranchRequest,
  BranchPaginatedResource,
} from '../../types'

/**
 * Service class to handle CRUD operations for branches.
 */
export default class BranchService {
  private static readonly BASE_ENDPOINT = '/branch'

  /**
   * Retrieves all branches with optional preloads.
   *
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<BranchResource[]>} - A promise that resolves to an array of branch resources.
   */
  public static async getAll(preloads?: string[]): Promise<BranchResource[]> {
    const query = preloads ? `?preloads=${preloads.join(',')}` : ''
    const endpoint = `${BranchService.BASE_ENDPOINT}${query}`
    const response = await UseServer.get<BranchResource[]>(endpoint)
    return response.data
  }

  /**
   * Retrieves a branch by its ID with optional preloads.
   *
   * @param {number} id - The ID of the branch to retrieve.
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<BranchResource>} - A promise that resolves to the branch resource.
   */
  public static async getById(id: number, preloads?: string[]): Promise<BranchResource> {
    const query = preloads ? `?preloads=${preloads.join(',')}` : ''
    const endpoint = `${BranchService.BASE_ENDPOINT}/${id}${query}`
    const response = await UseServer.get<BranchResource>(endpoint)
    return response.data
  }

  /**
   * Creates a new branch.
   *
   * @param {BranchRequest} branchData - The data for the new branch.
   * @returns {Promise<BranchResource>} - A promise that resolves to the created branch resource.
   */
  public static async create(
    branchData: BranchRequest
  ): Promise<BranchResource> {
    const response = await UseServer.post<BranchRequest, BranchResource>(
      BranchService.BASE_ENDPOINT,
      branchData
    )
    return response.data
  }

  /**
   * Deletes a branch by its ID.
   *
   * @param {number} id - The ID of the branch to delete.
   * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
   */
  public static async delete(id: number): Promise<void> {
    const endpoint = `${BranchService.BASE_ENDPOINT}/${id}`
    await UseServer.delete<void>(endpoint)
  }

  /**
   * Updates an existing branch by its ID with optional preloads.
   *
   * @param {number} id - The ID of the branch to update.
   * @param {BranchRequest} branchData - The updated data for the branch.
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<BranchResource>} - A promise that resolves to the updated branch resource.
   */
  public static async update(
    id: number,
    branchData: BranchRequest,
    preloads?: string[]
  ): Promise<BranchResource> {
    const query = preloads ? `?preloads=${preloads.join(',')}` : ''
    const endpoint = `${BranchService.BASE_ENDPOINT}/${id}${query}`
    const response = await UseServer.put<BranchRequest, BranchResource>(
      endpoint,
      branchData
    )
    return response.data
  }

  /**
   * Filters branches based on provided filters with optional preloads.
   *
   * @param {string} [filters] - The filters to apply for exporting branches.
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<BranchPaginatedResource>} - A promise that resolves to paginated branch resources.
   */
  public static async filter(
    filters?: string,
    preloads?: string[]
  ): Promise<BranchPaginatedResource> {
    const queryFilters = filters ? `filter=${encodeURIComponent(filters)}` : ''
    const queryPreloads = preloads ? `preloads=${preloads.join(',')}` : ''
    const query = [queryFilters, queryPreloads].filter(Boolean).join('&')
    const url = `${BranchService.BASE_ENDPOINT}/search?${query}`
    const response = await UseServer.get<BranchPaginatedResource>(url)
    return response.data
  }

  /**
   * Exports all branches.
   *
   * @returns {Promise<void>} - A promise that resolves when the export is complete.
   */
  public static async exportAll(): Promise<void> {
    const url = `${BranchService.BASE_ENDPOINT}/export`
    await downloadFile(url, 'all_branches_export.xlsx')
  }

  /**
   * Exports filtered branches.
   *
   * @param {string} [filters] - The filters to apply for exporting branches.
   * @returns {Promise<void>} - A promise that resolves when the export is complete.
   */
  public static async exportAllFiltered(filters?: string): Promise<void> {
    const url = `${BranchService.BASE_ENDPOINT}/export-search?filter=${encodeURIComponent(filters || '')}`
    await downloadFile(url, 'filtered_branches_export.xlsx')
  }

  /**
   * Exports selected branches.
   *
   * @param {number[]} ids - The IDs of the selected branches to export.
   * @returns {Promise<void>} - A promise that resolves when the export is complete.
   */
  public static async exportSelected(ids: number[]): Promise<void> {
    const query = ids.map((id) => `ids=${id}`).join('&')
    const url = `${BranchService.BASE_ENDPOINT}/export-selected?${query}`
    await downloadFile(url, 'selected_branches_export.xlsx')
  }

  /**
   * Exports the current page of branches.
   *
   * @param {number} page - The page number to export.
   * @returns {Promise<void>} - A promise that resolves when the export is complete.
   */
  public static async exportCurrentPage(page: number): Promise<void> {
    const url = `${BranchService.BASE_ENDPOINT}/export-current-page/${page}`
    await downloadFile(url, `current_page_branches_${page}_export.xlsx`)
  }
}
