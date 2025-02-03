import { ITimeStamps, TEntityId } from './common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberJointAccountsResource extends ITimeStamps {
    id: TEntityId
    membersProfileID: TEntityId
    description: string
    firstName: string
    lastName: string
    middleName?: string
    familyRelationship?: string
    membersProfile?: IMemberProfileResource
}
