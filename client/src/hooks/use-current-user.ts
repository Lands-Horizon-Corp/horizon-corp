import { useEffect } from 'react'
import useAuthStore from '@/store/auth-store'
import UserService from '@/horizon-corp/server/auth/UserService'

const useCurrentUser = ({ loadOnMount = false }: { loadOnMount?: boolean }) => {
    const { currentUser, setCurrentUser, loadingUser, setLoadingUser } =
        useAuthStore()

    useEffect(() => {
        if (!loadOnMount || loadingUser) return
        setLoadingUser(true)

        const fetchCurrentUser = async () => {
            try {
                const user = await UserService.currentUser()
                setCurrentUser(user.data)
            } catch (e) {
                setCurrentUser(null)
            } finally {
                setLoadingUser(false)
            }
        }

        fetchCurrentUser()
    }, [setCurrentUser, loadOnMount])

    return { currentUser, setCurrentUser, loadingUser }
}

export default useCurrentUser
