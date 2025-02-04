import { ITimeStamps, TEntityId } from './common'
import { IMemberProfileResource } from './member-profile'

export interface IMemberAddressResource extends ITimeStamps {
    id: TEntityId
    membersProfileID: TEntityId
    postalCode: string
    province: string
    city: string
    barangay: string
    region: string
    label: string
    membersProfile?: IMemberProfileResource
}

export interface IMemberAddressRequest {
    postalCode : string
    province : string
    city : string
    barangay : string
    label : string
}