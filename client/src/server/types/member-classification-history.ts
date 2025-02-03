import { TEntityId } from './common'
import { IMemberProfileResource } from './member-profile'
import { IMemberClassificationResource } from './member-classification'

export interface IMemberClassificationHistoryResource {
    id: TEntityId
    createdAt: string
    updatedAt: string
    deletedAt: string
    memberProfileID: TEntityId
    memberClassificationID: TEntityId
    memberProfile?: IMemberProfileResource
    memberClassification?: IMemberClassificationResource
}
