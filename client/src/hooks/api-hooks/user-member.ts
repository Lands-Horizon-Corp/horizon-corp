import {
    useQuery,
    useMutation,
    queryOptions,
    useQueryClient,
} from '@tanstack/react-query'
import { toast } from 'sonner'

import { toBase64, withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import MemberService from '@/server/api-service/member-service'

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
} from './types'

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
            data: [
                {
                    id: '0194b5f6-c5fd-7bfd-ae1b-542cf2b1055e',
                    accountType: 'Member',
                    username: 'john_smith',
                    description: 'Active community member',
                    isEmailVerified: true,
                    isContactVerified: true,
                    isSkipVerification: false,
                    role: undefined,
                    gender: { id: 1, name: 'Male' },
                    status: 'Verified',
                    footsteps: [],
                    firstName: 'John',
                    lastName: 'Smith',
                    middleName: 'A.',
                    fullName: 'John A. Smith',
                    birthDate: new Date('1990-05-10'),
                    email: 'john.smith@example.com',
                    contactNumber: '123-456-7890',
                    permanentAddress: '123 Main St, New York, NY',
                    mediaId: 1,
                    media: {
                        id: '101',
                        fileName: 'profile1.jpg',
                        fileSize: 512,
                        fileType: 'image/jpeg',
                        storageKey: 'media/profile1.jpg',
                        url: 'https://example.com/media/profile1.jpg',
                        bucketName: 'user-media',
                        createdAt: new Date().toISOString(),
                        updatedAt: new Date().toISOString(),
                        downloadURL: 'https://example.com/media/profile1.jpg',
                    },
                    longitude: -74.006,
                    latitude: 40.7128,
                    createdAt: new Date(),
                    updatedAt: new Date(),
                    deletedAt: null,
                },
                {
                    id: '0194b5f6-c5fd-7bf3-825b-f7323015d38a',
                    accountType: 'Member',
                    username: 'jane_doe',
                    description: 'Event organizer',
                    isEmailVerified: false,
                    isContactVerified: true,
                    isSkipVerification: false,
                    role: undefined,
                    gender: { id: 2, name: 'Female' },
                    status: 'Pending',
                    footsteps: [],
                    firstName: 'Jane',
                    lastName: 'Doe',
                    middleName: 'B.',
                    fullName: 'Jane B. Doe',
                    birthDate: new Date('1995-09-21'),
                    email: 'jane.doe@example.com',
                    contactNumber: '987-654-3210',
                    permanentAddress: '456 Broadway, Los Angeles, CA',
                    mediaId: 2,
                    media: {
                        id: '102',
                        fileName: 'profile2.jpg',
                        fileSize: 620,
                        fileType: 'image/jpeg',
                        storageKey: 'media/profile2.jpg',
                        url: 'https://example.com/media/profile2.jpg',
                        bucketName: 'user-media',
                        createdAt: new Date().toISOString(),
                        updatedAt: new Date().toISOString(),
                        downloadURL: 'https://example.com/media/profile2.jpg',
                    },
                    longitude: -118.2437,
                    latitude: 34.0522,
                    createdAt: new Date(),
                    updatedAt: new Date(),
                    deletedAt: null,
                },
                {
                    id: '0194b5f6-c5fd-7fa6-9ec8-4a77d0f5524f',
                    accountType: 'Member',
                    username: 'mike_jordan',
                    description: 'Basketball player',
                    isEmailVerified: true,
                    isContactVerified: false,
                    isSkipVerification: false,
                    role: undefined,
                    gender: { id: 1, name: 'Male' },
                    status: 'Not Allowed',
                    footsteps: [],
                    firstName: 'Michael',
                    lastName: 'Jordan',
                    middleName: 'C.',
                    fullName: 'Michael C. Jordan',
                    birthDate: new Date('1982-02-17'),
                    email: 'michael.jordan@example.com',
                    contactNumber: '555-432-6789',
                    permanentAddress: '789 Street, Chicago, IL',
                    mediaId: 3,
                    media: {
                        id: '103',
                        fileName: 'profile3.jpg',
                        fileSize: 700,
                        fileType: 'image/jpeg',
                        storageKey: 'media/profile3.jpg',
                        url: 'https://example.com/media/profile3.jpg',
                        bucketName: 'user-media',
                        createdAt: new Date().toISOString(),
                        updatedAt: new Date().toISOString(),
                        downloadURL: 'https://example.com/media/profile3.jpg',
                    },
                    longitude: -87.6298,
                    latitude: 41.8781,
                    createdAt: new Date(),
                    updatedAt: new Date(),
                    deletedAt: null,
                },
                {
                    id: '0194b5f6-c5fd-73b9-8135-e12613ecc40b',
                    accountType: 'Member',
                    username: 'chris_evans',
                    description: 'Captain America Fan',
                    isEmailVerified: true,
                    isContactVerified: true,
                    isSkipVerification: false,
                    role: undefined,
                    gender: { id: 1, name: 'Male' },
                    status: 'Verified',
                    footsteps: [],
                    firstName: 'Chris',
                    lastName: 'Evans',
                    middleName: 'D.',
                    fullName: 'Chris D. Evans',
                    birthDate: new Date('1981-06-13'),
                    email: 'chris.evans@example.com',
                    contactNumber: '321-678-9001',
                    permanentAddress: '987 Marvel St, Boston, MA',
                    mediaId: 4,
                    media: {
                        id: '104',
                        fileName: 'profile4.jpg',
                        fileSize: 550,
                        fileType: 'image/jpeg',
                        storageKey: 'media/profile4.jpg',
                        url: 'https://example.com/media/profile4.jpg',
                        bucketName: 'user-media',
                        createdAt: new Date().toISOString(),
                        updatedAt: new Date().toISOString(),
                        downloadURL: 'https://example.com/media/profile4.jpg',
                    },
                    longitude: -71.0589,
                    latitude: 42.3601,
                    createdAt: new Date(),
                    updatedAt: new Date(),
                    deletedAt: null,
                },
                {
                    id: '0194b5f6-c5fd-7f4f-83fa-65c7cca41955',
                    accountType: 'Member',
                    username: 'emma_watson',
                    description: 'Actress and activist',
                    isEmailVerified: true,
                    isContactVerified: true,
                    isSkipVerification: false,
                    role: undefined,
                    gender: { id: 2, name: 'Female' },
                    status: 'Verified',
                    footsteps: [],
                    firstName: 'Emma',
                    lastName: 'Watson',
                    middleName: 'E.',
                    fullName: 'Emma E. Watson',
                    birthDate: new Date('1990-04-15'),
                    email: 'emma.watson@example.com',
                    contactNumber: '888-555-1212',
                    permanentAddress: '221B Baker Street, London',
                    mediaId: 5,
                    media: {
                        id: '105',
                        fileName: 'profile5.jpg',
                        fileSize: 580,
                        fileType: 'image/jpeg',
                        storageKey: 'media/profile5.jpg',
                        url: 'https://example.com/media/profile5.jpg',
                        bucketName: 'user-media',
                        createdAt: new Date().toISOString(),
                        updatedAt: new Date().toISOString(),
                        downloadURL: 'https://example.com/media/profile5.jpg',
                    },
                    longitude: -0.1276,
                    latitude: 51.5074,
                    createdAt: new Date(),
                    updatedAt: new Date(),
                    deletedAt: null,
                },
            ],
            totalPage: 2,
            totalSize: 20,
            pages: [
                { page: '1', pageIndex: '1' },
                { page: '2', pageIndex: '2' },
            ],
            ...pagination,
        },
        enabled,
        retry: 1,
    })
}
