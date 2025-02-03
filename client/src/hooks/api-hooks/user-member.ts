import {
    useQuery,
    useMutation,
    queryOptions,
    useQueryClient,
} from '@tanstack/react-query'
import { toast } from 'sonner'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import MemberService from '@/server/api-service/member-service'

import {
    IQueryProps,
    IAPIPreloads,
    IOperationCallbacks,
    IFilterPaginatedHookProps,
} from './types'
import {
    TEntityId,
    IMemberRequest,
    IMemberResource,
    IMemberPaginatedResource,
} from '@/server/types'

export const memberLoader = (
    memberId: TEntityId,
    preloads: string[] = ['Media']
) =>
    queryOptions<IMemberResource>({
        queryKey: ['member', 'loader', memberId],
        queryFn: async () => {
            const data = await MemberService.getById(memberId, preloads)
            return data
        },
        retry: 0,
    })

export const useCreateMember = ({
    preloads = ['Media'],
    onSuccess,
    onError,
}:
    | undefined
    | (IOperationCallbacks<IMemberResource, string> & IAPIPreloads) = {}) => {
    const queryClient = useQueryClient()

    return useMutation<IMemberResource, string, IMemberRequest>({
        mutationKey: ['member', 'create'],
        mutationFn: async (data) => {
            const [error, newMember] = await withCatchAsync(
                MemberService.create(data, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['member', 'resource-query'],
            })

            queryClient.invalidateQueries({
                queryKey: ['member', newMember.id],
            })
            queryClient.removeQueries({
                queryKey: ['member', 'loader', newMember.id],
            })

            toast.success('New Member Account Created')
            onSuccess?.(newMember)

            return newMember
        },
    })
}

export const useDeleteMember = ({
    onSuccess,
    onError,
}: IOperationCallbacks) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, TEntityId>({
        mutationKey: ['member', 'delete'],
        mutationFn: async (memberId) => {
            const [error] = await withCatchAsync(MemberService.delete(memberId))

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw new Error(errorMessage)
            }

            queryClient.invalidateQueries({
                queryKey: ['member', 'resource-query'],
            })

            queryClient.invalidateQueries({ queryKey: ['member', memberId] })
            queryClient.removeQueries({
                queryKey: ['member', 'loader', memberId],
            })

            toast.success('Member deleted')
            onSuccess?.(undefined)
        },
    })
}

export const useFilteredPaginatedMembers = ({
    sort,
    enabled,
    filterPayload,
    preloads = [],
    pagination = { pageSize: 10, pageIndex: 1 },
}: IFilterPaginatedHookProps & IQueryProps = {}) => {
    return useQuery<IMemberPaginatedResource, string>({
        queryKey: ['member', 'resource-query', filterPayload, pagination, sort],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                MemberService.getMembers({
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
