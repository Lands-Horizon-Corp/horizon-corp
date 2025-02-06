import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'
import { IMemberEducationalAttainmentResource } from './member-educational-attainment'

export interface IMemberEducationalAttainmentHistoryResource
    extends ITimeStamps {
    id: TEntityId
    memberProfileId: TEntityId
    memberEducationalAttainmentId: TEntityId
    memberProfile?: IMemberProfileResource
    memberEducationalAttainment?: IMemberEducationalAttainmentResource
}
