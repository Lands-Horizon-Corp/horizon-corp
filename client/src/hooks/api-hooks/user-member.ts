import { toast } from 'sonner'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import MemberService from '@/server/api-service/member-service'

import { IMemberPaginatedResource } from '@/server/types'
import { IFilterPaginatedHookProps, IOperationCallbacks } from './types'

export const useDeleteMember = ({
    onSuccess,
    onError,
}: IOperationCallbacks) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, number>({
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
    filterPayload,
    preloads = [],
    pagination = { pageSize: 10, pageIndex: 1 },
}: IFilterPaginatedHookProps = {}) => {
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
        retry: 1,
    })
}
