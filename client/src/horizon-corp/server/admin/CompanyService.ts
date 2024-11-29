import UseServer from '../../request/server'
import { CompanyResource, CompanyRequest } from '../../types'

/**
 * Service class to handle CRUD operations for companies.
 */
export default class CompanyService {
  private static readonly BASE_ENDPOINT = '/companies'

  /**
   * Retrieves all companies.
   *
   * @returns {Promise<CompanyResource[]>} - A promise that resolves to an array of company resources.
   */
  public static async getAll(): Promise<CompanyResource[]> {
    const response = await UseServer.get<CompanyResource[]>(
      CompanyService.BASE_ENDPOINT
    )
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
    const response = await UseServer.post<
      CompanyRequest,
      CompanyResource
    >(CompanyService.BASE_ENDPOINT, companyData)
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
   * Updates an existing company by its ID.
   *
   * @param {number} id - The ID of the company to update.
   * @param {CompanyRequest} companyData - The updated data for the company.
   * @returns {Promise<CompanyResource>} - A promise that resolves to the updated company resource.
   */
  public static async update(
    id: number,
    companyData: CompanyRequest
  ): Promise<CompanyResource> {
    const endpoint = `${CompanyService.BASE_ENDPOINT}/${id}`
    const response = await UseServer.put<CompanyRequest, CompanyResource>(
      endpoint,
      companyData
    )
    return response.data
  }

  /**
   * Retrieves a company by its ID.
   *
   * @param {number} id - The ID of the company to retrieve.
   * @returns {Promise<CompanyResource>} - A promise that resolves to the company resource.
   */
  public static async getById(id: number): Promise<CompanyResource> {
    const endpoint = `${CompanyService.BASE_ENDPOINT}/${id}`
    const response = await UseServer.get<CompanyResource>(endpoint)
    return response.data
  }
}
