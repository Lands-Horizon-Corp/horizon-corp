import { FilterPages } from "../table"

export interface GenderResource {
  id: number
  name: string
  description: string
  createdAt: string
  updatedAt: string
}


export type GenderPaginatedResource = FilterPages<GenderResource>
