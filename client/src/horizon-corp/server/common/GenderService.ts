import UseServer from '../../request/server'
import { GendersResource, GendersRequest } from '../../types'

/**
 * Service class to handle CRUD operations for genders.
 */
export default class GenderService {
  private static readonly BASE_ENDPOINT = '/genders'

  /**
   * Retrieves all genders.
   *
   * @returns {Promise<GendersResource[]>} - A promise that resolves to an array of gender resources.
   */
  public static async getAll(): Promise<GendersResource[]> {
    const response = await UseServer.get<GendersResource[]>(GenderService.BASE_ENDPOINT)
    return response.data
  }

  /**
   * Creates a new gender.
   *
   * @param {GendersRequest} genderData - The data for the new gender.
   * @returns {Promise<GendersResource>} - A promise that resolves to the created gender resource.
   */
  public static async create(genderData: GendersRequest): Promise<GendersResource> {
    const response = await UseServer.post<GendersRequest, GendersResource>(
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
   * @returns {Promise<GendersResource>} - A promise that resolves to the updated gender resource.
   */
  public static async update(id: number, genderData: GendersRequest): Promise<GendersResource> {
    const endpoint = `${GenderService.BASE_ENDPOINT}/${id}`
    const response = await UseServer.put<GendersRequest, GendersResource>(
      endpoint,
      genderData
    )
    return response.data
  }

  /**
   * Retrieves a gender by its ID.
   *
   * @param {number} id - The ID of the gender to retrieve.
   * @returns {Promise<GendersResource>} - A promise that resolves to the gender resource.
   */
  public static async getById(id: number): Promise<GendersResource> {
    const endpoint = `${GenderService.BASE_ENDPOINT}/${id}`
    const response = await UseServer.get<GendersResource>(endpoint)
    return response.data
  }
}
