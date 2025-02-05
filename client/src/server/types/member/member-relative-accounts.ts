import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberRelativeAccountsResource extends ITimeStamps {
    id: TEntityId
    membersProfileID: TEntityId
    relativeProfileMemberID: TEntityId
    familyRelationship: string
    description: string
    memberProfile?: IMemberProfileResource
    relativeProfileMemberProfile?: IMemberProfileResource
}