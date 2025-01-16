import { IUserBase } from '../common'
import { PaginatedResult } from '../paginated-result'

export interface OwnerResource extends IUserBase {}

export type OwnerPaginatedResource = PaginatedResult<OwnerResource>
