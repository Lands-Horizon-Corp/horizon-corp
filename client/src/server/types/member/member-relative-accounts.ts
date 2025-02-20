import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberRelativeAccountsRequest {
    id?: TEntityId
    membersProfileId: TEntityId
    relativeProfileMemberId: TEntityId
    familyRelationship: string
    description: string
    memberProfile?: IMemberProfileResource
    relativeProfileMemberProfile?: IMemberProfileResource
}

export interface IMemberRelativeAccountsResource extends ITimeStamps {
    id: TEntityId
    membersProfileId: TEntityId
    relativeProfileMemberId: TEntityId
    familyRelationship: string
    description: string
    memberProfile?: IMemberProfileResource
    relativeProfileMemberProfile?: IMemberProfileResource
}
