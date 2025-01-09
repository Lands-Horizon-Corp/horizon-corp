import { IUserBase } from '../common'
import { MediaResource } from '../media'
import { RolesResource } from '../roles'
import { BranchResource } from '../branch'
import { CompanyResource } from '../company'
import { FootstepResource } from '../footstep'
import { TimesheetResource } from '../timesheet'

export interface MemberResource extends IUserBase {
    accountType: string
    longitude?: number
    latitude?: number
    branch?: BranchResource
    role?: RolesResource
    footsteps?: FootstepResource[]
}

export interface OwnerResource extends IUserBase {
    accountType: string
    media?: MediaResource
    companies?: CompanyResource[]
    footsteps?: FootstepResource[]
}

export interface AdminResource extends IUserBase {
    accountType: string
    description?: string
    role?: RolesResource
    footsteps?: FootstepResource[]
}

export interface EmployeeResource extends IUserBase {
    accountType: string
    branch?: BranchResource
    longitude?: number
    latitude?: number
    timesheets?: TimesheetResource[]
    role?: RolesResource
    footsteps?: FootstepResource[]
}
