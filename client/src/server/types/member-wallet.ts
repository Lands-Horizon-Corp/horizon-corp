import { ITimeStamps, TEntityId } from './common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberWalletResource extends ITimeStamps {
    id: TEntityId
    membersProfileID: TEntityId
    debit: number
    credit: number
    date: string
    description: string
    membersProfile?: IMemberProfileResource
}
