import { create } from 'zustand'
import { UserBase } from '@/types'

interface IAuthStore {
    loadingUser: boolean
    currentUser: UserBase | null
    setCurrentUser: (user: UserBase | null) => void
    setLoadingUser: (loading: boolean) => void
}

const useAuthStore = create<IAuthStore>((set) => ({
    loadingUser: false,
    currentUser: null,
    setCurrentUser: (user) => set({ currentUser: user }),
    setLoadingUser: (loading) => set({ loadingUser: loading }),
}))

export default useAuthStore
