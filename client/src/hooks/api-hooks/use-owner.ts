import { toast } from 'sonner'
import { useQuery, useQueryClient } from '@tanstack/react-query'

import { withCatchAsync } from '@/utils'
import { IOperationCallbacks } from './types'
import { serverRequestErrExtractor } from '@/helpers'
import { OwnerResource } from '@/horizon-corp/types/profile'

// Load Specific owner by ID
export const useOwner = ({
    ownerId,
    onError,
}: Omit<IOperationCallbacks<OwnerResource, string>, 'onSuccess'> & {
    ownerId: number
}) => {
    const queryClient = useQueryClient()

    return useQuery<OwnerResource>({
        queryKey: ['owner', ownerId],
        queryFn: async () => {
            // TODO: Replace below once OwnerService is implemented
            const [error, data] = await withCatchAsync(
                {
                    getById: async (a: number) => {
                        if (a % 2 === 0) throw new Error('')
                        return {} as OwnerResource
                    },
                }.getById(ownerId)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (onError) onError(errorMessage)
                else toast.error(errorMessage)
                throw errorMessage
            }

            queryClient.setQueryData(['owner', 'loader', ownerId], data)

            return data
        },
        enabled: ownerId !== undefined && ownerId !== null,
    })
}
