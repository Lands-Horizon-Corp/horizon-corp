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
    IAccountsPaginatedResource,
    IAccountsRequest,
    IAccountsResource,
} from '@/server/types/accounts/accounts'
import AccountsService from '@/server/api-service/accounting-services/accounts-service'
import { TEntityId } from '@/server'
import { SampleAccountsData } from '@/modules/owner/pages/transaction/transaction-dummy-data'

export const useFilteredPaginatedAccounts = ({
    sort,
    enabled = true,
    filterPayload,
    preloads,
    pagination = { pageSize: 20, pageIndex: 1 },
}: IFilterPaginatedHookProps & IQueryProps = {}) => {
    return useQuery<IAccountsPaginatedResource, string>({
        queryKey: [
            'accounts',
            'resource-query',
            filterPayload,
            pagination,
            sort,
        ],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                AccountsService.getAccounts({
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
            ...SampleAccountsData,
            ...pagination,
        },
        enabled,
        retry: 1,
    })
}

export const useDeleteAccounts = ({
    onSuccess,
    onError,
}: undefined | IOperationCallbacks<undefined> = {}) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, TEntityId>({
        mutationKey: ['accounts', 'delete'],
        mutationFn: async (accountsId) => {
            const [error] = await withCatchAsync(
                AccountsService.delete(accountsId)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw new Error(errorMessage)
            }

            queryClient.invalidateQueries({
                queryKey: ['accounts', 'resource-query'],
            })

            queryClient.invalidateQueries({
                queryKey: ['accounts', accountsId],
            })
            queryClient.removeQueries({
                queryKey: ['accounts', 'loader', accountsId],
            })

            toast.success('accounts deleted')
            onSuccess?.(undefined)
        },
    })
}

export const useCreateAccounts = ({
    preloads = [],
    onError,
    onSuccess,
}: IOperationCallbacks<IAccountsResource> & IAPIPreloads) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, IAccountsRequest>({
        mutationKey: ['accounts', 'create'],
        mutationFn: async (newAccountData) => {
            const [error, data] = await withCatchAsync(
                AccountsService.create(newAccountData, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.setQueryData<IAccountsResource>(
                ['accounts', data.id],
                data
            )

            queryClient.setQueryData<IAccountsResource>(
                ['accounts', 'loader', data.id],
                data
            )

            toast.success('Accounts Created')
            onSuccess?.(data)
        },
    })
}

export const useUpdateAccounts = ({
    preloads,
    onSuccess,
    onError,
}: IOperationCallbacks<IAccountsResource, string> & IAPIPreloads) => {
    const queryClient = useQueryClient()

    return useMutation<
        IAccountsResource,
        string,
        {
            id: TEntityId
            data: IAccountsRequest
        }
    >({
        mutationKey: ['accounts', 'update'],
        mutationFn: async ({ id, data }) => {
            const [error, response] = await withCatchAsync(
                AccountsService.update(id, data, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (onError) onError(errorMessage)
                else toast.error(errorMessage)
                throw errorMessage
            }

            queryClient.setQueriesData<IAccountsPaginatedResource>(
                { queryKey: ['accounts', 'resource-query'], exact: false },
                (oldData) => {
                    if (!oldData) return oldData

                    return {
                        ...oldData,
                        data: oldData.data.map((accounts) =>
                            accounts.id === id ? response : accounts
                        ),
                    }
                }
            )

            queryClient.setQueryData<IAccountsResource>(
                ['accounts', id],
                response
            )

            queryClient.setQueryData<IAccountsResource>(
                ['accounts', 'loader', id],
                response
            )

            toast.success('accounts updated')

            onSuccess?.(response)
            return response
        },
    })
}
