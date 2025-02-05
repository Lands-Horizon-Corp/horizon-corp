import { TEntityId } from '../common'
import { IMemberClassificationHistoryResource } from './member-classification-history'

export interface IMemberClassificationResource {
    id: TEntityId
    name: string
    createdAt: string
    updatedAt: string
    deletedAt: string
    description: string
    history?: IMemberClassificationHistoryResource[]
}
