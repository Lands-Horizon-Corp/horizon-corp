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
    focusTypePayment: string | null

    setSelectedMember: (member: IMemberResource | null) => void
    setORNumber: (orNumber: string) => void
    setSelectedPayments: (payments: IPaymentsEntry[]) => void
    setSelectedAccounts: (accounts: IAccountsRequest | null) => void
    deletePaymentByIndex: (index: number) => void
    setFocusTypePayment: (payment: string) => void
}

export const usePaymentsDataStore = create<PaymentsDataStore>((set, get) => ({
    selectedMember: null,
    selectedAccounts: null,
    ORNumber: '',
    selectedPayments: [],
    focusTypePayment: 'payment',

    setSelectedMember: (member) => set({ selectedMember: member }),
    setORNumber: (orNumber) => set({ ORNumber: orNumber }),
    setSelectedPayments: (payments) => {
        set({ selectedPayments: payments })
    },
    setSelectedAccounts: (accounts: IAccountsRequest | null) =>
        set({ selectedAccounts: accounts }),
    deletePaymentByIndex: (index: number) => {
        const current = get().selectedPayments
        set({ selectedPayments: current.filter((_, i) => i !== index) })
    },
    setFocusTypePayment: (payment) => set({ focusTypePayment: payment }),
}))

// Payments Modal
export interface PaymentsModalStore {
    transactionType: string | null
    openPaymentsEntryModal: boolean
    openCheckClearingFormModal: boolean
    openWithdrawFormModal: boolean
    openDepositCheckClearingFormModal: boolean

    setTransactionType: (paymentType: string) => void
    setOpenPaymentsEntryModal: (isOpen: boolean) => void
    setOpenCheckClearingFormModal: (isOpen: boolean) => void
}

export const usePaymentsModalStore = create<PaymentsModalStore>((set) => ({
    transactionType: 'payment',
    openCheckClearingFormModal: false,
    openWithdrawFormModal: false,
    openDepositCheckClearingFormModal: false,
    openPaymentsEntryModal: false,

    setTransactionType: (paymentType) => set({ transactionType: paymentType }),
    setOpenPaymentsEntryModal: (isOpen) =>
        set({ openPaymentsEntryModal: isOpen }),
    setOpenCheckClearingFormModal: (isOpen) =>
        set({ openCheckClearingFormModal: isOpen }),
}))
