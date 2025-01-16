import { PaginatedResult } from "../paginated-result"

export interface GenderResource {
  id: number
  name: string
  description: string
  createdAt: string
  updatedAt: string
}


export type GenderPaginatedResource = PaginatedResult<GenderResource>
