import { IPaginatedResult } from './paginated-result'

export interface IGendersRequest {
    name: string
    description?: string
}

export interface IGenderResource {
    id: number
    name: string
    description: string
    createdAt: string
    updatedAt: string
}

export type IGenderPaginatedResource = IPaginatedResult<IGenderResource>
