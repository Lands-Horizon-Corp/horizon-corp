import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberIncomeResource extends ITimeStamps {
    id: TEntityId
    membersProfileID: TEntityId
    name: string
    amount: number
    date: string
    description: string
    membersProfile?: IMemberProfileResource
}

export interface IMemberIncomeRequest {
    name: string
    amount: number
    date: string
    description: string
}
