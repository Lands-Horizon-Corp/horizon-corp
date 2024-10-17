import { useCallback, useEffect } from 'react'
import useAuthStore from '@/store/auth-store'
import UserService from '@/horizon-corp/server/auth/UserService'

const useCurrentUser = (opt?: { loadOnMount?: boolean }) => {
    const { currentUser, setCurrentUser, loadingUser, setLoadingUser } =
        useAuthStore()

    const fetchCurrentUser = useCallback(async () => {
        if (!opt?.loadOnMount || loadingUser) return
        setLoadingUser(true)

        try {
            const userData = await UserService.CurrentUser()
            setCurrentUser(userData.data)
        } catch (e) {
            setCurrentUser(null)
        } finally {
            setLoadingUser(false)
        }
    }, [opt])

    useEffect(() => {
        fetchCurrentUser()
    }, [])

    return { currentUser, setCurrentUser, loadingUser }
}

export default useCurrentUser
