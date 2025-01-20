import { IUserBase } from "./common";
import { IRolesResource } from "./role";
import { IBranchResource } from "./branch";
import { IFootstepResource } from "./footstep";
import { IPaginatedResult } from "./paginated-result";

export interface IMemberRequest {

}

export interface IMemberResource extends IUserBase {
    accountType: string
    longitude?: number
    latitude?: number
    branch?: IBranchResource
    role?: IRolesResource
    footsteps?: IFootstepResource[]
}

export type IMemberPaginatedResource = IPaginatedResult<IMemberResource>

