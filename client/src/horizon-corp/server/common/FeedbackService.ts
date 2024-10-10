import UseServer from '../../request/server'
import { FeedbacksResource, FeedbacksRequest } from '../../types'

/**
 * Service class to handle CRUD operations for feedbacks.
 */
export default class FeedbackService {
  private static readonly BASE_ENDPOINT = '/feedback'

  /**
   * Retrieves all feedbacks.
   *
   * @returns {Promise<FeedbacksResource[]>} - A promise that resolves to an array of feedback resources.
   */
  public static async getAll(): Promise<FeedbacksResource[]> {
    const response = await UseServer.get<FeedbacksResource[]>(
      FeedbackService.BASE_ENDPOINT
    )
    return response.data
  }

  /**
   * Creates a new feedback.
   *
   * @param {FeedbacksRequest} feedbackData - The data for the new feedback.
   * @returns {Promise<FeedbacksResource>} - A promise that resolves to the created feedback resource.
   */
  public static async create(
    feedbackData: FeedbacksRequest
  ): Promise<FeedbacksResource> {
    const response = await UseServer.post<
      FeedbacksRequest,
      FeedbacksResource
    >(FeedbackService.BASE_ENDPOINT, feedbackData)
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
   * Updates an existing feedback by its ID.
   *
   * @param {number} id - The ID of the feedback to update.
   * @param {FeedbacksRequest} feedbackData - The updated data for the feedback.
   * @returns {Promise<FeedbacksResource>} - A promise that resolves to the updated feedback resource.
   */
  public static async update(
    id: number,
    feedbackData: FeedbacksRequest
  ): Promise<FeedbacksResource> {
    const endpoint = `${FeedbackService.BASE_ENDPOINT}/${id}`
    const response = await UseServer.put<
      FeedbacksRequest,
      FeedbacksResource
    >(endpoint, feedbackData)
    return response.data
  }

  /**
   * Retrieves a feedback by its ID.
   *
   * @param {number} id - The ID of the feedback to retrieve.
   * @returns {Promise<FeedbacksResource>} - A promise that resolves to the feedback resource.
   */
  public static async getById(id: number): Promise<FeedbacksResource> {
    const endpoint = `${FeedbackService.BASE_ENDPOINT}/${id}`
    const response = await UseServer.get<FeedbacksResource>(endpoint)
    return response.data
  }
}
