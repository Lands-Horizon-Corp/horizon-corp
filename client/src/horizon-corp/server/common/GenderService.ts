import UseServer from '../../request/server'
import { GenderResource, GendersRequest } from '../../types'

/**
 * Service class to handle CRUD operations for genders.
 */
export default class GenderService {
  private static readonly BASE_ENDPOINT = '/gender'

  /**
   * Retrieves all genders.
   *
   * @returns {Promise<GenderResource[]>} - A promise that resolves to an array of gender resources.
   */
  public static async getAll(): Promise<GenderResource[]> {
    const response = await UseServer.get<GenderResource[]>(
      GenderService.BASE_ENDPOINT
    )
    return response.data
  }

  /**
   * Creates a new gender.
   *
   * @param {GendersRequest} genderData - The data for the new gender.
   * @returns {Promise<GenderResource>} - A promise that resolves to the created gender resource.
   */
  public static async create(
    genderData: GendersRequest
  ): Promise<GenderResource> {
    const response = await UseServer.post<GendersRequest, GenderResource>(
      GenderService.BASE_ENDPOINT,
      genderData
    )
    return response.data
  }

  /**
   * Deletes a gender by its ID.
   *
   * @param {number} id - The ID of the gender to delete.
   * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
   */
  public static async delete(id: number): Promise<void> {
    const endpoint = `${GenderService.BASE_ENDPOINT}/${id}`
    await UseServer.delete<void>(endpoint)
  }

  /**
   * Updates an existing gender by its ID.
   *
   * @param {number} id - The ID of the gender to update.
   * @param {GendersRequest} genderData - The updated data for the gender.
   * @returns {Promise<GenderResource>} - A promise that resolves to the updated gender resource.
   */
  public static async update(
    id: number,
    genderData: GendersRequest
  ): Promise<GenderResource> {
    const endpoint = `${GenderService.BASE_ENDPOINT}/${id}`
    const response = await UseServer.put<GendersRequest, GenderResource>(
      endpoint,
      genderData
    )
    return response.data
  }

  /**
   * Retrieves a gender by its ID.
   *
   * @param {number} id - The ID of the gender to retrieve.
   * @returns {Promise<GenderResource>} - A promise that resolves to the gender resource.
   */
  public static async getById(id: number): Promise<GenderResource> {
    const endpoint = `${GenderService.BASE_ENDPOINT}/${id}`
    const response = await UseServer.get<GenderResource>(endpoint)
    return response.data
  }
}
