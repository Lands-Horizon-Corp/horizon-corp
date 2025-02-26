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
    ITransactionPaymentTypePaginatedResource,
    ITransactionPaymentTypesRequest,
    ITransactionPaymentTypesResource,
} from '@/server/types/transactions/transaction-payment-types'
import TransactionPaymentTypesService from '@/server/api-service/transactions/transaction-payment-types'
import { TEntityId } from '@/server/types'

export const useFilteredPaginatedTransactionPaymentTypes = ({
    sort,
    enabled,
    filterPayload,
    preloads,
    pagination = { pageSize: 20, pageIndex: 1 },
}: IFilterPaginatedHookProps & IQueryProps = {}) => {
    return useQuery<ITransactionPaymentTypePaginatedResource, string>({
        queryKey: [
            'transaction-types',
            'resource-query',
            filterPayload,
            pagination,
            sort,
        ],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                TransactionPaymentTypesService.getTransactionTypes({
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

export const useDeleteTransactionType = ({
    onSuccess,
    onError,
}: undefined | IOperationCallbacks<undefined> = {}) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, TEntityId>({
        mutationKey: ['transaction-types', 'delete'],
        mutationFn: async (transactionTypeId) => {
            const [error] = await withCatchAsync(
                TransactionPaymentTypesService.delete(transactionTypeId)
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
}: IOperationCallbacks<ITransactionPaymentTypesResource> & IAPIPreloads) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, ITransactionPaymentTypesRequest>({
        mutationKey: ['transaction-types', 'create'],
        mutationFn: async (newTransactionTypeData) => {
            const [error, data] = await withCatchAsync(
                TransactionPaymentTypesService.create(
                    newTransactionTypeData,
                    preloads
                )
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.setQueryData<ITransactionPaymentTypesResource>(
                ['transaction-types', data.id],
                data
            )

            queryClient.setQueryData<ITransactionPaymentTypesResource>(
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
}: IOperationCallbacks<ITransactionPaymentTypesResource, string> &
    IAPIPreloads) => {
    const queryClient = useQueryClient()

    return useMutation<
        ITransactionPaymentTypesResource,
        string,
        {
            id: TEntityId
            data: ITransactionPaymentTypesRequest
        }
    >({
        mutationKey: ['transaction-types', 'update'],
        mutationFn: async ({ id, data }) => {
            const [error, response] = await withCatchAsync(
                TransactionPaymentTypesService.update(id, data, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (onError) onError(errorMessage)
                else toast.error(errorMessage)
                throw errorMessage
            }

            queryClient.setQueriesData<ITransactionPaymentTypePaginatedResource>(
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

            queryClient.setQueryData<ITransactionPaymentTypesResource>(
                ['transaction-types', id],
                response
            )

            queryClient.setQueryData<ITransactionPaymentTypesResource>(
                ['transaction-types', 'loader', id],
                response
            )

            toast.success('Transaction Type Updated')

            onSuccess?.(response)
            return response
        },
    })
}
