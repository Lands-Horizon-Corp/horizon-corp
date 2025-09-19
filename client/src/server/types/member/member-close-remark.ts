import { AccountClosureReasonType, ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberCloseRemarkRequest {
    id?: TEntityId
    category: AccountClosureReasonType
    membersProfileId: TEntityId
    description: string
}

export interface IMemberCloseRemarkResource extends ITimeStamps {
    id: TEntityId
    membersProfileId: TEntityId
    category: AccountClosureReasonType
    description: string
    membersProfile?: IMemberProfileResource
}
