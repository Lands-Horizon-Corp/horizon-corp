import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberAssetsResource extends ITimeStamps {
    id: TEntityId
    membersProfileID: TEntityId
    entryDate: string
    description: string
    name: string
    membersProfile?: IMemberProfileResource
}
