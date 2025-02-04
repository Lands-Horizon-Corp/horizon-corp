import { ITimeStamps, TEntityId } from './common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberDescriptionResource extends ITimeStamps {
    id: TEntityId
    membersProfileID: TEntityId
    date: string
    description: string
    name: string
    membersProfile?: IMemberProfileResource
}

export interface IMemberDescriptionRequest {
    name: string
    description: string
}
