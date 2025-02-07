import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberMutualFundsHistoryResource extends ITimeStamps {
    id: TEntityId
    membersProfileId: TEntityId
    description: string
    amount: number
    membersProfile?: IMemberProfileResource
}
