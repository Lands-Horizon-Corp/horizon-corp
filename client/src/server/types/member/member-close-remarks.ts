import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberCloseRemarksRequest {
    id?: TEntityId
    membersProfileId: TEntityId
    description: string
    membersProfile?: IMemberProfileResource
}

export interface IMemberCloseRemarksResource extends ITimeStamps {
    id: TEntityId
    membersProfileId: TEntityId
    description: string
    membersProfile?: IMemberProfileResource
}
