import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import {
    IAccountingLedgerPaginatedResource,
    IAccountingLedgerResource,
} from '@/server/types/accounts/accounting-ledger'
import {
    IAPIPreloads,
    IFilterPaginatedHookProps,
    IOperationCallbacks,
    IQueryProps,
} from '../types'
import IAccountingLedgerService from '@/server/api-service/transactions/accounting-ledger'
import { TEntityId } from '@/server'

export const useCreateAccountingLedger = ({
    preloads = [],
    onError,
    onSuccess,
}: IOperationCallbacks<IAccountingLedgerResource> & IAPIPreloads) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, IAccountingLedgerResource>({
        mutationKey: ['accounting-ledger', 'create'],
        mutationFn: async (newAccountingLedger) => {
            const [error, newLedger] = await withCatchAsync(
                IAccountingLedgerService.create(newAccountingLedger, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['accounting-ledger', 'resource-query'],
            })
            queryClient.invalidateQueries({
                queryKey: ['accounting-ledger', newLedger.id],
            })
            queryClient.removeQueries({
                queryKey: ['accounting-ledger', 'loader', newLedger.id],
            })

            onSuccess?.(newLedger)
        },
    })
}

export const useFilteredPaginatedIAccountingLedger = ({
    sort,
    enabled,
    filterPayload,
    preloads = [],
    showMessage = true,
    pagination = { pageSize: 10, pageIndex: 1 },
    memberProfileId,
}: IFilterPaginatedHookProps &
    IQueryProps & {
        memberProfileId?: TEntityId
    } = {}) => {
    return useQuery<IAccountingLedgerPaginatedResource, string>({
        queryKey: [
            'accounting-ledger',
            'resource-query',
            filterPayload,
            memberProfileId,
            pagination,
            sort,
        ],
        queryFn: async () => {
            const filters = {
                ...(filterPayload || {}),
                ...(memberProfileId && { memberProfileId: memberProfileId }),
            }

            const [error, result] = await withCatchAsync(
                IAccountingLedgerService.getLedgers({
                    preloads,
                    pagination,
                    sort: sort && toBase64(sort),
                    filters: Object.keys(filters).length
                        ? toBase64(JSON.stringify(filters))
                        : undefined,
                })
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (showMessage) toast.error(errorMessage)
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
