import UseServer from '../../request/server'
import { RolesResource, RolesRequest } from '../../types'

/**
 * Service class to handle CRUD operations for roles.
 */
export default class RoleService {
    private static readonly BASE_ENDPOINT = '/role'

    /**
     * Retrieves all roles.
     *
     * @returns {Promise<RolesResource[]>} - A promise that resolves to an array of role resources.
     */
    public static async getAll(): Promise<RolesResource[]> {
        const response = await UseServer.get<RolesResource[]>(
            RoleService.BASE_ENDPOINT
        )
        return response.data
    }

    /**
     * Creates a new role.
     *
     * @param {RolesRequest} roleData - The data for the new role.
     * @returns {Promise<RolesResource>} - A promise that resolves to the created role resource.
     */
    public static async create(roleData: RolesRequest): Promise<RolesResource> {
        const response = await UseServer.post<RolesRequest, RolesResource>(
            RoleService.BASE_ENDPOINT,
            roleData
        )
        return response.data
    }

    /**
     * Deletes a role by its ID.
     *
     * @param {number} id - The ID of the role to delete.
     * @returns {Promise<void>} - A promise that resolves when the deletion is complete.
     */
    public static async delete(id: number): Promise<void> {
        const endpoint = `${RoleService.BASE_ENDPOINT}/${id}`
        await UseServer.delete<void>(endpoint)
    }

    /**
     * Updates an existing role by its ID.
     *
     * @param {number} id - The ID of the role to update.
     * @param {RolesRequest} roleData - The updated data for the role.
     * @returns {Promise<RolesResource>} - A promise that resolves to the updated role resource.
     */
    public static async update(
        id: number,
        roleData: RolesRequest
    ): Promise<RolesResource> {
        const endpoint = `${RoleService.BASE_ENDPOINT}/${id}`
        const response = await UseServer.put<RolesRequest, RolesResource>(
            endpoint,
            roleData
        )
        return response.data
    }

    /**
     * Retrieves a role by its ID.
     *
     * @param {number} id - The ID of the role to retrieve.
     * @returns {Promise<RolesResource>} - A promise that resolves to the role resource.
     */
    public static async getById(id: number): Promise<RolesResource> {
        const endpoint = `${RoleService.BASE_ENDPOINT}/${id}`
        const response = await UseServer.get<RolesResource>(endpoint)
        return response.data
    }
}
