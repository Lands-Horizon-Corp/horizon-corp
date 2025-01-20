import APIService from './api-service'
import { IRolesResource, IRolesRequest } from '../types'

/**
 * Service class to handle CRUD operations for roles.
 */
export default class RoleService {
    private static readonly BASE_ENDPOINT = '/role'

    public static async getAll(): Promise<IRolesResource[]> {
        const response = await APIService.get<IRolesResource[]>(
            RoleService.BASE_ENDPOINT
        )
        return response.data
    }

    public static async create(
        roleData: IRolesRequest
    ): Promise<IRolesResource> {
        const response = await APIService.post<IRolesRequest, IRolesResource>(
            RoleService.BASE_ENDPOINT,
            roleData
        )
        return response.data
    }

    public static async delete(id: number): Promise<void> {
        const endpoint = `${RoleService.BASE_ENDPOINT}/${id}`
        await APIService.delete<void>(endpoint)
    }

    public static async update(
        id: number,
        roleData: IRolesRequest
    ): Promise<IRolesResource> {
        const endpoint = `${RoleService.BASE_ENDPOINT}/${id}`
        const response = await APIService.put<IRolesRequest, IRolesResource>(
            endpoint,
            roleData
        )
        return response.data
    }

    public static async getById(id: number): Promise<IRolesResource> {
        const endpoint = `${RoleService.BASE_ENDPOINT}/${id}`
        const response = await APIService.get<IRolesResource>(endpoint)
        return response.data
    }
}
