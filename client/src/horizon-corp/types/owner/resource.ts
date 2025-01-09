import { IUserBase } from '../common'
import { FilterPages } from '../table'

export interface OwnerResource extends IUserBase {}

export type OwnerPaginatedResource = FilterPages<OwnerResource>
