import {
    useQuery,
    useMutation,
    queryOptions,
    useQueryClient,
} from '@tanstack/react-query'
import { toast } from 'sonner'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import MemberService from '@/server/api-service/member-services/member-service'

import {
    TEntityId,
    IMemberRequest,
    IMemberResource,
    IMemberPaginatedResource,
} from '@/server/types'
import {
    IQueryProps,
    IAPIPreloads,
    IOperationCallbacks,
    IFilterPaginatedHookProps,
} from '../types'

export const memberLoader = (
    memberId: TEntityId,
    preloads: string[] = ['Media']
) =>
    queryOptions<IMemberResource>({
        queryKey: ['member', 'loader', memberId],
        queryFn: async () => {
            // const data = await MemberService.getById(memberId, preloads)
            // return data

            return {
                id: '550e8400-e29b-41d4-a716-446655440000',
                accountType: 'Member',
                username: 'john_doe',
                description: 'A passionate developer.',
                isEmailVerified: true,
                isContactVerified: false,
                isSkipVerification: false,
                status: 'Verified',
                firstName: 'John',
                lastName: 'Doe',
                middleName: 'A.',
                fullName: 'John A. Doe',
                birthDate: new Date('1990-05-15'),
                email: 'john.doe@example.com',
                contactNumber: '+1234567890',
                permanentAddress: '123 Main Street, Springfield, USA',
                mediaId: 1,
                media: {
                    id: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890',
                    fileName: 'profile.jpg',
                    fileSize: 204800,
                    fileType: 'image/jpeg',
                    storageKey: 'users/john_doe/profile.jpg',
                    url: 'https://example.com/storage/users/john_doe/profile.jpg',
                    bucketName: 'user-uploads',
                    createdAt: '2024-02-01T09:00:00Z',
                    updatedAt: '2024-02-01T09:00:00Z',
                    downloadURL:
                        'https://example.com/storage/users/john_doe/profile.jpg?download=true',
                },
                longitude: -73.935242,
                latitude: 40.73061,
                deletedAt: null,
                createdAt: '2024-02-01T08:30:00Z',
                updatedAt: '2024-02-01T08:30:00Z',
            }
        },
        retry: 0,
    })

export const useCreateMember = ({
    preloads = ['Media'],
    onSuccess,
    onError,
}:
    | undefined
    | (IOperationCallbacks<IMemberResource, string> & IAPIPreloads) = {}) => {
    const queryClient = useQueryClient()

    return useMutation<IMemberResource, string, IMemberRequest>({
        mutationKey: ['member', 'create'],
        mutationFn: async (data) => {
            const [error, newMember] = await withCatchAsync(
                MemberService.create(data, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['member', 'resource-query'],
            })

            queryClient.invalidateQueries({
                queryKey: ['member', newMember.id],
            })
            queryClient.removeQueries({
                queryKey: ['member', 'loader', newMember.id],
            })

            toast.success('New Member Account Created')
            onSuccess?.(newMember)

            return newMember
        },
    })
}

export const useDeleteMember = ({
    onSuccess,
    onError,
}: IOperationCallbacks) => {
    const queryClient = useQueryClient()

    return useMutation<void, string, TEntityId>({
        mutationKey: ['member', 'delete'],
        mutationFn: async (memberId) => {
            const [error] = await withCatchAsync(MemberService.delete(memberId))

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw new Error(errorMessage)
            }

            queryClient.invalidateQueries({
                queryKey: ['member', 'resource-query'],
            })

            queryClient.invalidateQueries({ queryKey: ['member', memberId] })
            queryClient.removeQueries({
                queryKey: ['member', 'loader', memberId],
            })

            toast.success('Member deleted')
            onSuccess?.(undefined)
        },
    })
}

export const useFilteredPaginatedMembers = ({
    sort,
    enabled,
    filterPayload,
    preloads = [],
    pagination = { pageSize: 10, pageIndex: 1 },
}: IFilterPaginatedHookProps & IQueryProps = {}) => {
    return useQuery<IMemberPaginatedResource, string>({
        queryKey: ['member', 'resource-query', filterPayload, pagination, sort],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                MemberService.getMembers({
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
            data: [ ],
            pages: [],
            totalSize: 0,
            totalPage: 1,
            ...pagination,
        },
        enabled,
        retry: 1,
    })
}
