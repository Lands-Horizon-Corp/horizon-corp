import { create } from 'zustand'
import { IMemberResource } from '@/server/types'
import { IPaymentsEntry } from '@/server/types/transactions/payments-entry'
import { IAccountsRequest } from '@/server/types/accounts/accounts'

// Payments Data
export interface PaymentsDataStore {
    selectedMember: IMemberResource | null
    ORNumber: string
    selectedPayments: IPaymentsEntry[]
    selectedAccounts: IAccountsRequest | null

    setSelectedMember: (member: IMemberResource | null) => void
    setORNumber: (orNumber: string) => void
    setSelectedPayments: (payments: IPaymentsEntry[]) => void
    setSelectedAccounts: (accounts: IAccountsRequest) => void
    deletePaymentByIndex: (index: number) => void
}

export const usePaymentsDataStore = create<PaymentsDataStore>((set, get) => ({
    selectedMember: null,
    selectedAccounts: null,
    ORNumber: '',
    selectedPayments: [],

    setSelectedMember: (member) => set({ selectedMember: member }),
    setORNumber: (orNumber) => set({ ORNumber: orNumber }),
    setSelectedPayments: (payments) => set({ selectedPayments: payments }),
    setSelectedAccounts: (accounts) => set({ selectedAccounts: accounts }),
    deletePaymentByIndex: (index: number) => {
        const current = get().selectedPayments
        set({ selectedPayments: current.filter((_, i) => i !== index) })
    },
}))

// Payments Modal
export interface PaymentsModalStore {
    transactionType: string | null
    openPaymentsEntryModal: boolean
    openDepositPaymentModal: boolean
    openCheckClearingFormModal: boolean
    openWithdrawFormModal: boolean
    openDepositCheckClearingFormModal: boolean

    setTransactionType: (paymentType: string) => void
    setOpenPaymentsEntryModal: (isOpen: boolean) => void
    setOpenDepositPaymentModal: (isOpen: boolean) => void
    setOpenCheckClearingFormModal: (isOpen: boolean) => void
    setOpenWithdrawFormModal: (isOpen: boolean) => void
    setOpenDepositCheckClearingFormModal: (isOpen: boolean) => void
}

export const usePaymentsModalStore = create<PaymentsModalStore>((set) => ({
    transactionType: null,
    openDepositPaymentModal: false,
    openCheckClearingFormModal: false,
    openWithdrawFormModal: false,
    openDepositCheckClearingFormModal: false,
    openPaymentsEntryModal: false,

    setTransactionType: (paymentType) => set({ transactionType: paymentType }),
    setOpenPaymentsEntryModal: (isOpen) =>
        set({ openPaymentsEntryModal: isOpen }),
    setOpenDepositPaymentModal: (isOpen) =>
        set({ openDepositPaymentModal: isOpen }),
    setOpenCheckClearingFormModal: (isOpen) =>
        set({ openCheckClearingFormModal: isOpen }),
    setOpenWithdrawFormModal: (isOpen) =>
        set({ openWithdrawFormModal: isOpen }),
    setOpenDepositCheckClearingFormModal: (isOpen) =>
        set({ openDepositCheckClearingFormModal: isOpen }),
}))
