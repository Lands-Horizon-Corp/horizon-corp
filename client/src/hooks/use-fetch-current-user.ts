import { AxiosError } from 'axios'
import { useQuery } from '@tanstack/react-query'

import { withCatchAsync } from '@/lib'
import { UserData } from '@/horizon-corp/types'
import { useUserAuthStore } from '@/store/user-auth-store'
import UserService from '@/horizon-corp/server/auth/UserService'

const useFetchCurrentUser = (options?: {
    onUnauthorized?: () => void
    onError?: (error: unknown) => void
    onSuccess?: (userData: UserData) => void
    retry?: number
    refetchOnWindowFocus?: boolean
}) => {
    const { currentUser, setAuthStatus, setCurrentUser } = useUserAuthStore()

    const query = useQuery<UserData | null>({
        queryKey: ['current-user'],
        queryFn: async () => {
            if (!currentUser) setAuthStatus('loading')

            const [error, response] = await withCatchAsync(
                UserService.CurrentUser()
            )

            if (error) {
                if (error instanceof AxiosError && error.status === 401) {
                    options?.onUnauthorized?.()
                    setCurrentUser(null)
                    setAuthStatus('authorized')
                    return null
                }

                if (error instanceof AxiosError && error.status === 500) {
                    options?.onError?.(error)
                    setAuthStatus('error')
                    return null
                }

                options?.onError?.(error)
                throw error
            }

            const userData = response.data
            setCurrentUser(userData)
            setAuthStatus('authorized')
            options?.onSuccess?.(userData)
            return userData
        },
        refetchOnWindowFocus: options?.refetchOnWindowFocus,
        retry: options?.retry ?? 0,
    })

    return { query }
}

export default useFetchCurrentUser
