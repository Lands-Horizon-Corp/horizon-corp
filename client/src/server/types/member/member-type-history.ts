import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'
import { IMemberTypeResource } from './member-type'

export interface IMemberTypeHistory extends ITimeStamps {
    id: TEntityId
    memberProfileId: TEntityId
    memberTypeId: TEntityId
    memberProfile?: IMemberProfileResource
    memberType?: IMemberTypeResource
}
