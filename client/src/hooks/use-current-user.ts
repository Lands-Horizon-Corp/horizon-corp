import { useQuery, useQueryClient } from '@tanstack/react-query'

import { withCatchAsync } from '@/lib'
import { UserData } from '@/horizon-corp/types'
import UserService from '@/horizon-corp/server/auth/UserService'

const useCurrentUser = (options?: {
    onError?: (error: unknown) => void
    onSuccess?: (userData: UserData) => void
}) => {
    const queryClient = useQueryClient()
    const query = useQuery<UserData | null>({
        queryKey: ['current-user'],
        queryFn: async () => {
            console.log('Fetching')
            const [error, response] = await withCatchAsync(
                UserService.CurrentUser()
            )

            if (error) {
                options?.onError?.(error)
                throw error
            }

            options?.onSuccess?.(response.data)
            return response.data
        },
        retry: 2,
    })

    const setCurrentUser = (newUserData: UserData | null) => {
        queryClient.setQueryData<UserData | null>(['current-user'], newUserData)
    }

    return { ...query, setCurrentUser }
}

export default useCurrentUser
