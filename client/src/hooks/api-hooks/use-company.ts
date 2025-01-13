import {
    useQuery,
    useMutation,
    queryOptions,
    useQueryClient,
} from '@tanstack/react-query'
import { toast } from 'sonner'

import {
    MediaRequest,
    CompanyResource,
    CompanyPaginatedResource,
    CompanyRequest,
} from '@/horizon-corp/types'
import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import { CompanyService } from '@/horizon-corp/services'
import {
    IApiPreloads,
    IFilterPaginatedHookProps,
    IOperationCallbacks,
} from './types'

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

// load/get company by id
export const useCompany = ({
    companyId,
    preloads = ['Media', 'Owner', 'Owner.Media'],
    onError,
    onSuccess,
}: { companyId: number } & IApiPreloads &
    IOperationCallbacks<CompanyResource, string>) => {
    return useQuery<CompanyResource, string>({
        queryKey: ['company', companyId],
        queryFn: async () => {
            const [error, data] = await withCatchAsync(
                CompanyService.getById(companyId, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            onSuccess?.(data)
            return data
        },
        enabled: companyId !== undefined && companyId !== null,
    })
}

// create company
export const useCreateCompany = ({
    onError,
    onSuccess,
}: IOperationCallbacks<CompanyResource>) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, CompanyRequest>({
        mutationKey: ['company', 'create'],
        mutationFn: async (newCompanyData) => {
            const [error, data] = await withCatchAsync(
                CompanyService.create(newCompanyData)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.setQueryData<CompanyResource>(
                ['company', data.id],
                data
            )

            queryClient.setQueryData<CompanyResource>(
                ['company', 'loader', data.id],
                data
            )

            toast.success('Company Created')
            onSuccess?.(data)
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

            queryClient.setQueriesData<CompanyPaginatedResource>(
                { queryKey: ['company', 'resource-query'], exact: false },
                (oldData) => {
                    if (!oldData) return oldData

                    return {
                        ...oldData,
                        data: oldData.data.map((company) =>
                            company.id === companyId ? data : company
                        ),
                    }
                }
            )

            queryClient.setQueryData<CompanyResource>(
                ['company', companyId],
                data
            )
            queryClient.setQueryData<CompanyResource>(
                ['company', 'loader', companyId],
                data
            )

            toast.success('Company approved')
            onSuccess?.(data)
        },
    })
}

// update company
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

            queryClient.setQueriesData<CompanyPaginatedResource>(
                { queryKey: ['company', 'resource-query'], exact: false },
                (oldData) => {
                    if (!oldData) return oldData

                    return {
                        ...oldData,
                        data: oldData.data.map((company) =>
                            company.id === id ? response : company
                        ),
                    }
                }
            )

            queryClient.setQueryData<CompanyResource>(['company', id], response)

            queryClient.setQueryData<CompanyResource>(
                ['company', 'loader', id],
                response
            )

            onSuccess?.(response)
            return response
        },
    })
}

// update company logo
export const useUpdateCompanyProfilePicture = ({
    onSuccess,
    onError,
}: IOperationCallbacks<CompanyResource>) => {
    const queryClient = useQueryClient()

    return useMutation<
        void,
        string,
        { companyId: number; mediaResource: MediaRequest }
    >({
        mutationKey: ['company', 'update', 'logo'],
        mutationFn: async ({ companyId, mediaResource }) => {
            const [error, data] = await withCatchAsync(
                CompanyService.ProfilePicture(companyId, mediaResource)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(`Failed to update company : ${errorMessage}`)
                onError?.(errorMessage)
                throw new Error(errorMessage)
            }

            queryClient.setQueriesData<CompanyPaginatedResource>(
                { queryKey: ['company', 'resource-query'], exact: false },
                (oldData) => {
                    if (!oldData) return oldData

                    return {
                        ...oldData,
                        data: oldData.data.map((company) =>
                            company.id === companyId ? data : company
                        ),
                    }
                }
            )

            queryClient.setQueryData<CompanyResource>(
                ['company', companyId],
                data
            )
            queryClient.setQueryData<CompanyResource>(
                ['company', companyId],
                data
            )

            toast.success('Company Logo Updated')
            onSuccess?.(data)
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

            queryClient.invalidateQueries({
                queryKey: ['company', 'resource-query'],
            })

            queryClient.invalidateQueries({ queryKey: ['company', companyId] })
            queryClient.removeQueries({
                queryKey: ['company', 'loader', companyId],
            })

            toast.success('Company deleted')
            onSuccess?.(undefined)
        },
    })
}

export const useFilteredPaginatedCompanies = ({
    sort,
    filterPayload,
    preloads = ['Media', 'Owner'],
    pagination = { pageSize: 10, pageIndex: 1 },
}: IFilterPaginatedHookProps = {}) => {
    return useQuery<CompanyPaginatedResource, string>({
        queryKey: ['company', 'resource-query', filterPayload, pagination, sort],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                CompanyService.getCompanies({
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
