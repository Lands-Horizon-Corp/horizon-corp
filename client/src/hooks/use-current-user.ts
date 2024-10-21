import { useCallback, useEffect } from 'react'

import { withCatchAsync } from '@/lib'
import useAuthStore from '@/store/auth-store'
import { UserData } from '@/horizon-corp/types'
import UserService from '@/horizon-corp/server/auth/UserService'

const useCurrentUser = (opt?: { loadOnMount?: boolean }) => {
    console.log('use current user hook')
    const { currentUser, setCurrentUser, authState, setAuthState } =
        useAuthStore()

    const fetchCurrentUser = useCallback(
        async ({
            shouldSetToLoading = false,
            onError,
            onSuccess,
        }: {
            shouldSetToLoading?: boolean
            onError?: (error: unknown) => void
            onSuccess?: (userData: UserData) => void
        }) => {
            if (!opt?.loadOnMount) return

            if (shouldSetToLoading) setAuthState('Loading')

            const [error, response] = await withCatchAsync(
                UserService.CurrentUser()
            )

            if (shouldSetToLoading) setAuthState('Authenticated')

            if (error) {
                setAuthState('Error')
                setCurrentUser(null)
                onError?.(error)
                return
            }

            setCurrentUser(response.data)
            onSuccess?.(response.data)
        },
        [opt?.loadOnMount, setAuthState, setCurrentUser]
    )

    useEffect(() => {
        fetchCurrentUser({ shouldSetToLoading: true })
    }, [fetchCurrentUser])

    return { currentUser, authState, fetchCurrentUser, setCurrentUser }
}

export default useCurrentUser
