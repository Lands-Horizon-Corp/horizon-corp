import { useQuery, useQueryClient } from '@tanstack/react-query'

import { withCatchAsync } from '@/lib'
import { UserData } from '@/horizon-corp/types'
import UserService from '@/horizon-corp/server/auth/UserService'
import { AxiosError } from 'axios'

const useCurrentUser = (options?: {
    onError?: (error: unknown) => void
    onUnauthorized?: () => void
    onSuccess?: (userData: UserData) => void
    retry?: number
}) => {
    const queryClient = useQueryClient()
    const query = useQuery<UserData | null>({
        queryKey: ['current-user'],
        queryFn: async () => {
            const [error, response] = await withCatchAsync(
                UserService.CurrentUser()
            )

            if (error) {
                if (error instanceof AxiosError && error.status === 401) {
                    options?.onUnauthorized?.()
                    return null
                }
                options?.onError?.(error)
                throw error
            }

            options?.onSuccess?.(response.data)
            return response.data
        },
        retry: options?.retry ?? 0,
    })

    const setCurrentUser = (newUserData: UserData | null) => {
        queryClient.setQueryData<UserData | null>(['current-user'], newUserData)
    }

    return { ...query, setCurrentUser }
}

export default useCurrentUser
