import { ITimeStamps, TEntityId } from './common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberExpensesResource extends ITimeStamps {
    id: TEntityId
    membersProfileID: TEntityId
    name: string
    date: string
    amount: number
    description: string
    membersProfile?: IMemberProfileResource
}
