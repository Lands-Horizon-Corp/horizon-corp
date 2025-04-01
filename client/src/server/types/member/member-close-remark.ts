import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberCloseRemarkRequest {
    id?: TEntityId
    membersProfileId: TEntityId
    description: string
    membersProfile?: IMemberProfileResource
}

export interface IMemberCloseRemarkResource extends ITimeStamps {
    id: TEntityId
    membersProfileId: TEntityId
    description: string
    membersProfile?: IMemberProfileResource
}
