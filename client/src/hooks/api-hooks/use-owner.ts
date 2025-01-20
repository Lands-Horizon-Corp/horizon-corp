import { toast } from 'sonner'
import { useQuery, useQueryClient } from '@tanstack/react-query'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import OwnerService from '@/server/api-service/owner-service'

import {
    IApiPreloads,
    IOperationCallbacks,
    IFilterPaginatedHookProps,
} from './types'
import { IOwnerPaginatedResource, IOwnerResource } from '@/server/types'

// Load Specific owner by ID
export const useOwner = ({
    ownerId,
    onError,
    preloads = ['Media'],
}: Omit<
    IOperationCallbacks<IOwnerResource, string> & IApiPreloads,
    'onSuccess'
> & {
    ownerId: number
}) => {
    const queryClient = useQueryClient()

    return useQuery<IOwnerResource>({
        queryKey: ['owner', ownerId],
        queryFn: async () => {
            const [error, data] = await withCatchAsync(
                OwnerService.getById(ownerId, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (onError) onError(errorMessage)
                else toast.error(errorMessage)
                throw errorMessage
            }

            queryClient.setQueryData(['owner', ownerId], data)
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
    return useQuery<IOwnerPaginatedResource, string>({
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
