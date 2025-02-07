import { toast } from 'sonner'
import { useMutation, useQueryClient } from '@tanstack/react-query'

import { withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'

import {
    IMemberProfileRequest,
    IMemberProfileResource,
} from '@/server/types/member/member-profile'
import { IAPIPreloads, IOperationCallbacks } from '../types'
import MemberProfileService from '@/server/api-service/member-services/member-profile-service'

export const useCreateMemberProfile = ({
    preloads = ['Media'],
    onSuccess,
    onError,
}:
    | undefined
    | (IOperationCallbacks<IMemberProfileResource, string> &
          IAPIPreloads) = {}) => {
    const queryClient = useQueryClient()

    return useMutation<IMemberProfileResource, string, IMemberProfileRequest>({
        mutationKey: ['member-profile', 'create'],
        mutationFn: async (data) => {
            const [error, newMember] = await withCatchAsync(
                MemberProfileService.create(data, preloads)
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
