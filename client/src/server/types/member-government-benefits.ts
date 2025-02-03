import { IMediaResource } from './media'
import { ITimeStamps, TEntityId } from './common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberGovernmentBenefitsResource extends ITimeStamps {
    id: TEntityId
    membersProfileID: TEntityId
    country: string
    name: string
    description: string
    value: number
    frontMediaID?: TEntityId
    backMediaID?: TEntityId
    membersProfile?: IMemberProfileResource
    frontMedia?: IMediaResource
    backMedia?: IMediaResource
}
