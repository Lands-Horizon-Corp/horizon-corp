import { ITimeStamps, TEntityId } from './common'
import { IMemberCenterHistoryResource } from './member-center-history'

export interface IMemberCenterResource extends ITimeStamps {
    id: TEntityId
    name: string
    description: string
    history?: IMemberCenterHistoryResource[]
}
