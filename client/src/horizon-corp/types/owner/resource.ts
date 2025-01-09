import { FilterPages } from '../table'
import { IUserBase } from '../common'

export interface OwnerResource extends IUserBase {}

export type OwnerPaginatedResource = FilterPages<OwnerResource>
