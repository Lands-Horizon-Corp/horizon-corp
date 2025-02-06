import { ITimeStamps, TEntityId } from '../common'
import { IMemberGenderResource } from './member-gender'
import { IMemberProfileResource } from './member-profile'

export interface IMemberGenderHistoryResource extends ITimeStamps {
    id: TEntityId
    memberProfileId: TEntityId
    memberGenderId: TEntityId
    memberProfile?: IMemberProfileResource
    memberGender?: IMemberGenderResource
}
