import { IOwnerResource } from './owner'
import { IAdminResource } from './admin'
import { IMemberResource } from './member'
import { IEmployeeResource } from './employee'
import { IPaginatedResult } from './paginated-result'

export interface IFootstepResource {
    id: string
    createdAt: string
    updatedAt: string
    deletedAt: string

    accountType: string
    module: string
    description: string
    activity: string
    latitude?: number
    longitude?: number
    timestamp: string
    isDeleted: boolean

    adminID?: string
    admin?: IAdminResource
    employeeID?: string
    employee?: IEmployeeResource
    ownerID?: string
    owner?: IOwnerResource
    memberID?: string
    member?: IMemberResource
}

export type IFootstepPaginatedResource = IPaginatedResult<IFootstepResource>
