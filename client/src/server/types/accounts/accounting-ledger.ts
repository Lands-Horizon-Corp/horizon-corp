import { ITimeStamps, TEntityId } from '../common'
import { IPaginatedResult } from '../paginated-result'

interface IAccountLedgerSummaryResource {
    debit: number
    credit: number
    balance: number
}

export type IAccountingLedgerRequest = {
    id?: TEntityId
    description: string
    notes: string
    company_id: TEntityId
    member_profile_id: TEntityId
    account_id: TEntityId
    employee_id: TEntityId
    or_number: string
    transaction_source: TransactionSource
    debit: number
    credit: number
    balance: number
    transaction_date: Date
    entry_date: Date
    line_number?: number | null
    transaction_type_id: string
    created_by: TEntityId
    created_at: Date
}

export interface IAccountingLedgerResource extends ITimeStamps {
    id: TEntityId
    description: string
    notes: string
    company_id: TEntityId
    member_profile_id: TEntityId
    account_id: TEntityId
    employee_id: TEntityId
    or_number: string
    transaction_source: TransactionSource
    debit: number
    credit: number
    balance: number
    transaction_date: Date
    entry_date: Date
    line_number?: number | null
    transaction_type_id: string
    created_by: TEntityId
    created_at: Date
    summary: IAccountLedgerSummaryResource
}

export interface IAccountingLedgerPaginatedResource
    extends IPaginatedResult<IAccountingLedgerResource> {}

export enum TransactionSource {
    withdrawal = 'withdrawal',
    deposit = 'deposit',
    payment = 'payment',
}

export type TransactionType = {
    id: TEntityId
    description: string
    cheque_id: TEntityId
}

export enum TransactionTypeEnum {
    CASH = 'CSH',
    CHEQUE = 'CHQ',
    FWB = 'Beginning',
}
