import { TEntityId } from '../common'

export interface IPaymentsEntryRequest {
    memberId: TEntityId
    ORNumber: string
    amount: number
    accountsId: TEntityId
    transactionType: TEntityId
    isPrinted: boolean
    notes: string
}

export interface IPaymentsEntryResource extends IPaymentsEntryRequest {
    id: TEntityId
    createdAt: Date
    updatedAt: Date
}
