import { TEntityId } from '@/server/types'

export interface IPaymentsEntryRequest {
    memberId: TEntityId
    ORNumber: string
    amount: number
    accountsId: TEntityId
    transactionType: TEntityId
    isPrinted: boolean
    notes: string
}
