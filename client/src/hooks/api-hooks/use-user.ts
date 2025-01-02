import { toast } from 'sonner'
import { useMutation, useQueryClient } from '@tanstack/react-query'

import { withCatchAsync } from '@/utils'
import { IOperationCallbacks } from './types'
import { serverRequestErrExtractor } from '@/helpers'
import { MediaResource, UserData } from '@/horizon-corp/types'
import ProfileService from '@/horizon-corp/server/auth/ProfileService'

export const useUserUpdateProfilePicture = ({
    onError,
    onSuccess,
    invalidateCurrentUser,
}: { invalidateCurrentUser?: boolean } & IOperationCallbacks<UserData>) => {
    const queryClient = useQueryClient()

    return useMutation<UserData, string, MediaResource>({
        mutationKey: ['user', 'profile'],
        mutationFn: async (data) => {
            const [error, response] = await withCatchAsync(
                ProfileService.ProfilePicture(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                onError?.(errorMessage)
                toast.error(errorMessage)
                throw errorMessage
            }

            if (invalidateCurrentUser)
                queryClient.invalidateQueries({ queryKey: ['current-user'] })

            onSuccess?.(response.data)
            return response.data
        },
    })
}
