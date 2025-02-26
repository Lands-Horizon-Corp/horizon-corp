import { ITimeStamps, TEntityId } from '../common'
import { IPaginatedResult } from '../paginated-result'

export interface ITransactionTypeRequest {
    id: TEntityId
    name: string
    description: TEntityId
    cheque_id: TEntityId
}

export interface ITransactionTypeResource extends ITimeStamps {
    id: TEntityId
    name: string
    description: TEntityId
    cheque_id: TEntityId
}

export type ITransactionTypePaginatedResource =
    IPaginatedResult<ITransactionTypeResource>
