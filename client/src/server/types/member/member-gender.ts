import { ITimeStamps, TEntityId } from '../common'
import { IMemberGenderHistoryResource } from './member-gender-history'

export interface IMemberGenderResource extends ITimeStamps {
    id: TEntityId
    name: string
    createdAt: string
    updatedAt: string
    deletedAt: string
    description: string
    history?: IMemberGenderHistoryResource[]
}
