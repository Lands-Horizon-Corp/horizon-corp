import CheckClearingService from '@/server/api-service/transactions/check-clearing'
import { ICheckClearingRequest } from '@/server/types/transactions/check-clearing'
import { useMutation } from '@tanstack/react-query'
import { toast } from 'sonner'

export const useCreateCheckClearing = ({
    onSuccess,
    onError,
}: {
    onSuccess?: (data: ICheckClearingRequest[]) => void
    onError?: (error: string) => void
}) => {
    return useMutation<
        ICheckClearingRequest[],
        string,
        ICheckClearingRequest[]
    >({
        mutationKey: ['check-clearing', 'create'],
        mutationFn: async (paymentData) => {
            try {
                console.log(
                    '⏳ Sending request to CheckClearingService.create...'
                )
                const response = await CheckClearingService.create(paymentData)
                toast.success('Check clearing successfully created')
                onSuccess?.(response)
                return response
            } catch (error) {
                const errorMessage = 'Failed to process check clearing'
                console.error(errorMessage, error)
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }
        },
    })
}
