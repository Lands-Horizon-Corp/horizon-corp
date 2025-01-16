import {
    AdminResource,
    OwnerResource,
    MemberResource,
    EmployeeResource,
} from '../profile'
import { PaginatedResult } from '../paginated-result'

export interface FootstepResource {
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
    admin?: AdminResource
    employeeID?: string
    employee?: EmployeeResource
    ownerID?: string
    owner?: OwnerResource
    memberID?: string
    member?: MemberResource
}

export type FootstepPaginatedResource = PaginatedResult<FootstepResource>
