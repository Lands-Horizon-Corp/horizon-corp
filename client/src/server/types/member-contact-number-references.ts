import { ITimeStamps, TEntityId } from './common'

export interface IMemberContactNumberReferencesResource extends ITimeStamps {
    id: TEntityId
    name: string
    description: string
    contactNumber: string
}

export interface IMemberContactNumberReferencesRequest {
    name: string
    description: string
    contactNumber: string
}

