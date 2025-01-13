import {
    useQuery,
    useMutation,
    queryOptions,
    useQueryClient,
} from '@tanstack/react-query'
import { toast } from 'sonner'

import {
    FeedbackRequest,
    FeedbackResource,
    FeedbackPaginatedResource,
} from '@/horizon-corp/types'
import { FeedbackService } from '@/horizon-corp/services'
import {
    IApiPreloads,
    IFilterPaginatedHookProps,
    IOperationCallbacks,
} from './types'
import { serverRequestErrExtractor } from '@/helpers'
import { withCatchAsync, toBase64 } from '@/utils'

// Only used by path preloader
export const feedbackLoader = (feedbackId: number) =>
    queryOptions<FeedbackResource>({
        queryKey: ['feedback', 'loader', feedbackId],
        queryFn: async () => {
            const data = await FeedbackService.getById(feedbackId, ['User'])
            return data
        },
        retry: 0,
    })

// Load/get feedback by id
export const useFeedback = ({
    feedbackId,
    preloads = ['User'],
    onError,
    onSuccess,
}: { feedbackId: number } & IApiPreloads &
    IOperationCallbacks<FeedbackResource, string>) => {
    return useQuery<FeedbackResource, string>({
        queryKey: ['feedback', feedbackId],
        queryFn: async () => {
            const [error, data] = await withCatchAsync(
                FeedbackService.getById(feedbackId, preloads)
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
        enabled: feedbackId !== undefined && feedbackId !== null,
    })
}

// Create feedback
export const useCreateFeedback = ({
    onError,
    onSuccess,
}: IOperationCallbacks<FeedbackResource>) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, FeedbackRequest>({
        mutationKey: ['feedback', 'create'],
        mutationFn: async (feedbackData) => {
            const [error, data] = await withCatchAsync(
                FeedbackService.create(feedbackData)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['feedback', 'resource-query'],
            })

            toast.success('Feedback Created')
            onSuccess?.(data)
        },
    })
}

// Delete feedback
export const useDeleteFeedback = ({
    onError,
    onSuccess,
}: IOperationCallbacks) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, number>({
        mutationKey: ['feedback', 'delete'],
        mutationFn: async (feedbackId) => {
            const [error] = await withCatchAsync(
                FeedbackService.delete(feedbackId)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw new Error(errorMessage)
            }

            queryClient.invalidateQueries({
                queryKey: ['feedback', 'resource-query'],
            })

            toast.success('Feedback Deleted')
            onSuccess?.(undefined)
        },
    })
}

// feedbacks with pagination, filter, sort
export const useFilteredPaginatedFeedbacks = ({
    sort,
    filterPayload,
    preloads = ['User'],
    pagination = { pageSize: 10, pageIndex: 1 },
}: IFilterPaginatedHookProps = {}) => {
    return useQuery<FeedbackPaginatedResource, string>({
        queryKey: ['feedback', 'resource-query', filterPayload, pagination, sort],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                FeedbackService.getFeedbacks({
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
