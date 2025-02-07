import {
    useQuery,
    useMutation,
    useQueryClient,
} from '@tanstack/react-query'
import { toast } from 'sonner'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import GenderService from '@/server/api-service/gender-service'

import {
    IQueryProps,
    IOperationCallbacks,
    IFilterPaginatedHookProps,
} from './types'
import {
    TEntityId,
    IGenderRequest,
    IGenderResource,
    IGenderPaginatedResource,
} from '@/server/types'

export const useCreateGender = ({
    onError,
    onSuccess,
}: IOperationCallbacks<IGenderResource, string> = {}) => {
    const queryClient = useQueryClient()

    return useMutation<IGenderResource, string, IGenderRequest>({
        mutationKey: ['gender', 'create'],
        mutationFn: async (data) => {
            const [error, newGender] = await withCatchAsync(
                GenderService.create(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['gender', 'resource-query'],
            })
            queryClient.invalidateQueries({
                queryKey: ['gender', newGender.id],
            })
            queryClient.removeQueries({
                queryKey: ['gender', 'loader', newGender.id],
            })

            toast.success('New Gender Created')
            onSuccess?.(newGender)

            return newGender
        },
    })
}

export const useUpdateGender = ({
    onError,
    onSuccess,
}: IOperationCallbacks<IGenderResource> = {}) => {
    const queryClient = useQueryClient()

    return useMutation<
        void,
        string,
        { genderId: TEntityId; data: IGenderRequest }
    >({
        mutationKey: ['gender', 'update'],
        mutationFn: async ({ genderId, data }) => {
            const [error, result] = await withCatchAsync(
                GenderService.update(genderId, data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['gender', 'resource-query'],
            })
            queryClient.invalidateQueries({
                queryKey: ['gender', genderId],
            })
            queryClient.removeQueries({
                queryKey: ['gender', 'loader', genderId],
            })

            toast.success('Gender updated')
            onSuccess?.(result)
        },
    })
}

export const useDeleteGender = ({
    onError,
    onSuccess,
}: IOperationCallbacks = {}) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, TEntityId>({
        mutationKey: ['gender', 'delete'],
        mutationFn: async (genderId) => {
            const [error] = await withCatchAsync(GenderService.delete(genderId))

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['gender', 'resource-query'],
            })
            queryClient.invalidateQueries({
                queryKey: ['gender', genderId],
            })
            queryClient.removeQueries({
                queryKey: ['gender', 'loader', genderId],
            })

            toast.success('Gender deleted')
            onSuccess?.(undefined)
        },
    })
}

export const useFilteredPaginatedGenders = ({
    sort,
    enabled,
    filterPayload,
    preloads = [],
    pagination = { pageSize: 10, pageIndex: 1 },
}: IFilterPaginatedHookProps & IQueryProps = {}) => {
    return useQuery<IGenderPaginatedResource, string>({
        queryKey: ['gender', 'resource-query', filterPayload, pagination, sort],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                GenderService.getGenders({
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
        enabled,
        retry: 1,
    })
}
