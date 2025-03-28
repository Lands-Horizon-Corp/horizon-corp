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
    const handleD = useMemo(
        () => () => {
            if (hadSelectedPayments) {
                setSelectedMember(null)
            }
        },
        [hadSelectedPayments, setSelectedMember]
    )

    const handleS = useMemo(
        () => () => {
            handleSubmitPayment(selectedPayments)
        },
        [handleSubmitPayment, selectedPayments]
    )

    const handleF1 = useMemo(
        () => () => {
            handleOpenCreateModal('payment')
        },
        [handleOpenCreateModal]
    )

    const handleF2 = useMemo(
        () => () => {
            handleOpenCreateModal('deposit')
        },
        [handleOpenCreateModal]
    )

    const handleF3 = useMemo(
        () => () => {
            handleOpenCreateModal('withdraw')
        },
        [handleOpenCreateModal]
    )

    useShortcut('d', handleD, {
        disableTextInputs: true,
        disableActiveButton: true,
    })
    useShortcut('s', handleS, {
        disableTextInputs: true,
        disableActiveButton: true,
    })
    useShortcut('f1', handleF1, {
        disableTextInputs: true,
        disableActiveButton: true,
    })
    useShortcut('f2', handleF2, {
        disableTextInputs: true,
        disableActiveButton: true,
    })
    useShortcut('f3', handleF3, {
        disableTextInputs: true,
        disableActiveButton: true,
    })
}

export default usePaymentsShortcuts
