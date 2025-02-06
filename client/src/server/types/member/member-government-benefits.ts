import { IMediaResource } from '../media'
import { ITimeStamps, TEntityId } from '../common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberGovernmentBenefitsResource extends ITimeStamps {
    id: TEntityId
    membersProfileId: TEntityId
    country: string
    name: string
    description: string
    value: number
    frontMediaId?: TEntityId
    backMediaId?: TEntityId
    membersProfile?: IMemberProfileResource
    frontMedia?: IMediaResource
    backMedia?: IMediaResource
}
