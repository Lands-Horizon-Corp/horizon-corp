import { TEntityId } from './common'
import { IMediaResource } from './media'
import { IMemberResource } from './member/member'
import { ICompanyResource } from './company'
import { IEmployeeResource } from './employee'
import { IPaginatedResult } from './paginated-result'

export interface IBranchRequest {
    name: string
    address?: string
    description?: string
    longitude?: number
    latitude?: number
    email: string
    contactNumber: string
    isAdminVerified: boolean
    mediaId?: TEntityId
    companyId?: TEntityId
}

export interface IBranchResource {
    id: TEntityId
    name: string
    address?: string
    description?: string
    longitude?: number
    latitude?: number
    email: string
    companyId?: TEntityId
    contactNumber: string
    isAdminVerified: boolean
    media?: IMediaResource
    company?: ICompanyResource
    employees?: IEmployeeResource[]
    members?: IMemberResource[]
    createdAt: string
    updatedAt: string
}

export type IBranchPaginatedResource = IPaginatedResult<IBranchResource>
