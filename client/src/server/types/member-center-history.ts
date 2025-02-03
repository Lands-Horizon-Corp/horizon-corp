import { ITimeStamps, TEntityId } from './common'
import { IMemberCenterResource } from './member-center'
import { IMemberProfileResource } from './member-profile'

export interface IMemberCenterHistoryResource extends ITimeStamps {
    id: TEntityId
    memberProfileID: TEntityId
    memberCenterID: TEntityId
    memberProfile?: IMemberProfileResource
    memberCenter?: IMemberCenterResource
}
