import { create } from 'zustand'
import { UserBase } from '@/types'
import { UserData } from '@/horizon-corp/types'

interface IAuthStore {
    loadingUser: boolean
    currentUser: UserData | null
    setCurrentUser: (user: UserData | null) => void
    setLoadingUser: (loading: boolean) => void
}

const useAuthStore = create<IAuthStore>((set) => ({
    loadingUser: false,
    currentUser: null,
    setCurrentUser: (user) => set({ currentUser: user }),
    setLoadingUser: (loading) => set({ loadingUser: loading }),
}))

export default useAuthStore
