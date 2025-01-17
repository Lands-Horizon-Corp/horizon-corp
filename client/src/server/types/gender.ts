import { PaginatedResult } from './paginated-result'

export interface IGenderResource {
    id: number
    name: string
    description: string
    createdAt: string
    updatedAt: string
}

export type IGenderPaginatedResource = PaginatedResult<IGenderResource>
