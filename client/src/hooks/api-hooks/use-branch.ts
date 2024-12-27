import { toast } from 'sonner'
import {
    queryOptions,
    useMutation,
    useQuery,
    useQueryClient,
} from '@tanstack/react-query'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import BranchService from '@/horizon-corp/server/admin/BranchService'
import {
    BranchPaginatedResource,
    BranchRequest,
    BranchResource,
    MediaRequest,
} from '@/horizon-corp/types'
import { IFilterPaginatedHookProps, IOperationCallbacks } from './types'

// for route pathParam loader
export const branchLoader = (companyId: number) =>
    queryOptions<BranchResource>({
        queryKey: ['branch', 'loader', companyId],
        queryFn: async () => {
            const data = await BranchService.getById(companyId, [
                'Owner',
                'Owner.Media',
                'Media',
            ])
            return data
        },
        retry: 0,
    })

// Update branch
export const useUpdateBranch = ({
    onSuccess,
    onError,
}: IOperationCallbacks<BranchResource, string>) => {
    const queryClient = useQueryClient()

    return useMutation<
        BranchResource,
        string,
        {
            id: number
            data: BranchRequest
        }
    >({
        mutationKey: ['branch', 'update'],
        mutationFn: async ({ id, data }) => {
            const [error, response] = await withCatchAsync(
                BranchService.update(id, data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (onError) onError(errorMessage)
                else toast.error(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({ queryKey: ['branch', 'table'] })
            queryClient.invalidateQueries({ queryKey: ['branch', id] })
            queryClient.invalidateQueries({
                queryKey: ['branch', 'loader', id],
            })

            onSuccess?.(response)
            return response
        },
    })
}

// Update branch logo
export const useUpdateBranchProfilePicture = ({
    onSuccess,
    onError,
}: IOperationCallbacks<BranchResource>) => {
    const queryClient = useQueryClient()

    return useMutation<
        void,
        string,
        { branchId: number; mediaResource: MediaRequest }
    >({
        mutationKey: ['branch', 'update', 'logo'],
        mutationFn: async ({ branchId, mediaResource }) => {
            const [error, data] = await withCatchAsync(
                BranchService.ProfilePicture(branchId, mediaResource)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(`Failed to update branch: ${errorMessage}`)
                onError?.(errorMessage)
                throw new Error(errorMessage)
            }

            queryClient.invalidateQueries({ queryKey: ['branch', 'table'] })
            queryClient.invalidateQueries({ queryKey: ['branch', branchId] })
            queryClient.invalidateQueries({
                queryKey: ['branch', 'loader', branchId],
            })

            toast.success('Branch Logo Updated')
            onSuccess?.(data)
        },
    })
}

// Delete branch
export const useDeleteBranch = ({
    onSuccess,
    onError,
}: IOperationCallbacks) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, number>({
        mutationKey: ['branch', 'delete'],
        mutationFn: async (branchId) => {
            const [error] = await withCatchAsync(BranchService.delete(branchId))

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw new Error(errorMessage)
            }

            queryClient.invalidateQueries({ queryKey: ['branch', 'table'] })
            queryClient.invalidateQueries({ queryKey: ['branch', branchId] })
            queryClient.removeQueries({
                queryKey: ['branch', 'loader', branchId],
            })

            toast.success('Branch deleted')
            onSuccess?.(undefined)
        },
    })
}

export const useFilteredPaginatedBranch = ({
    filterPayload,
    pagination = { pageSize: 10, pageIndex: 1 },
    preloads = ['Media', 'Owner'],
}: IFilterPaginatedHookProps = {}) => {
    return useQuery<BranchPaginatedResource, string>({
        queryKey: ['table', 'branches', filterPayload, pagination],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                BranchService.getBranches({
                    preloads,
                    pagination,
                    filters: filterPayload
                        ? toBase64(filterPayload)
                        : undefined,
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
