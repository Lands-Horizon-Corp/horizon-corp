// services/FeedbackService.ts
import { downloadFile } from '@/horizon-corp/helpers'
import UseServer from '../../request/server'
import {
  FeedbackResource,
  FeedbackRequest,
  FeedbackPaginatedResource,
} from '../../types'

/**
 * Service class to handle CRUD operations for feedbacks.
 */
export default class FeedbackService {
  private static readonly BASE_ENDPOINT = '/feedback'

  /**
   * Retrieves a feedback by its ID with optional preloads.
   *
   * @param {number} id - The ID of the feedback to retrieve.
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<FeedbackResource>} - A promise that resolves to the feedback resource.
   */
  public static async getById(id: number, preloads?: string[]): Promise<FeedbackResource> {
    // Construct each preload as a separate 'preloads' query parameter
    const preloadParams = preloads?.map(preload => `preloads=${encodeURIComponent(preload)}`).join('&') || '';
    const separator = preloadParams ? '?' : '';
    const endpoint = `${FeedbackService.BASE_ENDPOINT}/${id}${separator}${preloadParams}`;

    // Make the GET request with necessary headers
    const response = await UseServer.get<FeedbackResource>(endpoint, {
      headers: {
        'Authorization': `Bearer YOUR_TOKEN`, // Replace with actual token if needed
      },
    });

    return response.data;
  }

  /**
   * Creates a new feedback.
   *
   * @param {FeedbackRequest} feedbackData - The data for the new feedback.
   * @returns {Promise<FeedbackResource>} - A promise that resolves to the created feedback resource.
   */
  public static async create(
    feedbackData: FeedbackRequest
  ): Promise<FeedbackResource> {
    const response = await UseServer.post<FeedbackRequest, FeedbackResource>(
      FeedbackService.BASE_ENDPOINT,
      feedbackData
    )
    return response.data
  }

  /**
   * Deletes a feedback by its ID.
   *
   * @param {number} id - The ID of the feedback to delete.
   * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
   */
  public static async delete(id: number): Promise<void> {
    const endpoint = `${FeedbackService.BASE_ENDPOINT}/${id}`
    await UseServer.delete<void>(endpoint)
  }


  /**
   * Filters feedbacks based on provided filters with optional preloads.
   *
   * @param {string} [filters] - The filters to apply for exporting feedbacks.
   * @param {string[]} [preloads] - Optional array of relations to preload.
   * @returns {Promise<FeedbackPaginatedResource>} - A promise that resolves to paginated feedback resources.
   */
  public static async filter(
    filters?: string,
    preloads?: string[]
  ): Promise<FeedbackPaginatedResource> {
    // Construct 'filter' query parameter
    const filterParams = filters ? `filter=${encodeURIComponent(filters)}` : ''

    // Construct each preload as a separate 'preloads' query parameter
    const preloadParams = preloads?.map(preload => `preloads=${encodeURIComponent(preload)}`).join('&') || ''

    // Combine filter and preload parameters
    const query = [filterParams, preloadParams].filter(Boolean).join('&')

    // Construct the full endpoint URL
    const url = `${FeedbackService.BASE_ENDPOINT}/search${query ? `?${query}` : ''}`

    // Make the GET request
    const response = await UseServer.get<FeedbackPaginatedResource>(url)
    return response.data
  }

  /**
   * Exports all feedbacks.
   *
   * @returns {Promise<void>} - A promise that resolves when the export is complete.
   */
  public static async exportAll(): Promise<void> {
    const url = `${FeedbackService.BASE_ENDPOINT}/export`
    await downloadFile(url, 'all_feedbacks_export.csv')
  }

  /**
   * Exports filtered feedbacks.
   *
   * @param {string} [filters] - The filters to apply for exporting feedbacks.
   * @returns {Promise<void>} - A promise that resolves when the export is complete.
   */
  public static async exportAllFiltered(filters?: string): Promise<void> {
    const filterQuery = filters ? `filter=${encodeURIComponent(filters)}` : ''
    const url = `${FeedbackService.BASE_ENDPOINT}/export-search${filterQuery ? `?${filterQuery}` : ''}`
    await downloadFile(url, 'filtered_feedbacks_export.csv')
  }

  /**
   * Exports selected feedbacks.
   *
   * @param {number[]} ids - The IDs of the selected feedbacks to export.
   * @returns {Promise<void>} - A promise that resolves when the export is complete.
   */
  public static async exportSelected(ids: number[]): Promise<void> {
    if (ids.length === 0) {
      throw new Error('No feedback IDs provided for export.')
    }

    // Construct each preload as a separate 'preloads' query parameter if needed
    // (Assuming export-selected might also accept preloads; if not, you can omit this)

    const query = ids.map((id) => `ids=${encodeURIComponent(id)}`).join('&')
    const url = `${FeedbackService.BASE_ENDPOINT}/export-selected?${query}`
    await downloadFile(url, 'selected_feedbacks_export.csv')
  }


  /**
   * Deletes multiple feedbacks by their IDs.
   *
   * @param {number[]} ids - The IDs of the feedbacks to delete.
   * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
   */
  public static async deleteMany(ids: number[]): Promise<void> {
    const endpoint = `${FeedbackService.BASE_ENDPOINT}/bulk-delete`

    // Construct the request payload
    const payload = { ids }

    // Make the DELETE request with the payload
    await UseServer.delete<void>(endpoint, payload)
  }

}
