import { toast } from 'sonner'
import { useQuery, useQueryClient } from '@tanstack/react-query'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import { OwnerResource } from '@/horizon-corp/types/profile'
import { IFilterPaginatedHookProps, IOperationCallbacks } from './types'

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
                OwnerService.getById(ownerId)
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

// paginated/ filtered and sorted 
export const useFilteredPaginatedOwners = ({
    sort,
    filterPayload,
    preloads = ['Media', 'Company'],
    pagination = { pageSize: 10, pageIndex: 1 },
}: IFilterPaginatedHookProps = {}) => {
    return useQuery<OwnerPaginatedResource, string>({
        queryKey: ['owner', 'resource-query', filterPayload, pagination],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                OwnerService.getOwners({
                    preloads,
                    pagination,
                    sort: sort && toBase64(sort),
                    filters: filterPayload && toBase64(filterPayload),
                })
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            return result
        },
        initialData: {
            data: [],
            pages: [],
            totalSize: 0,
            totalPage: 1,
            ...pagination,
        },
        retry: 1,
    })
}
