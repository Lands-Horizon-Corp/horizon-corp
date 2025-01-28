import { TEntityId } from './common'
import { IPaginatedResult } from './paginated-result'

export interface IGenderRequest {
    name: string
    description?: string
}

export interface IGenderResource {
    id: TEntityId
    name: string
    description: string
    createdAt: string
    updatedAt: string
}

export type IGenderPaginatedResource = IPaginatedResult<IGenderResource>
