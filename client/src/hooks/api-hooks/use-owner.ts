import { toast } from 'sonner'
import { useQuery, useQueryClient } from '@tanstack/react-query'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import OwnerService from '@/server/api-service/owner-service'

import {
    IAPIPreloads,
    IOperationCallbacks,
    IFilterPaginatedHookProps,
} from './types'
import {
    IOwnerResource,
    ICompanyResource,
    IOwnerPaginatedResource,
} from '@/server/types'

// Load Specific owner by ID
export const useOwner = ({
    ownerId,
    onError,
    preloads = ['Media'],
}: Omit<
    IOperationCallbacks<IOwnerResource, string> & IAPIPreloads,
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

export const useOwnerCompany = ({
    ownerId,
    preloads = ['Owner', 'Owner.Media'],
    onError,
    onSuccess,
}: { ownerId: number } & IOperationCallbacks<ICompanyResource> &
    IAPIPreloads) => {
    const queryClient = useQueryClient()

    return useQuery<ICompanyResource>({
        queryKey: ['owner', 'company', ownerId],
        queryFn: async () => {
            const [error, response] = await withCatchAsync(
                OwnerService.getCompany(ownerId, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (onError) onError(errorMessage)
                else toast.error(errorMessage)
                throw errorMessage
            }

            queryClient.setQueryData<ICompanyResource>(
                ['company', response.id],
                response
            )
            queryClient.setQueryData<ICompanyResource>(
                ['company', 'loader', response.id],
                response
            )

            onSuccess?.(response)
            return response
        },
        retry: 0,
        enabled: ownerId !== null || ownerId !== undefined,
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

