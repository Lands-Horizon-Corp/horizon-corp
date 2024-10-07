import UseServer from "../../request/server";
import { AxiosResponse } from 'axios';
import type { RolesRequest, RolesResource } from '../../types'


class RolesService {
  private server: UseServer;

  constructor() {
    this.server = new UseServer();
  }
  /**
  * Gets all feedbacks.
  *
  * @returns {Promise<RolesResource[]>}
  */
  async getAll(): Promise<RolesResource[]> {
    const response: AxiosResponse<RolesResource[]> = await this.server.get('/roles');
    return response.data;
  }

  /**
   * Creates a new role.
   *
   * @param {RolesRequest} roleData - The role data to create.
   * @returns {Promise<RolesResource>} - The created role resource.
   */
  async create(roleData: RolesRequest): Promise<RolesResource> {
    const response: AxiosResponse<RolesResource> = await this.server.post('/roles', roleData);
    return response.data;
  }


  /**
   * Deletes a role by ID.
   *
   * @param {number} id - The ID of the role to delete.
   * @returns {Promise<void>} - A promise that resolves when the role is deleted.
   */
  async delete(id: number): Promise<void> {
    await this.server.delete(`/roles/${id}`);
  }

  /**
   * Updates an existing role by ID.
   *
   * @param {number} id - The ID of the role to update.
   * @param {RolesRequest} roleData - The updated role data.
   * @returns {Promise<RolesResource>} - The updated role resource.
   */
  async update(id: number, roleData: RolesRequest): Promise<RolesResource> {
    const response: AxiosResponse<RolesResource> = await this.server.put(`/roles/${id}`, roleData);
    return response.data;
  }

  /**
  * Gets a role by ID.
  *
  * @param {number} id - The ID of the role to retrieve.
  * @returns {Promise<RolesResource>} - The retrieved role resource.
  */
  async getById(id: number): Promise<RolesResource> {
    const response: AxiosResponse<RolesResource> = await this.server.get(`/roles/${id}`);
    return response.data;
  }
}

export default RolesService