import { IMediaResource } from '../media'
import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberJointAccountsRequest {
    id?: TEntityId
    lastName: string
    firstName: string
    description: string
    middleName?: string
    familyRelationship?: string
    membersProfileId?: TEntityId // This should be optional @Zalven
    membersProfile?: IMemberProfileResource

    // TODO:Zalven NOTE: Tell Zalven to add these properties bellow
    mediaId?: TEntityId
    media?: IMediaResource
    signatureMediaId?: TEntityId
    signatureMedia?: TEntityId
}

export interface IMemberJointAccountsResource extends ITimeStamps {
    id: TEntityId
    lastName: string
    firstName: string
    description: string
    middleName?: string
    familyRelationship?: string
    membersProfileId: TEntityId // This should be optional @Zalven
    membersProfile?: IMemberProfileResource

    // TODO:Zalven NOTE: Tell Zalven to add these properties bellow
    mediaId?: TEntityId
    media?: IMediaResource
    signatureMediaId?: TEntityId
    signatureMedia?: TEntityId
}
