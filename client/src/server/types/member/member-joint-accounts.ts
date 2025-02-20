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
    membersProfileId?: TEntityId
    membersProfile?: IMemberProfileResource

    mediaId?: TEntityId
    media?: IMediaResource
    signatureMediaId?: TEntityId
    signatureMedia?: IMediaResource
}

export interface IMemberJointAccountsResource extends ITimeStamps {
    id: TEntityId
    lastName: string
    firstName: string
    description: string
    middleName?: string
    familyRelationship?: string
    membersProfileId: TEntityId
    membersProfile: IMemberProfileResource

    mediaId?: TEntityId
    media?: IMediaResource
    signatureMediaId?: TEntityId
    signatureMedia?: TEntityId
}
