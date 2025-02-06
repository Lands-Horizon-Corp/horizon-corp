import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'
import { IMemberTypeResource } from './member-type'

export interface IMemberTypeHistory extends ITimeStamps {
    id: TEntityId
    memberProfileID: TEntityId
    memberTypeID: TEntityId
    memberProfile?: IMemberProfileResource
    memberType?: IMemberTypeResource
}
