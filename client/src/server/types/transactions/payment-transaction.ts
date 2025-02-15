import { TEntityId } from '../common'

export interface IPaymentTransactionRequest {
    memberId: TEntityId
    ORNumber: string
    memberType: string
    idNumber: string
    accountNumber: TEntityId
    amount: number
}

export interface IPaymentTransactionResource
    extends IPaymentTransactionRequest {
    id: TEntityId
    createdAt: Date
    updatedAt: Date
}
