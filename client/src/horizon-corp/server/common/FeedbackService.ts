import UseServer from "../../request/server";
import { AxiosResponse } from 'axios';
import type { FeedbacksRequest, FeedbacksResource } from '../../types'


export default class FeedbackService {
  private server: UseServer;

  constructor() {
    this.server = new UseServer();
  }

  /**
   * Gets all feedbacks.
   *
   * @returns {Promise<FeedbacksResource[]>} - An array of feedback resources.
   */
  async getAll(): Promise<FeedbacksResource[]> {
    const response: AxiosResponse<FeedbacksResource[]> = await this.server.get('/feedbacks');
    return response.data;
  }

  /**
   * Creates a new feedback.
   *
   * @param {FeedbacksRequest} feedbackData - The feedback data to create.
   * @returns {Promise<FeedbacksResource>} - The created feedback resource.
   */
  async create(feedbackData: FeedbacksRequest): Promise<FeedbacksResource> {
    const response: AxiosResponse<FeedbacksResource> = await this.server.post('/feedbacks', feedbackData);
    return response.data;
  }


  /**
   * Deletes a feedback by ID.
   *
   * @param {number} id - The ID of the feedback to delete.
   * @returns {Promise<void>} - A promise that resolves when the feedback is deleted.
   */
  async delete(id: number): Promise<void> {
    await this.server.delete(`/feedbacks/${id}`);
  }

  /**
   * Updates an existing feedback by ID.
   *
   * @param {number} id - The ID of the feedback to update.
   * @param {FeedbacksRequest} feedbackData - The updated feedback data.
   * @returns {Promise<FeedbacksResource>} - The updated feedback resource.
   */
  async update(id: number, feedbackData: FeedbacksRequest): Promise<FeedbacksResource> {
    const response: AxiosResponse<FeedbacksResource> = await this.server.put(`/feedbacks/${id}`, feedbackData);
    return response.data;
  }

  /**
  * Gets a feedback by ID.
  *
  * @param {number} id - The ID of the feedback to retrieve.
  * @returns {Promise<FeedbacksResource>} - The retrieved feedback resource.
  */
  async getById(id: number): Promise<FeedbacksResource> {
    const response: AxiosResponse<FeedbacksResource> = await this.server.get(`/feedbacks/${id}`);
    return response.data;
  }

}