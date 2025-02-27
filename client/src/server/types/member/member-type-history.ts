import { ITimeStamps, TEntityId } from '../common'
import { IMemberTypeResource } from './member-type'
import { IMemberProfileResource } from './member-profile'

export interface IMemberTypeHistory extends ITimeStamps {
    id: TEntityId
    memberProfileId: TEntityId
    memberTypeId: TEntityId
    memberProfile?: IMemberProfileResource
    memberType?: IMemberTypeResource
}
