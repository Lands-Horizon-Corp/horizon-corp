import { IUserBase } from "./common"
import { IFootstepResource } from "./footstep"
import { IRolesResource } from "./role"

export interface IAdminResource extends IUserBase {
    accountType: string
    description?: string
    role?: IRolesResource
    footsteps?: IFootstepResource[]
}