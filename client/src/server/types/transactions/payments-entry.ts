import { TEntityId } from '@/server/types'
import { IAccountsRequest } from '@/server/types/accounts/accounts'

/**
 * Represents the finalized shape of a payment entry for submission or storage.
 */
export interface IPaymentsEntryRequest {
    ORNumber: string
    memberId: TEntityId
    amount: number
    accountsId: TEntityId
    transactionType: TEntityId
    isPrinted: boolean
    notes?: string
}

/**
 * Represents the client-side form data including the full account object.
 */
export interface IPaymentsEntry
    extends Omit<IPaymentsEntryRequest, 'accountsId'> {
    account: IAccountsRequest
}
