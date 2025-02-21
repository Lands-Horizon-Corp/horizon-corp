import { toast } from 'sonner'
import { useMutation, useQueryClient } from '@tanstack/react-query'

import { withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'

import {
    IMemberProfileRequest,
    IMemberProfileResource,
} from '@/server/types/member/member-profile'
import { IAPIHook, IQueryProps } from '../types'
import MemberProfileService from '@/server/api-service/member-services/member-profile-service'
import { TEntityId } from '@/server'

export const useCreateMemberProfile = ({
    showMessage = true,
    preloads = ['Media'],
    onSuccess,
    onError,
}:
    | undefined
    | (IAPIHook<IMemberProfileResource, string> & IQueryProps) = {}) => {
    const queryClient = useQueryClient()

    return useMutation<IMemberProfileResource, string, IMemberProfileRequest>({
        mutationKey: ['member-profile', 'create'],
        mutationFn: async (data) => {
            const [error, newMember] = await withCatchAsync(
                MemberProfileService.create(data, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (showMessage) toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['member', 'resource-query'],
            })

            queryClient.setQueryData<IMemberProfileResource>(
                ['member-profile', newMember.id],
                newMember
            )

            queryClient.invalidateQueries({
                queryKey: ['member', newMember.id],
            })

            queryClient.removeQueries({
                queryKey: ['member', 'loader', newMember.id],
            })

            if (showMessage) toast.success('Member Profile Created')
            onSuccess?.(newMember)

            return newMember
        },
    })
}

export const useUpdateMemberProfile = ({
    showMessage = true,
    preloads = ['Media'],
    onSuccess,
    onError,
}:
    | undefined
    | (IAPIHook<IMemberProfileResource, string> & IQueryProps) = {}) => {
    const queryClient = useQueryClient()

    return useMutation<
        IMemberProfileResource,
        string,
        { id: TEntityId; data: IMemberProfileRequest }
    >({
        mutationKey: ['member-profile', 'update'],
        mutationFn: async ({ id, data }) => {
            const [error, updatedMember] = await withCatchAsync(
                MemberProfileService.update(id, data, preloads)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (showMessage) toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            queryClient.invalidateQueries({
                queryKey: ['member', 'resource-query'],
            })

            queryClient.setQueryData<IMemberProfileResource>(
                ['member-profile', updatedMember.id],
                updatedMember
            )

            queryClient.invalidateQueries({
                queryKey: ['member', updatedMember.id],
            })
            queryClient.removeQueries({
                queryKey: ['member', 'loader', updatedMember.id],
            })

            if (showMessage) toast.success('Member Profile Updated')
            onSuccess?.(updatedMember)

            return updatedMember
        },
    })
}
