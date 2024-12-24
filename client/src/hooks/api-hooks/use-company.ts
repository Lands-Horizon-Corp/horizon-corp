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

// Only used by path preloader
export const companyLoader = (companyId: number) =>
    queryOptions<CompanyResource>({
        queryKey: ['company', companyId],
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

export const useUpdateCompany = (
    onSuccess?: (data: CompanyResource) => void,
    onError?: (error: string) => void,
) => {
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

            onSuccess?.(response)
            return response
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
