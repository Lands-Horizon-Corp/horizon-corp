import { TEntityId } from '../common'

export interface IPaymentsEntryRequest {
    memberId: TEntityId
    ORNumber: string
    memberType: string
    idNumber: string
    accountNumber: TEntityId
    amount: number
}

export interface IPaymentsEntryResource
    extends IPaymentsEntryRequest {
    id: TEntityId
    createdAt: Date
    updatedAt: Date
}
