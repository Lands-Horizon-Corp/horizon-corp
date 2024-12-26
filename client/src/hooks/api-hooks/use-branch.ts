import { toast } from 'sonner'
import { queryOptions, useMutation, useQuery, useQueryClient } from '@tanstack/react-query'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import BranchService from '@/horizon-corp/server/admin/BranchService'
import { BranchPaginatedResource, BranchResource, MediaRequest } from '@/horizon-corp/types'
import { IOperationCallbacks } from './types'

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
            data: Omit<BranchResource, 'id' | 'createdAt' | 'updatedAt'>
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

export const useFilteredPaginatedBranch = (
    filterState: { finalFilters: Record<string, unknown> },
    pagination: { pageIndex: number; pageSize: number }
) => {
    return useQuery<BranchPaginatedResource, string>({
        queryKey: ['table', 'branches', filterState.finalFilters, pagination],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                BranchService.filter(
                    toBase64({
                        preloads: ['Media', 'Owner'],
                        ...pagination,
                        ...filterState.finalFilters,
                    })
                )
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
