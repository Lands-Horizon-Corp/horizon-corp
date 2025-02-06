import {
    useQuery,
    useMutation,
    queryOptions,
    useQueryClient,
} from '@tanstack/react-query'
import { toast } from 'sonner'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import MemberTypeService from '@/server/api-service/member-services/member-type-service'

import {
    IQueryProps,
    IAPIPreloads,
    IOperationCallbacks,
    IFilterPaginatedHookProps,
} from '../types'
import {
    TEntityId,
    IMemberTypeRequest,
    IMemberTypeResource,
    TMemberTypePaginatedResource,
} from '@/server/types'

export const memberTypeLoader = (
    memberTypeId: TEntityId,
    preloads: string[] = []
) =>
    queryOptions<IMemberTypeResource>({
        queryKey: ['member-type', 'loader', memberTypeId],
        queryFn: async () => {
            const data = await MemberTypeService.getById(memberTypeId, preloads)
            return data
        },
        retry: 0,
    })

export const useCreateMemberType = ({
    preloads = [],
    onSuccess,
    onError,
}:
    | undefined
    | (IOperationCallbacks<IMemberTypeResource, string> &
          IAPIPreloads) = {}) => {
    const queryClient = useQueryClient()

    return useMutation<IMemberTypeResource, string, IMemberTypeRequest>({
        mutationKey: ['member-type', 'create'],
        mutationFn: async (data) => {
            const [error, newMemberType] = await withCatchAsync(
                MemberTypeService.create(data, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['member-type', 'resource-query'],
            })

            queryClient.invalidateQueries({
                queryKey: ['member-type', newMemberType.id],
            })
            queryClient.removeQueries({
                queryKey: ['member-type', 'loader', newMemberType.id],
            })

            toast.success('New Member Type Created')
            onSuccess?.(newMemberType)

            return newMemberType
        },
    })
}

export const useUpdateMemberType = ({
    preloads = ['Owner', 'Media', 'Owner.Media'],
    onSuccess,
    onError,
}: IOperationCallbacks<IMemberTypeResource> & IAPIPreloads) => {
    const queryClient = useQueryClient()

    return useMutation<
        void,
        string,
        { memberTypeId: TEntityId; data: IMemberTypeRequest }
    >({
        mutationKey: ['member-type', 'update'],
        mutationFn: async ({ memberTypeId, data }) => {
            const [error, result] = await withCatchAsync(
                MemberTypeService.update(memberTypeId, data, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['member-type', 'resource-query'],
            })

            queryClient.invalidateQueries({
                queryKey: ['member-type', memberTypeId],
            })
            queryClient.removeQueries({
                queryKey: ['member-type', 'loader', memberTypeId],
            })

            toast.success('Member Type updated')
            onSuccess?.(result)
        },
    })
}

export const useDeleteMemberType = ({
    onSuccess,
    onError,
}: IOperationCallbacks) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, TEntityId>({
        mutationKey: ['member-type', 'delete'],
        mutationFn: async (memberTypeId) => {
            const [error] = await withCatchAsync(
                MemberTypeService.delete(memberTypeId)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['member-type', 'resource-query'],
            })

            queryClient.invalidateQueries({
                queryKey: ['member-type', memberTypeId],
            })
            queryClient.removeQueries({
                queryKey: ['member-type', 'loader', memberTypeId],
            })

            toast.success('Member Type deleted')
            onSuccess?.(undefined)
        },
    })
}

export const useFilteredPaginatedMemberTypes = ({
    sort,
    enabled,
    filterPayload,
    preloads = [],
    pagination = { pageSize: 10, pageIndex: 1 },
}: IFilterPaginatedHookProps & IQueryProps = {}) => {
    return useQuery<TMemberTypePaginatedResource, string>({
        queryKey: [
            'member-type',
            'resource-query',
            filterPayload,
            pagination,
            sort,
        ],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                MemberTypeService.getMemberTypes({
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
