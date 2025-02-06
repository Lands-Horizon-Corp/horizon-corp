import { ITimeStamps, TEntityId } from '../common'

export interface IMemberTypeRequest {
    name: string
    prefix: string
    description: string
}

export interface IMemberTypeResource extends ITimeStamps {
    id: TEntityId

    name: string
    description: string
    prefix: string
}
