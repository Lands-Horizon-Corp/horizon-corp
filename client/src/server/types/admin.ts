import { IUserBase } from './common'
import { IRolesResource } from './role'
import { IFootstepResource } from './footstep'

export interface IAdminResource extends IUserBase {
  accountType: string
  description?: string
  role?: IRolesResource
  footsteps?: IFootstepResource[]
}
