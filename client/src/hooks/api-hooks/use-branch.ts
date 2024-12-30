import {
    useQuery,
    useMutation,
    queryOptions,
    useQueryClient,
} from '@tanstack/react-query'
import { toast } from 'sonner'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import BranchService from '@/horizon-corp/server/admin/BranchService'
import {
    MediaRequest,
    BranchRequest,
    BranchResource,
    BranchPaginatedResource,
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

            queryClient.setQueryData<BranchResource[]>(
                ['branch', 'table'],
                (oldData) => {
                    if (!oldData) return oldData
                    return oldData.map((branch) =>
                        branch.id === id ? response : branch
                    )
                }
            )
            queryClient.setQueryData<BranchResource>(['branch', id], response)
            queryClient.setQueryData<BranchResource>(
                ['branch', 'loader', id],
                response
            )

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

            queryClient.setQueryData<BranchResource[]>(
                ['branch', 'table'],
                (oldData) => {
                    if (!oldData) return oldData
                    return oldData.map((branch) =>
                        branch.id === branchId ? data : branch
                    )
                }
            )
            queryClient.setQueryData<BranchResource>(['branch', branchId], data)
            queryClient.setQueryData<BranchResource>(
                ['branch', 'loader', branchId],
                data
            )

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

            queryClient.setQueryData<BranchResource[]>(
                ['branch', 'table'],
                (oldData) => {
                    if (!oldData) return oldData
                    return oldData.filter((branch) => branch.id !== branchId)
                }
            )
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
        queryKey: ['branch', 'table', filterPayload, pagination],
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
