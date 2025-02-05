import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberJointAccountsResource extends ITimeStamps {
    id: TEntityId
    lastName: string
    firstName: string
    description: string
    middleName?: string
    familyRelationship?: string
    membersProfileID: TEntityId
    membersProfile?: IMemberProfileResource
}
