import { toast } from 'sonner'
import { useMutation, useQueryClient } from '@tanstack/react-query'

import { withCatchAsync } from '@/utils'
import { IOperationCallbacks } from './types'
import { serverRequestErrExtractor } from '@/helpers'
import { IMediaResource, IUserData } from '@/server/types'
import ProfileService from '@/horizon-corp/services/auth/ProfileService'

export const useUserUpdateProfilePicture = ({
    onError,
    onSuccess,
    invalidateCurrentUser,
}: { invalidateCurrentUser?: boolean } & IOperationCallbacks<IUserData>) => {
    const queryClient = useQueryClient()

    return useMutation<IUserData, string, IMediaResource>({
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
