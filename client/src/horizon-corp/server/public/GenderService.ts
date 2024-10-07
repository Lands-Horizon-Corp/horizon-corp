import UseServer from "../../request/server";
import { AxiosResponse } from 'axios';
import type { GendersRequest, GendersResource } from '../../types'


export class GenderService {
  private server: UseServer;

  constructor() {
    this.server = new UseServer();
  }

  /**
   * Gets all genders.
   *
   * @returns {Promise<GendersResource[]>} - An array of feedback resources.
   */
  async getAll(): Promise<GendersResource[]> {
    const response: AxiosResponse<GendersResource[]> = await this.server.get('/genders');
    return response.data;
  }


  /**
   * Creates a new gender.
   *
   * @param {GendersRequest} genderData - The gender data to create.
   * @returns {Promise<GenderResource>} - The created gender resource.
   */
  async create(genderData: GendersRequest): Promise<GendersResource> {
    const response: AxiosResponse<GendersResource> = await this.server.post('/genders', genderData);
    return response.data;
  }


  /**
   * Deletes a gender by ID.
   *
   * @param {number} id - The ID of the gender to delete.
   * @returns {Promise<void>} - A promise that resolves when the gender is deleted.
   */
  async delete(id: number): Promise<void> {
    await this.server.delete(`/genders/${id}`);
  }

  /**
   * Updates an existing gender by ID.
   *
   * @param {number} id - The ID of the gender to update.
   * @param {GendersRequest} genderData - The updated gender data.
   * @returns {Promise<GenderResource>} - The updated gender resource.
   */
  async update(id: number, genderData: GendersRequest): Promise<GendersResource> {
    const response: AxiosResponse<GendersResource> = await this.server.put(`/genders/${id}`, genderData);
    return response.data;
  }

  /**
  * Gets a gender by ID.
  *
  * @param {number} id - The ID of the gender to retrieve.
  * @returns {Promise<GenderResource>} - The retrieved gender resource.
  */
  async getById(id: number): Promise<GendersResource> {
    const response: AxiosResponse<GendersResource> = await this.server.get(`/genders/${id}`);
    return response.data;
  }
}