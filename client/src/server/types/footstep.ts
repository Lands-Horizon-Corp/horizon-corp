import { IOwnerResource } from './owner'
import { IAdminResource } from './admin'
import { IMemberResource } from './member'
import { IEmployeeResource } from './employee'
import { TAccountType, TEntityId } from './common'
import { IPaginatedResult } from './paginated-result'

export interface IFootstepResource {
    id: TEntityId
    createdAt: string
    updatedAt: string
    deletedAt: string

    accountType: TAccountType
    module: string
    description: string
    activity: string
    latitude?: number
    longitude?: number
    timestamp: string
    isDeleted: boolean

    adminID?: TEntityId
    admin?: IAdminResource
    employeeID?: TEntityId
    employee?: IEmployeeResource
    ownerID?: TEntityId
    owner?: IOwnerResource
    memberID?: TEntityId
    member?: IMemberResource
}

export type IFootstepPaginatedResource = IPaginatedResult<IFootstepResource>
