import { IMediaResource } from './media'
import { IOwnerResource } from './owner'
import { IBranchResource } from './branch'
import { IPaginatedResult } from './paginated-result'
import { TEntityId } from './common'

export interface ICompanyRequest {
    name: string
    description?: string
    address?: string
    longitude?: number
    latitude?: number
    email?: string
    contactNumber: string
    ownerId?: number
    companyId?: number
    mediaId?: number
    isAdminVerified?: boolean
}

export interface ICompanyResource {
    id: TEntityId
    name: string
    description?: string
    address?: string
    longitude?: number
    latitude?: number
    contactNumber: string
    isAdminVerified: boolean
    owner?: IOwnerResource
    media?: IMediaResource
    branches?: IBranchResource[]
    createdAt: string
    updatedAt: string
}

export type ICompanyPaginatedResource = IPaginatedResult<ICompanyResource>
