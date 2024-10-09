import { create } from 'zustand'
import { UserBase } from '@/types'

interface IAuthStore {
    currentUser: UserBase | null
    setCurrentUser: (user: UserBase | null) => void
}

const useAuthStore = create<IAuthStore>((set) => ({
    currentUser: null,
    setCurrentUser: (user) => set({ currentUser: user }),
}))

export default useAuthStore
