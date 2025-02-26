import { serverRequestErrExtractor } from '@/helpers'
import { toBase64, withCatchAsync } from '@/utils'
import { useQueryClient, useMutation, useQuery } from '@tanstack/react-query'
import { toast } from 'sonner'
import {
    IOperationCallbacks,
    IAPIPreloads,
    IFilterPaginatedHookProps,
    IQueryProps,
} from '../types'
import {
    ITransactionTypePaginatedResource,
    ITransactionTypeRequest,
    ITransactionTypeResource,
} from '@/server/types/transactions/transaction-type'
import TransactionTypeService from '@/server/api-service/transactions/transaction-type'
import { TEntityId } from '@/server/types'
import { SampleTransactionTypeData } from '@/modules/owner/pages/transaction/transaction-dummy-data'

export const useFilteredPaginatedTransactionTypes = ({
    sort,
    enabled,
    filterPayload,
    preloads,
    pagination = { pageSize: 20, pageIndex: 1 },
}: IFilterPaginatedHookProps & IQueryProps = {}) => {
    return useQuery<ITransactionTypePaginatedResource, string>({
        queryKey: [
            'transaction-types',
            'resource-query',
            filterPayload,
            pagination,
            sort,
        ],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                TransactionTypeService.getTransactionTypes({
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
           ...SampleTransactionTypeData,
            ...pagination,
        },
        enabled,
        retry: 1,
    })
}

export const useDeleteTransactionType = ({
    onSuccess,
    onError,
}: undefined | IOperationCallbacks<undefined> = {}) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, TEntityId>({
        mutationKey: ['transaction-types', 'delete'],
        mutationFn: async (transactionTypeId) => {
            const [error] = await withCatchAsync(
                TransactionTypeService.delete(transactionTypeId)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw new Error(errorMessage)
            }

            queryClient.invalidateQueries({
                queryKey: ['transaction-types', 'resource-query'],
            })

            queryClient.invalidateQueries({
                queryKey: ['transaction-types', transactionTypeId],
            })
            queryClient.removeQueries({
                queryKey: ['transaction-types', 'loader', transactionTypeId],
            })

            toast.success('Transaction type deleted')
            onSuccess?.(undefined)
        },
    })
}

export const useCreateTransactionType = ({
    preloads = [],
    onError,
    onSuccess,
}: IOperationCallbacks<ITransactionTypeResource> & IAPIPreloads) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, ITransactionTypeRequest>({
        mutationKey: ['transaction-types', 'create'],
        mutationFn: async (newTransactionTypeData) => {
            const [error, data] = await withCatchAsync(
                TransactionTypeService.create(newTransactionTypeData, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.setQueryData<ITransactionTypeResource>(
                ['transaction-types', data.id],
                data
            )

            queryClient.setQueryData<ITransactionTypeResource>(
                ['transaction-types', 'loader', data.id],
                data
            )

            toast.success('Transaction Type Created')
            onSuccess?.(data)
        },
    })
}

export const useUpdateTransactionType = ({
    preloads,
    onSuccess,
    onError,
}: IOperationCallbacks<ITransactionTypeResource, string> & IAPIPreloads) => {
    const queryClient = useQueryClient()

    return useMutation<
        ITransactionTypeResource,
        string,
        {
            id: TEntityId
            data: ITransactionTypeRequest
        }
    >({
        mutationKey: ['transaction-types', 'update'],
        mutationFn: async ({ id, data }) => {
            const [error, response] = await withCatchAsync(
                TransactionTypeService.update(id, data, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (onError) onError(errorMessage)
                else toast.error(errorMessage)
                throw errorMessage
            }

            queryClient.setQueriesData<ITransactionTypePaginatedResource>(
                {
                    queryKey: ['transaction-types', 'resource-query'],
                    exact: false,
                },
                (oldData) => {
                    if (!oldData) return oldData

                    return {
                        ...oldData,
                        data: oldData.data.map((transType) =>
                            transType.id === id ? response : transType
                        ),
                    }
                }
            )

            queryClient.setQueryData<ITransactionTypeResource>(
                ['transaction-types', id],
                response
            )

            queryClient.setQueryData<ITransactionTypeResource>(
                ['transaction-types', 'loader', id],
                response
            )

            toast.success('Transaction Type Updated')

            onSuccess?.(response)
            return response
        },
    })
}
