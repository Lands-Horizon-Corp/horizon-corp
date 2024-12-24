// services/BranchService.ts
import { downloadFile } from '@/horizon-corp/helpers'
import UseServer from '../../request/server'
import {
  BranchResource,
  BranchRequest,
  BranchPaginatedResource,
  MediaRequest,
} from '../../types'
import { AxiosResponse } from 'axios';


/**
 * Service class to handle CRUD operations for branches.
 */
export default class BranchService {
  private static readonly BASE_ENDPOINT = '/branch'

  /**
   * Retrieves a branch by its ID with optional preloads.
   *
   * @param {number} id - The ID of the branch to retrieve.
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<BranchResource>} - A promise that resolves to the branch resource.
   */
  public static async getById(id: number, preloads?: string[]): Promise<BranchResource> {
    // Construct each preload as a separate 'preloads' query parameter
    const preloadParams = preloads?.map(preload => `preloads=${encodeURIComponent(preload)}`).join('&') || '';
    const separator = preloadParams ? '?' : '';
    const endpoint = `${BranchService.BASE_ENDPOINT}/${id}${separator}${preloadParams}`;

    // Make the GET request with necessary headers
    const response = await UseServer.get<BranchResource>(endpoint, {
      headers: {
        'Authorization': `Bearer YOUR_TOKEN`, // Replace with actual token if needed
      },
    });

    return response.data;
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
    // Construct each preload as a separate 'preloads' query parameter
    const preloadParams = preloads?.map(preload => `preloads=${encodeURIComponent(preload)}`).join('&') || '';
    const separator = preloadParams ? '?' : '';
    const endpoint = `${BranchService.BASE_ENDPOINT}/${id}${separator}${preloadParams}`;

    // Make the PUT request with necessary headers
    const response = await UseServer.put<BranchRequest, BranchResource>(
      endpoint,
      branchData,
      {
        headers: {
          'Authorization': `Bearer YOUR_TOKEN`, // Replace with actual token if needed
        },
      }
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
    // Construct 'filter' query parameter
    const filterParams = filters ? `filter=${encodeURIComponent(filters)}` : ''

    // Construct each preload as a separate 'preloads' query parameter
    const preloadParams = preloads?.map(preload => `preloads=${encodeURIComponent(preload)}`).join('&') || ''

    // Combine filter and preload parameters
    const query = [filterParams, preloadParams].filter(Boolean).join('&')

    // Construct the full endpoint URL
    const url = `${BranchService.BASE_ENDPOINT}/search${query ? `?${query}` : ''}`

    // Make the GET request
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
    await downloadFile(url, 'all_branches_export.csv')
  }

  /**
   * Exports filtered branches.
   *
   * @param {string} [filters] - The filters to apply for exporting branches.
   * @returns {Promise<void>} - A promise that resolves when the export is complete.
   */
  public static async exportAllFiltered(filters?: string): Promise<void> {
    const filterQuery = filters ? `filter=${encodeURIComponent(filters)}` : ''
    const url = `${BranchService.BASE_ENDPOINT}/export-search${filterQuery ? `?${filterQuery}` : ''}`
    await downloadFile(url, 'filtered_branches_export.csv')
  }

  /**
   * Exports selected branches.
   *
   * @param {number[]} ids - The IDs of the selected branches to export.
   * @returns {Promise<void>} - A promise that resolves when the export is complete.
   */
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


  /**
   * Deletes multiple branches by their IDs.
   *
   * @param {number[]} ids - The IDs of the branches to delete.
   * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
   */
  public static async deleteMany(ids: number[]): Promise<void> {
    const endpoint = `${BranchService.BASE_ENDPOINT}/bulk-delete`

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
  public static async verify(id: number): Promise<BranchResource> {
    const endpoint = `${BranchService.BASE_ENDPOINT}/verify/${id}`;
    const response = await UseServer.post<void, BranchResource>(endpoint);
    return response.data;
  }

  /**
     * Uploads a profile picture for the company.
     * 
     * @async
     * @function
     * @param {MediaRequest} data - The request payload containing the media information.
     * @param {string[]} [preloads=["Media"]] - Optional array of relations to preload.
     * @returns {Promise<AxiosResponse<BranchResource>>} - The response containing the updated company resource.
     */
  public static async ProfilePicture(
    data: MediaRequest, preloads: string[] = ["Media"]
  ): Promise<AxiosResponse<BranchResource>> {
    const preloadParams = preloads?.map(preload => `preloads=${encodeURIComponent(preload)}`).join('&') || '';
    const separator = preloadParams ? '?' : '';
    const endpoint = `${BranchService.BASE_ENDPOINT}/profile-picture${separator}${preloadParams}`;
    return await UseServer.post<MediaRequest, BranchResource>(endpoint, data);
  }
}
