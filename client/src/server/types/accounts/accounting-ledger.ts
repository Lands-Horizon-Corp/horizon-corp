import { TEntityId } from "../common";

export type AccountingLedger = {
  id: TEntityId;
  description: string;
  notes: string;
  company_id: TEntityId;
  member_profile_id: TEntityId;
  account_id: TEntityId;
  employee_id: TEntityId;
  or_number: string;
  transaction_source: TransactionSource;
  debit: number;
  credit: number;
  balance: Date;
  transaction_date: Date;
  entry_date: Date;
  line_number?: number | null;
  transaction_type_id: string;
  created_by: TEntityId;
  created_at: Date;
};

export enum TransactionSource {
  withdrawal = "withdrawal",
  deposit = "deposit",
  payment = "payment",
}

export type TransactionType = {
  id: TEntityId;
  description: string;
  cheque_id: TEntityId;
};

export enum TransactionTypeEnum {
  CASH = "CSH",
  CHEQUE = "CHQ",
  FWB = "Beginning",
}

