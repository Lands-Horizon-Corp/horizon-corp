import { TEntityId } from '@/server/types'

export interface ICheckClearingDepositRequest {
    checkNo: string
    bankId: TEntityId
    picture?: TEntityId
    amount: number
    checkDepositDate: string
}
