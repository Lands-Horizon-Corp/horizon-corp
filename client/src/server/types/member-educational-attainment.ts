import { ITimeStamps, TEntityId } from './common'
import { IMemberEducationalAttainmentHistoryResource } from './member-educational-attainment-history'

export interface IMemberEducationalAttainmentResource extends ITimeStamps {
    id: TEntityId
    name: string
    description: string
    history?: IMemberEducationalAttainmentHistoryResource[]
}
