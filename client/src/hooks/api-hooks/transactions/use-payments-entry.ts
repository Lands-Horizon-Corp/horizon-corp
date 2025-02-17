import PaymentsEntryService from '@/server/api-service/transactions/payments-entry'
import { IPaymentsEntryRequest } from '@/server/types/transactions/payments-Entry'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'

export const useProcessPayment = () => {
    const queryClient = useQueryClient()

    const createPayments = useMutation<void, string, IPaymentsEntryRequest>({
        mutationKey: ['payment', 'process'],
        mutationFn: async (paymentData) => {
            try {
                await PaymentsEntryService.create(paymentData)

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
