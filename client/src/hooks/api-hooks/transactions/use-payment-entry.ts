import PaymentService from '@/server/api-service/transactions/payments-entry'
import { IPaymentTransactionRequest } from '@/server/types/transactions/payment-transaction'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'

export const useProcessPayment = () => {
    const queryClient = useQueryClient()

    const createPayments = useMutation<
        void,
        string,
        IPaymentTransactionRequest
    >({
        mutationKey: ['payment', 'process'],
        mutationFn: async (paymentData) => {
            try {
                await PaymentService.create(paymentData)

                // Invalidate queries to update affected ledgers
                queryClient.invalidateQueries({
                    queryKey: ['daily-transaction-ledger'],
                })
                queryClient.invalidateQueries({ queryKey: ['transactions'] })
                queryClient.invalidateQueries({
                    queryKey: ['specific-account-ledger'],
                })
                queryClient.invalidateQueries({
                    queryKey: ['specific-member-ledger'],
                })
                queryClient.invalidateQueries({ queryKey: ['accounts'] })
                queryClient.invalidateQueries({
                    queryKey: ['general-accounts-history'],
                })

                toast.success('Payment successfully recorded')
            } catch (error) {
                toast.error('Payment processing failed')
                throw error
            }
        },
    })

    return {
        isPending: createPayments.isPending,
        isSuccess: createPayments.isSuccess,
        processPayment: createPayments.mutate,
    }
}
