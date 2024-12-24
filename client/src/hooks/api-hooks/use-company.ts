import { toast } from 'sonner'
import {
    queryOptions,
    useMutation,
    useQuery,
    useQueryClient,
} from '@tanstack/react-query'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'
import { CompanyPaginatedResource, CompanyResource } from '@/horizon-corp/types'

interface IOperationCallbacks<TDataSuccess = unknown, TError = unknown> {
    onSuccess?: (data: TDataSuccess) => void
    onError?: (error: TError) => void
}

// Only used by path preloader
export const companyLoader = (companyId: number) =>
    queryOptions<CompanyResource>({
        queryKey: ['company', 'loader', companyId],
        queryFn: async () => {
            const data = await CompanyService.getById(companyId, [
                'Owner',
                'Owner.Media',
                'Media',
            ])
            return data
        },
        retry: 0,
    })

export const useUpdateCompany = ({
    onSuccess,
    onError,
}: IOperationCallbacks<CompanyResource, string>) => {
    const queryClient = useQueryClient()

    return useMutation<
        CompanyResource,
        string,
        {
            id: number
            data: Omit<CompanyResource, 'id' | 'createdAt' | 'updatedAt'>
        }
    >({
        mutationKey: ['company', 'update'],
        mutationFn: async ({ id, data }) => {
            const [error, response] = await withCatchAsync(
                CompanyService.update(id, data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (onError) onError(errorMessage)
                else toast.error(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({ queryKey: ['company', 'table'] })
            queryClient.invalidateQueries({ queryKey: ['company', id] })
            queryClient.invalidateQueries({
                queryKey: ['company', 'loader', id],
            })

            onSuccess?.(response)
            return response
        },
    })
}

// delete company
export const useDeleteCompany = ({
    onSuccess,
    onError,
}: IOperationCallbacks) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, number>({
        mutationKey: ['company', 'delete'],
        mutationFn: async (companyId) => {
            const [error] = await withCatchAsync(
                CompanyService.delete(companyId)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw new Error(errorMessage)
            }

            queryClient.invalidateQueries({ queryKey: ['company', 'table'] })
            queryClient.invalidateQueries({ queryKey: ['company', companyId] })
            queryClient.removeQueries({
                queryKey: ['company', 'loader', companyId],
            })

            toast.success('Company deleted')
            onSuccess?.(undefined)
        },
    })
}

// approve company
export const useApproveCompany = ({
    onSuccess,
    onError,
}: IOperationCallbacks<CompanyResource, string>) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, number>({
        mutationKey: ['company', 'approve'],
        mutationFn: async (companyId) => {
            const [error, data] = await withCatchAsync(
                CompanyService.verify(companyId)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({ queryKey: ['company', 'table'] })
            queryClient.invalidateQueries({ queryKey: ['company', companyId] })
            queryClient.invalidateQueries({
                queryKey: ['company', 'loader', companyId],
            })

            toast.success('Company approved')
            onSuccess?.(data)
        },
    })
}

// Used by tables with filter + pagination
export const useFilteredPaginatedCompanies = (
    filterState: { finalFilters: Record<string, unknown> },
    pagination: { pageIndex: number; pageSize: number }
) => {
    return useQuery<CompanyPaginatedResource, string>({
        queryKey: ['table', 'company', filterState.finalFilters, pagination],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                CompanyService.filter(
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
