import { useShortcut } from '@/components/use-shorcuts'
import { IPaymentsEntry } from '@/server/types/transactions/payments-entry'
import { PaymentsDataStore } from '@/store/transaction/payments-entry-store'
import { useMemo } from 'react'

type usePaymentsShortcutsTypes = Pick<
    PaymentsDataStore,
    'setSelectedMember' | 'selectedPayments'
> & {
    hadSelectedPayments: boolean
    handleSubmitPayment: (payments: IPaymentsEntry[]) => void
    handleOpenCreateModal: (type: 'payment' | 'deposit' | 'withdraw') => void
}

const usePaymentsShortcuts = ({
    hadSelectedPayments,
    setSelectedMember,
    handleSubmitPayment,
    selectedPayments,
    handleOpenCreateModal,
}: usePaymentsShortcutsTypes) => {
    const paymentsEntryShortcuts = useMemo(
        () => ({
            d: () => {
                if (hadSelectedPayments) {
                    setSelectedMember(null)
                }
            },
            s: () => {
                handleSubmitPayment(selectedPayments)
            },
            f1: () => {
                handleOpenCreateModal('payment')
            },
            f2: () => {
                handleOpenCreateModal('deposit')
            },
            f3: () => {
                handleOpenCreateModal('withdraw')
            },
        }),
        [
            hadSelectedPayments,
            setSelectedMember,
            handleSubmitPayment,
            selectedPayments,
            handleOpenCreateModal,
        ]
    )

    Object.entries(paymentsEntryShortcuts).forEach(([shorcut, callback]) => {
        useShortcut(shorcut, callback, {
            disableTextInputs: true,
            disableActiveButton: true,
        })
    })
}

export default usePaymentsShortcuts
