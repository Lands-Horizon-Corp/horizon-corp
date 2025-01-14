import { IUserBase } from '../common'
import { PaginatedResult } from '../paginated-result'

export interface MemberResource extends IUserBase {}

export type MemberPaginatedResource = PaginatedResult<MemberResource>
